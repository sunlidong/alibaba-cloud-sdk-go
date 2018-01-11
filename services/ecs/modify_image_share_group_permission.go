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

func (client *Client) ModifyImageShareGroupPermission(request *ModifyImageShareGroupPermissionRequest) (response *ModifyImageShareGroupPermissionResponse, err error) {
	response = CreateModifyImageShareGroupPermissionResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyImageShareGroupPermissionWithChan(request *ModifyImageShareGroupPermissionRequest) (<-chan *ModifyImageShareGroupPermissionResponse, <-chan error) {
	responseChan := make(chan *ModifyImageShareGroupPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyImageShareGroupPermission(request)
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

func (client *Client) ModifyImageShareGroupPermissionWithCallback(request *ModifyImageShareGroupPermissionRequest, callback func(response *ModifyImageShareGroupPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyImageShareGroupPermissionResponse
		var err error
		defer close(result)
		response, err = client.ModifyImageShareGroupPermission(request)
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

type ModifyImageShareGroupPermissionRequest struct {
	*requests.RpcRequest
	RemoveGroup1         string           `position:"Query" name:"RemoveGroup.1"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AddGroup1            string           `position:"Query" name:"AddGroup.1"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ImageId              string           `position:"Query" name:"ImageId"`
}

type ModifyImageShareGroupPermissionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyImageShareGroupPermissionRequest() (request *ModifyImageShareGroupPermissionRequest) {
	request = &ModifyImageShareGroupPermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyImageShareGroupPermission", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

func CreateModifyImageShareGroupPermissionResponse() (response *ModifyImageShareGroupPermissionResponse) {
	response = &ModifyImageShareGroupPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
