package alidns

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

func (client *Client) DeleteDomainRecord(request *DeleteDomainRecordRequest) (response *DeleteDomainRecordResponse, err error) {
	response = CreateDeleteDomainRecordResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteDomainRecordWithChan(request *DeleteDomainRecordRequest) (<-chan *DeleteDomainRecordResponse, <-chan error) {
	responseChan := make(chan *DeleteDomainRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDomainRecord(request)
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

func (client *Client) DeleteDomainRecordWithCallback(request *DeleteDomainRecordRequest, callback func(response *DeleteDomainRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDomainRecordResponse
		var err error
		defer close(result)
		response, err = client.DeleteDomainRecord(request)
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

type DeleteDomainRecordRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RecordId     string `position:"Query" name:"RecordId"`
}

type DeleteDomainRecordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	RecordId  string `json:"RecordId" xml:"RecordId"`
}

func CreateDeleteDomainRecordRequest() (request *DeleteDomainRecordRequest) {
	request = &DeleteDomainRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DeleteDomainRecord", "", "openAPI")
	request.Method = requests.POST
	return
}

func CreateDeleteDomainRecordResponse() (response *DeleteDomainRecordResponse) {
	response = &DeleteDomainRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
