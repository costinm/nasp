package server

import (
	"encoding/json"
	"io"
	"net"
	"sync"
	"time"

	"emperror.dev/errors"
	"github.com/pborman/uuid"
	"github.com/xtaci/smux"

	"github.com/cisco-open/nasp/pkg/tunnel/api"
)

type session struct {
	server     *server
	session    *smux.Session
	ctrlStream api.ControlStream

	ports map[int]*port

	backChannels map[string]chan io.ReadWriteCloser

	mu  sync.Mutex
	mu2 sync.RWMutex
}

func NewSession(srv *server, sess *smux.Session) *session {
	return &session{
		server:       srv,
		session:      sess,
		ports:        make(map[int]*port),
		mu:           sync.Mutex{},
		backChannels: make(map[string]chan io.ReadWriteCloser),
	}
}

func (s *session) Handle() error {
	for {
		stream, err := s.session.AcceptStream()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return errors.WrapIf(err, "could not accept stream")
		}

		go func() {
			if err := s.handleStream(stream); err != nil && !errors.Is(err, io.EOF) {
				s.server.logger.Error(err, "error during session handling")
			}
		}()
	}
}

func (s *session) Close() error {
	s.server.logger.V(2).Info("close session")

	s.mu.Lock()
	for p, mp := range s.ports {
		s.server.logger.V(2).Info("release port", "port", mp.servicePort)
		s.server.portProvider.ReleasePort(p)
		if err := mp.Close(); err != nil {
			s.server.logger.Error(err, "could not gracefully close managed port")
		}
		delete(s.ports, p)
	}
	s.mu.Unlock()

	return s.session.Close()
}

func (s *session) handleStream(stream *smux.Stream) error {
	s.server.logger.V(2).Info("handle stream", "id", stream.ID())

	var msg api.Message
	if err := json.NewDecoder(stream).Decode(&msg); err != nil {
		return err
	}

	switch msg.Type { //nolint:exhaustive
	case api.OpenControlStreamMessageType:
		if s.ctrlStream != nil {
			if _, _, err := api.SendMessage(stream, api.ErrCtrlStreamAlreadyExistsMessageType, nil); err != nil {
				return errors.WrapIfWithDetails(err, "could not send message", "type", api.ErrCtrlStreamAlreadyExistsMessageType)
			}

			return errors.New("control stream already established")
		}

		if _, _, err := api.SendMessage(stream, api.StreamOpenedResponseMessageType, nil); err != nil {
			return errors.WrapIfWithDetails(err, "could not send message", "type", api.StreamOpenedResponseMessageType)
		}

		s.ctrlStream = NewControlStream(s, stream)

		return s.ctrlStream.Handle()
	case api.OpenTCPStreamMessageType:
		var m api.OpenTCPStreamMessage
		if err := msg.Decode(&m); err != nil {
			return errors.WrapIfWithDetails(err, "could not decode message", "type", api.OpenTCPStreamMessageType)
		}

		s.mu2.Lock()
		if c, ok := s.backChannels[m.ID]; ok {
			delete(s.backChannels, m.ID)
			s.mu2.Unlock()
			c <- stream
		}

		if _, _, err := api.SendMessage(stream, api.StreamOpenedResponseMessageType, nil); err != nil {
			return errors.WrapIfWithDetails(err, "could not send message", "type", api.StreamOpenedResponseMessageType)
		}

		return nil
	default:
		if _, _, err := api.SendMessage(stream, api.ErrInvalidStreamTypeMessageType, nil); err != nil {
			return errors.WrapIfWithDetails(err, "could not send message", "type", api.ErrInvalidStreamTypeMessageType)
		}

		return errors.New("invalid stream type")
	}
}

func (s *session) AddPort(port int) int {
	p := s.server.portProvider.GetFreePort()

	s.server.logger.V(2).Info("get free port", "targetPort", port, "servicePort", p)

	mp := NewPort(s, p, port)

	s.mu.Lock()
	s.ports[p] = mp
	s.mu.Unlock()

	go func() {
		if err := mp.Listen(); err != nil {
			s.server.logger.Error(err, "could not listen")
		}
	}()

	return p
}

func (s *session) RequestConn(port int, c net.Conn) error {
	id := uuid.NewUUID().String()
	_, _, err := api.SendMessage(s.ctrlStream, api.RequestConnectionMessageType, &api.RequestConnectionMessage{
		Port:       port,
		Identifier: id,
	})
	if err != nil {
		return errors.WrapIfWithDetails(err, "could not send message", "type", api.RequestConnectionMessageType)
	}

	ch := make(chan io.ReadWriteCloser, 1)

	s.mu2.Lock()
	s.backChannels[id] = ch
	s.mu2.Unlock()

	var tcpStream io.ReadWriteCloser

	select {
	case tcpStream = <-ch:
	case <-time.After(time.Second * 1):
		return errors.WithStackIf(api.ErrSessionTimeout)
	}

	go func() {
		defer c.Close()

		var wg sync.WaitGroup
		wg.Add(2)

		go s.proxy(&wg, c, tcpStream)
		go s.proxy(&wg, tcpStream, c)

		wg.Wait()
	}()

	return err
}

func (s *session) proxy(wg *sync.WaitGroup, dst, src io.ReadWriteCloser) {
	defer wg.Done()

	if _, err := io.Copy(dst, src); err != nil {
		s.server.logger.Error(err, "could not copy data")
	}
}