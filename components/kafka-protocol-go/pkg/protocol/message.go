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

package protocol

import (
	"bytes"

	typesbytes "github.com/cisco-open/libnasp/components/kafka-protocol-go/pkg/protocol/types/bytes"
)

type MessageData interface {
	Read(r *bytes.Reader, messageVersion int16) error
	Write(w *typesbytes.SliceWriter, messageVersion int16) error

	ApiKey() int16
	SizeInBytes(messageVersion int16) (int, error)
	Release()
	String() string
}