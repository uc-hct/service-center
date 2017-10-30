package service

import (
	"crypto/sha1"
	"errors"
	_ "fmt"
	"strings"

	"github.com/ServiceComb/service-center/pkg/util"
	apt "github.com/ServiceComb/service-center/server/core"
	pb "github.com/ServiceComb/service-center/server/core/proto"
	uServiceUtil "github.com/ServiceComb/service-center/server/service/util"
	"golang.org/x/net/context"
)

type BrokerController struct {
}

func (s *BrokerController) GetBrokerHome(ctx context.Context,
	in *pb.BaseBrokerRequest) (*pb.BrokerHomeResponse, error) {

	if in == nil || len(in.HostAddress) == 0 {
		util.Logger().Errorf(nil, "Get Participant versions request failed: invalid params.")
		return &pb.BrokerHomeResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Request format invalid."),
		}, nil
	}

	return uServiceUtil.CreateBrokerHomeResponse(in.HostAddress, in.Scheme), nil
}

func (s *BrokerController) GetParticipantVersions(ctx context.Context,
	in *pb.GetParticipantsRequest) (*pb.GetParticipantVersionsResponse, error) {

	if in == nil || in.PartyReqType == pb.GetParticipantsRequest_ALL_PARTIES || len(in.ParticipantId) == 0 {
		util.Logger().Errorf(nil, "Get Participant versions request failed: invalid params.")
		return &pb.GetParticipantVersionsResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Request format invalid."),
		}, nil
	}
	participant, microservice, errBroker, errService :=
		uServiceUtil.GetBrokerPartyFromServiceId(ctx, in.ParticipantId)

	if errService != nil {
		util.Logger().Errorf(nil,
			"get participant failed, microService cannot be searched %s.", in.ParticipantId)
		return &pb.GetParticipantVersionsResponse{
			Response: pb.CreateResponse(pb.Response_FAIL,
				"get participant version, Microservice cannot be searched."),
		}, errService
	}
	if errBroker != nil || participant == nil {
		util.Logger().Errorf(nil,
			"participant does not exist for the serviceId (%s).", in.ParticipantId)
		return &pb.GetParticipantVersionsResponse{
			Response: pb.CreateResponse(pb.Response_FAIL,
				"Particpant does not exist."),
		}, errBroker
	}
	tenant := util.ParseTenantProject(ctx)
	//serviceVersions, err := getServiceAllVersions(ctx, "tets", "test", "lol")
	partyVersions, err := uServiceUtil.GetAllBrokerPartyVersions(ctx, tenant,
		participant.Id)
	if err != nil {
		return &pb.GetParticipantVersionsResponse{
			Response: pb.CreateResponse(pb.Response_FAIL,
				"error finding participant versions"),
		}, err
	}
	//check if we are looking for latest version or not
	return uServiceUtil.CreatePartyVersionsResponse(participant, partyVersions,
		in.BrokerInfo.HostAddress, in.BrokerInfo.Scheme, microservice.GetServiceId()), nil

}

func (s *BrokerController) GetParticipants(ctx context.Context,
	in *pb.GetParticipantsRequest) (*pb.GetParticipantsResponse, error) {

	if in == nil || in.PartyReqType != pb.GetParticipantsRequest_ALL_PARTIES {
		util.Logger().Errorf(nil, "Get Participants request failed: invalid params.")
		return &pb.GetParticipantsResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Request format invalid."),
		}, nil
	}
	tenant := util.ParseTenantProject(ctx)
	dataParticpants, err := uServiceUtil.GetBrokerParticipants(ctx, tenant)
	if err != nil {
		util.Logger().Errorf(err, "could not get the particpants")
		return &pb.GetParticipantsResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, err.Error()),
		}, nil
	}

	participantsResp := uServiceUtil.CreateParticipantsResponse(dataParticpants,
		in.BrokerInfo.HostAddress, in.BrokerInfo.Scheme)

	return participantsResp, nil
}

func (s *BrokerController) GetParticipant(ctx context.Context,
	in *pb.GetParticipantsRequest) (*pb.GetParticipantResponse, error) {

	if in == nil || in.PartyReqType == pb.GetParticipantsRequest_ALL_PARTIES || len(in.ParticipantId) == 0 {

		util.Logger().Errorf(nil, "Get Participant request failed: invalid params.")
		return &pb.GetParticipantResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Request format invalid."),
		}, nil
	}

	participant, _, errBroker, errService :=
		uServiceUtil.GetBrokerPartyFromServiceId(ctx, in.ParticipantId)

	if errService != nil {
		util.Logger().Errorf(nil,
			"get participant failed, microService cannot be searched %s.", in.ParticipantId)
		return &pb.GetParticipantResponse{
			Response: pb.CreateResponse(pb.Response_FAIL,
				"get participant, Microservice cannot be searched."),
		}, errService
	}
	if errBroker != nil || participant == nil {
		util.Logger().Errorf(nil,
			"participant does not exist for the serviceId (%s).", in.ParticipantId)
		return &pb.GetParticipantResponse{
			Response: pb.CreateResponse(pb.Response_FAIL,
				"Particpant does not exist."),
		}, errBroker
	}

	return &pb.GetParticipantResponse{
		Response: pb.CreateResponse(pb.Response_SUCCESS, "Get Participant successfully."),
		ParticipantInfo: uServiceUtil.CreateParticipantResponse(participant,
			in.BrokerInfo.HostAddress, in.BrokerInfo.Scheme, in.ParticipantId),
	}, nil
}

