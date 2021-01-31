//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

type TransportSize int8

type ITransportSize interface {
    Supported_S7_300() bool
    Supported_LOGO() bool
    Code() uint8
    SizeInBytes() uint8
    Supported_S7_400() bool
    Supported_S7_1200() bool
    ShortName() uint8
    Supported_S7_1500() bool
    DataTransportSize() DataTransportSize
    BaseType() TransportSize
    DataProtocolId() string
    Serialize(io utils.WriteBuffer) error
}

const(
    TransportSize_BOOL TransportSize = 0x01
    TransportSize_BYTE TransportSize = 0x02
    TransportSize_WORD TransportSize = 0x03
    TransportSize_DWORD TransportSize = 0x04
    TransportSize_LWORD TransportSize = 0x05
    TransportSize_INT TransportSize = 0x06
    TransportSize_UINT TransportSize = 0x07
    TransportSize_SINT TransportSize = 0x08
    TransportSize_USINT TransportSize = 0x09
    TransportSize_DINT TransportSize = 0x0A
    TransportSize_UDINT TransportSize = 0x0B
    TransportSize_LINT TransportSize = 0x0C
    TransportSize_ULINT TransportSize = 0x0D
    TransportSize_REAL TransportSize = 0x0E
    TransportSize_LREAL TransportSize = 0x0F
    TransportSize_CHAR TransportSize = 0x10
    TransportSize_WCHAR TransportSize = 0x11
    TransportSize_STRING TransportSize = 0x12
    TransportSize_WSTRING TransportSize = 0x13
    TransportSize_TIME TransportSize = 0x14
    TransportSize_LTIME TransportSize = 0x16
    TransportSize_DATE TransportSize = 0x17
    TransportSize_TIME_OF_DAY TransportSize = 0x18
    TransportSize_TOD TransportSize = 0x19
    TransportSize_DATE_AND_TIME TransportSize = 0x1A
    TransportSize_DT TransportSize = 0x1B
)


func (e TransportSize) Supported_S7_300() bool {
    switch e  {
        case 0x01: { /* '0x01' */
            return true
        }
        case 0x02: { /* '0x02' */
            return true
        }
        case 0x03: { /* '0x03' */
            return true
        }
        case 0x04: { /* '0x04' */
            return true
        }
        case 0x05: { /* '0x05' */
            return false
        }
        case 0x06: { /* '0x06' */
            return true
        }
        case 0x07: { /* '0x07' */
            return false
        }
        case 0x08: { /* '0x08' */
            return false
        }
        case 0x09: { /* '0x09' */
            return false
        }
        case 0x0A: { /* '0x0A' */
            return true
        }
        case 0x0B: { /* '0x0B' */
            return false
        }
        case 0x0C: { /* '0x0C' */
            return false
        }
        case 0x0D: { /* '0x0D' */
            return false
        }
        case 0x0E: { /* '0x0E' */
            return true
        }
        case 0x0F: { /* '0x0F' */
            return false
        }
        case 0x10: { /* '0x10' */
            return true
        }
        case 0x11: { /* '0x11' */
            return false
        }
        case 0x12: { /* '0x12' */
            return true
        }
        case 0x13: { /* '0x13' */
            return false
        }
        case 0x14: { /* '0x14' */
            return true
        }
        case 0x16: { /* '0x16' */
            return false
        }
        case 0x17: { /* '0x17' */
            return true
        }
        case 0x18: { /* '0x18' */
            return true
        }
        case 0x19: { /* '0x19' */
            return true
        }
        case 0x1A: { /* '0x1A' */
            return true
        }
        case 0x1B: { /* '0x1B' */
            return true
        }
        default: {
            return false
        }
    }
}

