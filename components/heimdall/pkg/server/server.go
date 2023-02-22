// Copyright (c) 2022 Cisco and/or its affiliates. All rights reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       https://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	clientconfig "sigs.k8s.io/controller-runtime/pkg/client/config"

	cluster_registry "github.com/cisco-open/cluster-registry-controller/api/v1alpha1"
	istio_ca "github.com/cisco-open/nasp/pkg/ca/istio"
	"github.com/cisco-open/nasp/pkg/environment"
)

var podNamespace = os.Getenv("POD_NAMESPACE")
var clusterID = os.Getenv("NASP_CLUSTER_ID")
var istioVersion = os.Getenv("NASP_ISTIO_VERSION")
var istioRevision = os.Getenv("NASP_ISTIO_REVISION")

var ErrClientNotFound = errors.New("client not found in database")
var ErrClusterIDNotFound = errors.New("clusterID not found")
var ErrClientOrClientSecretInvalid = errors.New("invalid ClientID or ClientSecret")
var ClientDatabaseConfigMap = types.NamespacedName{Namespace: podNamespace, Name: "heimdall-client-database"}

type ConfigRequest struct {
	ClientID     string `binding:"required"`
	ClientSecret string `binding:"required"`
	Version      string
}

type Client struct {
	ClientID           string
	ClientSecret       string
	WorkloadName       string
	PodNamespace       string
	Network            string
	MeshID             string
	ServiceName        string
	Version            string
	ServiceAccountName string
}

type ClientDatabase interface {
	Lookup(ClientID string) (*Client, error)
}

type ConfigMapClientDatabase struct {
	c client.Client
}

func NewConfigMapClientDatabase() (ClientDatabase, error) {
	kubeconfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	c, err := client.New(kubeconfig, client.Options{})
	if err != nil {
		return nil, err
	}

	// Sanity check
	var configMap corev1.ConfigMap
	err = c.Get(context.Background(), ClientDatabaseConfigMap, &configMap)
	if err != nil {
		return nil, err
	}

	return &ConfigMapClientDatabase{c: c}, nil
}

func (db *ConfigMapClientDatabase) Lookup(clientID string) (*Client, error) {
	var configMap corev1.ConfigMap
	err := db.c.Get(context.Background(), ClientDatabaseConfigMap, &configMap)
	if err != nil {
		return nil, err
	}

	if clientData, ok := configMap.Data[clientID]; ok {
		var client Client
		err = json.Unmarshal([]byte(clientData), &client)
		if err != nil {
			return nil, err
		}
		return &client, nil
	}

	return nil, ErrClientNotFound
}

type server struct {
	logger logr.Logger
}

func New(logger logr.Logger) *server {
	return &server{
		logger: logger,
	}
}

func (s *server) Run(ctx context.Context) error {
	srv, err := s.run()
	if err != nil {
		return err
	}

	<-ctx.Done()

	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (s *server) run() (*http.Server, error) {
	clientDb, err := NewConfigMapClientDatabase()
	if err != nil {
		return nil, err
	}

	if clusterID == "" {
		clusterID, err = getClusterIDFromRegistry()
		if err != nil {
			return nil, err
		}
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {})

	r.POST("/config", func(c *gin.Context) {
		var configRequest ConfigRequest
		if err := c.ShouldBindJSON(&configRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log := s.logger.WithValues("clientID", configRequest.ClientID)
		log.Info("client requesting config")

		client, err := clientDb.Lookup(configRequest.ClientID)
		if err != nil {
			log.Error(nil, "could not find client id")
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrClientOrClientSecretInvalid.Error()})
			return
		}

		if client.ClientSecret != configRequest.ClientSecret {
			log.Error(nil, "client secret mismatch")
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrClientOrClientSecretInvalid.Error()})
			return
		}

		client.Version = configRequest.Version
		if client.ServiceAccountName == "" {
			client.ServiceAccountName = "default"
		}

		var e IstioCAClientConfigAndEnvironment

		e.CAClientConfig, err = istio_ca.GetIstioCAClientConfigWithKubeConfig(clusterID, istioRevision, nil, &types.NamespacedName{
			Name:      client.ServiceAccountName,
			Namespace: client.PodNamespace,
		})
		if err != nil {
			_ = c.AbortWithError(500, err)
			return
		}

		clientIP := c.ClientIP()
		if clientIP == "::1" {
			clientIP = "127.0.0.1"
		}

		e.Environment = environment.IstioEnvironment{
			Type:              "sidecar",
			PodName:           client.WorkloadName + "-" + configRequest.ClientID,
			PodNamespace:      client.PodNamespace,
			PodOwner:          fmt.Sprintf("kubernetes://apis/v1/namespaces/%s/pods/%s", client.PodNamespace, client.WorkloadName),
			PodServiceAccount: client.ServiceAccountName,
			WorkloadName:      client.WorkloadName,
			AppContainers:     nil,
			InstanceIPs:       []string{clientIP},
			Labels: map[string]string{
				"security.istio.io/tlsMode":           "istio",
				"service.istio.io/canonical-revision": "latest",
				"istio.io/rev":                        e.CAClientConfig.Revision,
				"topology.istio.io/network":           e.Environment.Network,
				"k8s-app":                             client.WorkloadName,
				"service.istio.io/canonical-name":     client.ServiceName,
				"app":                                 client.WorkloadName,
				"version":                             client.Version,
			},
			PlatformMetadata: nil,
			Network:          client.Network,
			SearchDomains:    []string{"svc.cluster.local", "cluster.local"},
			ClusterID:        e.CAClientConfig.ClusterID,
			DNSDomain:        "cluster.local",
			MeshID:           client.MeshID,
			IstioCAAddress:   e.CAClientConfig.CAEndpoint,
			IstioVersion:     istioVersion,
			IstioRevision:    istioRevision,
		}

		c.JSON(http.StatusOK, e)
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server: %s\n", err)
		}
	}()

	return srv, nil
}

func getClusterIDFromRegistry() (string, error) {
	config, err := clientconfig.GetConfig()
	if err != nil {
		return "", err
	}

	k8sClient, err := client.New(config, client.Options{})
	if err != nil {
		return "", err
	}

	var clusters cluster_registry.ClusterList
	err = k8sClient.List(context.Background(), &clusters)
	if err != nil {
		return "", err
	}

	for _, cluster := range clusters.Items {
		if cluster.Status.Type == cluster_registry.ClusterTypeLocal {
			return cluster.Name, nil
		}
	}

	return "", ErrClusterIDNotFound
}

type IstioCAClientConfigAndEnvironment struct {
	CAClientConfig istio_ca.IstioCAClientConfig
	Environment    environment.IstioEnvironment
}
