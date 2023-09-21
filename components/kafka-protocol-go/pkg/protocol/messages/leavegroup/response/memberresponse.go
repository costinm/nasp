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

var memberResponseMemberId = fields.Context{
	SpecName:                    "MemberId",
	LowestSupportedVersion:      3,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  4,
	HighestSupportedFlexVersion: 32767,
}
var memberResponseGroupInstanceId = fields.Context{
	SpecName:                        "GroupInstanceId",
	LowestSupportedVersion:          3,
	HighestSupportedVersion:         32767,
	LowestSupportedFlexVersion:      4,
	HighestSupportedFlexVersion:     32767,
	LowestSupportedNullableVersion:  3,
	HighestSupportedNullableVersion: 32767,
}
var memberResponseErrorCode = fields.Context{
	SpecName:                    "ErrorCode",
	LowestSupportedVersion:      3,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  4,
	HighestSupportedFlexVersion: 32767,
}

type MemberResponse struct {
	unknownTaggedFields []fields.RawTaggedField
	memberId            fields.NullableString
	groupInstanceId     fields.NullableString
	errorCode           int16
	isNil               bool
}

func (o *MemberResponse) MemberId() fields.NullableString {
	return o.memberId
}

func (o *MemberResponse) SetMemberId(val fields.NullableString) {
	o.isNil = false
	o.memberId = val
}

func (o *MemberResponse) GroupInstanceId() fields.NullableString {
	return o.groupInstanceId
}

func (o *MemberResponse) SetGroupInstanceId(val fields.NullableString) {
	o.isNil = false
	o.groupInstanceId = val
}

func (o *MemberResponse) ErrorCode() int16 {
	return o.errorCode
}

func (o *MemberResponse) SetErrorCode(val int16) {
	o.isNil = false
	o.errorCode = val
}

func (o *MemberResponse) UnknownTaggedFields() []fields.RawTaggedField {
	return o.unknownTaggedFields
}

func (o *MemberResponse) SetUnknownTaggedFields(val []fields.RawTaggedField) {
	o.unknownTaggedFields = val
}

func (o *MemberResponse) Read(buf *bytes.Reader, version int16) error {
	o.SetDefault()

	memberIdField := fields.String{Context: memberResponseMemberId}
	if err := memberIdField.Read(buf, version, &o.memberId); err != nil {
		return errors.WrapIf(err, "couldn't set \"memberId\" field")
	}

	groupInstanceIdField := fields.String{Context: memberResponseGroupInstanceId}
	if err := groupInstanceIdField.Read(buf, version, &o.groupInstanceId); err != nil {
		return errors.WrapIf(err, "couldn't set \"groupInstanceId\" field")
	}

	errorCodeField := fields.Int16{Context: memberResponseErrorCode}
	if err := errorCodeField.Read(buf, version, &o.errorCode); err != nil {
		return errors.WrapIf(err, "couldn't set \"errorCode\" field")
	}

	// process tagged fields

	if version < MemberResponseLowestSupportedFlexVersion() || version > MemberResponseHighestSupportedFlexVersion() {
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

func (o *MemberResponse) Write(buf *typesbytes.SliceWriter, version int16) error {
	if o.IsNil() {
		return nil
	}
	if err := o.validateNonIgnorableFields(version); err != nil {
		return err
	}

	memberIdField := fields.String{Context: memberResponseMemberId}
	if err := memberIdField.Write(buf, version, o.memberId); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"memberId\" field")
	}
	groupInstanceIdField := fields.String{Context: memberResponseGroupInstanceId}
	if err := groupInstanceIdField.Write(buf, version, o.groupInstanceId); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"groupInstanceId\" field")
	}
	errorCodeField := fields.Int16{Context: memberResponseErrorCode}
	if err := errorCodeField.Write(buf, version, o.errorCode); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"errorCode\" field")
	}

	// serialize tagged fields
	numTaggedFields := o.getTaggedFieldsCount(version)
	if version < MemberResponseLowestSupportedFlexVersion() || version > MemberResponseHighestSupportedFlexVersion() {
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

func (o *MemberResponse) String() string {
	s, err := o.MarshalJSON()
	if err != nil {
		return err.Error()
	}

	return string(s)
}

func (o *MemberResponse) MarshalJSON() ([]byte, error) {
	if o == nil || o.IsNil() {
		return []byte("null"), nil
	}

	s := make([][]byte, 0, 4)
	if b, err := fields.MarshalPrimitiveTypeJSON(o.memberId); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"memberId\""), b}, []byte(": ")))
	}
	if b, err := fields.MarshalPrimitiveTypeJSON(o.groupInstanceId); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"groupInstanceId\""), b}, []byte(": ")))
	}
	if b, err := fields.MarshalPrimitiveTypeJSON(o.errorCode); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"errorCode\""), b}, []byte(": ")))
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

