package util_test

import (
	"fmt"
	"testing"

	"github.com/ServiceComb/service-center/server/core/registry"
	serviceUtil "github.com/ServiceComb/service-center/server/service/util"
	"golang.org/x/net/context"
)

func TestBrokerGetParticipantUtils(t *testing.T) {
	participant1, err := serviceUtil.GetBrokerParticipantUtils(context.Background(),
		"", "", "", registry.WithCacheOnly())
	if err != nil || participant1 != nil {
		fmt.Printf("GetBrokerParticipantUtilsfor for non-existing WithCacheOnly failed: " + err.Error())
		t.FailNow()
	}

	// participant2, err := serviceUtil.GetBrokerParticipantUtils(context.Background(),
	// 	"", "", "", registry.WithNoCache())
	// if err != nil || participant2 != nil {
	// 	fmt.Printf("GetBrokerParticipantUtilsfor for non-existing WithNoCache failed: " + err.Error())
	// 	t.FailNow()
	// }

}

func TestBrokerGetPartyVersionUtils(t *testing.T) {
}

func TestBrokerGetPactUtils(t *testing.T) {
}

func TestBrokerGetPactPubUtils(t *testing.T) {
	var content []byte
	copy(content[:], "pact..content")
	pact, err := serviceUtil.GetBrokerPact(context.Background(), "", 0, 0,
		content, registry.WithCacheOnly())
	if err != nil || pact != nil {
		fmt.Printf("TestBrokerGetPactPubUtils for non-existing WithCacheOnly failed: " + err.Error())
		t.FailNow()
	}
}

func TestBrokerGetLastVersionOrder(t *testing.T) {
	order, err := serviceUtil.GetBrokerParticipantVersionOrder(context.Background(),
		"", 0, registry.WithCacheOnly())
	if err != nil || order != int32(0) {
		fmt.Printf("TestBrokerGetLastVersionOrdert for non-existing WithCacheOnly failed: " + err.Error())
		t.FailNow()
	}
}
