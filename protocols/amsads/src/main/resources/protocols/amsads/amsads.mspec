//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

////////////////////////////////////////////////////////////////
// AMS/TCP Packet
////////////////////////////////////////////////////////////////

[type 'AmsTCPPacket'
    // AMS/TCP Header	6 bytes	contains the tcpLength of the data packet.
    // These bytes must be set to 0.
    [reserved   uint       16       '0x0000'                            ]
    // This array contains the length of the data packet.
    // It consists of the AMS-Header and the enclosed ADS data. The unit is bytes.
    [implicit   uint       32       'length'  'userdata.lengthInBytes'  ]
    // The AMS packet to be sent.
    [simple AmsPacket    'userdata'                                     ]
]

////////////////////////////////////////////////////////////////
// AMS/Serial Packet
////////////////////////////////////////////////////////////////

// If an AMS serial frame has been received and the frame is OK (magic cookie OK, CRC OK, correct fragment number etc.),
// then the receiver has to send an acknowledge frame, to inform the transmitter that the frame has arrived.
//
// @see <a href="https://infosys.beckhoff.com/content/1033/tcadsamsserialspec/html/tcamssericalspec_amsframe.htm?id=8115637053270715044">TwinCAT AMS via RS232 Specification</a>
[type 'AmsSerialAcknowledgeFrame'
    // Id for detecting an AMS serial frame.
    [simple     uint        16  'magicCookie'        ]
    // Address of the sending participant. This value can always be set to 0 for an RS232 communication,
    // since it is a 1 to 1 connection and hence the participants are unique.
    [simple     int          8  'transmitterAddress' ]
    // Receiver’s address. This value can always be set to 0 for an RS232 communication, since it is a 1 to 1
    // connection and hence the participants are unique.
    [simple     int          8  'receiverAddress'    ]
    // Number of the frame sent. Once the number 255 has been sent, it starts again from 0. The receiver checks this
    // number with an internal counter.
    [simple     int          8  'fragmentNumber'     ]
    // The max. length of the AMS packet to be sent is 255. If larger AMS packets are to be sent then they have to be
    // fragmented (not published at the moment).
    [simple     int          8  'length'             ]
    [simple     uint        16  'crc'                ]
]

// An AMS packet can be transferred via RS232 with the help of an AMS serial frame.
// The actual AMS packet is in the user data field of the frame.
// The max. length of the AMS packet is limited to 255 bytes.
// Therefore the max. size of an AMS serial frame is 263 bytes.
// The fragment number is compared with an internal counter by the receiver.
// The frame number is simply accepted and not checked when receiving the first AMS frame or in case a timeout is
// exceeded. The CRC16 algorithm is used for calculating the checksum.
// @see <a href="https://infosys.beckhoff.com/content/1033/tcadsamsserialspec/html/tcamssericalspec_amsframe.htm?id=8115637053270715044">TwinCAT AMS via RS232 Specification</a>
[type 'AmsSerialFrame'
    // Id for detecting an AMS serial frame.
    [simple     uint        16  'magicCookie'        ]
    // Address of the sending participant. This value can always be set to 0 for an RS232 communication,
    // since it is a 1 to 1 connection and hence the participants are unique.
    [simple     int          8  'transmitterAddress' ]
    // Receiver’s address. This value can always be set to 0 for an RS232 communication, since it is a 1 to 1
    // connection and hence the participants are unique.
    [simple     int          8  'receiverAddress'    ]
    // Number of the frame sent. Once the number 255 has been sent, it starts again from 0. The receiver checks this
    // number with an internal counter.
    [simple     int          8  'fragmentNumber'     ]
    // The max. length of the AMS packet to be sent is 255. If larger AMS packets are to be sent then they have to be
    // fragmented (not published at the moment).
    [simple     int          8  'length'             ]
    // The AMS packet to be sent.
    [simple AmsPacket           'userdata'           ]
    [simple     uint        16  'crc'                ]
]

