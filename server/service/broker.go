package service

import (
	"crypto/sha1"
	_ "fmt"

	"github.com/ServiceComb/service-center/pkg/util"
	apt "github.com/ServiceComb/service-center/server/core"
	pb "github.com/ServiceComb/service-center/server/core/proto"
	uServiceUtil "github.com/ServiceComb/service-center/server/service/util"
	"golang.org/x/net/context"
)

type BrokerController struct {
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
	tenant := util.ParseTenantProject(ctx)

	provider, err := uServiceUtil.GetService(ctx, tenant, in.ProviderId)
	if err != nil {
		util.Logger().Errorf(err,
			"pact publish failed, providerId is %s: query provider failed.", in.ProviderId)
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Query provider failed."),
		}, err
	}

	if provider == nil {
		util.Logger().Errorf(nil,
			"pact publish failed, providerId is %s: provider not exist.", in.ProviderId)
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Provider does not exist."),
		}, nil
	}

	consumer, err := uServiceUtil.GetService(ctx, tenant, in.ConsumerId)

	if err != nil {
		util.Logger().Errorf(err,
			"pact publish failed, consumerId is %s: query consumer failed.", in.ProviderId)
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Query provider failed."),
		}, err
	}

	if consumer == nil {
		util.Logger().Errorf(nil,
			"pact publish failed, consumerId is %s: consumer not exist.", in.ProviderId)
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "consumer does not exist."),
		}, nil
	}

	// Get or create provider participant
	providerParticipant, err := uServiceUtil.GetBrokerParticipantUtils(ctx, tenant,
		provider.AppId, provider.ServiceName)

	if err != nil {
		util.Logger().Errorf(nil,
			"pact publish failed, provider participant cannot be searched.", in.ProviderId)
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL,
				"Provider participant cannot be searched."),
		}, err
	}
	if providerParticipant == nil {
		providerParticipant, err := uServiceUtil.AddBrokerParticipantIntoETCD(ctx, tenant,
			provider.AppId, provider.ServiceName)
		if err != nil {
			return &pb.PublishPactResponse{
				Response: pb.CreateResponse(pb.Response_FAIL, "could not create provider."),
			}, err
		}
	}

	// Get or create provider participant
	consumerParticipant, err := uServiceUtil.GetBrokerParticipantUtils(ctx, tenant,
		consumer.AppId, consumer.ServiceName)

	if err != nil {
		util.Logger().Errorf(nil,
			"pact publish failed, consumer participant cannot be searched.", in.ProviderId)
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL,
				"Consumer participant cannot be searched."),
		}, err
	}
	if consumerParticipant == nil {
		consumerParticipant, err := uServiceUtil.AddBrokerParticipantIntoETCD(ctx, tenant,
			consumer.AppId, consumer.ServiceName)
		if err != nil {
			return &pb.PublishPactResponse{
				Response: pb.CreateResponse(pb.Response_FAIL, "could not create provider."),
			}, err
		}
	}

	// Get or create version
	// TODO: should we validate that the version is already in the service center?
	consumerVersion, err := uServiceUtil.GetBrokerParticipantVersion(ctx, tenant,
		consumerParticipant.Id, in.Version)
	if err != nil {
		util.Logger().Errorf(nil, "consumerVersion could not be retrieved", in.ProviderId)
		return &pb.PublishPactResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Consumer version cannot be searched."),
		}, err
	}

	if consumerVersion == nil {
		consumerVersion, err := uServiceUtil.AddBrokerParticipantVersionIntoETCD(ctx, tenant,
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

	pactPubEntry := &pb.PactPublication{}
	if pactEntry == nil {
		pactEntry, err := uServiceUtil.AddBrokerPactIntoETCD(ctx, tenant,
			consumerParticipant.Id, providerParticipant.Id, sha, in.Pact)
		if err != nil {
			return &pb.PublishPactResponse{
				Response: pb.CreateResponse(pb.Response_FAIL, "could not create Pact."),
			}, err
		}
		// crate pact publicaion
		pactPubEntry, err := uServiceUtil.AddBrokerPctPublication(ctx, tenant,
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
