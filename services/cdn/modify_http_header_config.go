package cdn

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

func (client *Client) ModifyHttpHeaderConfig(request *ModifyHttpHeaderConfigRequest) (response *ModifyHttpHeaderConfigResponse, err error) {
	response = CreateModifyHttpHeaderConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyHttpHeaderConfigWithChan(request *ModifyHttpHeaderConfigRequest) (<-chan *ModifyHttpHeaderConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyHttpHeaderConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyHttpHeaderConfig(request)
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

func (client *Client) ModifyHttpHeaderConfigWithCallback(request *ModifyHttpHeaderConfigRequest, callback func(response *ModifyHttpHeaderConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyHttpHeaderConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyHttpHeaderConfig(request)
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

type ModifyHttpHeaderConfigRequest struct {
	*requests.RpcRequest
	HeaderKey     string           `position:"Query" name:"HeaderKey"`
	HeaderValue   string           `position:"Query" name:"HeaderValue"`
	DomainName    string           `position:"Query" name:"DomainName"`
	ConfigID      string           `position:"Query" name:"ConfigID"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type ModifyHttpHeaderConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyHttpHeaderConfigRequest() (request *ModifyHttpHeaderConfigRequest) {
	request = &ModifyHttpHeaderConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "ModifyHttpHeaderConfig", "", "openAPI")
	request.Method = requests.POST
	return
}

func CreateModifyHttpHeaderConfigResponse() (response *ModifyHttpHeaderConfigResponse) {
	response = &ModifyHttpHeaderConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
