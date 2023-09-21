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
package topicdata

import (
	"bytes"
	"strconv"
	"strings"

	"emperror.dev/errors"
	typesbytes "github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/bytes"
	"github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/fields"
	"github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/varint"
)

var partitionDataPartitionIndex = fields.Context{
	SpecName:                    "PartitionIndex",
	LowestSupportedVersion:      0,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  0,
	HighestSupportedFlexVersion: 32767,
}
var partitionDataErrorCode = fields.Context{
	SpecName:                    "ErrorCode",
	LowestSupportedVersion:      0,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  0,
	HighestSupportedFlexVersion: 32767,
}
var partitionDataLeaderId = fields.Context{
	SpecName:                    "LeaderId",
	LowestSupportedVersion:      0,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  0,
	HighestSupportedFlexVersion: 32767,
}
var partitionDataLeaderEpoch = fields.Context{
	SpecName:                    "LeaderEpoch",
	LowestSupportedVersion:      0,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  0,
	HighestSupportedFlexVersion: 32767,
}
var partitionDataIsr = fields.Context{
	SpecName:                    "Isr",
	LowestSupportedVersion:      0,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  0,
	HighestSupportedFlexVersion: 32767,
}
var partitionDataLeaderRecoveryState = fields.Context{
	SpecName:                    "LeaderRecoveryState",
	LowestSupportedVersion:      1,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  0,
	HighestSupportedFlexVersion: 32767,
}
var partitionDataPartitionEpoch = fields.Context{
	SpecName:                    "PartitionEpoch",
	LowestSupportedVersion:      0,
	HighestSupportedVersion:     32767,
	LowestSupportedFlexVersion:  0,
	HighestSupportedFlexVersion: 32767,
}

type PartitionData struct {
	isr                 []int32
	unknownTaggedFields []fields.RawTaggedField
	partitionIndex      int32
	leaderId            int32
	leaderEpoch         int32
	partitionEpoch      int32
	errorCode           int16
	leaderRecoveryState int8
	isNil               bool
}

func (o *PartitionData) PartitionIndex() int32 {
	return o.partitionIndex
}

func (o *PartitionData) SetPartitionIndex(val int32) {
	o.isNil = false
	o.partitionIndex = val
}

func (o *PartitionData) ErrorCode() int16 {
	return o.errorCode
}

func (o *PartitionData) SetErrorCode(val int16) {
	o.isNil = false
	o.errorCode = val
}

func (o *PartitionData) LeaderId() int32 {
	return o.leaderId
}

func (o *PartitionData) SetLeaderId(val int32) {
	o.isNil = false
	o.leaderId = val
}

func (o *PartitionData) LeaderEpoch() int32 {
	return o.leaderEpoch
}

func (o *PartitionData) SetLeaderEpoch(val int32) {
	o.isNil = false
	o.leaderEpoch = val
}

func (o *PartitionData) Isr() []int32 {
	return o.isr
}

func (o *PartitionData) SetIsr(val []int32) {
	o.isNil = false
	o.isr = val
}

func (o *PartitionData) LeaderRecoveryState() int8 {
	return o.leaderRecoveryState
}

func (o *PartitionData) SetLeaderRecoveryState(val int8) {
	o.isNil = false
	o.leaderRecoveryState = val
}

func (o *PartitionData) PartitionEpoch() int32 {
	return o.partitionEpoch
}

func (o *PartitionData) SetPartitionEpoch(val int32) {
	o.isNil = false
	o.partitionEpoch = val
}

func (o *PartitionData) UnknownTaggedFields() []fields.RawTaggedField {
	return o.unknownTaggedFields
}

func (o *PartitionData) SetUnknownTaggedFields(val []fields.RawTaggedField) {
	o.unknownTaggedFields = val
}

