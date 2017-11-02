//Copyright 2017 Huawei Technologies Co., Ltd
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
package util

import (
	"net/url"
	"strings"
	"time"

	"github.com/ServiceComb/service-center/pkg/cache"
	pb "github.com/ServiceComb/service-center/server/core/proto"
)

const (
	BROKER_HOME_URL                      = "/broker/"
	BROKER_PARTICIPANTS_URL              = "/broker/participants"
	BROKER_PARTICIPANT_URL               = "/broker/participants/:participantId"
	BROKER_PARTY_VERSIONS_URL            = "/broker/participants/:participantId/versions"
	BROKER_PARTY_LATEST_VERSION_URL      = "/broker/participants/:participantId/versions/latest"
	BROKER_PARTY_VERSION_URL             = "/broker/participants/:participantId/versions/:number"
	BROKER_PROVIDER_URL                  = "/broker/pacts/provider"
	BROKER_PROVIDER_LATEST_PACTS_URL     = "/broker/pacts/provider/:providerId/latest"
	BROKER_PROVIDER_LATEST_PACTS_TAG_URL = "/broker/pacts/provider/:providerId/latest/:tag"
	BROKER_PACTS_LATEST_URL              = "/broker/pacts/latest"

	BROKER_PUBLISH_URL  = "/broker/pacts/provider/:providerId/consumer/:consumerId/version/:number"
	BROKER_WEBHOOHS_URL = "/broker/webhooks"

	BROKER_CURIES_URL = "/broker/doc/:rel"
)

var brokerAPILinksValues = map[string]string{
	"self":                              BROKER_HOME_URL,
	"pb:publish-pact":                   BROKER_PUBLISH_URL,
	"pb:latest-pact-versions":           BROKER_PACTS_LATEST_URL,
	"pb:pacticipants":                   BROKER_PARTICIPANTS_URL,
	"pb:latest-provider-pacts":          BROKER_PROVIDER_LATEST_PACTS_URL,
	"pb:latest-provider-pacts-with-tag": BROKER_PROVIDER_LATEST_PACTS_TAG_URL,
	"pb:webhooks":                       BROKER_WEBHOOHS_URL,
}

var brokerAPILinksTempl = map[string]bool{
	"self":                              false,
	"pb:publish-pact":                   true,
	"pb:latest-pact-versions":           false,
	"pb:pacticipants":                   false,
	"pb:latest-provider-pacts":          true,
	"pb:latest-provider-pacts-with-tag": true,
	"pb:webhooks":                       false,
}

var brokerAPILinksTitles = map[string]string{
	"self":                              "Index",
	"pb:publish-pact":                   "Publish a pact",
	"pb:latest-pact-versions":           "Latest pact versions",
	"pb:pacticipants":                   "Pacticipants",
	"pb:latest-provider-pacts":          "Latest pacts by provider",
	"pb:latest-provider-pacts-with-tag": "Latest pacts by provider with a specified tag",
	"pb:webhooks":                       "Webhooks",
}

var brokerCache *cache.Cache

func init() {
	d, _ := time.ParseDuration("2m")
	brokerCache = cache.New(d, d)
}

//GenerateBrokerAPIPath creates the API link from the constant template
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

//GetBrokerHomeLinksAPIS return the generated Home links
func GetBrokerHomeLinksAPIS(scheme string, host string, apiKey string) string {
	return GenerateBrokerAPIPath(scheme, host, brokerAPILinksValues[apiKey],
		strings.NewReplacer(":providerId", "{provider}",
			":consumerId", "{consumer}",
			":number", "{consumerApplicationVersion}",
			":tag", "{tag}"))
}

//CreateBrokerHomeResponse create the templated broker home response
func CreateBrokerHomeResponse(host string, scheme string) *pb.BrokerHomeResponse {

	var apiEntries map[string]*pb.BrokerAPIInfoEntry
	apiEntries = make(map[string]*pb.BrokerAPIInfoEntry)

	for k := range brokerAPILinksValues {
		apiEntries[k] = &pb.BrokerAPIInfoEntry{
			Href:      GetBrokerHomeLinksAPIS(scheme, host, k),
			Title:     brokerAPILinksTitles[k],
			Templated: brokerAPILinksTempl[k],
		}
	}

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

//GetBrokerHomeResponse gets the homeResponse from cache if it exists
func GetBrokerHomeResponse(host string, scheme string) *pb.BrokerHomeResponse {
	brokerResponse, ok := brokerCache.Get(host + "://" + host)
	if !ok {
		brokerResp := CreateBrokerHomeResponse(host, scheme)
		if brokerResp == nil {
			return nil
		}
		d, _ := time.ParseDuration("10m")
		brokerCache.Set(host+"://"+host, brokerResp, d)
		return brokerResp
	}
	return brokerResponse.(*pb.BrokerHomeResponse)
}
