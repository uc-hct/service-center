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
package service

import (
	"service-center/util"

	pb "github.com/ServiceComb/service-center/server/core/proto"
	uServiceUtil "github.com/ServiceComb/service-center/server/service/util"
	"golang.org/x/net/context"
)

type BrokerController struct {
}

func (s *BrokerController) GetBrokerHome(ctx context.Context, in *pb.BaseBrokerRequest) (*pb.BrokerHomeResponse, error) {
	if in == nil || len(in.HostAddress) == 0 {
		util.Logger().Errorf(nil, "Get Participant versions request failed: invalid params.")
		return &pb.BrokerHomeResponse{
			Response: pb.CreateResponse(pb.Response_FAIL, "Request format invalid."),
		}, nil
	}

	return uServiceUtil.GetBrokerHomeResponse(in.HostAddress, in.Scheme), nil
}

func (s *BrokerController) PublishPact(ctx context.Context, in *pb.PublishPactRequest) (*pb.PublishPactResponse, error) {
	return nil, nil
}