func (o *PartitionData) Read(buf *bytes.Reader, version int16) error {
	o.SetDefault()

	partitionIndexField := fields.Int32{Context: partitionDataPartitionIndex}
	if err := partitionIndexField.Read(buf, version, &o.partitionIndex); err != nil {
		return errors.WrapIf(err, "couldn't set \"partitionIndex\" field")
	}

	errorCodeField := fields.Int16{Context: partitionDataErrorCode}
	if err := errorCodeField.Read(buf, version, &o.errorCode); err != nil {
		return errors.WrapIf(err, "couldn't set \"errorCode\" field")
	}

	leaderIdField := fields.Int32{Context: partitionDataLeaderId}
	if err := leaderIdField.Read(buf, version, &o.leaderId); err != nil {
		return errors.WrapIf(err, "couldn't set \"leaderId\" field")
	}

	leaderEpochField := fields.Int32{Context: partitionDataLeaderEpoch}
	if err := leaderEpochField.Read(buf, version, &o.leaderEpoch); err != nil {
		return errors.WrapIf(err, "couldn't set \"leaderEpoch\" field")
	}

	isrField := fields.Array[int32, *fields.Int32]{
		Context:          partitionDataIsr,
		ElementProcessor: &fields.Int32{Context: partitionDataIsr}}

	isr, err := isrField.Read(buf, version)
	if err != nil {
		return errors.WrapIf(err, "couldn't set \"isr\" field")
	}
	o.isr = isr

	leaderRecoveryStateField := fields.Int8{Context: partitionDataLeaderRecoveryState}
	if err := leaderRecoveryStateField.Read(buf, version, &o.leaderRecoveryState); err != nil {
		return errors.WrapIf(err, "couldn't set \"leaderRecoveryState\" field")
	}

	partitionEpochField := fields.Int32{Context: partitionDataPartitionEpoch}
	if err := partitionEpochField.Read(buf, version, &o.partitionEpoch); err != nil {
		return errors.WrapIf(err, "couldn't set \"partitionEpoch\" field")
	}

	// process tagged fields

	if version < PartitionDataLowestSupportedFlexVersion() || version > PartitionDataHighestSupportedFlexVersion() {
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

func (o *PartitionData) Write(buf *typesbytes.SliceWriter, version int16) error {
	if o.IsNil() {
		return nil
	}
	if err := o.validateNonIgnorableFields(version); err != nil {
		return err
	}

	partitionIndexField := fields.Int32{Context: partitionDataPartitionIndex}
	if err := partitionIndexField.Write(buf, version, o.partitionIndex); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"partitionIndex\" field")
	}
	errorCodeField := fields.Int16{Context: partitionDataErrorCode}
	if err := errorCodeField.Write(buf, version, o.errorCode); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"errorCode\" field")
	}
	leaderIdField := fields.Int32{Context: partitionDataLeaderId}
	if err := leaderIdField.Write(buf, version, o.leaderId); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"leaderId\" field")
	}
	leaderEpochField := fields.Int32{Context: partitionDataLeaderEpoch}
	if err := leaderEpochField.Write(buf, version, o.leaderEpoch); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"leaderEpoch\" field")
	}
	isrField := fields.Array[int32, *fields.Int32]{
		Context:          partitionDataIsr,
		ElementProcessor: &fields.Int32{Context: partitionDataIsr}}
	if err := isrField.Write(buf, version, o.isr); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"isr\" field")
	}

	leaderRecoveryStateField := fields.Int8{Context: partitionDataLeaderRecoveryState}
	if err := leaderRecoveryStateField.Write(buf, version, o.leaderRecoveryState); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"leaderRecoveryState\" field")
	}
	partitionEpochField := fields.Int32{Context: partitionDataPartitionEpoch}
	if err := partitionEpochField.Write(buf, version, o.partitionEpoch); err != nil {
		return errors.WrapIf(err, "couldn't serialize \"partitionEpoch\" field")
	}

	// serialize tagged fields
	numTaggedFields := o.getTaggedFieldsCount(version)
	if version < PartitionDataLowestSupportedFlexVersion() || version > PartitionDataHighestSupportedFlexVersion() {
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

func (o *PartitionData) String() string {
	s, err := o.MarshalJSON()
	if err != nil {
		return err.Error()
	}

	return string(s)
}

func (o *PartitionData) MarshalJSON() ([]byte, error) {
	if o == nil || o.IsNil() {
		return []byte("null"), nil
	}

	s := make([][]byte, 0, 8)
	if b, err := fields.MarshalPrimitiveTypeJSON(o.partitionIndex); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"partitionIndex\""), b}, []byte(": ")))
	}
	if b, err := fields.MarshalPrimitiveTypeJSON(o.errorCode); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"errorCode\""), b}, []byte(": ")))
	}
	if b, err := fields.MarshalPrimitiveTypeJSON(o.leaderId); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"leaderId\""), b}, []byte(": ")))
	}
	if b, err := fields.MarshalPrimitiveTypeJSON(o.leaderEpoch); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"leaderEpoch\""), b}, []byte(": ")))
	}
	if b, err := fields.ArrayMarshalJSON("isr", o.isr); err != nil {
		return nil, err
	} else {
		s = append(s, b)
	}
	if b, err := fields.MarshalPrimitiveTypeJSON(o.leaderRecoveryState); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"leaderRecoveryState\""), b}, []byte(": ")))
	}
	if b, err := fields.MarshalPrimitiveTypeJSON(o.partitionEpoch); err != nil {
		return nil, err
	} else {
		s = append(s, bytes.Join([][]byte{[]byte("\"partitionEpoch\""), b}, []byte(": ")))
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

func (o *PartitionData) IsNil() bool {
	return o.isNil
}

func (o *PartitionData) Clear() {
	o.Release()
	o.isNil = true

	o.isr = nil
	o.unknownTaggedFields = nil
}

func (o *PartitionData) SetDefault() {
	for i := range o.unknownTaggedFields {
		o.unknownTaggedFields[i].Release()
	}
	o.unknownTaggedFields = nil
	o.partitionIndex = 0
	o.errorCode = 0
	o.leaderId = 0
	o.leaderEpoch = 0
	o.isr = nil
	o.leaderRecoveryState = 0
	o.partitionEpoch = 0

	o.isNil = false
}

func (o *PartitionData) Equal(that *PartitionData) bool {
	if !fields.RawTaggedFieldsEqual(o.unknownTaggedFields, that.unknownTaggedFields) {
		return false
	}

	if o.partitionIndex != that.partitionIndex {
		return false
	}
	if o.errorCode != that.errorCode {
		return false
	}
	if o.leaderId != that.leaderId {
		return false
	}
	if o.leaderEpoch != that.leaderEpoch {
		return false
	}
	if !fields.PrimitiveTypeSliceEqual(o.isr, that.isr) {
		return false
	}
	if o.leaderRecoveryState != that.leaderRecoveryState {
		return false
	}
	if o.partitionEpoch != that.partitionEpoch {
		return false
	}

	return true
}

// SizeInBytes returns the size of this data structure in bytes when it's serialized
func (o *PartitionData) SizeInBytes(version int16) (int, error) {
	if o.IsNil() {
		return 0, nil
	}

	if err := o.validateNonIgnorableFields(version); err != nil {
		return 0, err
	}

	size := 0
	fieldSize := 0
	var err error

	partitionIndexField := fields.Int32{Context: partitionDataPartitionIndex}
	fieldSize, err = partitionIndexField.SizeInBytes(version, o.partitionIndex)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"partitionIndex\" field")
	}
	size += fieldSize

	errorCodeField := fields.Int16{Context: partitionDataErrorCode}
	fieldSize, err = errorCodeField.SizeInBytes(version, o.errorCode)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"errorCode\" field")
	}
	size += fieldSize

	leaderIdField := fields.Int32{Context: partitionDataLeaderId}
	fieldSize, err = leaderIdField.SizeInBytes(version, o.leaderId)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"leaderId\" field")
	}
	size += fieldSize

	leaderEpochField := fields.Int32{Context: partitionDataLeaderEpoch}
	fieldSize, err = leaderEpochField.SizeInBytes(version, o.leaderEpoch)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"leaderEpoch\" field")
	}
	size += fieldSize

	isrField := fields.Array[int32, *fields.Int32]{
		Context:          partitionDataIsr,
		ElementProcessor: &fields.Int32{Context: partitionDataIsr}}
	fieldSize, err = isrField.SizeInBytes(version, o.isr)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"isr\" field")
	}
	size += fieldSize

	leaderRecoveryStateField := fields.Int8{Context: partitionDataLeaderRecoveryState}
	fieldSize, err = leaderRecoveryStateField.SizeInBytes(version, o.leaderRecoveryState)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"leaderRecoveryState\" field")
	}
	size += fieldSize

	partitionEpochField := fields.Int32{Context: partitionDataPartitionEpoch}
	fieldSize, err = partitionEpochField.SizeInBytes(version, o.partitionEpoch)
	if err != nil {
		return 0, errors.WrapIf(err, "couldn't compute size of \"partitionEpoch\" field")
	}
	size += fieldSize

	// tagged fields
	numTaggedFields := int64(o.getTaggedFieldsCount(version))
	if numTaggedFields > 0xffffffff {
		return 0, errors.New(strings.Join([]string{"invalid tagged fields count:", strconv.Itoa(int(numTaggedFields))}, " "))
	}
	if version < PartitionDataLowestSupportedFlexVersion() || version > PartitionDataHighestSupportedFlexVersion() {
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
func (o *PartitionData) Release() {
	if o.IsNil() {
		return
	}

	for i := range o.unknownTaggedFields {
		o.unknownTaggedFields[i].Release()
	}
	o.unknownTaggedFields = nil

	o.isr = nil
}

func (o *PartitionData) getTaggedFieldsCount(version int16) int {
	numTaggedFields := len(o.unknownTaggedFields)

	return numTaggedFields
}

// validateNonIgnorableFields throws an error if any non-ignorable field not supported by current version is set to
// non-default value
func (o *PartitionData) validateNonIgnorableFields(version int16) error {
	return nil
}

func PartitionDataLowestSupportedVersion() int16 {
	return 0
}

func PartitionDataHighestSupportedVersion() int16 {
	return 32767
}

func PartitionDataLowestSupportedFlexVersion() int16 {
	return 0
}

func PartitionDataHighestSupportedFlexVersion() int16 {
	return 32767
}

func PartitionDataDefault() PartitionData {
	var d PartitionData
	d.SetDefault()

	return d
}
