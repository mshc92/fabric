/*
Copyright IBM Corp. 2016 All Rights Reserved.

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

package config

import (
	"testing"

	"github.com/hyperledger/fabric/common/configvalues/channel"
	"github.com/hyperledger/fabric/common/configvalues/msp"

	logging "github.com/op/go-logging"
	"github.com/stretchr/testify/assert"
)

func init() {
	logging.SetLevel(logging.DEBUG, "")
}

func TestBeginBadRoot(t *testing.T) {
	r := NewRoot(&msp.MSPConfigHandler{})

	_, err := r.BeginValueProposals([]string{channel.GroupKey, channel.GroupKey})
	assert.Error(t, err, "Only one root element allowed")

	_, err = r.BeginValueProposals([]string{"foo"})
	assert.Error(t, err, "Non %s group not allowed", channel.GroupKey)
}

func TestProposeValue(t *testing.T) {
	r := NewRoot(&msp.MSPConfigHandler{})

	err := r.ProposeValue("foo", nil)
	assert.Error(t, err, "ProposeValue should return error")
}
