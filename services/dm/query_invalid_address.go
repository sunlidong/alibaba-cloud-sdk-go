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

func (client *Client) QueryInvalidAddress(request *QueryInvalidAddressRequest) (response *QueryInvalidAddressResponse, err error) {
	response = CreateQueryInvalidAddressResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryInvalidAddressWithChan(request *QueryInvalidAddressRequest) (<-chan *QueryInvalidAddressResponse, <-chan error) {
	responseChan := make(chan *QueryInvalidAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryInvalidAddress(request)
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

func (client *Client) QueryInvalidAddressWithCallback(request *QueryInvalidAddressRequest, callback func(response *QueryInvalidAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryInvalidAddressResponse
		var err error
		defer close(result)
		response, err = client.QueryInvalidAddress(request)
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

type QueryInvalidAddressRequest struct {
	*requests.RpcRequest
	EndTime              string           `position:"Query" name:"EndTime"`
	NextStart            string           `position:"Query" name:"NextStart"`
	StartTime            string           `position:"Query" name:"StartTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	KeyWord              string           `position:"Query" name:"KeyWord"`
	Length               requests.Integer `position:"Query" name:"Length"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type QueryInvalidAddressResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	NextStart  int    `json:"NextStart" xml:"NextStart"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	Data       struct {
		MailDetail []struct {
			LastUpdateTime    string `json:"LastUpdateTime" xml:"LastUpdateTime"`
			UtcLastUpdateTime int    `json:"UtcLastUpdateTime" xml:"UtcLastUpdateTime"`
			ToAddress         string `json:"ToAddress" xml:"ToAddress"`
		} `json:"mailDetail" xml:"mailDetail"`
	} `json:"data" xml:"data"`
}

func CreateQueryInvalidAddressRequest() (request *QueryInvalidAddressRequest) {
	request = &QueryInvalidAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "QueryInvalidAddress", "", "openAPI")
	request.Method = requests.POST
	return
}

func CreateQueryInvalidAddressResponse() (response *QueryInvalidAddressResponse) {
	response = &QueryInvalidAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
