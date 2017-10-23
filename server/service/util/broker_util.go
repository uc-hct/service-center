package util

import (
	"context"
	"encoding/json"
	"math"

	"github.com/ServiceComb/service-center/pkg/util"
	apt "github.com/ServiceComb/service-center/server/core"
	pb "github.com/ServiceComb/service-center/server/core/proto"
	"github.com/ServiceComb/service-center/server/core/registry"
	"github.com/ServiceComb/service-center/server/core/registry/store"
)

var participantIds int32 = 0
var participantVersionIds int32 = 0
var pactIds int32 = 0
var pactPubIds int32 = 0

func GetBrokerParticipantUtils(ctx context.Context, tenant string, appId string,
	serviceName string, opts ...registry.PluginOpOption) (*pb.Participant, error) {
	key := apt.GenerateBrokerParticipantKey(tenant, appId, serviceName)
	opts = append(opts, registry.WithStrKey(key))
	participants, err := store.Store().BrokerParticipant().Search(ctx, opts...)

	if err != nil {
		return nil, err
	}

	if len(participants.Kvs) == 0 {
		util.Logger().Info("GetParticipant found no participant")
		return nil, nil
	}

	participant := &pb.Participant{}
	err = json.Unmarshal(participants.Kvs[0].Value, participant)
	if err != nil {
		return nil, err
	}
	util.Logger().Infof("GetParticipant: (%d, %s, %s)", participant.Id, participant.AppId,
		participant.ServiceName)
	return participant, nil
}

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
