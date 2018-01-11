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

func (client *Client) CopyImage(request *CopyImageRequest) (response *CopyImageResponse, err error) {
	response = CreateCopyImageResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CopyImageWithChan(request *CopyImageRequest) (<-chan *CopyImageResponse, <-chan error) {
	responseChan := make(chan *CopyImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopyImage(request)
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

func (client *Client) CopyImageWithCallback(request *CopyImageRequest, callback func(response *CopyImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopyImageResponse
		var err error
		defer close(result)
		response, err = client.CopyImage(request)
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

type CopyImageRequest struct {
	*requests.RpcRequest
	DestinationRegionId    string           `position:"Query" name:"DestinationRegionId"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	Tag5Key                string           `position:"Query" name:"Tag.5.Key"`
	ImageId                string           `position:"Query" name:"ImageId"`
	Encrypted              requests.Boolean `position:"Query" name:"Encrypted"`
	Tag5Value              string           `position:"Query" name:"Tag.5.Value"`
	Tag3Key                string           `position:"Query" name:"Tag.3.Key"`
	DestinationImageName   string           `position:"Query" name:"DestinationImageName"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	Tag1Key                string           `position:"Query" name:"Tag.1.Key"`
	Tag2Key                string           `position:"Query" name:"Tag.2.Key"`
	Tag1Value              string           `position:"Query" name:"Tag.1.Value"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Tag4Value              string           `position:"Query" name:"Tag.4.Value"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	Tag3Value              string           `position:"Query" name:"Tag.3.Value"`
	DestinationDescription string           `position:"Query" name:"DestinationDescription"`
	Tag2Value              string           `position:"Query" name:"Tag.2.Value"`
	Tag4Key                string           `position:"Query" name:"Tag.4.Key"`
}

type CopyImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ImageId   string `json:"ImageId" xml:"ImageId"`
}

func CreateCopyImageRequest() (request *CopyImageRequest) {
	request = &CopyImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CopyImage", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

func CreateCopyImageResponse() (response *CopyImageResponse) {
	response = &CopyImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