// In case the transmitter does not receive a valid acknowledgement after multiple transmission, then a reset frame is
// sent. In this way the receiver is informed that a new communication is running and the receiver then accepts the
// fragment number during the next AMS-Frame, without carrying out a check.
[type 'AmsSerialResetFrame'
    // Id for detecting an AMS serial frame.
    [simple     uint        16  'magicCookie'        ]
    // Address of the sending participant. This value can always be set to 0 for an RS232 communication,
    // since it is a 1 to 1 connection and hence the participants are unique.
    [simple     int          8  'transmitterAddress' ]
    // Receiver’s address. This value can always be set to 0 for an RS232 communication, since it is a 1 to 1
    // connection and hence the participants are unique.
    [simple     int          8  'receiverAddress'    ]
    // Number of the frame sent. Once the number 255 has been sent, it starts again from 0. The receiver checks this
    // number with an internal counter.
    [simple     int          8  'fragmentNumber'     ]
    // The max. length of the AMS packet to be sent is 255. If larger AMS packets are to be sent then they have to be
    // fragmented (not published at the moment).
    [simple     int          8  'length'     ]
    [simple     uint        16  'crc'                ]
]

////////////////////////////////////////////////////////////////
// AMS Common
////////////////////////////////////////////////////////////////

[type 'AmsPacket'
    // AMS Header	32 bytes	The AMS/TCP-Header contains the addresses of the transmitter and receiver. In addition the AMS error code , the ADS command Id and some other information.
    // This is the AmsNetId of the station, for which the packet is intended. Remarks see below.
    [simple     AmsNetId        'targetAmsNetId'                            ]
    // This is the AmsPort of the station, for which the packet is intended.
    [simple     uint        16  'targetAmsPort'                             ]
    // This contains the AmsNetId of the station, from which the packet was sent.
    [simple     AmsNetId        'sourceAmsNetId'                            ]
    // This contains the AmsPort of the station, from which the packet was sent.
    [simple     uint        16  'sourceAmsPort'                             ]
    // 2 bytes.
    [enum       CommandId       'commandId'                                 ]
    // 2 bytes.
    [simple     State           'state'                                     ]
    // 4 bytes	Size of the data range. The unit is byte.
    [implicit   uint        32  'length'   'data.lengthInBytes'             ]
    // 4 bytes	AMS error number. See ADS Return Codes.
    [simple     uint        32  'errorCode'                                 ]
    // free usable field of 4 bytes
    // 4 bytes	Free usable 32 bit array. Usually this array serves to send an Id. This Id makes is possible to assign a received response to a request, which was sent before.
    [simple      uint        32  'invokeId'                                 ]
    // The payload
    [simple     AdsData    'data'   ['commandId', 'state.response']         ]
]

[enum uint 16 'CommandId'
    ['0x0000' INVALID]
    ['0x0001' ADS_READ_DEVICE_INFO]
    ['0x0002' ADS_READ]
    ['0x0003' ADS_WRITE]
    ['0x0004' ADS_READ_STATE]
    ['0x0005' ADS_WRITE_CONTROL]
    ['0x0006' ADS_ADD_DEVICE_NOTIFICATION]
    ['0x0007' ADS_DELETE_DEVICE_NOTIFICATION]
    ['0x0008' ADS_DEVICE_NOTIFICATION]
    ['0x0009' ADS_READ_WRITE]
]

[type 'State'
    [simple     bit 'initCommand'           ]
    [simple     bit 'updCommand'            ]
    [simple     bit 'timestampAdded'        ]
    [simple     bit 'highPriorityCommand'   ]
    [simple     bit 'systemCommand'         ]
    [simple     bit 'adsCommand'            ]
    [simple     bit 'noReturn'              ]
    [simple     bit 'response'              ]
    [simple     bit 'broadcast'             ]
    [reserved   int 7 '0x0'                 ]
]

