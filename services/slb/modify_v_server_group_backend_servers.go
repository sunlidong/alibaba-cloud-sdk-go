package slb

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

func (client *Client) ModifyVServerGroupBackendServers(request *ModifyVServerGroupBackendServersRequest) (response *ModifyVServerGroupBackendServersResponse, err error) {
	response = CreateModifyVServerGroupBackendServersResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyVServerGroupBackendServersWithChan(request *ModifyVServerGroupBackendServersRequest) (<-chan *ModifyVServerGroupBackendServersResponse, <-chan error) {
	responseChan := make(chan *ModifyVServerGroupBackendServersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVServerGroupBackendServers(request)
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

func (client *Client) ModifyVServerGroupBackendServersWithCallback(request *ModifyVServerGroupBackendServersRequest, callback func(response *ModifyVServerGroupBackendServersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVServerGroupBackendServersResponse
		var err error
		defer close(result)
		response, err = client.ModifyVServerGroupBackendServers(request)
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

type ModifyVServerGroupBackendServersRequest struct {
	*requests.RpcRequest
	NewBackendServers    string           `position:"Query" name:"NewBackendServers"`
	VServerGroupId       string           `position:"Query" name:"VServerGroupId"`
	Tags                 string           `position:"Query" name:"Tags"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	OldBackendServers    string           `position:"Query" name:"OldBackendServers"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type ModifyVServerGroupBackendServersResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	VServerGroupId string `json:"VServerGroupId" xml:"VServerGroupId"`
	BackendServers struct {
		BackendServer []struct {
			ServerId string `json:"ServerId" xml:"ServerId"`
			Port     int    `json:"Port" xml:"Port"`
			Weight   int    `json:"Weight" xml:"Weight"`
		} `json:"BackendServer" xml:"BackendServer"`
	} `json:"BackendServers" xml:"BackendServers"`
}

func CreateModifyVServerGroupBackendServersRequest() (request *ModifyVServerGroupBackendServersRequest) {
	request = &ModifyVServerGroupBackendServersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "ModifyVServerGroupBackendServers", "slb", "openAPI")
	request.Method = requests.POST
	return
}

func CreateModifyVServerGroupBackendServersResponse() (response *ModifyVServerGroupBackendServersResponse) {
	response = &ModifyVServerGroupBackendServersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
