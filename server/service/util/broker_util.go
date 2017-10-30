package util

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"math"
	"net/url"
	"strings"

	"github.com/ServiceComb/service-center/pkg/util"
	apt "github.com/ServiceComb/service-center/server/core"
	pb "github.com/ServiceComb/service-center/server/core/proto"
	"github.com/ServiceComb/service-center/server/core/registry"
	"github.com/ServiceComb/service-center/server/core/registry/store"
)

const (
	BROKER_HOME_URL                      = "/broker/v3/"
	BROKER_PARTICIPANTS_URL              = "/broker/v3/participants"
	BROKER_PARTICIPANT_URL               = "/broker/v3/participants/:participantId"
	BROKER_PARTY_VERSIONS_URL            = "/broker/v3/participants/:participantId/versions"
	BROKER_PARTY_LATEST_VERSION_URL      = "/broker/v3/participants/:participantId/versions/latest"
	BROKER_PARTY_VERSION_URL             = "/broker/v3/participants/:participantId/versions/:number"
	BROKER_PROVIDER_URL                  = "/broker/v3/pacts/provider"
	BROKER_PROVIDER_LATEST_PACTS_URL     = "/broker/v3/pacts/provider/:providerId/latest"
	BROKER_PROVIDER_LATEST_PACTS_TAG_URL = "/broker/v3/pacts/provider/:providerId/latest/:tag"
	BROKER_PACTS_LATEST_URL              = "/broker/v3/pacts/latest"

	BROKER_PUBLISH_URL  = "/broker/v3/pacts/provider/:providerId/consumer/:consumerId/version/:number"
	BROKER_WEBHOOHS_URL = "/broker/v3/webhooks"

	BROKER_CURIES_URL = "/broker/v3/doc/:rel"
)

var brokerPartyApiLinksValues = map[string]string{
	"self":           BROKER_PARTICIPANT_URL,
	"latest-version": BROKER_PARTY_LATEST_VERSION_URL,
	"versions":       BROKER_PARTY_VERSIONS_URL,
}

var brokerPartyVerApiLinksValues = map[string]string{
	"self":                                                         BROKER_PARTY_VERSION_URL,
	"pb:pacticipant":                                               BROKER_PARTICIPANT_URL,
	"pb:latest-verification-results-where-pacticipant-is-consumer": BROKER_PARTY_LATEST_VERSION_URL,
}

var brokerPartyVerApiLinksTitles = map[string]string{
	"self":                                                         "Version",
	"pb:pacticipant":                                               "Participant",
	"pb:latest-verification-results-where-pacticipant-is-consumer": "Latest verification results for consumer version",
}

var brokerApiLinksValues = map[string]string{
	"self":                              BROKER_HOME_URL,
	"pb:publish-pact":                   BROKER_PUBLISH_URL,
	"pb:latest-pact-versions":           BROKER_PACTS_LATEST_URL,
	"pb:pacticipants":                   BROKER_PARTICIPANTS_URL,
	"pb:latest-provider-pacts":          BROKER_PROVIDER_LATEST_PACTS_URL,
	"pb:latest-provider-pacts-with-tag": BROKER_PROVIDER_LATEST_PACTS_TAG_URL,
	"pb:webhooks":                       BROKER_WEBHOOHS_URL,
}

var brokerApiLinksTitles = map[string]string{
	"self":                              "Index",
	"pb:publish-pact":                   "Publish a pact",
	"pb:latest-pact-versions":           "Latest pact versions",
	"pb:pacticipants":                   "Pacticipants",
	"pb:latest-provider-pacts":          "Latest pacts by provider",
	"pb:latest-provider-pacts-with-tag": "Latest pacts by provider with a specified tag",
	"pb:webhooks":                       "Webhooks",
}

var brokerApiLinksTempl = map[string]bool{
	"self":                              false,
	"pb:publish-pact":                   true,
	"pb:latest-pact-versions":           false,
	"pb:pacticipants":                   false,
	"pb:latest-provider-pacts":          true,
	"pb:latest-provider-pacts-with-tag": true,
	"pb:webhooks":                       false,
}