func (s *BrokerController) PublishPact(ctx context.Context,
	in *pb.PublishPactRequest) (*pb.PublishPactResponse, error) {

	if in == nil || len(in.ProviderId) == 0 || len(in.ConsumerId) == 0 || len(in.Version) == 0 || len(in.Pact) == 0 {
		util.Logger().Errorf(nil, "pact publish request failed: invalid params.")
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Request format invalid."),
		}, nil
	}

	//TODO: add validator here
	err := apt.Validate(in)

	if err != nil {
		util.Logger().Errorf(err,
			"publish pact failed, providerId %s, consumerId %s: invalid parameters.",
			in.GetProviderId(), in.GetConsumerId())
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, err.Error()),
		}, nil
	}

	consumerParticipant, consumer, _, errConsumerService :=
		uServiceUtil.GetBrokerPartyFromServiceId(ctx, in.ConsumerId)

	if errConsumerService != nil {
		util.Logger().Errorf(errConsumerService,
			"pact publish failed, providerId is %s: query provider failed.", in.ConsumerId)
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Query consumer failed."),
		}, errConsumerService
	}

	if consumer == nil {
		util.Logger().Errorf(nil,
			"pact publish failed, consumerId is %s: consumer not exist.", in.ConsumerId)
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Consumer does not exist."),
		}, errors.New("consumer does not exist.")
	}

	// check that the consumer has that vesion in the url
	if strings.Compare(consumer.GetVersion(), in.Version) != 0 {
		util.Logger().Errorf(nil,
			"pact publish failed, version (%s) does not exist for consmer", in.Version)
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Version does not exist."),
		}, nil
	}

	providerParticipant, provider, _, errProviderService :=
		uServiceUtil.GetBrokerPartyFromServiceId(ctx, in.ProviderId)

	if errProviderService != nil {
		util.Logger().Errorf(errProviderService,
			"pact publish failed, providerId is %s: query provider failed.", in.ProviderId)
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Query provider failed."),
		}, errProviderService
	}

	if provider == nil {
		util.Logger().Errorf(nil,
			"pact publish failed, providerId is %s: provider not exist.", in.ProviderId)
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Provider does not exist."),
		}, errors.New("provider does not exist.")
	}

	tenant := util.ParseTenantProject(ctx)
	//create provider particpant if it does not exist
	if providerParticipant == nil {
		providerParticipant, err = uServiceUtil.AddBrokerParticipantIntoETCD(ctx, tenant,
			provider.AppId, provider.ServiceName)
		if err != nil {
			return &pb.PublishPactResponse{
				Response: pb.CreateResponse(pb.Response_FAIL, "could not create provider."),
			}, err
		}
	}
	//create consumer particpant if it does not exist
	if consumerParticipant == nil {
		consumerParticipant, err = uServiceUtil.AddBrokerParticipantIntoETCD(ctx, tenant,
			consumer.AppId, consumer.ServiceName)
		if err != nil {
			return &pb.PublishPactResponse{
				Response: pb.CreateResponse(pb.Response_FAIL, "could not create consumer."),
			}, err
		}
	}

	// Get or create version
	consumerVersion, err := uServiceUtil.GetBrokerParticipantVersion(ctx, tenant,
		consumerParticipant.Id, in.Version)
	if err != nil {
		util.Logger().Errorf(nil, "consumerVersion could not be retrieved for cosnumer ", in.ConsumerId)
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Consumer version cannot be searched."),
		}, err
	}

	if consumerVersion == nil {
		consumerVersion, err = uServiceUtil.AddBrokerParticipantVersionIntoETCD(ctx, tenant,
			consumerParticipant.Id, in.Version)
		if err != nil {
			return &pb.PublishPactResponse{
				Response: pb.CreateResponse(pb.Response_FAIL, "could not create consumer Version."),
			}, err
		}
	}

	//get or create pact and pactPublication
	sha1 := sha1.Sum(in.Pact)
	var sha []byte = sha1[:]
	pactEntry, err := uServiceUtil.GetBrokerPact(ctx, tenant, consumerParticipant.Id,
		providerParticipant.Id, sha)

	if err != nil {
		util.Logger().Errorf(nil, "pact publish failed, Pact cannot be searched.")
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Pact cannot be searched."),
		}, err
	}
	//pactPubEntry := &pb.PactPublication{}
	if pactEntry == nil {
		pactEntry, err = uServiceUtil.AddBrokerPactIntoETCD(ctx, tenant,
			consumerParticipant.Id, providerParticipant.Id, sha, in.Pact)
		if err != nil {
			return &pb.PublishPactResponse{
				Response: pb.CreateResponse(pb.Response_FAIL, "could not create Pact."),
			}, err
		}
		// crate pact publicaion
		_, err := uServiceUtil.AddBrokerPctPublication(ctx, tenant,
			consumerParticipant.Id, providerParticipant.Id, pactEntry.Id)
		if err != nil {
			util.Logger().Errorf(err,
				"pact publish failed, could not create pact Publications")
			return &pb.PublishPactResponse{
				Response: pb.CreateResponse(pb.Response_FAIL, "publication Error."),
			}, err
		}
	}

	return &pb.PublishPactResponse{
		Response: pb.CreateResponse(pb.Response_SUCCESS, "Pact published successfully."),
	}, nil
}
