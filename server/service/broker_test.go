package service_test

import (
	"fmt"

	pb "github.com/ServiceComb/service-center/server/core/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	TEST_BROKER_NO_SERVICE_ID      = "noServiceId"
	TEST_BROKER_NO_VERSION         = "noVersion"
	TEST_BROKER_TOO_LONG_SERVICEID = "addasdfasaddasdfasaddasdfasaddasdfasaddasdfasaddasdfasaddasdfasadafd"
	//Consumer
	TEST_BROKER_CONSUMER_VERSION = "4.0.0"
	TEST_BROKER_CONSUMER_NAME    = "broker_name_consumer"
	TEST_BROKER_CONSUMER_APP     = "broker_group_consumer"
	//Provider
	TEST_BROKER_PROVIDER_VERSION = "3.0.0"
	TEST_BROKER_PROVIDER_NAME    = "broker_name_provider"
	TEST_BROKER_PROVIDER_APP     = "broker_group_provider"
)

var consumerServiceId string
var providerServiceId string

var _ = Describe("BrokerController", func() {
	Describe("brokerDependency", func() {
		Context("normal", func() {
			It("PublishPact", func() {
				fmt.Println("UT===========PublishPact")

				//(1) create consumer service
				resp, err := serviceResource.Create(getContext(), &pb.CreateServiceRequest{
					Service: &pb.MicroService{
						ServiceName: TEST_BROKER_CONSUMER_NAME,
						AppId:       TEST_BROKER_CONSUMER_APP,
						Version:     TEST_BROKER_CONSUMER_VERSION,
						Level:       "FRONT",
						Schemas: []string{
							"xxxxxxxx",
						},
						Status: "UP",
					},
				})

				Expect(err).To(BeNil())
				consumerServiceId = resp.ServiceId
				Expect(resp.GetResponse().Code).To(Equal(pb.Response_SUCCESS))

				//(2) create provider service
				resp, err = serviceResource.Create(getContext(), &pb.CreateServiceRequest{
					Service: &pb.MicroService{
						ServiceName: TEST_BROKER_PROVIDER_NAME,
						AppId:       TEST_BROKER_PROVIDER_APP,
						Version:     TEST_BROKER_PROVIDER_VERSION,
						Level:       "FRONT",
						Schemas: []string{
							"xxxxxxxx",
						},
						Status:     "UP",
						Properties: map[string]string{"allowCrossApp": "true"},
					},
				})
				Expect(err).To(BeNil())
				providerServiceId = resp.ServiceId
				Expect(resp.GetResponse().Code).To(Equal(pb.Response_SUCCESS))

				//(3) publish a pact between two services
				respPublishPact, err := brokerResource.PublishPact(getContext(), &pb.PublishPactRequest{
					ProviderId: providerServiceId,
					ConsumerId: consumerServiceId,
					Version:    TEST_BROKER_CONSUMER_VERSION,
					Pact:       []byte("hello"),
				})

				Expect(err).To(BeNil())
				Expect(respPublishPact.GetResponse().Code).To(Equal(pb.Response_SUCCESS))
			})

			It("PublishPact-noProviderServiceId", func() {
				fmt.Println("UT===========PublishPact, no provider serviceID")

				//(3) publish a pact between two services
				respPublishPact, _ := brokerResource.PublishPact(getContext(), &pb.PublishPactRequest{
					ProviderId: TEST_BROKER_NO_SERVICE_ID,
					ConsumerId: consumerServiceId,
					Version:    TEST_BROKER_CONSUMER_VERSION,
					Pact:       []byte("hello"),
				})

				Expect(respPublishPact.GetResponse().Code).To(Equal(pb.Response_FAIL))
			})

			It("PublishPact-noConumerServiceId", func() {
				fmt.Println("UT===========PublishPact, no consumer serviceID")

				//(3) publish a pact between two services
				respPublishPact, _ := brokerResource.PublishPact(getContext(), &pb.PublishPactRequest{
					ProviderId: providerServiceId,
					ConsumerId: TEST_BROKER_NO_SERVICE_ID,
					Version:    TEST_BROKER_CONSUMER_VERSION,
					Pact:       []byte("hello"),
				})

				Expect(respPublishPact.GetResponse().Code).To(Equal(pb.Response_FAIL))
			})

			It("PublishPact-noConumerVersion", func() {
				fmt.Println("UT===========PublishPact, no consumer Version")

				//(3) publish a pact between two services
				respPublishPact, _ := brokerResource.PublishPact(getContext(), &pb.PublishPactRequest{
					ProviderId: providerServiceId,
					ConsumerId: consumerServiceId,
					Version:    TEST_BROKER_NO_VERSION,
					Pact:       []byte("hello"),
				})

				Expect(respPublishPact.GetResponse().Code).To(Equal(pb.Response_FAIL))
			})

			It("PublishPact-wrong-formats", func() {
				fmt.Println("UT===========PublishPact, no consumer Version")

				//(3) publish a pact between two services
				respPublishPact, _ := brokerResource.PublishPact(getContext(), &pb.PublishPactRequest{
					ProviderId: providerServiceId,
					ConsumerId: consumerServiceId,
					Version:    TEST_BROKER_NO_VERSION,
					Pact:       []byte("hello"),
				})

				Expect(respPublishPact.GetResponse().Code).To(Equal(pb.Response_FAIL))
			})

			It("GetParticipant", func() {
				fmt.Println("UT===========GetParticipant")

				//(1) get Participant of the consumer
				respGetConsumerParty, _ := brokerResource.GetParticipant(getContext(),
					&pb.GetParticipantsRequest{
						PartyReqType:  pb.GetParticipantsRequest_SINGLE_PARTY,
						ParticipantId: consumerServiceId,
						UrlPrefix:     "http://localhost/",
					})
				Expect(respGetConsumerParty.GetResponse().Code).To(Equal(pb.Response_SUCCESS))

				//(2) get Participant of the provider
				respGetProviderParty, _ := brokerResource.GetParticipant(getContext(),
					&pb.GetParticipantsRequest{
						PartyReqType:  pb.GetParticipantsRequest_SINGLE_PARTY,
						ParticipantId: providerServiceId,
						UrlPrefix:     "http://localhost/",
					})

				Expect(respGetProviderParty.GetResponse().Code).To(Equal(pb.Response_SUCCESS))
				//fmt.Println("UT===========GetParticipant----->," + respGetParty.ParticipantInfo.GetName())
			})

			It("GetParticipant-noServiceId", func() {
				fmt.Println("UT===========GetParticipant, non-exising serviceId")

				//(1) get Participant of the non-existingId
				respGetParty, _ := brokerResource.GetParticipant(getContext(), &pb.GetParticipantsRequest{
					PartyReqType:  pb.GetParticipantsRequest_SINGLE_PARTY,
					ParticipantId: TEST_BROKER_NO_SERVICE_ID,
					UrlPrefix:     "http://localhost/",
				})
				Expect(respGetParty.GetResponse().Code).To(Equal(pb.Response_FAIL))

			})
		})
	})
})