func (e TransportSize) Supported_LOGO() bool {
    switch e  {
        case 0x01: { /* '0x01' */
            return true
        }
        case 0x02: { /* '0x02' */
            return true
        }
        case 0x03: { /* '0x03' */
            return true
        }
        case 0x04: { /* '0x04' */
            return true
        }
        case 0x05: { /* '0x05' */
            return false
        }
        case 0x06: { /* '0x06' */
            return true
        }
        case 0x07: { /* '0x07' */
            return true
        }
        case 0x08: { /* '0x08' */
            return true
        }
        case 0x09: { /* '0x09' */
            return true
        }
        case 0x0A: { /* '0x0A' */
            return true
        }
        case 0x0B: { /* '0x0B' */
            return true
        }
        case 0x0C: { /* '0x0C' */
            return false
        }
        case 0x0D: { /* '0x0D' */
            return false
        }
        case 0x0E: { /* '0x0E' */
            return true
        }
        case 0x0F: { /* '0x0F' */
            return false
        }
        case 0x10: { /* '0x10' */
            return true
        }
        case 0x11: { /* '0x11' */
            return true
        }
        case 0x12: { /* '0x12' */
            return true
        }
        case 0x13: { /* '0x13' */
            return true
        }
        case 0x14: { /* '0x14' */
            return true
        }
        case 0x16: { /* '0x16' */
            return false
        }
        case 0x17: { /* '0x17' */
            return true
        }
        case 0x18: { /* '0x18' */
            return true
        }
        case 0x19: { /* '0x19' */
            return true
        }
        case 0x1A: { /* '0x1A' */
            return false
        }
        case 0x1B: { /* '0x1B' */
            return false
        }
        default: {
            return false
        }
    }
}

func (e TransportSize) Code() uint8 {
    switch e  {
        case 0x01: { /* '0x01' */
            return 0x01
        }
        case 0x02: { /* '0x02' */
            return 0x02
        }
        case 0x03: { /* '0x03' */
            return 0x04
        }
        case 0x04: { /* '0x04' */
            return 0x06
        }
        case 0x05: { /* '0x05' */
            return 0x00
        }
        case 0x06: { /* '0x06' */
            return 0x05
        }
        case 0x07: { /* '0x07' */
            return 0x05
        }
        case 0x08: { /* '0x08' */
            return 0x02
        }
        case 0x09: { /* '0x09' */
            return 0x02
        }
        case 0x0A: { /* '0x0A' */
            return 0x07
        }
        case 0x0B: { /* '0x0B' */
            return 0x07
        }
        case 0x0C: { /* '0x0C' */
            return 0x00
        }
        case 0x0D: { /* '0x0D' */
            return 0x00
        }
        case 0x0E: { /* '0x0E' */
            return 0x08
        }
        case 0x0F: { /* '0x0F' */
            return 0x30
        }
        case 0x10: { /* '0x10' */
            return 0x03
        }
        case 0x11: { /* '0x11' */
            return 0x13
        }
        case 0x12: { /* '0x12' */
            return 0x03
        }
        case 0x13: { /* '0x13' */
            return 0x00
        }
        case 0x14: { /* '0x14' */
            return 0x0B
        }
        case 0x16: { /* '0x16' */
            return 0x00
        }
        case 0x17: { /* '0x17' */
            return 0x09
        }
        case 0x18: { /* '0x18' */
            return 0x06
        }
        case 0x19: { /* '0x19' */
            return 0x06
        }
        case 0x1A: { /* '0x1A' */
            return 0x0F
        }
        case 0x1B: { /* '0x1B' */
            return 0x0F
        }
        default: {
            return 0
        }
    }
}

