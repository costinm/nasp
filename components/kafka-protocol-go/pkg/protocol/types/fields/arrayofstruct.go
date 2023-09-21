//  Copyright (c) 2023 Cisco and/or its affiliates. All rights reserved.
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//        https://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package fields

import (
	"bytes"
	"encoding/json"
	"math"
	"strconv"
	"strings"

	typesbytes "github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/bytes"

	"github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types"

	"github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/varint"

	"emperror.dev/errors"
)

type ArrayOfStruct[T any, P StructType] struct {
	Context
}

func (f *ArrayOfStruct[T, P]) Read(buf *bytes.Reader, version int16) ([]T, error) {
	if !f.IsSupportedVersion(version) {
		return nil, nil
	}

	if f.IsFlexibleVersion(version) {
		return f.readCompactArray(buf, version)
	}

	return f.readArray(buf, version)
}

// SizeInBytes returns the size of data in bytes when it's serialized
func (f *ArrayOfStruct[T, P]) SizeInBytes(version int16, data []T) (int, error) {
	if !f.IsSupportedVersion(version) {
		return 0, nil
	}

	if data == nil && !f.IsNullableVersion(version) {
		return 0, errors.New("non-nullable array field was set to null")
	}

	if len(data) == 0 {
		var length int32
		if data == nil {
			length = -1
		}

		if f.IsFlexibleVersion(version) {
			return varint.Uint32Size(uint32(length + 1)), nil // bytes needed to serialize the value -1 or 0 as varint
		}
		return 4, nil // bytes needed to serialize the value -1 or 0 as int32
	}

	dataLen := len(data)
	if f.IsFlexibleVersion(version) {
		if int64(dataLen)+1 > math.MaxUint32 {
			return 0, errors.New(strings.Join([]string{"field of type array has invalid length:", strconv.Itoa(dataLen)}, " "))
		}
	} else {
		if int64(dataLen) > math.MaxInt32 {
			return 0, errors.New(strings.Join([]string{"field of type array has invalid length:", strconv.Itoa(dataLen)}, " "))
		}
	}

	itemsSize := 0
	for i := range data {
		p, ok := any(&data[i]).(P)
		if !ok {
			return 0, errors.New(strings.Join([]string{"array item", strconv.Itoa(i), "doesn't implement StructType interface"}, " "))
		}
		itemSize, err := p.SizeInBytes(version)
		if err != nil {
			return 0, errors.WrapIf(err, strings.Join([]string{"couldn't compute size of array item", strconv.Itoa(i)}, " "))
		}

		itemsSize += itemSize
	}

	if f.IsFlexibleVersion(version) {
		dataLenLength := varint.Uint32Size(uint32(dataLen) + 1)
		return dataLenLength + itemsSize, nil
	}

	dataLenLength := 4
	return dataLenLength + itemsSize, nil
}

func (f *ArrayOfStruct[T, P]) Write(w *typesbytes.SliceWriter, version int16, data []T) error {
	if !f.IsSupportedVersion(version) {
		return nil
	}

	if data == nil && !f.IsNullableVersion(version) {
		return errors.New("non-nullable array field was set to null")
	}

	if len(data) == 0 {
		var length int32
		if data == nil {
			length = -1
		}

		var err error
		if f.IsFlexibleVersion(version) {
			err = varint.WriteUint32(w, uint32(length+1))
		} else {
			err = types.WriteInt32(w, length)
		}

		if err != nil {
			return errors.WrapIf(err, "couldn't write length of null array field")
		}
		return nil
	}

	if f.IsFlexibleVersion(version) {
		return f.writeCompactArray(w, version, data)
	}

	return f.writeArray(w, version, data)
}

func (f *ArrayOfStruct[T, P]) readArray(buf *bytes.Reader, version int16) ([]T, error) {
	var length int32
	err := types.ReadInt32(buf, &length)
	if err != nil {
		return nil, errors.WrapIf(err, "unable to read array field length")
	}

	return f.readArrayValue(buf, version, int(length))
}

func (f *ArrayOfStruct[T, P]) readCompactArray(buf *bytes.Reader, version int16) ([]T, error) {
	length, err := varint.ReadUint32(buf)
	if err != nil {
		return nil, errors.WrapIf(err, "unable to read compact array field length")
	}

	return f.readArrayValue(buf, version, int(length)-1)
}

func (f *ArrayOfStruct[T, P]) readArrayValue(
	buf *bytes.Reader,
	version int16,
	length int) ([]T, error) {
	if length < 0 {
		if f.IsNullableVersion(version) {
			return []T(nil), nil
		}

		return nil, errors.New("non-nullable array field was serialized as null")
	}
	items := make([]T, length)

	for i := range items {
		p, ok := any(&items[i]).(P)
		if !ok {
			return nil, errors.New(strings.Join([]string{"array item", strconv.Itoa(i), "doesn't implement StructType interface"}, " "))
		}
		err := p.Read(buf, version)
		if err != nil {
			return nil, errors.WrapIf(err, strings.Join([]string{"unable to read array item", strconv.Itoa(i)}, " "))
		}
	}

	return items, nil
}

func (f *ArrayOfStruct[T, P]) writeArray(w *typesbytes.SliceWriter, version int16, data []T) error {
	length := len(data)
	err := types.WriteInt32(w, int32(length))
	if err != nil {
		return errors.WrapIf(err, "couldn't write length to byte buffer")
	}
	if length == 0 {
		return nil
	}

	return f.writeArrayValue(w, version, data)
}

func (f *ArrayOfStruct[T, P]) writeCompactArray(w *typesbytes.SliceWriter, version int16, data []T) error {
	length := len(data)
	err := varint.WriteUint32(w, uint32(length+1))
	if err != nil {
		return errors.WrapIf(err, "couldn't write length to byte buffer")
	}
	if length == 0 {
		return nil
	}

	return f.writeArrayValue(w, version, data)
}

func (f *ArrayOfStruct[T, P]) writeArrayValue(w *typesbytes.SliceWriter, version int16, data []T) error {
	for i := range data {
		p, ok := any(&data[i]).(P)
		if !ok {
			return errors.New(strings.Join([]string{"array item", strconv.Itoa(i), "doesn't implement StructType interface"}, " "))
		}

		if err := p.Write(w, version); err != nil {
			return errors.WrapIf(err, strings.Join([]string{"couldn't serialize array item", strconv.Itoa(i)}, " "))
		}
	}

	return nil
}

func ArrayOfStructMarshalJSON[T any](key string, data []T) ([]byte, error) {
	if data == nil {
		return []byte("\"" + key + "\": null"), nil
	}

	a := make([][]byte, 0, len(data))
	for i := range data {
		p, ok := any(&data[i]).(json.Marshaler)
		if !ok {
			return nil, errors.New(strings.Join([]string{"array item", strconv.Itoa(i), "doesn't implement json.Marshaler interface"}, " "))
		}
		j, err := p.MarshalJSON()
		if err != nil {
			return nil, err
		}
		a = append(a, j)
	}
	var arr bytes.Buffer
	if _, err := arr.WriteString("\"" + key + "\": "); err != nil {
		return nil, err
	}
	if err := arr.WriteByte('['); err != nil {
		return nil, err
	}
	if _, err := arr.Write(bytes.Join(a, []byte(", "))); err != nil {
		return nil, err
	}
	if err := arr.WriteByte(']'); err != nil {
		return nil, err
	}

	return arr.Bytes(), nil
}