// It is not only possible to exchange data between TwinCAT modules on one PC, it is even possible to do so by ADS
// methods between multiple TwinCAT PC's on the network.
// <p>
// Every PC on the network can be uniquely identified by a TCP/IP address, such as "172.1.2.16". The AdsAmsNetId is an
// extension of the TCP/IP address and identifies a TwinCAT message router, e.g. "172.1.2.16.1.1". TwinCAT message
// routers exist on every TwinCAT PC, and on every Beckhoff BCxxxx bus controller (e.g. BC3100, BC8100, BC9000, ...).
// <p>
// The AmsNetId consists of 6 bytes and addresses the transmitter or receiver. One possible AmsNetId would be e.g.
// "172.16.17.10.1.1". The storage arrangement in this example is as follows:
// <p>
// _____0     1     2     3     4     5
// __+-----------------------------------+
// 0 | 127 |  16 |  17 |  10 |   1 |   1 |
// __+-----------------------------------+
// <p>
// The AmsNetId is purely logical and has usually no relation to the IP address. The AmsNetId is configured at the
// target system. At the PC for this the TwinCAT System Control is used. If you use other hardware, see the considering
// documentation for notes about settings of the AMS NetId.
// @see <a href="https://infosys.beckhoff.com/content/1033/tcadscommon/html/tcadscommon_identadsdevice.htm?id=3991659524769593444">ADS device identification</a>
[type 'AmsNetId'
    [simple     uint        8   'octet1'            ]
    [simple     uint        8   'octet2'            ]
    [simple     uint        8   'octet3'            ]
    [simple     uint        8   'octet4'            ]
    [simple     uint        8   'octet5'            ]
    [simple     uint        8   'octet6'            ]
]

