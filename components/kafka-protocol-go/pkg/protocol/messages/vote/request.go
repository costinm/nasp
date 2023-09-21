// Code generated by kafka-protocol-go. DO NOT EDIT.

// Copyright (c) 2023 Cisco and/or its affiliates. All rights reserved.
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//	     https://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License.
package vote

import (
	"bytes"
	"strconv"
	"strings"

	"emperror.dev/errors"
	"github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/messages/vote/request"
	typesbytes "github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/bytes"
	"github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/fields"
	"github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/varint"
)

var requestClusterId = fields.Context{
	SpecName:           "ClusterId",
	CustomDefaultValue: "null",

	LowestSupportedVersion:          0,
	HighestSupportedVersion:         32767,
	LowestSupportedFlexVersion:      0,
	HighestSupportedFlexVersion:     32767,
	LowestSupportedNullableVersion:  0,
	HighestSupportedNullableVersion: 32767,
}
var requestTopics = fields.Context{
	SpecName:                    "Topics",
	LowestSupportedVersion:      0,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  0,
	HighestSupportedFlexVersion: 32767,
}

type Request struct {
	topics              []request.TopicData
	unknownTaggedFields []fields.RawTaggedField
	clusterId           fields.NullableString
	isNil               bool
}

func (o *Request) ClusterId() fields.NullableString {
	return o.clusterId
}

func (o *Request) SetClusterId(val fields.NullableString) {
	o.isNil = false
	o.clusterId = val
}

func (o *Request) Topics() []request.TopicData {
	return o.topics
}

func (o *Request) SetTopics(val []request.TopicData) {
	o.isNil = false
	o.topics = val
}

func (o *Request) ApiKey() int16 {
	return 52
}

func (o *Request) UnknownTaggedFields() []fields.RawTaggedField {
	return o.unknownTaggedFields
}

func (o *Request) SetUnknownTaggedFields(val []fields.RawTaggedField) {
	o.unknownTaggedFields = val
}

func (o *Request) Read(buf *bytes.Reader, version int16) error {
	o.SetDefault()

	clusterIdField := fields.String{Context: requestClusterId}
	if err := clusterIdField.Read(buf, version, &o.clusterId); err != nil {
		return errors.WrapIf(err, "couldn't set \"clusterId\" field")
	}

	topicsField := fields.ArrayOfStruct[request.TopicData, *request.TopicData]{Context: requestTopics}
	topics, err := topicsField.Read(buf, version)
	if err != nil {
		return errors.WrapIf(err, "couldn't set \"topics\" field")
	}
	o.topics = topics

	// process tagged fields

	if version < RequestLowestSupportedFlexVersion() || version > RequestHighestSupportedFlexVersion() {
		// tagged fields are only supported by flexible versions
		o.isNil = false
		return nil
	}

	if buf.Len() == 0 {
		o.isNil = false
		return nil
	}

	rawTaggedFields, err := fields.ReadRawTaggedFields(buf)
	if err != nil {
		return err
	}

	o.unknownTaggedFields = rawTaggedFields

	o.isNil = false
	return nil
}

func (o *Request) Write(buf *typesbytes.SliceWriter, version int16) error {
	if o.IsNil() {
		return nil
	}
	if err := o.validateNonIgnorableFields(version); err != nil {
		return err
	}

	clusterIdField := fields.String{Context: requestClusterId}
	if err := clusterIdField.Write(buf, version, o.clusterId); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"clusterId\" field")
	}

	topicsField := fields.ArrayOfStruct[request.TopicData, *request.TopicData]{Context: requestTopics}
	if err := topicsField.Write(buf, version, o.topics); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"topics\" field")
	}

	// serialize tagged fields
	numTaggedFields := o.getTaggedFieldsCount(version)
	if version < RequestLowestSupportedFlexVersion() || version > RequestHighestSupportedFlexVersion() {
		if numTaggedFields > 0 {
			return errors.New(strings.Join([]string{"tagged fields were set, but version", strconv.Itoa(int(version)), "of this message does not support them"}, " "))
		}

		return nil
	}

	rawTaggedFields := make([]fields.RawTaggedField, 0, numTaggedFields)
	rawTaggedFields = append(rawTaggedFields, o.unknownTaggedFields...)

	if err := fields.WriteRawTaggedFields(buf, rawTaggedFields); err != nil {
		return errors.WrapIf(err, "couldn't serialize tagged fields")
	}

	return nil
}

func (o *Request) String() string {
	s, err := o.MarshalJSON()
	if err != nil {
		return err.Error()
	}

	return string(s)
}