func GenerateBrokerAPIPath(scheme string, host string, apiPath string,
	replacer *strings.Replacer) string {
	genPath := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   apiPath,
	}
	if replacer != nil {
		return replacer.Replace(genPath.String())
	}
	return genPath.String()
}

func GetBrokerHomeLinksAPIS(scheme string, host string, apiKey string) string {
	return GenerateBrokerAPIPath(scheme, host, brokerApiLinksValues[apiKey],
		strings.NewReplacer(":providerId", "{provider}",
			":consumerId", "{consumer}",
			":number", "{consumerApplicationVersion}",
			":tag", "{tag}"))
}

func GetBrokerParticipantsLinksAPIS(scheme string, host string,
	partidipantId string) map[string]*pb.BrokerAPIInfoEntry {

	replacer := strings.NewReplacer(":participantId", partidipantId)
	var apiEntries map[string]*pb.BrokerAPIInfoEntry
	apiEntries = make(map[string]*pb.BrokerAPIInfoEntry)

	for linkKey := range brokerPartyApiLinksValues {

		apiEntries[linkKey] = &pb.BrokerAPIInfoEntry{
			Href: GenerateBrokerAPIPath(scheme, host,
				brokerPartyApiLinksValues[linkKey], replacer),
			Title: brokerPartyVerApiLinksTitles[linkKey],
		}
	}
	return apiEntries
}

func GetBrokerPartyVersLinksAPIS(scheme string, host string,
	partidipantId string, partyName string,
	versionNum string) map[string]*pb.BrokerAPIInfoEntry {

	replacer := strings.NewReplacer(":participantId", partidipantId, ":number", versionNum)
	var apiEntries map[string]*pb.BrokerAPIInfoEntry
	apiEntries = make(map[string]*pb.BrokerAPIInfoEntry)
	apiEntry := &pb.BrokerAPIInfoEntry{}

	for linkKey := range brokerPartyVerApiLinksValues {
		hrefVal := GenerateBrokerAPIPath(scheme, host,
			brokerPartyVerApiLinksValues[linkKey], replacer)
		title := brokerPartyVerApiLinksTitles[linkKey]
		apiEntry = &pb.BrokerAPIInfoEntry{
			Href:  hrefVal,
			Title: title,
		}
		if strings.Compare(linkKey, "self") == 0 {
			apiEntry.Name = versionNum
		} else if strings.Compare(linkKey, "pb:pacticipant") != 0 {
			apiEntry.Name = partyName
		}
		apiEntries[linkKey] = apiEntry
	}

	return apiEntries
}

var participantIds int32 = 0
var participantVersionIds int32 = 0
var pactIds int32 = 0
var pactPubIds int32 = 0

//GetBrokerParticipantUtils returns the participant from ETCD
func GetBrokerParticipantUtils(ctx context.Context, tenant string, appId string,
	serviceName string, opts ...registry.PluginOpOption) (*pb.Participant, error) {

	key := apt.GenerateBrokerParticipantKey(tenant, appId, serviceName)
	opts = append(opts, registry.WithStrKey(key))
	participants, err := store.Store().BrokerParticipant().Search(ctx, opts...)

	if err != nil {
		util.Logger().Errorf(nil, "pact publish failed, participant with, could not be searched.")
		return nil, err
	}

	if len(participants.Kvs) == 0 {
		util.Logger().Info("GetParticipant found no participant")
		return nil, nil
	}

	participant := &pb.Participant{}
	err = json.Unmarshal(participants.Kvs[0].Value, participant)
	if err != nil {
		util.Logger().Errorf(nil, "pact publish failed, participant with id %d, could not be searched.")
		return nil, err
	}
	util.Logger().Infof("GetParticipant: (%d, %s, %s)", participant.Id, participant.AppId,
		participant.ServiceName)
	return participant, nil
}