func (o *MemberResponse) IsNil() bool {
	return o.isNil
}

func (o *MemberResponse) Clear() {
	o.Release()
	o.isNil = true

	o.unknownTaggedFields = nil
}

func (o *MemberResponse) SetDefault() {
	for i := range o.unknownTaggedFields {
		o.unknownTaggedFields[i].Release()
	}
	o.unknownTaggedFields = nil
	o.memberId.SetValue("")
	o.groupInstanceId.SetValue("")
	o.errorCode = 0

	o.isNil = false
}

func (o *MemberResponse) Equal(that *MemberResponse) bool {
	if !fields.RawTaggedFieldsEqual(o.unknownTaggedFields, that.unknownTaggedFields) {
		return false
	}

	if !o.memberId.Equal(&that.memberId) {
		return false
	}
	if !o.groupInstanceId.Equal(&that.groupInstanceId) {
		return false
	}
	if o.errorCode != that.errorCode {
		return false
	}

	return true
}

// SizeInBytes returns the size of this data structure in bytes when it's serialized
func (o *MemberResponse) SizeInBytes(version int16) (int, error) {
	if o.IsNil() {
		return 0, nil
	}

	if err := o.validateNonIgnorableFields(version); err != nil {
		return 0, err
	}

	size := 0
	fieldSize := 0
	var err error

	memberIdField := fields.String{Context: memberResponseMemberId}
	fieldSize, err = memberIdField.SizeInBytes(version, o.memberId)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"memberId\" field")
	}
	size += fieldSize

	groupInstanceIdField := fields.String{Context: memberResponseGroupInstanceId}
	fieldSize, err = groupInstanceIdField.SizeInBytes(version, o.groupInstanceId)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"groupInstanceId\" field")
	}
	size += fieldSize

	errorCodeField := fields.Int16{Context: memberResponseErrorCode}
	fieldSize, err = errorCodeField.SizeInBytes(version, o.errorCode)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"errorCode\" field")
	}
	size += fieldSize

	// tagged fields
	numTaggedFields := int64(o.getTaggedFieldsCount(version))
	if numTaggedFields > 0xffffffff {
		return 0, errors.New(strings.Join([]string{"invalid tagged fields count:", strconv.Itoa(int(numTaggedFields))}, " "))
	}
	if version < MemberResponseLowestSupportedFlexVersion() || version > MemberResponseHighestSupportedFlexVersion() {
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
func (o *MemberResponse) Release() {
	if o.IsNil() {
		return
	}

	for i := range o.unknownTaggedFields {
		o.unknownTaggedFields[i].Release()
	}
	o.unknownTaggedFields = nil

	o.memberId.Release()
	o.groupInstanceId.Release()
}

func (o *MemberResponse) getTaggedFieldsCount(version int16) int {
	numTaggedFields := len(o.unknownTaggedFields)

	return numTaggedFields
}

// validateNonIgnorableFields throws an error if any non-ignorable field not supported by current version is set to
// non-default value
func (o *MemberResponse) validateNonIgnorableFields(version int16) error {
	if !memberResponseMemberId.IsSupportedVersion(version) {
		if o.memberId.Bytes() != nil {
			return errors.New(strings.Join([]string{"attempted to write non-default \"memberId\" at version", strconv.Itoa(int(version))}, " "))
		}
	}
	if !memberResponseGroupInstanceId.IsSupportedVersion(version) {
		if o.groupInstanceId.Bytes() != nil {
			return errors.New(strings.Join([]string{"attempted to write non-default \"groupInstanceId\" at version", strconv.Itoa(int(version))}, " "))
		}
	}
	if !memberResponseErrorCode.IsSupportedVersion(version) {
		if o.errorCode != 0 {
			return errors.New(strings.Join([]string{"attempted to write non-default \"errorCode\" at version", strconv.Itoa(int(version))}, " "))
		}
	}
	return nil
}

func MemberResponseLowestSupportedVersion() int16 {
	return 3
}

func MemberResponseHighestSupportedVersion() int16 {
	return 32767
}

func MemberResponseLowestSupportedFlexVersion() int16 {
	return 4
}

func MemberResponseHighestSupportedFlexVersion() int16 {
	return 32767
}

func MemberResponseDefault() MemberResponse {
	var d MemberResponse
	d.SetDefault()

	return d
}
