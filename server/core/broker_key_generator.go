//Copyright 2017 Huawei Technologies Co., Ltd
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
package core

import (
	"strconv"

	"github.com/ServiceComb/service-center/pkg/util"
)

const (
	BROKER_ROOT_KEY                 = "cse-broker"
	BROKER_PARTICIPANTS_KEY         = "participants"
	BROKER_PARTICIPANTS_VERSION_KEY = "versions"
	BROKER_PARTIES_TAG_KEY          = "tags"

	BROKER_PACTS_KEY    = "pacts"
	BOKER_PACTS_PUB_KEY = "pact-pub"
)

// GetBrokerRootKey returns url (/cse-broker)
func GetBrokerRootKey() string {
	return util.StringJoin([]string{
		"",
		BROKER_ROOT_KEY,
	}, "/")
}

// GetBrokerParticipantRootKey returns url (/cse-broker/participants/TENANT)
func GetBrokerParticipantRootKey(tenant string) string {
	return util.StringJoin([]string{
		GetBrokerRootKey(),
		BROKER_PARTICIPANTS_KEY,
		tenant,
	}, "/")
}

// GetBrokerPartiesVersionRootKey returns url (/cse-broker/versions/TENANT)
func GetBrokerPartiesVersionRootKey(tenant string) string {
	return util.StringJoin([]string{
		GetBrokerRootKey(),
		BROKER_PARTICIPANTS_VERSION_KEY,
		tenant,
	}, "/")
}

// GenerateBrokerParticipantKey returns url (/cse-broker/participants/TENANT/APPID/SERVICE_NAME)
func GenerateBrokerParticipantKey(tenant string, appID string, serviceName string) string {
	return util.StringJoin([]string{
		GetBrokerParticipantRootKey(tenant),
		appID,
		serviceName,
	}, "/")
}

// GenerateBrokerPartiesVersionKey returns url (/cse-broker/versions/TENANT/PARTY_ID/NUMBER)
func GenerateBrokerPartiesVersionKey(tenant string, participantID int32, number string) string {
	return util.StringJoin([]string{
		GetBrokerPartiesVersionRootKey(tenant),
		strconv.Itoa(int(participantID)),
		number,
	}, "/")
}