[discriminatedType 'AdsData' [CommandId 'commandId', bit 'response']
    [typeSwitch 'commandId', 'response'
        ['CommandId.INVALID', 'false' AdsInvalidRequest]
        ['CommandId.INVALID', 'true' AdsInvalidResponse]

        ['CommandId.ADS_READ_DEVICE_INFO', 'false' AdsReadDeviceInfoRequest]
        ['CommandId.ADS_READ_DEVICE_INFO', 'true' AdsReadDeviceInfoResponse
            // 4 bytes	ADS error number.
            [simple uint 32 'result']
            // Version	1 byte	Major version number
            [simple uint 8  'majorVersion']
            // Version	1 byte	Minor version number
            [simple uint 8  'minorVersion']
            // Build	2 bytes	Build number
            [simple uint 16  'version']
            // Name	16 bytes	Name of ADS device
            [array int 8  'device' count '16']
        ]

        ['CommandId.ADS_READ', 'false' AdsReadRequest
            // 4 bytes	Index Group of the data which should be read.
            [simple uint 32 'indexGroup']
            // 4 bytes	Index Offset of the data which should be read.
            [simple uint 32 'indexOffset']
            // 4 bytes	Length of the data (in bytes) which should be read.
            [simple uint 32 'length']
        ]
        ['CommandId.ADS_READ', 'true' AdsReadResponse
            // 4 bytes	ADS error number
            [simple uint 32 'result']
            // 4 bytes	Length of data which are supplied back.
            [implicit uint 32 'length' 'COUNT(data)']
            // n bytes	Data which are supplied back.
            [array int 8 'data' count 'length']
        ]

        ['CommandId.ADS_WRITE', 'false' AdsWriteRequest
            // 4 bytes	Index Group of the data which should be written.
            [simple uint 32 'indexGroup']
            // 4 bytes	Index Offset of the data which should be written.
            [simple uint 32 'indexOffset']
            // 4 bytes	Length of the data (in bytes) which should be written.
            [implicit uint 32 'length' 'COUNT(data)']
            // n bytes	Data which are written in the ADS device.
            [array int 8 'data' count 'length']
        ]
        ['CommandId.ADS_WRITE', 'true' AdsWriteResponse
            // 4 bytes	ADS error number
            [simple uint 32 'result']
        ]

        ['CommandId.ADS_READ_STATE', 'false' AdsReadStateRequest]
        ['CommandId.ADS_READ_STATE', 'true' AdsReadStateResponse
            // 4 bytes	ADS error number
            [simple uint 32 'result']
            // 2 bytes	New ADS status (see data type ADSSTATE of the ADS-DLL).
            [simple uint 16 'adsState']
            // 2 bytes	New device status.
            [simple uint 16 'deviceState']
        ]

        ['CommandId.ADS_WRITE_CONTROL', 'false' AdsWriteControlRequest
            // 2 bytes	New ADS status (see data type ADSSTATE of the ADS-DLL).
            [simple uint 16 'adsState']
            // 2 bytes	New device status.
            [simple uint 16 'deviceState']
            // 4 bytes	Length of data in byte.
            [implicit uint 32 'length' 'COUNT(data)']
            // n bytes	Additional data which are sent to the ADS device
            [array int 8 'data' count 'length']
        ]
        ['CommandId.ADS_WRITE_CONTROL', 'true' AdsWriteControlResponse
            // 4 bytes	ADS error number
            [simple uint 32 'result']
        ]

        ['CommandId.ADS_ADD_DEVICE_NOTIFICATION', 'false' AdsAddDeviceNotificationRequest
            // 4 bytes	Index Group of the data, which should be sent per notification.
            [simple uint 32 'indexGroup']
            // 4 bytes	Index Offset of the data, which should be sent per notification.
            [simple uint 32 'indexOffset']
            // 4 bytes	Index Offset of the data, which should be sent per notification.
            // 4 bytes	Length of data in bytes, which should be sent per notification.
            [simple uint 32 'length']
            // 4 bytes	See description of the structure ADSTRANSMODE at the ADS-DLL.
            [simple uint 32 'transmissionMode']
            // 4 bytes	At the latest after this time, the ADS Device Notification is called. The unit is 1ms.
            [simple uint 32 'maxDelay']
            // 4 bytes	The ADS server checks if the value changes in this time slice. The unit is 1ms
            [simple uint 32 'cycleTime']
            // 16bytes	Must be set to 0
            [reserved   uint       128       '0x0000' ]
        ]
        ['CommandId.ADS_ADD_DEVICE_NOTIFICATION', 'true' AdsAddDeviceNotificationResponse
            // 4 bytes	ADS error number
            [simple uint 32 'result']
            // 4 bytes	Handle of notification
            [simple uint 32 'notificationHandle']
        ]

        ['CommandId.ADS_DELETE_DEVICE_NOTIFICATION', 'false' AdsDeleteDeviceNotificationRequest
            // 4 bytes	Handle of notification
            [simple uint 32 'notificationHandle']
        ]
        ['CommandId.ADS_DELETE_DEVICE_NOTIFICATION', 'true' AdsDeleteDeviceNotificationResponse
            // 4 bytes	ADS error number
            [simple uint 32 'result']
        ]

        ['CommandId.ADS_DEVICE_NOTIFICATION', 'false' AdsDeviceNotificationRequest
            // 4 bytes	Size of data in byte.
            [simple uint 32 'length']
            // 4 bytes	Number of elements of type AdsStampHeader.
            [simple uint 32 'stamps']
            // n bytes	Array with elements of type AdsStampHeader.
            [array AdsStampHeader 'adsStampHeaders' count 'stamps']
        ]
        ['CommandId.ADS_DEVICE_NOTIFICATION', 'true' AdsDeviceNotificationResponse]

        ['CommandId.ADS_READ_WRITE', 'false' AdsReadWriteRequest
            // 4 bytes	Index Group of the data which should be written.
            [simple uint 32 'indexGroup']
            // 4 bytes	Index Offset of the data which should be written.
            [simple uint 32 'indexOffset']
            // 4 bytes	Length of data in bytes, which should be read.
            [simple uint 32 'readLength']
            // 4 bytes	Length of the data (in bytes) which should be written.
            [implicit uint 32 'writeLength' '(COUNT(items) * ((indexGroup == ReservedIndexGroups.ADSIGRP_MULTIPLE_READ_WRITE.value) ? 16 : 12)) + COUNT(data)']
            // Only if the indexGroup implies a sum-read response, will the indexOffset indicate the number of elements.
            [array  AdsMultiRequestItem 'items' COUNT '((indexGroup == ReservedIndexGroups.ADSIGRP_MULTIPLE_READ.value) || (indexGroup == ReservedIndexGroups.ADSIGRP_MULTIPLE_WRITE.value) || (indexGroup == ReservedIndexGroups.ADSIGRP_MULTIPLE_READ_WRITE.value)) ? indexOffset : 0' ['indexGroup']]
            // n bytes	Data which are written in the ADS device.
            [array int 8 'data' count 'writeLength - (COUNT(items) * 12)']
        ]
        ['CommandId.ADS_READ_WRITE', 'true' AdsReadWriteResponse
            // 4 bytes	ADS error number
            [simple uint 32 'result']
            // 4 bytes	Length of data in byte.
            [implicit uint 32 'length'  'COUNT(data)']
            // n bytes Additional data which are sent to the ADS device
            [array int 8 'data' count 'length']
        ]
    ]
]

