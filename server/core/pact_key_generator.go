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

// GetBrokerPartiesTagRootKey returns url (/cse-broker/tags/TENANT)
func GetBrokerPartiesTagRootKey(tenant string) string {
	return util.StringJoin([]string{
		GetBrokerRootKey(),
		BROKER_PARTIES_TAG_KEY,
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

// GenerateBrokerPartiesTagKey returns url (/cse-broker/tags/TENANT/PARTICIPANT_ID/VERSION_ID)
func GenerateBrokerPartiesTagKey(tenant string, participantID int32, versionID int32) string {
	return util.StringJoin([]string{
		GetBrokerPartiesTagRootKey(tenant),
		strconv.Itoa(int(participantID)),
		strconv.Itoa(int(versionID)),
	}, "/")
}

// GetBrokerPactsRootKey returns url (/cse-broker/pacts/TENANT)
func GetBrokerPactsRootKey(tenant string) string {
	return util.StringJoin([]string{
		GetBrokerRootKey(),
		BROKER_PACTS_KEY,
		tenant,
	}, "/")
}

// GetBrokerPactsPubRootKey returns url (/cse-broker/pact-pub/TENANT)
func GetBrokerPactsPubRootKey(tenant string) string {
	return util.StringJoin([]string{
		GetBrokerRootKey(),
		BOKER_PACTS_PUB_KEY,
		tenant,
	}, "/")
}

// GenerateBrokerPactKey returns url (/cse-broker/pacts/TENANT/CONSUMER_ID/PROVIDER_ID/SHA
func GenerateBrokerPactKey(tenant string, consumerID int32, providerID int32, sha []byte) string {
	return util.StringJoin([]string{
		GetBrokerPactsRootKey(tenant),
		strconv.Itoa(int(consumerID)),
		strconv.Itoa(int(providerID)),
		string(sha),
	}, "/")
}

// GenerateBrokerPactPubKey returns url (/cse-broker/pact-pub/TENANT/PARTY_ID/PCT_ID
func GenerateBrokerPactPubKey(tenant string, participantVersionID int32, pactID int32) string {
	return util.StringJoin([]string{
		GetBrokerPactsRootKey(tenant),
		strconv.Itoa(int(participantVersionID)),
		strconv.Itoa(int(pactID)),
	}, "/")
}