func (e TransportSize) SizeInBytes() uint8 {
    switch e  {
        case 0x01: { /* '0x01' */
            return 1
        }
        case 0x02: { /* '0x02' */
            return 1
        }
        case 0x03: { /* '0x03' */
            return 2
        }
        case 0x04: { /* '0x04' */
            return 4
        }
        case 0x05: { /* '0x05' */
            return 8
        }
        case 0x06: { /* '0x06' */
            return 2
        }
        case 0x07: { /* '0x07' */
            return 2
        }
        case 0x08: { /* '0x08' */
            return 1
        }
        case 0x09: { /* '0x09' */
            return 1
        }
        case 0x0A: { /* '0x0A' */
            return 4
        }
        case 0x0B: { /* '0x0B' */
            return 4
        }
        case 0x0C: { /* '0x0C' */
            return 8
        }
        case 0x0D: { /* '0x0D' */
            return 16
        }
        case 0x0E: { /* '0x0E' */
            return 4
        }
        case 0x0F: { /* '0x0F' */
            return 8
        }
        case 0x10: { /* '0x10' */
            return 1
        }
        case 0x11: { /* '0x11' */
            return 2
        }
        case 0x12: { /* '0x12' */
            return 1
        }
        case 0x13: { /* '0x13' */
            return 2
        }
        case 0x14: { /* '0x14' */
            return 4
        }
        case 0x16: { /* '0x16' */
            return 8
        }
        case 0x17: { /* '0x17' */
            return 2
        }
        case 0x18: { /* '0x18' */
            return 4
        }
        case 0x19: { /* '0x19' */
            return 4
        }
        case 0x1A: { /* '0x1A' */
            return 12
        }
        case 0x1B: { /* '0x1B' */
            return 12
        }
        default: {
            return 0
        }
    }
}

func (e TransportSize) Supported_S7_400() bool {
    switch e  {
        case 0x01: { /* '0x01' */
            return true
        }
        case 0x02: { /* '0x02' */
            return true
        }
        case 0x03: { /* '0x03' */
            return true
        }
        case 0x04: { /* '0x04' */
            return true
        }
        case 0x05: { /* '0x05' */
            return false
        }
        case 0x06: { /* '0x06' */
            return true
        }
        case 0x07: { /* '0x07' */
            return false
        }
        case 0x08: { /* '0x08' */
            return false
        }
        case 0x09: { /* '0x09' */
            return false
        }
        case 0x0A: { /* '0x0A' */
            return true
        }
        case 0x0B: { /* '0x0B' */
            return false
        }
        case 0x0C: { /* '0x0C' */
            return false
        }
        case 0x0D: { /* '0x0D' */
            return false
        }
        case 0x0E: { /* '0x0E' */
            return true
        }
        case 0x0F: { /* '0x0F' */
            return false
        }
        case 0x10: { /* '0x10' */
            return true
        }
        case 0x11: { /* '0x11' */
            return false
        }
        case 0x12: { /* '0x12' */
            return true
        }
        case 0x13: { /* '0x13' */
            return false
        }
        case 0x14: { /* '0x14' */
            return true
        }
        case 0x16: { /* '0x16' */
            return false
        }
        case 0x17: { /* '0x17' */
            return true
        }
        case 0x18: { /* '0x18' */
            return true
        }
        case 0x19: { /* '0x19' */
            return true
        }
        case 0x1A: { /* '0x1A' */
            return true
        }
        case 0x1B: { /* '0x1B' */
            return true
        }
        default: {
            return false
        }
    }
}

func (e TransportSize) Supported_S7_1200() bool {
    switch e  {
        case 0x01: { /* '0x01' */
            return true
        }
        case 0x02: { /* '0x02' */
            return true
        }
        case 0x03: { /* '0x03' */
            return true
        }
        case 0x04: { /* '0x04' */
            return true
        }
        case 0x05: { /* '0x05' */
            return false
        }
        case 0x06: { /* '0x06' */
            return true
        }
        case 0x07: { /* '0x07' */
            return true
        }
        case 0x08: { /* '0x08' */
            return true
        }
        case 0x09: { /* '0x09' */
            return true
        }
        case 0x0A: { /* '0x0A' */
            return true
        }
        case 0x0B: { /* '0x0B' */
            return true
        }
        case 0x0C: { /* '0x0C' */
            return false
        }
        case 0x0D: { /* '0x0D' */
            return false
        }
        case 0x0E: { /* '0x0E' */
            return true
        }
        case 0x0F: { /* '0x0F' */
            return true
        }
        case 0x10: { /* '0x10' */
            return true
        }
        case 0x11: { /* '0x11' */
            return true
        }
        case 0x12: { /* '0x12' */
            return true
        }
        case 0x13: { /* '0x13' */
            return true
        }
        case 0x14: { /* '0x14' */
            return true
        }
        case 0x16: { /* '0x16' */
            return false
        }
        case 0x17: { /* '0x17' */
            return true
        }
        case 0x18: { /* '0x18' */
            return true
        }
        case 0x19: { /* '0x19' */
            return true
        }
        case 0x1A: { /* '0x1A' */
            return false
        }
        case 0x1B: { /* '0x1B' */
            return false
        }
        default: {
            return false
        }
    }
}