//GetBrokerPartyFromServiceId returns the participant and the service from ETCD
func GetBrokerPartyFromServiceId(ctx context.Context, serviceId string) (*pb.Participant,
	*pb.MicroService, error, error) {

	tenant := util.ParseTenantProject(ctx)
	serviceParticipant, err := GetService(ctx, tenant, serviceId)
	if err != nil {
		util.Logger().Errorf(err,
			"get participant failed, serviceId is %s: query provider failed.", serviceId)
		return nil, nil, nil, err
	}
	if serviceParticipant == nil {
		util.Logger().Errorf(nil,
			"get participant failed, serviceId is %s: service not exist.", serviceId)
		return nil, nil, nil, errors.New("get participant, serviceId not exist.")
	}
	// Get or create provider participant
	participant, errBroker := GetBrokerParticipantUtils(ctx, tenant, serviceParticipant.AppId,
		serviceParticipant.ServiceName)
	if errBroker != nil {
		util.Logger().Errorf(errBroker,
			"get participant failed, serviceId %s: query participant failed.", serviceId)
		return nil, serviceParticipant, errBroker, err
	}
	if participant == nil {
		util.Logger().Errorf(nil,
			"get participant failed, particpant does not exist for serviceId %s", serviceId)
		return nil, serviceParticipant, errors.New("particpant does not exist for serviceId."), err
	}

	return participant, serviceParticipant, errBroker, nil
}

func GetBrokerPartyFromService(ctx context.Context,
	microservice *pb.MicroService) (*pb.Participant, error) {
	if microservice == nil {
		return nil, nil
	}
	tenant := util.ParseTenantProject(ctx)
	participant, errBroker := GetBrokerParticipantUtils(ctx, tenant, microservice.AppId,
		microservice.ServiceName)
	if errBroker != nil {
		util.Logger().Errorf(errBroker,
			"get participant failed, serviceId %s: query participant failed.",
			microservice.ServiceId)
		return nil, errBroker
	}
	return participant, errBroker
}

//AddBrokerParticipantIntoETCD adds the participant into ETCD
func AddBrokerParticipantIntoETCD(ctx context.Context, tenant string, appId string,
	serviceName string) (*pb.Participant, error) {

	newParticipant := &pb.Participant{Id: participantIds, AppId: appId,
		ServiceName: serviceName}
	data, err := json.Marshal(newParticipant)
	if err != nil {
		util.Logger().Errorf(nil, "pact publish failed, participant cannot be created.")
		return nil, err
	}
	key := apt.GenerateBrokerParticipantKey(tenant, appId, serviceName)
	_, err = registry.GetRegisterCenter().Do(ctx,
		registry.PUT,
		registry.WithStrKey(key),
		registry.WithValue(data))
	if err != nil {
		util.Logger().Errorf(nil, "pact publish failed, participant cannot be added to ETCD.")
		return nil, err
	}
	//TODO: Should we lock before incrementing the Ids?
	participantIds++
	return newParticipant, err
}

func GetAllBrokerPartyVersions(ctx context.Context, tenant string, participantId int32,
	opts ...registry.PluginOpOption) ([]*pb.ParticipantVersion, error) {
	key := apt.GetBrokerAllPartyVersionsKey(tenant, participantId)
	opts = append(opts, registry.WithStrKey(key))
	opts = append(opts, registry.WithPrefix())
	participantVersions, err := store.Store().BrokerVersion().Search(ctx, opts...)
	if err != nil {
		return nil, err
	}
	partyVersions := []*pb.ParticipantVersion{}
	if len(participantVersions.Kvs) == 0 {
		util.Logger().Infof("No versions found for participant (%d)", participantId)
		return nil, nil
	}
	for i := 0; i < len(participantVersions.Kvs); i++ {
		versionItem := &pb.ParticipantVersion{}
		err = json.Unmarshal(participantVersions.Kvs[i].Value, &versionItem)
		if err != nil {
			return nil, err
		}
		partyVersions = append(partyVersions, versionItem)
	}
	return partyVersions, nil
}

