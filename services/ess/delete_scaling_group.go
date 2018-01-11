package ess

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DeleteScalingGroup(request *DeleteScalingGroupRequest) (response *DeleteScalingGroupResponse, err error) {
	response = CreateDeleteScalingGroupResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteScalingGroupWithChan(request *DeleteScalingGroupRequest) (<-chan *DeleteScalingGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteScalingGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteScalingGroup(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DeleteScalingGroupWithCallback(request *DeleteScalingGroupRequest, callback func(response *DeleteScalingGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteScalingGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteScalingGroup(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DeleteScalingGroupRequest struct {
	*requests.RpcRequest
	ForceDelete          requests.Boolean `position:"Query" name:"ForceDelete"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
}

type DeleteScalingGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteScalingGroupRequest() (request *DeleteScalingGroupRequest) {
	request = &DeleteScalingGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DeleteScalingGroup", "ess", "openAPI")
	request.Method = requests.POST
	return
}

func CreateDeleteScalingGroupResponse() (response *DeleteScalingGroupResponse) {
	response = &DeleteScalingGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
