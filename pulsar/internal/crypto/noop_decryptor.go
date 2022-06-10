// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package crypto

import (
	"fmt"

	pb "github.com/skulkarni-ns/pulsar-client-go/pulsar/internal/pulsar_proto"
)

type noopDecryptor struct{}

func NewNoopDecryptor() Decryptor {
	return &noopDecryptor{}
}

// Decrypt noop decryptor
func (d *noopDecryptor) Decrypt(payload []byte,
	msgID *pb.MessageIdData,
	msgMetadata *pb.MessageMetadata) ([]byte, error) {
	if len(msgMetadata.GetEncryptionKeys()) > 0 {
		return payload, fmt.Errorf("incoming message payload is encrypted, consumer is not configured to decrypt")
	}
	return payload, nil
}
