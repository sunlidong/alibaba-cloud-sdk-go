package ecs

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

func (client *Client) TerminatePhysicalConnection(request *TerminatePhysicalConnectionRequest) (response *TerminatePhysicalConnectionResponse, err error) {
	response = CreateTerminatePhysicalConnectionResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) TerminatePhysicalConnectionWithChan(request *TerminatePhysicalConnectionRequest) (<-chan *TerminatePhysicalConnectionResponse, <-chan error) {
	responseChan := make(chan *TerminatePhysicalConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TerminatePhysicalConnection(request)
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

func (client *Client) TerminatePhysicalConnectionWithCallback(request *TerminatePhysicalConnectionRequest, callback func(response *TerminatePhysicalConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TerminatePhysicalConnectionResponse
		var err error
		defer close(result)
		response, err = client.TerminatePhysicalConnection(request)
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

type TerminatePhysicalConnectionRequest struct {
	*requests.RpcRequest
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	UserCidr             string           `position:"Query" name:"UserCidr"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PhysicalConnectionId string           `position:"Query" name:"PhysicalConnectionId"`
}

type TerminatePhysicalConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateTerminatePhysicalConnectionRequest() (request *TerminatePhysicalConnectionRequest) {
	request = &TerminatePhysicalConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "TerminatePhysicalConnection", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

func CreateTerminatePhysicalConnectionResponse() (response *TerminatePhysicalConnectionResponse) {
	response = &TerminatePhysicalConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