[discriminatedType 'AdsMultiRequestItem' [uint 32 'indexGroup']
    [typeSwitch 'indexGroup'
        // ReservedIndexGroups.ADSIGRP_MULTIPLE_READ
        ['61568L' AdsMultiRequestItemRead
            // 4 bytes	Index Group of the data which should be written.
            [simple uint 32 'itemIndexGroup']
            // 4 bytes	Index Offset of the data which should be written.
            [simple uint 32 'itemIndexOffset']
            // 4 bytes	Length of data in bytes, which should be read.
            [simple uint 32 'itemReadLength']
        ]
        // ReservedIndexGroups.ADSIGRP_MULTIPLE_WRITE
        ['61569L' AdsMultiRequestItemWrite
            // 4 bytes	Index Group of the data which should be written.
            [simple uint 32 'itemIndexGroup']
            // 4 bytes	Index Offset of the data which should be written.
            [simple uint 32 'itemIndexOffset']
            // 4 bytes	Length of the data (in bytes) which should be written.
            [simple uint 32 'itemWriteLength']
        ]
        // ReservedIndexGroups.ADSIGRP_MULTIPLE_READ_WRITE
        ['61570L' AdsMultiRequestItemReadWrite
            // 4 bytes	Index Group of the data which should be written.
            [simple uint 32 'itemIndexGroup']
            // 4 bytes	Index Offset of the data which should be written.
            [simple uint 32 'itemIndexOffset']
            // 4 bytes	Length of data in bytes, which should be read.
            [simple uint 32 'itemReadLength']
            // 4 bytes	Length of the data (in bytes) which should be written.
            [simple uint 32 'itemWriteLength']
        ]
    ]
]

[type 'AdsStampHeader'
    // 8 bytes	The timestamp is coded after the Windows FILETIME format. I.e. the value contains the number of the nano seconds, which passed since 1.1.1601. In addition, the local time change is not considered. Thus the time stamp is present as universal Coordinated time (UTC).
    [simple uint 64 'timestamp']
    // 4 bytes	Number of elements of type AdsNotificationSample.
    [simple uint 32 'samples']
    // n bytes	Array with elements of type AdsNotificationSample.
    [array AdsNotificationSample 'adsNotificationSamples' count 'samples']
]

[type 'AdsNotificationSample'
    // 4 bytes	Handle of notification
    [simple uint 32 'notificationHandle']
    // 4 Bytes	Size of data range in bytes.
    [simple uint 32 'sampleSize']
    // n Bytes	Data
    [array int 8 'data' count 'sampleSize']
]