func GetLatestBrokerParticipantVersion(ctx context.Context, tenant string, participantId int32,
	opts ...registry.PluginOpOption) (*pb.ParticipantVersion, error) {

	key := apt.GetBrokerAllPartyVersionsKey(tenant, participantId)
	opts = append(opts, registry.WithStrKey(key))
	participantVersions, err := store.Store().BrokerVersion().Search(ctx, opts...)
	if err != nil {
		return nil, err
	}
	if len(participantVersions.Kvs) == 0 {
		util.Logger().Infof("No versions found for participant (%d)", participantId)
		return nil, nil
	}
	order := int32(math.MinInt32)
	participantVersion := pb.ParticipantVersion{}
	for i := 0; i < len(participantVersions.Kvs); i++ {
		versionItem := &pb.ParticipantVersion{}
		err = json.Unmarshal(participantVersions.Kvs[i].Value, &versionItem)
		if err != nil {
			return nil, err
		}
		if versionItem.Order > order {
			order = versionItem.Order
			participantVersion = *versionItem
		}
	}
	return &participantVersion, nil
}

func GetBrokerParticipantVersionOrder(ctx context.Context, tenant string, participantId int32,
	opts ...registry.PluginOpOption) (int32, error) {
	key := apt.GetBrokerAllPartyVersionsKey(tenant, participantId)
	opts = append(opts, registry.WithStrKey(key))
	participantVersions, err := store.Store().BrokerVersion().Search(ctx, opts...)
	if err != nil {
		return -1, err
	}
	return int32(len(participantVersions.Kvs)), nil
}

func GetBrokerParticipants(ctx context.Context, tenant string,
	opts ...registry.PluginOpOption) (map[string]*pb.Participant, error) {

	//first get the services by tenant
	servicesArr, err := GetAllServiceUtil(ctx, opts...)

	if err != nil {
		return nil, err
	}

	systemParticipants := make(map[string]*pb.Participant)

	for i := 0; i < len(servicesArr); i++ {
		partyItem, err := GetBrokerPartyFromService(ctx, servicesArr[i])
		if err != nil {
			return systemParticipants, err
		}
		if partyItem != nil {
			systemParticipants[servicesArr[i].ServiceId] = partyItem
		}
	}

	return systemParticipants, nil
}

func GetBrokerParticipantVersion(ctx context.Context, tenant string, participantId int32,
	number string, opts ...registry.PluginOpOption) (*pb.ParticipantVersion, error) {
	key := apt.GenerateBrokerPartiesVersionKey(tenant, participantId, number)
	opts = append(opts, registry.WithStrKey(key))
	participantVersions, err := store.Store().BrokerVersion().Search(ctx, opts...)
	if err != nil {
		return nil, err
	}
	if len(participantVersions.Kvs) == 0 {
		util.Logger().Infof("No versions found for participant (%d)", participantId)
		return nil, nil
	}
	participantVersion := &pb.ParticipantVersion{}
	err = json.Unmarshal(participantVersions.Kvs[0].Value, participantVersion)
	if err != nil {
		util.Logger().Infof("Unmarshalling partcipantVersion error of participant (%d)",
			participantId)
		return nil, err
	}
	return participantVersion, nil
}

func AddBrokerParticipantVersionIntoETCD(ctx context.Context, tenant string,
	participantId int32, number string,
	opts ...registry.PluginOpOption) (*pb.ParticipantVersion, error) {
	//TODO: How to set the order?

	orderValue, err := GetBrokerParticipantVersionOrder(ctx, tenant, participantId)

	if err != nil {
		util.Logger().Errorf(nil, "pact publish failed, could not query max order value.")
		return nil, err
	}

	newParticipantVersion := &pb.ParticipantVersion{Id: participantVersionIds,
		Number: number, ParticipantId: participantId, Order: orderValue}

	data, err := json.Marshal(newParticipantVersion)

	if err != nil {
		util.Logger().Errorf(nil, "pact publish failed, participantVersion cannot be created.")
		return nil, err
	}

	key := apt.GenerateBrokerPartiesVersionKey(tenant, participantId, number)
	_, err = registry.GetRegisterCenter().Do(ctx,
		registry.PUT,
		registry.WithStrKey(key),
		registry.WithValue(data))

	if err != nil {
		util.Logger().Errorf(nil, "pact publish failed, participantVersion cannot be added to ETCD.")
		return nil, err
	}

	participantVersionIds++
	return newParticipantVersion, nil
}

