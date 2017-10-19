package v3

import (
	"net/http"

	"github.com/ServiceComb/service-center/pkg/rest"
	"github.com/ServiceComb/service-center/server/rest/controller"
)

//BrokerService handles requests and responses to the PactBroker
type BrokerService struct {
	//
}

func (this *BrokerService) URLPatterns() []rest.Route {
	return []rest.Route{
		{rest.HTTP_METHOD_GET, "/broker/v3/", this.GetBrokerHome},
		{rest.HTTP_METHOD_GET, "/broker/v3/participants", this.GetParticipants},
		{rest.HTTP_METHOD_GET, "/broker/v3/participants/:participantId", this.GetParticipant},
		{rest.HTTP_METHOD_GET, "/broker/v3/participants/:participantId/versions", this.GetParticipantVersions},
		{rest.HTTP_METHOD_GET, "/broker/v3/participants/:participantId/versions/latest", this.GetParticipantLatstVersion},
		{rest.HTTP_METHOD_GET, "/broker/v3/pacts/latest", this.GetPacts},
		{rest.HTTP_METHOD_GET, "/broker/v3/pacts/provider/:providerId/latest", this.GetProviderPacts},
		{rest.HTTP_METHOD_GET, "/broker/v3/pacts/provider/:providerId/consumer/:consumerId/latest", this.RetrievePact},
		{rest.HTTP_METHOD_GET, "/broker/v3/pacts/provider/:providerId/consumer/:consumerId/latest/:tag", this.RetrieveTaggedPact},
		{rest.HTTP_METHOD_GET, "/broker/v3/pacts/provider/:providerId/consumer/:consumerId/latest-untagged", this.RetrieveUntaggedPact},
		{rest.HTTP_METHOD_POST, "/broker/v3/pacts/provider/:providerId/consumer/:consumerId/version/:number", this.PublishPact},
	}
}

//GetBrokerHome returns home response
func (this *BrokerService) GetBrokerHome(w http.ResponseWriter, r *http.Request) {
	homeResponse := "{\r\n" +
		"  \"_links\": {\r\n" +
		"    \"pb:latest-provider-pacts\": {\r\n" +
		"      \"href\": \"http://localhost:30100/broker/v3/pacts/provider/{provider}/latest\",\r\n" +
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

}

//GetParticipant returns a participant  in the broker.
//Participant is a party that participates in a pact (ie. a Consumer or a Provider).
func (this *BrokerService) GetParticipant(w http.ResponseWriter, r *http.Request) {

}

//GetParticipantVersions returns all the participant versions.
func (this *BrokerService) GetParticipantVersions(w http.ResponseWriter, r *http.Request) {

}

//GetParticipantLatstVersion returns the latest ersion of a participant.
func (this *BrokerService) GetParticipantLatstVersion(w http.ResponseWriter, r *http.Request) {

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

}
