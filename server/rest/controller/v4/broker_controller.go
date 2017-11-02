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
package v4

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"service-center/util"

	"github.com/ServiceComb/service-center/pkg/rest"
	"github.com/ServiceComb/service-center/server/core"
	pb "github.com/ServiceComb/service-center/server/core/proto"
	"github.com/ServiceComb/service-center/server/rest/controller"
	uServiceUtil "github.com/ServiceComb/service-center/server/service/util"
)

type BrokerService struct {
	//
}

func (this *BrokerService) URLPatterns() []rest.Route {
	return []rest.Route{
		{rest.HTTP_METHOD_GET, uServiceUtil.BROKER_HOME_URL, this.GetBrokerHome},
		{rest.HTTP_METHOD_POST, uServiceUtil.BROKER_PUBLISH_URL, this.PublishPact},
	}
}

//GetBrokerHome implementation of the GetHome request
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

//PublishPact handles publication of a pact taking the pact as a json.
//A pact is published to the broker using a combination of the provider name, the
//consumer name, and the consumer application version.
func (this *BrokerService) PublishPact(w http.ResponseWriter, r *http.Request) {
	message, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.Logger().Error("body err", err)
		controller.WriteText(http.StatusInternalServerError, fmt.Sprintf("body error %s",
			err.Error()), w)
		return
	}

	request := &pb.PublishPactRequest{
		ProviderId: r.URL.Query().Get(":providerId"),
		ConsumerId: r.URL.Query().Get(":consumerId"),
		Version:    r.URL.Query().Get(":number"),
		Pact:       message,
	}

	resp, _ := core.BrokerServiceAPI.PublishPact(r.Context(), request)
	if resp.GetResponse().GetCode() != pb.Response_SUCCESS {
		controller.WriteJson(http.StatusBadRequest, []byte(resp.Response.GetMessage()), w)
	} else {
		controller.WriteJsonObject(http.StatusOK, resp, w)
	}
}