func GetBrokerPact(ctx context.Context, tenant string, consumerParticipantId int32,
	producerParticipantId int32, sha []byte,
	opts ...registry.PluginOpOption) (*pb.Pact, error) {

	key := apt.GenerateBrokerPactKey(tenant, consumerParticipantId, producerParticipantId, sha)
	opts = append(opts, registry.WithStrKey(key))
	pactEntries, err := store.Store().BrokerPact().Search(ctx, opts...)
	if err != nil {
		return nil, err
	}
	if len(pactEntries.Kvs) == 0 {
		util.Logger().Infof("No pact was stored between consumer (%d) and provider (%d)",
			consumerParticipantId, producerParticipantId)
		return nil, nil
	}

	pactEntry := &pb.Pact{}
	err = json.Unmarshal(pactEntries.Kvs[0].Value, pactEntry)

	if err != nil {
		util.Logger().Infof("Unmarshalling Pact error between consumer (%d) and provider (%d)",
			consumerParticipantId, producerParticipantId)
		return nil, err
	}

	return pactEntry, nil
}

func AddBrokerPactIntoETCD(ctx context.Context, tenant string, consumerParticipantId int32,
	producerParticipantId int32, sha []byte, content []byte,
	opts ...registry.PluginOpOption) (*pb.Pact, error) {

	newPactEntry := &pb.Pact{Id: pactIds, ConsumerParticipantId: consumerParticipantId,
		ProviderParticipantId: producerParticipantId, Sha: sha, Content: content}
	key := apt.GenerateBrokerPactKey(tenant, consumerParticipantId, producerParticipantId, sha)
	data, err := json.Marshal(newPactEntry)
	_, err = registry.GetRegisterCenter().Do(ctx,
		registry.PUT,
		registry.WithStrKey(key),
		registry.WithValue(data))

	if err != nil {
		util.Logger().Errorf(nil, "pact publish failed, pact cannot be added to ETCD.")
		return nil, err
	}

	pactIds++
	return newPactEntry, nil
}

func AddBrokerPctPublication(ctx context.Context, tenant string,
	consumerPartyVersiontId int32, providerPartyId int32,
	pactId int32) (*pb.PactPublication, error) {

	newPubPactEntry := &pb.PactPublication{Id: pactPubIds,
		ParticipantVersionId: consumerPartyVersiontId,
		PactId:               pactId,
		ParticipantId:        providerPartyId,
	}
	key := apt.GenerateBrokerPactPubKey(tenant, consumerPartyVersiontId, providerPartyId,
		pactId)
	data, err := json.Marshal(newPubPactEntry)
	if err != nil {
		_, err = registry.GetRegisterCenter().Do(ctx,
			registry.PUT,
			registry.WithStrKey(key),
			registry.WithValue(data))
	}

	if err != nil {
		util.Logger().Errorf(nil, "pact publish failed, pactPubs cannot be added to ETCD.")
		return nil, err
	}

	pactPubIds++
	return newPubPactEntry, nil
}

func CreateParticipantVersionResponse(participant *pb.Participant,
	participantVer *pb.ParticipantVersion, host string, scheme string,
	serviceId string) *pb.BrokerPartyVersionResponse {

	return &pb.BrokerPartyVersionResponse{
		Number: participantVer.Number,
		XLinks: GetBrokerPartyVersLinksAPIS(scheme, host,
			serviceId, participant.AppId+"/"+participant.ServiceName,
			participantVer.Number),
	}
}