func (e TransportSize) ShortName() uint8 {
    switch e  {
        case 0x01: { /* '0x01' */
            return 'X'
        }
        case 0x02: { /* '0x02' */
            return 'B'
        }
        case 0x03: { /* '0x03' */
            return 'W'
        }
        case 0x04: { /* '0x04' */
            return 'D'
        }
        case 0x05: { /* '0x05' */
            return 'X'
        }
        case 0x06: { /* '0x06' */
            return 'W'
        }
        case 0x07: { /* '0x07' */
            return 'W'
        }
        case 0x08: { /* '0x08' */
            return 'B'
        }
        case 0x09: { /* '0x09' */
            return 'B'
        }
        case 0x0A: { /* '0x0A' */
            return 'D'
        }
        case 0x0B: { /* '0x0B' */
            return 'D'
        }
        case 0x0C: { /* '0x0C' */
            return 'X'
        }
        case 0x0D: { /* '0x0D' */
            return 'X'
        }
        case 0x0E: { /* '0x0E' */
            return 'D'
        }
        case 0x0F: { /* '0x0F' */
            return 'X'
        }
        case 0x10: { /* '0x10' */
            return 'B'
        }
        case 0x11: { /* '0x11' */
            return 'X'
        }
        case 0x12: { /* '0x12' */
            return 'X'
        }
        case 0x13: { /* '0x13' */
            return 'X'
        }
        case 0x14: { /* '0x14' */
            return 'X'
        }
        case 0x16: { /* '0x16' */
            return 'X'
        }
        case 0x17: { /* '0x17' */
            return 'X'
        }
        case 0x18: { /* '0x18' */
            return 'X'
        }
        case 0x19: { /* '0x19' */
            return 'X'
        }
        case 0x1A: { /* '0x1A' */
            return 'X'
        }
        case 0x1B: { /* '0x1B' */
            return 'X'
        }
        default: {
            return 0
        }
    }
}

func (e TransportSize) Supported_S7_1500() bool {
    switch e  {
        case 0x01: { /* '0x01' */
            return true
        }
        case 0x02: { /* '0x02' */
            return true
        }
        case 0x03: { /* '0x03' */
            return true
        }
        case 0x04: { /* '0x04' */
            return true
        }
        case 0x05: { /* '0x05' */
            return true
        }
        case 0x06: { /* '0x06' */
            return true
        }
        case 0x07: { /* '0x07' */
            return true
        }
        case 0x08: { /* '0x08' */
            return true
        }
        case 0x09: { /* '0x09' */
            return true
        }
        case 0x0A: { /* '0x0A' */
            return true
        }
        case 0x0B: { /* '0x0B' */
            return true
        }
        case 0x0C: { /* '0x0C' */
            return true
        }
        case 0x0D: { /* '0x0D' */
            return true
        }
        case 0x0E: { /* '0x0E' */
            return true
        }
        case 0x0F: { /* '0x0F' */
            return true
        }
        case 0x10: { /* '0x10' */
            return true
        }
        case 0x11: { /* '0x11' */
            return true
        }
        case 0x12: { /* '0x12' */
            return true
        }
        case 0x13: { /* '0x13' */
            return true
        }
        case 0x14: { /* '0x14' */
            return true
        }
        case 0x16: { /* '0x16' */
            return true
        }
        case 0x17: { /* '0x17' */
            return true
        }
        case 0x18: { /* '0x18' */
            return true
        }
        case 0x19: { /* '0x19' */
            return true
        }
        case 0x1A: { /* '0x1A' */
            return true
        }
        case 0x1B: { /* '0x1B' */
            return true
        }
        default: {
            return false
        }
    }
}