func (o *Request) MarshalJSON() ([]byte, error) {
	if o == nil || o.IsNil() {
		return []byte("null"), nil
	}

	s := make([][]byte, 0, 3)
	if b, err := fields.MarshalPrimitiveTypeJSON(o.clusterId); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"clusterId\""), b}, []byte(": ")))
	}
	if b, err := fields.ArrayOfStructMarshalJSON("topics", o.topics); err != nil {
		return nil, err
	} else {
		s = append(s, b)
	}

	if b, err := fields.ArrayOfStructMarshalJSON("unknownTaggedFields", o.unknownTaggedFields); err != nil {
		return nil, err
	} else {
		s = append(s, b)
	}

	var b bytes.Buffer
	if err := b.WriteByte('{'); err != nil {
		return nil, err
	}
	if _, err := b.Write(bytes.Join(s, []byte(", "))); err != nil {
		return nil, err
	}
	if err := b.WriteByte('}'); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (o *Request) IsNil() bool {
	return o.isNil
}

func (o *Request) Clear() {
	o.Release()
	o.isNil = true

	o.topics = nil
	o.unknownTaggedFields = nil
}

func (o *Request) SetDefault() {
	for i := range o.unknownTaggedFields {
		o.unknownTaggedFields[i].Release()
	}
	o.unknownTaggedFields = nil
	o.clusterId.SetValue("null")
	for i := range o.topics {
		o.topics[i].Release()
	}
	o.topics = nil

	o.isNil = false
}

func (o *Request) Equal(that *Request) bool {
	if !fields.RawTaggedFieldsEqual(o.unknownTaggedFields, that.unknownTaggedFields) {
		return false
	}

	if !o.clusterId.Equal(&that.clusterId) {
		return false
	}
	if len(o.topics) != len(that.topics) {
		return false
	}
	for i := range o.topics {
		if !o.topics[i].Equal(&that.topics[i]) {
			return false
		}
	}

	return true
}

// SizeInBytes returns the size of this data structure in bytes when it's serialized
func (o *Request) SizeInBytes(version int16) (int, error) {
	if o.IsNil() {
		return 0, nil
	}

	if err := o.validateNonIgnorableFields(version); err != nil {
		return 0, err
	}

	size := 0
	fieldSize := 0
	var err error

	clusterIdField := fields.String{Context: requestClusterId}
	fieldSize, err = clusterIdField.SizeInBytes(version, o.clusterId)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"clusterId\" field")
	}
	size += fieldSize

	topicsField := fields.ArrayOfStruct[request.TopicData, *request.TopicData]{Context: requestTopics}
	fieldSize, err = topicsField.SizeInBytes(version, o.topics)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"topics\" field")
	}
	size += fieldSize

	// tagged fields
	numTaggedFields := int64(o.getTaggedFieldsCount(version))
	if numTaggedFields > 0xffffffff {
		return 0, errors.New(strings.Join([]string{"invalid tagged fields count:", strconv.Itoa(int(numTaggedFields))}, " "))
	}
	if version < RequestLowestSupportedFlexVersion() || version > RequestHighestSupportedFlexVersion() {
		if numTaggedFields > 0 {
			return 0, errors.New(strings.Join([]string{"tagged fields were set, but version", strconv.Itoa(int(version)), "of this message does not support them"}, " "))
		}

		return size, nil
	}

	taggedFieldsSize := varint.Uint32Size(uint32(numTaggedFields)) // bytes for serializing the number of tagged fields

	for i := range o.unknownTaggedFields {
		length := len(o.unknownTaggedFields[i].Value())
		if int64(length) > 0xffffffff {
			return 0, errors.New(strings.Join([]string{"invalid field value length:", strconv.Itoa(length), ", tag:", strconv.Itoa(int(o.unknownTaggedFields[i].Tag()))}, " "))
		}
		taggedFieldsSize += varint.Uint32Size(o.unknownTaggedFields[i].Tag()) // bytes for serializing the tag of the unknown tag
		taggedFieldsSize += varint.Uint32Size(uint32(length))                 // bytes for serializing the length of the unknown tagged field
		taggedFieldsSize += length
	}

	size += taggedFieldsSize

	return size, nil
}

// Release releases the dynamically allocated fields of this object by returning then to object pools
func (o *Request) Release() {
	if o.IsNil() {
		return
	}

	for i := range o.unknownTaggedFields {
		o.unknownTaggedFields[i].Release()
	}
	o.unknownTaggedFields = nil

	o.clusterId.Release()
	for i := range o.topics {
		o.topics[i].Release()
	}
	o.topics = nil
}

func (o *Request) getTaggedFieldsCount(version int16) int {
	numTaggedFields := len(o.unknownTaggedFields)

	return numTaggedFields
}

// validateNonIgnorableFields throws an error if any non-ignorable field not supported by current version is set to
// non-default value
func (o *Request) validateNonIgnorableFields(version int16) error {
	return nil
}

func RequestLowestSupportedVersion() int16 {
	return 0
}

func RequestHighestSupportedVersion() int16 {
	return 0
}

func RequestLowestSupportedFlexVersion() int16 {
	return 0
}

func RequestHighestSupportedFlexVersion() int16 {
	return 32767
}

func RequestDefault() Request {
	var d Request
	d.SetDefault()

	return d
}