[dataIo 'DataItem' [AdsDataType 'adsDataType']
    [typeSwitch 'adsDataType'
        // -----------------------------------------
        // Bit
        // -----------------------------------------
        ['AdsDataType.BOOL' Boolean
            [reserved uint 7 '0x00']
            [simple   bit    'value']
        ]
        ['AdsDataType.BIT' Boolean
            [reserved uint 7 '0x00']
            [simple   bit    'value']
        ]
        ['AdsDataType.BIT8' Boolean
            [reserved uint 7 '0x00']
            [simple   bit    'value']
        ]

        // -----------------------------------------
        // Bit-strings
        // -----------------------------------------
        // 1 byte
        ['AdsDataType.BYTE' List
            [array bit 'value' count '8']
        ]
        ['AdsDataType.BITARR8' List
            [array bit 'value' count '8']
        ]
        // 2 byte (16 bit)
        ['AdsDataType.WORD' List
            [array bit 'value' count '16']
        ]
        ['AdsDataType.BITARR16' List
            [array bit 'value' count '16']
        ]
        // 4 byte (32 bit)
        ['AdsDataType.DWORD' List
            [array bit 'value' count '32']
        ]
        ['AdsDataType.BITARR32' List
            [array bit 'value' count '32']
        ]

        // -----------------------------------------
        // Integers
        // -----------------------------------------
        // 8 bit:
        ['AdsDataType.SINT' Integer
            [simple int 8 'value']
        ]
        ['AdsDataType.INT8' Integer
            [simple int 8 'value']
        ]
        ['AdsDataType.USINT' Integer
            [simple uint 8 'value']
        ]
        ['AdsDataType.UINT8' Integer
            [simple uint 8 'value']
        ]
        // 16 bit:
        ['AdsDataType.INT' Integer
            [simple int 16 'value']
        ]
        ['AdsDataType.INT16' Integer
            [simple int 16 'value']
        ]
        ['AdsDataType.UINT' Integer
            [simple uint 16 'value']
        ]
        ['AdsDataType.UINT16' Integer
            [simple uint 16 'value']
        ]
        // 32 bit:
        ['AdsDataType.DINT' Integer
            [simple int 32 'value']
        ]
        ['AdsDataType.INT32' Integer
            [simple int 32 'value']
        ]
        ['AdsDataType.UDINT' Long
            [simple uint 32 'value']
        ]
        ['AdsDataType.UINT32' Long
            [simple uint 32 'value']
        ]
        // 64 bit:
        ['AdsDataType.LINT' Long
            [simple int 64 'value']
        ]
        ['AdsDataType.INT64' Long
            [simple int 64 'value']
        ]
        ['AdsDataType.ULINT' BigInteger
            [simple uint 64 'value']
        ]
        ['AdsDataType.UINT64' BigInteger
            [simple uint 64 'value']
        ]

        // -----------------------------------------
        // Floating point values
        // -----------------------------------------
        ['AdsDataType.REAL' Float
            [simple float 8.23  'value']
        ]
        ['AdsDataType.FLOAT' Float
            [simple float 8.23  'value']
        ]
        ['AdsDataType.LREAL' Double
            [simple float 11.52 'value']
        ]
        ['AdsDataType.DOUBLE' Double
            [simple float 11.52 'value']
        ]

        // -----------------------------------------
        // Characters & Strings
        // -----------------------------------------
        ['AdsDataType.STRING' String
//            [manual string 'UTF-8' 'value' 'STATIC_CALL("org.apache.plc4x.java.amsads.utils.StaticHelper.parseAmsString", io, _type.encoding)' 'STATIC_CALL("org.apache.plc4x.java.amsads.utils.StaticHelper.serializeAmsString", io, _value, _type.encoding)' '_value.length + 2']
        ]
    ]
]

