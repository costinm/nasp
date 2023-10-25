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
package response

import (
	"bytes"
	"strconv"
	"strings"

	"emperror.dev/errors"
	typesbytes "github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/bytes"
	"github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/fields"
	"github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/varint"
)

var alterConfigsResourceResponseErrorCode = fields.Context{
	SpecName:                    "ErrorCode",
	LowestSupportedVersion:      0,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  2,
	HighestSupportedFlexVersion: 32767,
}
var alterConfigsResourceResponseErrorMessage = fields.Context{
	SpecName:                        "ErrorMessage",
	LowestSupportedVersion:          0,
	HighestSupportedVersion:         32767,
	LowestSupportedFlexVersion:      2,
	HighestSupportedFlexVersion:     32767,
	LowestSupportedNullableVersion:  0,
	HighestSupportedNullableVersion: 32767,
}
var alterConfigsResourceResponseResourceType = fields.Context{
	SpecName:                    "ResourceType",
	LowestSupportedVersion:      0,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  2,
	HighestSupportedFlexVersion: 32767,
}
var alterConfigsResourceResponseResourceName = fields.Context{
	SpecName:                    "ResourceName",
	LowestSupportedVersion:      0,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  2,
	HighestSupportedFlexVersion: 32767,
}

type AlterConfigsResourceResponse struct {
	unknownTaggedFields []fields.RawTaggedField
	errorMessage        fields.NullableString
	resourceName        fields.NullableString
	errorCode           int16
	resourceType        int8
	isNil               bool
}

func (o *AlterConfigsResourceResponse) ErrorCode() int16 {
	return o.errorCode
}

func (o *AlterConfigsResourceResponse) SetErrorCode(val int16) {
	o.isNil = false
	o.errorCode = val
}

func (o *AlterConfigsResourceResponse) ErrorMessage() fields.NullableString {
	return o.errorMessage
}

func (o *AlterConfigsResourceResponse) SetErrorMessage(val fields.NullableString) {
	o.isNil = false
	o.errorMessage = val
}

func (o *AlterConfigsResourceResponse) ResourceType() int8 {
	return o.resourceType
}

func (o *AlterConfigsResourceResponse) SetResourceType(val int8) {
	o.isNil = false
	o.resourceType = val
}

func (o *AlterConfigsResourceResponse) ResourceName() fields.NullableString {
	return o.resourceName
}

func (o *AlterConfigsResourceResponse) SetResourceName(val fields.NullableString) {
	o.isNil = false
	o.resourceName = val
}

func (o *AlterConfigsResourceResponse) UnknownTaggedFields() []fields.RawTaggedField {
	return o.unknownTaggedFields
}

func (o *AlterConfigsResourceResponse) SetUnknownTaggedFields(val []fields.RawTaggedField) {
	o.unknownTaggedFields = val
}

func (o *AlterConfigsResourceResponse) Read(buf *bytes.Reader, version int16) error {
	o.SetDefault()

	errorCodeField := fields.Int16{Context: alterConfigsResourceResponseErrorCode}
	if err := errorCodeField.Read(buf, version, &o.errorCode); err != nil {
		return errors.WrapIf(err, "couldn't set \"errorCode\" field")
	}

	errorMessageField := fields.String{Context: alterConfigsResourceResponseErrorMessage}
	if err := errorMessageField.Read(buf, version, &o.errorMessage); err != nil {
		return errors.WrapIf(err, "couldn't set \"errorMessage\" field")
	}

	resourceTypeField := fields.Int8{Context: alterConfigsResourceResponseResourceType}
	if err := resourceTypeField.Read(buf, version, &o.resourceType); err != nil {
		return errors.WrapIf(err, "couldn't set \"resourceType\" field")
	}

	resourceNameField := fields.String{Context: alterConfigsResourceResponseResourceName}
	if err := resourceNameField.Read(buf, version, &o.resourceName); err != nil {
		return errors.WrapIf(err, "couldn't set \"resourceName\" field")
	}

	// process tagged fields

	if version < AlterConfigsResourceResponseLowestSupportedFlexVersion() || version > AlterConfigsResourceResponseHighestSupportedFlexVersion() {
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

func (o *AlterConfigsResourceResponse) Write(buf *typesbytes.SliceWriter, version int16) error {
	if o.IsNil() {
		return nil
	}
	if err := o.validateNonIgnorableFields(version); err != nil {
		return err
	}

	errorCodeField := fields.Int16{Context: alterConfigsResourceResponseErrorCode}
	if err := errorCodeField.Write(buf, version, o.errorCode); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"errorCode\" field")
	}
	errorMessageField := fields.String{Context: alterConfigsResourceResponseErrorMessage}
	if err := errorMessageField.Write(buf, version, o.errorMessage); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"errorMessage\" field")
	}
	resourceTypeField := fields.Int8{Context: alterConfigsResourceResponseResourceType}
	if err := resourceTypeField.Write(buf, version, o.resourceType); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"resourceType\" field")
	}
	resourceNameField := fields.String{Context: alterConfigsResourceResponseResourceName}
	if err := resourceNameField.Write(buf, version, o.resourceName); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"resourceName\" field")
	}

	// serialize tagged fields
	numTaggedFields := o.getTaggedFieldsCount(version)
	if version < AlterConfigsResourceResponseLowestSupportedFlexVersion() || version > AlterConfigsResourceResponseHighestSupportedFlexVersion() {
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

func (o *AlterConfigsResourceResponse) String() string {
	s, err := o.MarshalJSON()
	if err != nil {
		return err.Error()
	}

	return string(s)
}

func (o *AlterConfigsResourceResponse) MarshalJSON() ([]byte, error) {
	if o == nil || o.IsNil() {
		return []byte("null"), nil
	}

	s := make([][]byte, 0, 5)
	if b, err := fields.MarshalPrimitiveTypeJSON(o.errorCode); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"errorCode\""), b}, []byte(": ")))
	}
	if b, err := fields.MarshalPrimitiveTypeJSON(o.errorMessage); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"errorMessage\""), b}, []byte(": ")))
	}
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

