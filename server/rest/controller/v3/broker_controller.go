package v3

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ServiceComb/service-center/pkg/rest"
	"github.com/ServiceComb/service-center/pkg/util"
	"github.com/ServiceComb/service-center/server/core"
	pb "github.com/ServiceComb/service-center/server/core/proto"
	"github.com/ServiceComb/service-center/server/rest/controller"
)

//BrokerService handles requests and responses to the PactBroker
type BrokerService struct {
	//
}

const (
	BROKER_HOME_URL  = "/broker/v3/"
	PARTICIPANTS_URL = "/broker/v3/participants"
	PROVIDER_URL     = "/broker/v3/pacts/provider"
)

func (this *BrokerService) URLPatterns() []rest.Route {
	return []rest.Route{
		{rest.HTTP_METHOD_GET, BROKER_HOME_URL, this.GetBrokerHome},
		{rest.HTTP_METHOD_GET, PARTICIPANTS_URL, this.GetParticipants},
		{rest.HTTP_METHOD_GET, PARTICIPANTS_URL + "/:participantId", this.GetParticipant},
		{rest.HTTP_METHOD_GET, PARTICIPANTS_URL + "/:participantId/versions",
			this.GetParticipantVersions},
		{rest.HTTP_METHOD_GET, PARTICIPANTS_URL + "/:participantId/versions/latest",
			this.GetParticipantLastVersion},
		{rest.HTTP_METHOD_GET, BROKER_HOME_URL + "pacts/latest", this.GetPacts},
		{rest.HTTP_METHOD_GET, PROVIDER_URL + "/:providerId/latest",
			this.GetProviderPacts},
		{rest.HTTP_METHOD_GET, PROVIDER_URL + "/:providerId/consumer/:consumerId/latest",
			this.RetrievePact},
		{rest.HTTP_METHOD_GET, PROVIDER_URL + "/:providerId/consumer/:consumerId/latest/:tag",
			this.RetrieveTaggedPact},
		{rest.HTTP_METHOD_GET, PROVIDER_URL + "/:providerId/consumer/:consumerId/latest-untagged",
			this.RetrieveUntaggedPact},
		{rest.HTTP_METHOD_POST, PROVIDER_URL + "/:providerId/consumer/:consumerId/version/:number",
			this.PublishPact},
	}
}

//GetBrokerHome returns home response
func (this *BrokerService) GetBrokerHome(w http.ResponseWriter, r *http.Request) {
	homeResponse := "{\r\n" +
		"  \"_links\": {\r\n" +
		"    \"pb:latest-provider-pacts\": {\r\n" +
		"      \"href\": \"http://" + r.Host + PROVIDER_URL + "/{provider}/latest\",\r\n" +
		"      \"title\": \"Latest pacts by provider\",\r\n" +
		"      \"templated\": true\r\n" +
		"    }\r\n" +
		"  }\r\n" +
		"}"
	controller.WriteJsonObject(http.StatusOK, []byte(homeResponse), w)
}

//GetParticipants returns all the participants in the broker.
//Participant is a party that participates in a pact (ie. a Consumer or a Provider).
func (this *BrokerService) GetParticipants(w http.ResponseWriter, r *http.Request) {
	request := &pb.GetParticipantsRequest{
		PartyReqType: pb.GetParticipantsRequest_ALL_PARTIES,
		UrlPrefix:    "http://" + r.Host + PARTICIPANTS_URL,
	}
	resp, err := core.BrokerServiceAPI.GetParticipants(r.Context(), request)
	controller.WriteTextResponse(resp.GetResponse(), err, "Get Participants response", w)
}

//GetParticipant returns a participant  in the broker.
//Participant is a party that participates in a pact (ie. a Consumer or a Provider).
func (this *BrokerService) GetParticipant(w http.ResponseWriter, r *http.Request) {
	request := &pb.GetParticipantsRequest{
		PartyReqType:  pb.GetParticipantsRequest_SINGLE_PARTY,
		ParticipantId: r.URL.Query().Get(":participantId"),
		UrlPrefix:     "http://" + r.Host + PARTICIPANTS_URL,
	}
	resp, err := core.BrokerServiceAPI.GetParticipant(r.Context(), request)
	controller.WriteTextResponse(resp.GetResponse(), err, "get participant", w)
}

//GetParticipantVersions returns all the participant versions.
func (this *BrokerService) GetParticipantVersions(w http.ResponseWriter, r *http.Request) {
	request := &pb.GetParticipantsRequest{
		PartyReqType:  pb.GetParticipantsRequest_VERSIONS_PARTY,
		ParticipantId: r.URL.Query().Get(":participantId"),
		UrlPrefix:     "http://" + r.Host + PARTICIPANTS_URL,
	}
	resp, err := core.BrokerServiceAPI.GetParticipantVersions(r.Context(), request)
	controller.WriteTextResponse(resp.GetResponse(), err, "get participant versions", w)

}

//GetParticipantLatstVersion returns the latest ersion of a participant.
func (this *BrokerService) GetParticipantLastVersion(w http.ResponseWriter, r *http.Request) {
	request := &pb.GetParticipantsRequest{
		PartyReqType:  pb.GetParticipantsRequest_LATEST_PARTY,
		ParticipantId: r.URL.Query().Get(":participantId"),
		UrlPrefix:     "http://" + r.Host + PARTICIPANTS_URL,
	}
	resp, err := core.BrokerServiceAPI.GetParticipantVersions(r.Context(), request)
	controller.WriteTextResponse(resp.GetResponse(), err, "get participant versions", w)
}

//GetPacts returns all the latest pacts
func (this *BrokerService) GetPacts(w http.ResponseWriter, r *http.Request) {

}

//GetProviderPacts returns all Latest pacts for a provider
func (this *BrokerService) GetProviderPacts(w http.ResponseWriter, r *http.Request) {

}

//RetrievePact returns the latest pact between a specified consumer and provider.
func (this *BrokerService) RetrievePact(w http.ResponseWriter, r *http.Request) {

}

//RetrieveUntaggedPact returns the latest pact without any tag between a specified consumer and provider.
func (this *BrokerService) RetrieveUntaggedPact(w http.ResponseWriter, r *http.Request) {

}

//RetrieveTaggedPact returns the latest pact with a tag between a specified consumer and provider.
func (this *BrokerService) RetrieveTaggedPact(w http.ResponseWriter, r *http.Request) {

}

//PublishPact handles publication of a pact taking the pact as a json.
//A pact is published to the broker using a combination of the provider name, the
//consumer name, and the consumer application version.
func (this *BrokerService) PublishPact(w http.ResponseWriter, r *http.Request) {
	message, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.Logger().Error("body err", err)
		controller.WriteText(http.StatusInternalServerError, fmt.Sprintf("body error %s", err.Error()), w)
		return
	}

	request := &pb.PublishPactRequest{
		ProviderId: r.URL.Query().Get(":providerId"),
		ConsumerId: r.URL.Query().Get(":consumerId"),
		Version:    r.URL.Query().Get(":number"),
		Pact:       message,
	}

	resp, err := core.BrokerServiceAPI.PublishPact(r.Context(), request)
	controller.WriteTextResponse(resp.GetResponse(), err, "pact publish success", w)
}
