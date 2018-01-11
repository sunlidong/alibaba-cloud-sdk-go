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

func (client *Client) DeleteScalingConfiguration(request *DeleteScalingConfigurationRequest) (response *DeleteScalingConfigurationResponse, err error) {
	response = CreateDeleteScalingConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteScalingConfigurationWithChan(request *DeleteScalingConfigurationRequest) (<-chan *DeleteScalingConfigurationResponse, <-chan error) {
	responseChan := make(chan *DeleteScalingConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteScalingConfiguration(request)
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

func (client *Client) DeleteScalingConfigurationWithCallback(request *DeleteScalingConfigurationRequest, callback func(response *DeleteScalingConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteScalingConfigurationResponse
		var err error
		defer close(result)
		response, err = client.DeleteScalingConfiguration(request)
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

type DeleteScalingConfigurationRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	ScalingConfigurationId string           `position:"Query" name:"ScalingConfigurationId"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
}

type DeleteScalingConfigurationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteScalingConfigurationRequest() (request *DeleteScalingConfigurationRequest) {
	request = &DeleteScalingConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DeleteScalingConfiguration", "ess", "openAPI")
	request.Method = requests.POST
	return
}

func CreateDeleteScalingConfigurationResponse() (response *DeleteScalingConfigurationResponse) {
	response = &DeleteScalingConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
