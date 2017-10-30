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
	uServiceUtil "github.com/ServiceComb/service-center/server/service/util"
)

//BrokerService handles requests and responses to the PactBroker
type BrokerService struct {
	//
}

func (this *BrokerService) URLPatterns() []rest.Route {
	return []rest.Route{
		{rest.HTTP_METHOD_GET, uServiceUtil.BROKER_HOME_URL, this.GetBrokerHome},
		{rest.HTTP_METHOD_GET, uServiceUtil.BROKER_PARTICIPANTS_URL, this.GetParticipants},
		{rest.HTTP_METHOD_GET, uServiceUtil.BROKER_PARTICIPANT_URL, this.GetParticipant},
		{rest.HTTP_METHOD_GET, uServiceUtil.BROKER_PARTY_VERSIONS_URL,
			this.GetParticipantVersions},
		{rest.HTTP_METHOD_GET, uServiceUtil.BROKER_PARTY_VERSION_URL,
			this.GetParticipantVersion},
		// {rest.HTTP_METHOD_GET, uServiceUtil.BROKER_PARTICIPANTS_URL + "/:participantId/versions/latest",
		// 	this.GetParticipantLastVersion},
		{rest.HTTP_METHOD_GET, uServiceUtil.BROKER_PACTS_LATEST_URL, this.GetPacts},
		{rest.HTTP_METHOD_GET, uServiceUtil.BROKER_PROVIDER_LATEST_PACTS_URL,
			this.GetProviderPacts},
		{rest.HTTP_METHOD_GET, uServiceUtil.BROKER_PROVIDER_URL + "/:providerId/consumer/:consumerId/latest",
			this.RetrievePact},
		{rest.HTTP_METHOD_GET, uServiceUtil.BROKER_PROVIDER_URL + "/:providerId/consumer/:consumerId/latest/:tag",
			this.RetrieveTaggedPact},
		{rest.HTTP_METHOD_GET, uServiceUtil.BROKER_PROVIDER_URL + "/:providerId/consumer/:consumerId/latest-untagged",
			this.RetrieveUntaggedPact},
		{rest.HTTP_METHOD_POST, uServiceUtil.BROKER_PUBLISH_URL, this.PublishPact},
	}
}

//GetBrokerHome returns home response
func (this *BrokerService) GetBrokerHome(w http.ResponseWriter, r *http.Request) {
	request := &pb.BaseBrokerRequest{
		HostAddress: r.Host,
		Scheme:      r.URL.Scheme,
	}
	resp, _ := core.BrokerServiceAPI.GetBrokerHome(r.Context(), request)
	if resp.GetResponse().GetCode() != pb.Response_SUCCESS {
		controller.WriteJson(http.StatusBadRequest, []byte(resp.Response.GetMessage()), w)
	} else {
		controller.WriteJsonObject(http.StatusOK, resp, w)
	}
}

//GetParticipants returns all the participants in the broker.
//Participant is a party that participates in a pact (ie. a Consumer or a Provider).
func (this *BrokerService) GetParticipants(w http.ResponseWriter, r *http.Request) {
	request := &pb.GetParticipantsRequest{
		PartyReqType: pb.GetParticipantsRequest_ALL_PARTIES,
		BrokerInfo: &pb.BaseBrokerRequest{
			HostAddress: r.Host,
			Scheme:      r.URL.Scheme,
		},
	}
	resp, err := core.BrokerServiceAPI.GetParticipants(r.Context(), request)
	if resp.GetResponse().GetCode() != pb.Response_SUCCESS {
		controller.WriteJson(http.StatusBadRequest, []byte(resp.Response.GetMessage()), w)
	} else {
		controller.WriteJsonObject(http.StatusOK, resp, w)
	}
}

//GetParticipant returns a participant  in the broker.
//Participant is a party that participates in a pact (ie. a Consumer or a Provider).
func (this *BrokerService) GetParticipant(w http.ResponseWriter, r *http.Request) {
	request := &pb.GetParticipantsRequest{
		PartyReqType:  pb.GetParticipantsRequest_SINGLE_PARTY,
		ParticipantId: r.URL.Query().Get(":participantId"),
		BrokerInfo: &pb.BaseBrokerRequest{
			HostAddress: r.Host,
			Scheme:      r.URL.Scheme,
		},
	}
	resp, err := core.BrokerServiceAPI.GetParticipant(r.Context(), request)
	if resp.GetResponse().GetCode() != pb.Response_SUCCESS {
		controller.WriteJson(http.StatusBadRequest, []byte(resp.Response.GetMessage()), w)
	} else {
		controller.WriteJsonObject(http.StatusOK, resp, w)
	}
}

//GetParticipantVersions returns all the participant versions.
func (this *BrokerService) GetParticipantVersions(w http.ResponseWriter, r *http.Request) {
	request := &pb.GetParticipantsRequest{
		PartyReqType:  pb.GetParticipantsRequest_VERSIONS_PARTY,
		ParticipantId: r.URL.Query().Get(":participantId"),
		BrokerInfo: &pb.BaseBrokerRequest{
			HostAddress: r.Host,
			Scheme:      r.URL.Scheme,
		},
	}
	resp, err := core.BrokerServiceAPI.GetParticipantVersions(r.Context(), request)
	if resp.GetResponse().GetCode() != pb.Response_SUCCESS {
		controller.WriteJson(http.StatusBadRequest, []byte(resp.Response.GetMessage()), w)
	} else {
		controller.WriteJsonObject(http.StatusOK, resp, w)
	}
}

//GetParticipantVersion returns all the participant versions.
func (this *BrokerService) GetParticipantVersion(w http.ResponseWriter, r *http.Request) {

}

//GetParticipantLatstVersion returns the latest ersion of a participant.
func (this *BrokerService) GetParticipantLastVersion(w http.ResponseWriter, r *http.Request) {
	request := &pb.GetParticipantsRequest{
		PartyReqType:  pb.GetParticipantsRequest_LATEST_PARTY,
		ParticipantId: r.URL.Query().Get(":participantId"),
		BrokerInfo: &pb.BaseBrokerRequest{
			HostAddress: r.Host,
			Scheme:      r.URL.Scheme,
		},
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
	if resp.GetResponse().GetCode() != pb.Response_SUCCESS {
		controller.WriteJson(http.StatusBadRequest, []byte(resp.Response.GetMessage()), w)
	} else {
		controller.WriteJsonObject(http.StatusOK, resp, w)
	}
}
