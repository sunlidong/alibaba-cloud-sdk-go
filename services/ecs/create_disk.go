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

func (client *Client) CreateDisk(request *CreateDiskRequest) (response *CreateDiskResponse, err error) {
	response = CreateCreateDiskResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateDiskWithChan(request *CreateDiskRequest) (<-chan *CreateDiskResponse, <-chan error) {
	responseChan := make(chan *CreateDiskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDisk(request)
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

func (client *Client) CreateDiskWithCallback(request *CreateDiskRequest, callback func(response *CreateDiskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDiskResponse
		var err error
		defer close(result)
		response, err = client.CreateDisk(request)
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

type CreateDiskRequest struct {
	*requests.RpcRequest
	ZoneId               string           `position:"Query" name:"ZoneId"`
	Tag3Key              string           `position:"Query" name:"Tag.3.Key"`
	Tag5Value            string           `position:"Query" name:"Tag.5.Value"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	SnapshotId           string           `position:"Query" name:"SnapshotId"`
	Description          string           `position:"Query" name:"Description"`
	Tag1Key              string           `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string           `position:"Query" name:"Tag.1.Value"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Tag4Value            string           `position:"Query" name:"Tag.4.Value"`
	DiskName             string           `position:"Query" name:"DiskName"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tag5Key              string           `position:"Query" name:"Tag.5.Key"`
	DiskCategory         string           `position:"Query" name:"DiskCategory"`
	Size                 requests.Integer `position:"Query" name:"Size"`
	Encrypted            requests.Boolean `position:"Query" name:"Encrypted"`
	Tag2Key              string           `position:"Query" name:"Tag.2.Key"`
	Tag3Value            string           `position:"Query" name:"Tag.3.Value"`
	Tag4Key              string           `position:"Query" name:"Tag.4.Key"`
	Tag2Value            string           `position:"Query" name:"Tag.2.Value"`
}

type CreateDiskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	DiskId    string `json:"DiskId" xml:"DiskId"`
}

func CreateCreateDiskRequest() (request *CreateDiskRequest) {
	request = &CreateDiskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateDisk", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

func CreateCreateDiskResponse() (response *CreateDiskResponse) {
	response = &CreateDiskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