func CreatePartyVersionsResponse(participant *pb.Participant,
	partyVersions []*pb.ParticipantVersion, host string, scheme string,
	serviceId string) *pb.GetParticipantVersionsResponse {
	versionResponseArr := []*pb.BrokerPartyVersionResponse{}
	pactVersionsArr := []*pb.BrokerAPIInfoEntry{}
	selfInfo := &pb.BrokerAPIInfoEntry{}
	for i := 0; i < len(partyVersions); i++ {
		versionItem := CreateParticipantVersionResponse(participant, partyVersions[i],
			host, scheme, serviceId)
		versionResponseArr = append(versionResponseArr, versionItem)
		selfInfo = versionItem.XLinks["self"]
		if selfInfo != nil {
			pactVersionsArr = append(pactVersionsArr, &pb.BrokerAPIInfoEntry{
				Href: selfInfo.Href,
			})
		}
	}

	return &pb.GetParticipantVersionsResponse{
		Response: pb.CreateResponse(pb.Response_SUCCESS,
			"success getting particpant versions"),
		XEmbedded: &pb.BrokerPartyEmbeddedVersions{
			Versions: versionResponseArr,
		},
	}
}

func CreateParticipantResponse(participant *pb.Participant, host string, scheme string,
	serviceId string) *pb.ParticipantResponse {
	var bufferName bytes.Buffer
	bufferName.WriteString(participant.GetAppId())
	bufferName.WriteString("/")
	bufferName.WriteString(participant.GetServiceName())

	partyAPIEntries := GetBrokerParticipantsLinksAPIS(scheme, host, serviceId)

	resp := &pb.ParticipantResponse{
		Name:   bufferName.String(),
		XLinks: partyAPIEntries,
	}
	return resp
}

func CreateParticipantsResponse(participants map[string]*pb.Participant, host string,
	scheme string) *pb.GetParticipantsResponse {

	partiesResponses := []*pb.ParticipantResponse{}
	summaryEntries := []*pb.BrokerAPIInfoEntry{}
	selfInfo := &pb.BrokerAPIInfoEntry{}

	for serviceID, partValue := range participants {
		partyResponseItem := CreateParticipantResponse(partValue, host, scheme, serviceID)
		partiesResponses = append(partiesResponses, partyResponseItem)
		selfInfo = partyResponseItem.XLinks["self"]
		summaryEntries = append(summaryEntries, &pb.BrokerAPIInfoEntry{
			Href:  selfInfo.Href,
			Title: partyResponseItem.Name,
		})
	}

	resp := &pb.GetParticipantsResponse{
		Response:     pb.CreateResponse(pb.Response_SUCCESS, "success getting particpants"),
		Participants: partiesResponses,
		XLinks: &pb.ParticipantsSummaryLinks{
			Self: &pb.BrokerAPIInfoEntry{
				Href: GenerateBrokerAPIPath(scheme, host,
					brokerApiLinksValues["pb:pacticipants"], nil),
			},
			Participants: summaryEntries,
		},
	}
	return resp
}

func CreateBrokerHomeResponse(host string, scheme string) *pb.BrokerHomeResponse {
	var apiEntries map[string]*pb.BrokerAPIInfoEntry
	apiEntries = make(map[string]*pb.BrokerAPIInfoEntry)

	for k := range brokerApiLinksValues {
		apiEntries[k] = &pb.BrokerAPIInfoEntry{
			Href:      GetBrokerHomeLinksAPIS(scheme, host, k),
			Title:     brokerApiLinksTitles[k],
			Templated: brokerApiLinksTempl[k],
		}
	}

	// apiEntries["self"] = CreateBrokerAPIEntry(urlPrefix, BROKER_HOME_URL, "Index", false)
	// apiEntries["pb:publish-pact"] = CreateBrokerAPIEntry(urlPrefix, BROKER_PROVIDER_URL,
	// 	"Publish a pact", true)
	// apiEntries["pb:latest-provider-pacts"] = CreateBrokerAPIEntry(urlPrefix,
	// 	BROKER_PROVIDER_URL+"/{provider}/latest", "Latest pacts by provider", true)
	//
	curies := []*pb.BrokerAPIInfoEntry{}
	curies = append(curies, &pb.BrokerAPIInfoEntry{
		Name: "pb",
		Href: GenerateBrokerAPIPath(scheme, host, BROKER_CURIES_URL,
			strings.NewReplacer(":rel", "{rel}")),
	})

	return &pb.BrokerHomeResponse{
		Response: pb.CreateResponse(pb.Response_SUCCESS, "Broker Home."),
		XLinks:   apiEntries,
		Curies:   curies,
	}
}
