package rds

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

func (client *Client) RemoveTagsFromResource(request *RemoveTagsFromResourceRequest) (response *RemoveTagsFromResourceResponse, err error) {
	response = CreateRemoveTagsFromResourceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RemoveTagsFromResourceWithChan(request *RemoveTagsFromResourceRequest) (<-chan *RemoveTagsFromResourceResponse, <-chan error) {
	responseChan := make(chan *RemoveTagsFromResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveTagsFromResource(request)
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

func (client *Client) RemoveTagsFromResourceWithCallback(request *RemoveTagsFromResourceRequest, callback func(response *RemoveTagsFromResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveTagsFromResourceResponse
		var err error
		defer close(result)
		response, err = client.RemoveTagsFromResource(request)
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

type RemoveTagsFromResourceRequest struct {
	*requests.RpcRequest
	Tags                 string           `position:"Query" name:"Tags"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ProxyId              string           `position:"Query" name:"proxyId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tag5Key              string           `position:"Query" name:"Tag.5.key"`
	Tag5Value            string           `position:"Query" name:"Tag.5.value"`
	Tag3Key              string           `position:"Query" name:"Tag.3.key"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Tag1Key              string           `position:"Query" name:"Tag.1.key"`
	Tag2Key              string           `position:"Query" name:"Tag.2.key"`
	Tag1Value            string           `position:"Query" name:"Tag.1.value"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Tag4Value            string           `position:"Query" name:"Tag.4.value"`
	Tag3Value            string           `position:"Query" name:"Tag.3.value"`
	Tag2Value            string           `position:"Query" name:"Tag.2.value"`
	Tag4Key              string           `position:"Query" name:"Tag.4.key"`
}

type RemoveTagsFromResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateRemoveTagsFromResourceRequest() (request *RemoveTagsFromResourceRequest) {
	request = &RemoveTagsFromResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "RemoveTagsFromResource", "rds", "openAPI")
	request.Method = requests.POST
	return
}

func CreateRemoveTagsFromResourceResponse() (response *RemoveTagsFromResourceResponse) {
	response = &RemoveTagsFromResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