[enum 'AdsDataType' [uint 8 'numBytes']
    [BOOL       ['1']]
    [BIT        ['1']]
    [BIT8       ['1']]
    // -----------------------------------------
    // Bit-strings
    // -----------------------------------------
    // 1 byte
    [BYTE       ['1']]
    [BITARR8    ['1']]
    // 2 byte (16 bit)
    [WORD       ['2']]
    [BITARR16   ['2']]
    // 4 byte (32 bit)
    [DWORD      ['4']]
    [BITARR32   ['4']]
    // -----------------------------------------
    // Integers
    // -----------------------------------------
    // 8 bit:
    [SINT       ['1']]
    [INT8       ['1']]
    [USINT      ['1']]
    [UINT8      ['1']]
    // 16 bit:
    [INT        ['2']]
    [INT16      ['2']]
    [UINT       ['2']]
    [UINT16     ['2']]
    // 32 bit:
    [DINT       ['4']]
    [INT32      ['4']]
    [UDINT      ['4']]
    [UINT32     ['4']]
    // 64 bit:
    [LINT       ['8']]
    [INT64      ['8']]
    [ULINT      ['8']]
    [UINT64     ['8']]
    // -----------------------------------------
    // Floating point values
    // -----------------------------------------
    [REAL       ['4']]
    [FLOAT      ['4']]
    [LREAL      ['8']]
    [DOUBLE     ['8']]
    // -----------------------------------------
    // Characters & Strings
    // -----------------------------------------
    [STRING     ['9']]
]

[enum uint 32 'ReservedIndexGroups'
    ['0x0000F000' ADSIGRP_SYMTAB]
    ['0x0000F001' ADSIGRP_SYMNAME]
    ['0x0000F002' ADSIGRP_SYMVAL]
    ['0x0000F003' ADSIGRP_SYM_HNDBYNAME]
    ['0x0000F004' ADSIGRP_SYM_VALBYNAME]
    ['0x0000F005' ADSIGRP_SYM_VALBYHND]
    ['0x0000F006' ADSIGRP_SYM_RELEASEHND]
    ['0x0000F007' ADSIGRP_SYM_INFOBYNAME]
    ['0x0000F008' ADSIGRP_SYM_VERSION]
    ['0x0000F009' ADSIGRP_SYM_INFOBYNAMEEX]
    ['0x0000F00A' ADSIGRP_SYM_DOWNLOAD]
    ['0x0000F00B' ADSIGRP_SYM_UPLOAD]
    ['0x0000F00C' ADSIGRP_SYM_UPLOADINFO]
    ['0x0000F010' ADSIGRP_SYMNOTE]
    ['0x0000F020' ADSIGRP_IOIMAGE_RWIB]
    ['0x0000F021' ADSIGRP_IOIMAGE_RWIX]
    ['0x0000F025' ADSIGRP_IOIMAGE_RISIZE]
    ['0x0000F030' ADSIGRP_IOIMAGE_RWOB]
    ['0x0000F031' ADSIGRP_IOIMAGE_RWOX]
    ['0x0000F035' ADSIGRP_IOIMAGE_RWOSIZE]
    ['0x0000F040' ADSIGRP_IOIMAGE_CLEARI]
    ['0x0000F050' ADSIGRP_IOIMAGE_CLEARO]
    ['0x0000F060' ADSIGRP_IOIMAGE_RWIOB]
    ['0x0000F080' ADSIGRP_MULTIPLE_READ]
    ['0x0000F081' ADSIGRP_MULTIPLE_WRITE]
    ['0x0000F082' ADSIGRP_MULTIPLE_READ_WRITE]
    ['0x0000F083' ADSIGRP_MULTIPLE_RELEASE_HANDLE]
    ['0x0000F100' ADSIGRP_DEVICE_DATA]
    ['0x00000000' ADSIOFFS_DEVDATA_ADSSTATE]
    ['0x00000002' ADSIOFFS_DEVDATA_DEVSTATE]
]