func (e TransportSize) DataTransportSize() DataTransportSize {
    switch e  {
        case 0x01: { /* '0x01' */
            return DataTransportSize_BIT
        }
        case 0x02: { /* '0x02' */
            return DataTransportSize_BYTE_WORD_DWORD
        }
        case 0x03: { /* '0x03' */
            return DataTransportSize_BYTE_WORD_DWORD
        }
        case 0x04: { /* '0x04' */
            return DataTransportSize_BYTE_WORD_DWORD
        }
        case 0x05: { /* '0x05' */
            return 0
        }
        case 0x06: { /* '0x06' */
            return DataTransportSize_INTEGER
        }
        case 0x07: { /* '0x07' */
            return DataTransportSize_INTEGER
        }
        case 0x08: { /* '0x08' */
            return DataTransportSize_BYTE_WORD_DWORD
        }
        case 0x09: { /* '0x09' */
            return DataTransportSize_BYTE_WORD_DWORD
        }
        case 0x0A: { /* '0x0A' */
            return DataTransportSize_INTEGER
        }
        case 0x0B: { /* '0x0B' */
            return DataTransportSize_INTEGER
        }
        case 0x0C: { /* '0x0C' */
            return 0
        }
        case 0x0D: { /* '0x0D' */
            return 0
        }
        case 0x0E: { /* '0x0E' */
            return DataTransportSize_BYTE_WORD_DWORD
        }
        case 0x0F: { /* '0x0F' */
            return 0
        }
        case 0x10: { /* '0x10' */
            return DataTransportSize_BYTE_WORD_DWORD
        }
        case 0x11: { /* '0x11' */
            return 0
        }
        case 0x12: { /* '0x12' */
            return DataTransportSize_BYTE_WORD_DWORD
        }
        case 0x13: { /* '0x13' */
            return 0
        }
        case 0x14: { /* '0x14' */
            return 0
        }
        case 0x16: { /* '0x16' */
            return 0
        }
        case 0x17: { /* '0x17' */
            return DataTransportSize_BYTE_WORD_DWORD
        }
        case 0x18: { /* '0x18' */
            return DataTransportSize_BYTE_WORD_DWORD
        }
        case 0x19: { /* '0x19' */
            return DataTransportSize_BYTE_WORD_DWORD
        }
        case 0x1A: { /* '0x1A' */
            return 0
        }
        case 0x1B: { /* '0x1B' */
            return 0
        }
        default: {
            return 0
        }
    }
}

func (e TransportSize) BaseType() TransportSize {
    switch e  {
        case 0x01: { /* '0x01' */
            return 0
        }
        case 0x02: { /* '0x02' */
            return 0
        }
        case 0x03: { /* '0x03' */
            return 0
        }
        case 0x04: { /* '0x04' */
            return TransportSize_WORD
        }
        case 0x05: { /* '0x05' */
            return 0
        }
        case 0x06: { /* '0x06' */
            return 0
        }
        case 0x07: { /* '0x07' */
            return TransportSize_INT
        }
        case 0x08: { /* '0x08' */
            return TransportSize_INT
        }
        case 0x09: { /* '0x09' */
            return TransportSize_INT
        }
        case 0x0A: { /* '0x0A' */
            return TransportSize_INT
        }
        case 0x0B: { /* '0x0B' */
            return TransportSize_INT
        }
        case 0x0C: { /* '0x0C' */
            return TransportSize_INT
        }
        case 0x0D: { /* '0x0D' */
            return TransportSize_INT
        }
        case 0x0E: { /* '0x0E' */
            return 0
        }
        case 0x0F: { /* '0x0F' */
            return TransportSize_REAL
        }
        case 0x10: { /* '0x10' */
            return 0
        }
        case 0x11: { /* '0x11' */
            return 0
        }
        case 0x12: { /* '0x12' */
            return 0
        }
        case 0x13: { /* '0x13' */
            return 0
        }
        case 0x14: { /* '0x14' */
            return 0
        }
        case 0x16: { /* '0x16' */
            return TransportSize_TIME
        }
        case 0x17: { /* '0x17' */
            return 0
        }
        case 0x18: { /* '0x18' */
            return 0
        }
        case 0x19: { /* '0x19' */
            return 0
        }
        case 0x1A: { /* '0x1A' */
            return 0
        }
        case 0x1B: { /* '0x1B' */
            return 0
        }
        default: {
            return 0
        }
    }
}

