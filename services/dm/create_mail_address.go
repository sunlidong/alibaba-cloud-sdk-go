package dm

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

func (client *Client) CreateMailAddress(request *CreateMailAddressRequest) (response *CreateMailAddressResponse, err error) {
	response = CreateCreateMailAddressResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateMailAddressWithChan(request *CreateMailAddressRequest) (<-chan *CreateMailAddressResponse, <-chan error) {
	responseChan := make(chan *CreateMailAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMailAddress(request)
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

func (client *Client) CreateMailAddressWithCallback(request *CreateMailAddressRequest, callback func(response *CreateMailAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMailAddressResponse
		var err error
		defer close(result)
		response, err = client.CreateMailAddress(request)
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

type CreateMailAddressRequest struct {
	*requests.RpcRequest
	AccountName          string           `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ReplyAddress         string           `position:"Query" name:"ReplyAddress"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Sendtype             string           `position:"Query" name:"Sendtype"`
}

type CreateMailAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateCreateMailAddressRequest() (request *CreateMailAddressRequest) {
	request = &CreateMailAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "CreateMailAddress", "", "openAPI")
	request.Method = requests.POST
	return
}

func CreateCreateMailAddressResponse() (response *CreateMailAddressResponse) {
	response = &CreateMailAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
