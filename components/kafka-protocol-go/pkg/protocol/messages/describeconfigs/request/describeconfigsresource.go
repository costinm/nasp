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
package request

import (
	"bytes"
	"strconv"
	"strings"

	"emperror.dev/errors"
	typesbytes "github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/bytes"
	"github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/fields"
	"github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/varint"
)

var describeConfigsResourceResourceType = fields.Context{
	SpecName:                    "ResourceType",
	LowestSupportedVersion:      0,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  4,
	HighestSupportedFlexVersion: 32767,
}
var describeConfigsResourceResourceName = fields.Context{
	SpecName:                    "ResourceName",
	LowestSupportedVersion:      0,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  4,
	HighestSupportedFlexVersion: 32767,
}
var describeConfigsResourceConfigurationKeys = fields.Context{
	SpecName:                        "ConfigurationKeys",
	LowestSupportedVersion:          0,
	HighestSupportedVersion:         32767,
	LowestSupportedFlexVersion:      4,
	HighestSupportedFlexVersion:     32767,
	LowestSupportedNullableVersion:  0,
	HighestSupportedNullableVersion: 32767,
}

type DescribeConfigsResource struct {
	configurationKeys   []fields.NullableString
	unknownTaggedFields []fields.RawTaggedField
	resourceName        fields.NullableString
	resourceType        int8
	isNil               bool
}

func (o *DescribeConfigsResource) ResourceType() int8 {
	return o.resourceType
}

func (o *DescribeConfigsResource) SetResourceType(val int8) {
	o.isNil = false
	o.resourceType = val
}

func (o *DescribeConfigsResource) ResourceName() fields.NullableString {
	return o.resourceName
}

func (o *DescribeConfigsResource) SetResourceName(val fields.NullableString) {
	o.isNil = false
	o.resourceName = val
}

func (o *DescribeConfigsResource) ConfigurationKeys() []fields.NullableString {
	return o.configurationKeys
}

func (o *DescribeConfigsResource) SetConfigurationKeys(val []fields.NullableString) {
	o.isNil = false
	o.configurationKeys = val
}

func (o *DescribeConfigsResource) UnknownTaggedFields() []fields.RawTaggedField {
	return o.unknownTaggedFields
}

func (o *DescribeConfigsResource) SetUnknownTaggedFields(val []fields.RawTaggedField) {
	o.unknownTaggedFields = val
}