func (o *AlterConfigsResourceResponse) IsNil() bool {
	return o.isNil
}

func (o *AlterConfigsResourceResponse) Clear() {
	o.Release()
	o.isNil = true

	o.unknownTaggedFields = nil
}

func (o *AlterConfigsResourceResponse) SetDefault() {
	for i := range o.unknownTaggedFields {
		o.unknownTaggedFields[i].Release()
	}
	o.unknownTaggedFields = nil
	o.errorCode = 0
	o.errorMessage.SetValue("")
	o.resourceType = 0
	o.resourceName.SetValue("")

	o.isNil = false
}

func (o *AlterConfigsResourceResponse) Equal(that *AlterConfigsResourceResponse) bool {
	if !fields.RawTaggedFieldsEqual(o.unknownTaggedFields, that.unknownTaggedFields) {
		return false
	}

	if o.errorCode != that.errorCode {
		return false
	}
	if !o.errorMessage.Equal(&that.errorMessage) {
		return false
	}
	if o.resourceType != that.resourceType {
		return false
	}
	if !o.resourceName.Equal(&that.resourceName) {
		return false
	}

	return true
}

// SizeInBytes returns the size of this data structure in bytes when it's serialized
func (o *AlterConfigsResourceResponse) SizeInBytes(version int16) (int, error) {
	if o.IsNil() {
		return 0, nil
	}

	if err := o.validateNonIgnorableFields(version); err != nil {
		return 0, err
	}

	size := 0
	fieldSize := 0
	var err error

	errorCodeField := fields.Int16{Context: alterConfigsResourceResponseErrorCode}
	fieldSize, err = errorCodeField.SizeInBytes(version, o.errorCode)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"errorCode\" field")
	}
	size += fieldSize

	errorMessageField := fields.String{Context: alterConfigsResourceResponseErrorMessage}
	fieldSize, err = errorMessageField.SizeInBytes(version, o.errorMessage)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"errorMessage\" field")
	}
	size += fieldSize

	resourceTypeField := fields.Int8{Context: alterConfigsResourceResponseResourceType}
	fieldSize, err = resourceTypeField.SizeInBytes(version, o.resourceType)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"resourceType\" field")
	}
	size += fieldSize

	resourceNameField := fields.String{Context: alterConfigsResourceResponseResourceName}
	fieldSize, err = resourceNameField.SizeInBytes(version, o.resourceName)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"resourceName\" field")
	}
	size += fieldSize

	// tagged fields
	numTaggedFields := int64(o.getTaggedFieldsCount(version))
	if numTaggedFields > 0xffffffff {
		return 0, errors.New(strings.Join([]string{"invalid tagged fields count:", strconv.Itoa(int(numTaggedFields))}, " "))
	}
	if version < AlterConfigsResourceResponseLowestSupportedFlexVersion() || version > AlterConfigsResourceResponseHighestSupportedFlexVersion() {
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
func (o *AlterConfigsResourceResponse) Release() {
	if o.IsNil() {
		return
	}

	for i := range o.unknownTaggedFields {
		o.unknownTaggedFields[i].Release()
	}
	o.unknownTaggedFields = nil

	o.errorMessage.Release()
	o.resourceName.Release()
}

func (o *AlterConfigsResourceResponse) getTaggedFieldsCount(version int16) int {
	numTaggedFields := len(o.unknownTaggedFields)

	return numTaggedFields
}

// validateNonIgnorableFields throws an error if any non-ignorable field not supported by current version is set to
// non-default value
func (o *AlterConfigsResourceResponse) validateNonIgnorableFields(version int16) error {
	return nil
}

func AlterConfigsResourceResponseLowestSupportedVersion() int16 {
	return 0
}

func AlterConfigsResourceResponseHighestSupportedVersion() int16 {
	return 32767
}

func AlterConfigsResourceResponseLowestSupportedFlexVersion() int16 {
	return 2
}

func AlterConfigsResourceResponseHighestSupportedFlexVersion() int16 {
	return 32767
}

func AlterConfigsResourceResponseDefault() AlterConfigsResourceResponse {
	var d AlterConfigsResourceResponse
	d.SetDefault()

	return d
}