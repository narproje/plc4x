/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 */
package org.apache.plc4x.java.amsads;

import org.apache.plc4x.java.amsads.configuration.AdsConfiguration;
import org.apache.plc4x.java.amsads.field.AdsFieldHandler;
import org.apache.plc4x.java.amsads.protocol.AdsProtocolLogic;
import org.apache.plc4x.java.amsads.readwrite.AmsTCPPacket;
import org.apache.plc4x.java.amsads.readwrite.io.AmsTCPPacketIO;
import org.apache.plc4x.java.spi.configuration.Configuration;
import org.apache.plc4x.java.spi.connection.GeneratedDriverBase;
import org.apache.plc4x.java.spi.connection.ProtocolStackConfigurer;
import org.apache.plc4x.java.spi.connection.SingleProtocolStackConfigurer;

/**
 * Implementation of the ADS protocol, based on:
 * - ADS Protocol
 * - TCP
 * - Serial
 */
public class AMSADSPlcDriver extends GeneratedDriverBase<AmsTCPPacket> {

    public static final int TCP_PORT = 48898;

    @Override
    public String getProtocolCode() {
        return "ads";
    }

    @Override
    public String getProtocolName() {
        return "Beckhoff TwinCat ADS";
    }

    @Override
    protected boolean canRead() {
        return true;
    }

    @Override
    protected boolean canWrite() {
        return true;
    }

    @Override
    protected Class<? extends Configuration> getConfigurationType() {
        return AdsConfiguration.class;
    }

    @Override
    protected String getDefaultTransport() {
        return "tcp";
    }

    @Override
    protected AdsFieldHandler getFieldHandler() {
        return new AdsFieldHandler();
    }

    @Override
    protected ProtocolStackConfigurer<AmsTCPPacket> getStackConfigurer() {
        return SingleProtocolStackConfigurer.builder(AmsTCPPacket.class, AmsTCPPacketIO.class)
            .withProtocol(AdsProtocolLogic.class)
            .littleEndian()
            .build();
    }

}