func (e TransportSize) DataProtocolId() string {
    switch e  {
        case 0x01: { /* '0x01' */
            return "IEC61131_BOOL"
        }
        case 0x02: { /* '0x02' */
            return "IEC61131_BYTE"
        }
        case 0x03: { /* '0x03' */
            return "IEC61131_WORD"
        }
        case 0x04: { /* '0x04' */
            return "IEC61131_DWORD"
        }
        case 0x05: { /* '0x05' */
            return "IEC61131_LWORD"
        }
        case 0x06: { /* '0x06' */
            return "IEC61131_INT"
        }
        case 0x07: { /* '0x07' */
            return "IEC61131_UINT"
        }
        case 0x08: { /* '0x08' */
            return "IEC61131_SINT"
        }
        case 0x09: { /* '0x09' */
            return "IEC61131_USINT"
        }
        case 0x0A: { /* '0x0A' */
            return "IEC61131_DINT"
        }
        case 0x0B: { /* '0x0B' */
            return "IEC61131_UDINT"
        }
        case 0x0C: { /* '0x0C' */
            return "IEC61131_LINT"
        }
        case 0x0D: { /* '0x0D' */
            return "IEC61131_ULINT"
        }
        case 0x0E: { /* '0x0E' */
            return "IEC61131_REAL"
        }
        case 0x0F: { /* '0x0F' */
            return "IEC61131_LREAL"
        }
        case 0x10: { /* '0x10' */
            return "IEC61131_CHAR"
        }
        case 0x11: { /* '0x11' */
            return "IEC61131_WCHAR"
        }
        case 0x12: { /* '0x12' */
            return "IEC61131_STRING"
        }
        case 0x13: { /* '0x13' */
            return "IEC61131_WSTRING"
        }
        case 0x14: { /* '0x14' */
            return "IEC61131_TIME"
        }
        case 0x16: { /* '0x16' */
            return "IEC61131_LTIME"
        }
        case 0x17: { /* '0x17' */
            return "IEC61131_DATE"
        }
        case 0x18: { /* '0x18' */
            return "IEC61131_TIME_OF_DAY"
        }
        case 0x19: { /* '0x19' */
            return "IEC61131_TIME_OF_DAY"
        }
        case 0x1A: { /* '0x1A' */
            return "IEC61131_DATE_AND_TIME"
        }
        case 0x1B: { /* '0x1B' */
            return "IEC61131_DATE_AND_TIME"
        }
        default: {
            return ""
        }
    }
}
func TransportSizeByValue(value int8) TransportSize {
    switch value {
        case 0x01:
            return TransportSize_BOOL
        case 0x02:
            return TransportSize_BYTE
        case 0x03:
            return TransportSize_WORD
        case 0x04:
            return TransportSize_DWORD
        case 0x05:
            return TransportSize_LWORD
        case 0x06:
            return TransportSize_INT
        case 0x07:
            return TransportSize_UINT
        case 0x08:
            return TransportSize_SINT
        case 0x09:
            return TransportSize_USINT
        case 0x0A:
            return TransportSize_DINT
        case 0x0B:
            return TransportSize_UDINT
        case 0x0C:
            return TransportSize_LINT
        case 0x0D:
            return TransportSize_ULINT
        case 0x0E:
            return TransportSize_REAL
        case 0x0F:
            return TransportSize_LREAL
        case 0x10:
            return TransportSize_CHAR
        case 0x11:
            return TransportSize_WCHAR
        case 0x12:
            return TransportSize_STRING
        case 0x13:
            return TransportSize_WSTRING
        case 0x14:
            return TransportSize_TIME
        case 0x16:
            return TransportSize_LTIME
        case 0x17:
            return TransportSize_DATE
        case 0x18:
            return TransportSize_TIME_OF_DAY
        case 0x19:
            return TransportSize_TOD
        case 0x1A:
            return TransportSize_DATE_AND_TIME
        case 0x1B:
            return TransportSize_DT
    }
    return 0
}

