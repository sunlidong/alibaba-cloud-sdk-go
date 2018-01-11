package cs

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

func (client *Client) AttachInstances(request *AttachInstancesRequest) (response *AttachInstancesResponse, err error) {
	response = CreateAttachInstancesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AttachInstancesWithChan(request *AttachInstancesRequest) (<-chan *AttachInstancesResponse, <-chan error) {
	responseChan := make(chan *AttachInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachInstances(request)
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

func (client *Client) AttachInstancesWithCallback(request *AttachInstancesRequest, callback func(response *AttachInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachInstancesResponse
		var err error
		defer close(result)
		response, err = client.AttachInstances(request)
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

type AttachInstancesRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

type AttachInstancesResponse struct {
	*responses.BaseResponse
}

func CreateAttachInstancesRequest() (request *AttachInstancesRequest) {
	request = &AttachInstancesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "AttachInstances", "/clusters/[ClusterId]/attach", "", "openAPI")
	request.Method = requests.POST
	return
}

func CreateAttachInstancesResponse() (response *AttachInstancesResponse) {
	response = &AttachInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