func (o *DescribeConfigsResource) Read(buf *bytes.Reader, version int16) error {
	o.SetDefault()

	resourceTypeField := fields.Int8{Context: describeConfigsResourceResourceType}
	if err := resourceTypeField.Read(buf, version, &o.resourceType); err != nil {
		return errors.WrapIf(err, "couldn't set \"resourceType\" field")
	}

	resourceNameField := fields.String{Context: describeConfigsResourceResourceName}
	if err := resourceNameField.Read(buf, version, &o.resourceName); err != nil {
		return errors.WrapIf(err, "couldn't set \"resourceName\" field")
	}

	configurationKeysField := fields.Array[fields.NullableString, *fields.String]{
		Context:          describeConfigsResourceConfigurationKeys,
		ElementProcessor: &fields.String{Context: describeConfigsResourceConfigurationKeys}}

	configurationKeys, err := configurationKeysField.Read(buf, version)
	if err != nil {
		return errors.WrapIf(err, "couldn't set \"configurationKeys\" field")
	}
	o.configurationKeys = configurationKeys

	// process tagged fields

	if version < DescribeConfigsResourceLowestSupportedFlexVersion() || version > DescribeConfigsResourceHighestSupportedFlexVersion() {
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

func (o *DescribeConfigsResource) Write(buf *typesbytes.SliceWriter, version int16) error {
	if o.IsNil() {
		return nil
	}
	if err := o.validateNonIgnorableFields(version); err != nil {
		return err
	}

	resourceTypeField := fields.Int8{Context: describeConfigsResourceResourceType}
	if err := resourceTypeField.Write(buf, version, o.resourceType); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"resourceType\" field")
	}
	resourceNameField := fields.String{Context: describeConfigsResourceResourceName}
	if err := resourceNameField.Write(buf, version, o.resourceName); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"resourceName\" field")
	}
	configurationKeysField := fields.Array[fields.NullableString, *fields.String]{
		Context:          describeConfigsResourceConfigurationKeys,
		ElementProcessor: &fields.String{Context: describeConfigsResourceConfigurationKeys}}
	if err := configurationKeysField.Write(buf, version, o.configurationKeys); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"configurationKeys\" field")
	}

	// serialize tagged fields
	numTaggedFields := o.getTaggedFieldsCount(version)
	if version < DescribeConfigsResourceLowestSupportedFlexVersion() || version > DescribeConfigsResourceHighestSupportedFlexVersion() {
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

func (o *DescribeConfigsResource) String() string {
	s, err := o.MarshalJSON()
	if err != nil {
		return err.Error()
	}

	return string(s)
}

func (o *DescribeConfigsResource) MarshalJSON() ([]byte, error) {
	if o == nil || o.IsNil() {
		return []byte("null"), nil
	}

	s := make([][]byte, 0, 4)
	if b, err := fields.MarshalPrimitiveTypeJSON(o.resourceType); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"resourceType\""), b}, []byte(": ")))
	}
	if b, err := fields.MarshalPrimitiveTypeJSON(o.resourceName); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"resourceName\""), b}, []byte(": ")))
	}
	if b, err := fields.ArrayMarshalJSON("configurationKeys", o.configurationKeys); err != nil {
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

func (o *DescribeConfigsResource) IsNil() bool {
	return o.isNil
}

func (o *DescribeConfigsResource) Clear() {
	o.Release()
	o.isNil = true

	o.configurationKeys = nil
	o.unknownTaggedFields = nil
}

func (o *DescribeConfigsResource) SetDefault() {
	for i := range o.unknownTaggedFields {
		o.unknownTaggedFields[i].Release()
	}
	o.unknownTaggedFields = nil
	o.resourceType = 0
	o.resourceName.SetValue("")
	for i := range o.configurationKeys {
		o.configurationKeys[i].Release()
	}
	o.configurationKeys = nil

	o.isNil = false
}

func (o *DescribeConfigsResource) Equal(that *DescribeConfigsResource) bool {
	if !fields.RawTaggedFieldsEqual(o.unknownTaggedFields, that.unknownTaggedFields) {
		return false
	}

	if o.resourceType != that.resourceType {
		return false
	}
	if !o.resourceName.Equal(&that.resourceName) {
		return false
	}
	if !fields.NullableStringSliceEqual(o.configurationKeys, that.configurationKeys) {
		return false
	}

	return true
}

// SizeInBytes returns the size of this data structure in bytes when it's serialized
func (o *DescribeConfigsResource) SizeInBytes(version int16) (int, error) {
	if o.IsNil() {
		return 0, nil
	}

	if err := o.validateNonIgnorableFields(version); err != nil {
		return 0, err
	}

	size := 0
	fieldSize := 0
	var err error

	resourceTypeField := fields.Int8{Context: describeConfigsResourceResourceType}
	fieldSize, err = resourceTypeField.SizeInBytes(version, o.resourceType)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"resourceType\" field")
	}
	size += fieldSize

	resourceNameField := fields.String{Context: describeConfigsResourceResourceName}
	fieldSize, err = resourceNameField.SizeInBytes(version, o.resourceName)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"resourceName\" field")
	}
	size += fieldSize

	configurationKeysField := fields.Array[fields.NullableString, *fields.String]{
		Context:          describeConfigsResourceConfigurationKeys,
		ElementProcessor: &fields.String{Context: describeConfigsResourceConfigurationKeys}}
	fieldSize, err = configurationKeysField.SizeInBytes(version, o.configurationKeys)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"configurationKeys\" field")
	}
	size += fieldSize

	// tagged fields
	numTaggedFields := int64(o.getTaggedFieldsCount(version))
	if numTaggedFields > 0xffffffff {
		return 0, errors.New(strings.Join([]string{"invalid tagged fields count:", strconv.Itoa(int(numTaggedFields))}, " "))
	}
	if version < DescribeConfigsResourceLowestSupportedFlexVersion() || version > DescribeConfigsResourceHighestSupportedFlexVersion() {
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
func (o *DescribeConfigsResource) Release() {
	if o.IsNil() {
		return
	}

	for i := range o.unknownTaggedFields {
		o.unknownTaggedFields[i].Release()
	}
	o.unknownTaggedFields = nil

	o.resourceName.Release()
	for i := range o.configurationKeys {
		o.configurationKeys[i].Release()
	}
	o.configurationKeys = nil
}

func (o *DescribeConfigsResource) getTaggedFieldsCount(version int16) int {
	numTaggedFields := len(o.unknownTaggedFields)

	return numTaggedFields
}

// validateNonIgnorableFields throws an error if any non-ignorable field not supported by current version is set to
// non-default value
func (o *DescribeConfigsResource) validateNonIgnorableFields(version int16) error {
	return nil
}

func DescribeConfigsResourceLowestSupportedVersion() int16 {
	return 0
}

func DescribeConfigsResourceHighestSupportedVersion() int16 {
	return 32767
}

func DescribeConfigsResourceLowestSupportedFlexVersion() int16 {
	return 4
}

func DescribeConfigsResourceHighestSupportedFlexVersion() int16 {
	return 32767
}

func DescribeConfigsResourceDefault() DescribeConfigsResource {
	var d DescribeConfigsResource
	d.SetDefault()

	return d
}