func TransportSizeByName(value string) TransportSize {
    switch value {
    case "BOOL":
        return TransportSize_BOOL
    case "BYTE":
        return TransportSize_BYTE
    case "WORD":
        return TransportSize_WORD
    case "DWORD":
        return TransportSize_DWORD
    case "LWORD":
        return TransportSize_LWORD
    case "INT":
        return TransportSize_INT
    case "UINT":
        return TransportSize_UINT
    case "SINT":
        return TransportSize_SINT
    case "USINT":
        return TransportSize_USINT
    case "DINT":
        return TransportSize_DINT
    case "UDINT":
        return TransportSize_UDINT
    case "LINT":
        return TransportSize_LINT
    case "ULINT":
        return TransportSize_ULINT
    case "REAL":
        return TransportSize_REAL
    case "LREAL":
        return TransportSize_LREAL
    case "CHAR":
        return TransportSize_CHAR
    case "WCHAR":
        return TransportSize_WCHAR
    case "STRING":
        return TransportSize_STRING
    case "WSTRING":
        return TransportSize_WSTRING
    case "TIME":
        return TransportSize_TIME
    case "LTIME":
        return TransportSize_LTIME
    case "DATE":
        return TransportSize_DATE
    case "TIME_OF_DAY":
        return TransportSize_TIME_OF_DAY
    case "TOD":
        return TransportSize_TOD
    case "DATE_AND_TIME":
        return TransportSize_DATE_AND_TIME
    case "DT":
        return TransportSize_DT
    }
    return 0
}

func CastTransportSize(structType interface{}) TransportSize {
    castFunc := func(typ interface{}) TransportSize {
        if sTransportSize, ok := typ.(TransportSize); ok {
            return sTransportSize
        }
        return 0
    }
    return castFunc(structType)
}

func (m TransportSize) LengthInBits() uint16 {
    return 8
}

func (m TransportSize) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func TransportSizeParse(io *utils.ReadBuffer) (TransportSize, error) {
    val, err := io.ReadInt8(8)
    if err != nil {
        return 0, nil
    }
    return TransportSizeByValue(val), nil
}

func (e TransportSize) Serialize(io utils.WriteBuffer) error {
    err := io.WriteInt8(8, int8(e))
    return err
}

func (e TransportSize) String() string {
    switch e {
    case TransportSize_BOOL:
        return "BOOL"
    case TransportSize_BYTE:
        return "BYTE"
    case TransportSize_WORD:
        return "WORD"
    case TransportSize_DWORD:
        return "DWORD"
    case TransportSize_LWORD:
        return "LWORD"
    case TransportSize_INT:
        return "INT"
    case TransportSize_UINT:
        return "UINT"
    case TransportSize_SINT:
        return "SINT"
    case TransportSize_USINT:
        return "USINT"
    case TransportSize_DINT:
        return "DINT"
    case TransportSize_UDINT:
        return "UDINT"
    case TransportSize_LINT:
        return "LINT"
    case TransportSize_ULINT:
        return "ULINT"
    case TransportSize_REAL:
        return "REAL"
    case TransportSize_LREAL:
        return "LREAL"
    case TransportSize_CHAR:
        return "CHAR"
    case TransportSize_WCHAR:
        return "WCHAR"
    case TransportSize_STRING:
        return "STRING"
    case TransportSize_WSTRING:
        return "WSTRING"
    case TransportSize_TIME:
        return "TIME"
    case TransportSize_LTIME:
        return "LTIME"
    case TransportSize_DATE:
        return "DATE"
    case TransportSize_TIME_OF_DAY:
        return "TIME_OF_DAY"
    case TransportSize_TOD:
        return "TOD"
    case TransportSize_DATE_AND_TIME:
        return "DATE_AND_TIME"
    case TransportSize_DT:
        return "DT"
    }
    return ""
}
