/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package configtx

import (
	"fmt"

	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/utils"

	"github.com/golang/protobuf/proto"
)

// UnmarshalConfig attempts to unmarshal bytes to a *cb.Config
func UnmarshalConfig(data []byte) (*cb.Config, error) {
	config := &cb.Config{}
	err := proto.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// UnmarshalConfigurationEnvelope attempts to unmarshal bytes to a *cb.ConfigurationEnvelope
func UnmarshalConfigurationEnvelope(data []byte) (*cb.ConfigurationEnvelope, error) {
	configEnv := &cb.ConfigurationEnvelope{}
	err := proto.Unmarshal(data, configEnv)
	if err != nil {
		return nil, err
	}
	return configEnv, nil
}

func ConfigurationEnvelopeFromBlock(block *cb.Block) (*cb.ConfigurationEnvelope, error) {
	if block.Data == nil || len(block.Data.Data) != 1 {
		return nil, fmt.Errorf("Not a configuration block, must contain exactly one tx")
	}

	envelope, err := utils.UnmarshalEnvelope(block.Data.Data[0])
	if err != nil {
		return nil, err
	}

	payload, err := utils.UnmarshalPayload(envelope.Payload)
	if err != nil {
		return nil, err
	}

	return UnmarshalConfigurationEnvelope(payload.Data)
}
