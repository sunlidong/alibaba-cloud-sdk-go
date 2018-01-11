package cloudphoto

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

func (client *Client) LikePhoto(request *LikePhotoRequest) (response *LikePhotoResponse, err error) {
	response = CreateLikePhotoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) LikePhotoWithChan(request *LikePhotoRequest) (<-chan *LikePhotoResponse, <-chan error) {
	responseChan := make(chan *LikePhotoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LikePhoto(request)
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

func (client *Client) LikePhotoWithCallback(request *LikePhotoRequest, callback func(response *LikePhotoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LikePhotoResponse
		var err error
		defer close(result)
		response, err = client.LikePhoto(request)
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

type LikePhotoRequest struct {
	*requests.RpcRequest
	LibraryId string           `position:"Query" name:"LibraryId"`
	StoreName string           `position:"Query" name:"StoreName"`
	PhotoId   requests.Integer `position:"Query" name:"PhotoId"`
}

type LikePhotoResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
}

func CreateLikePhotoRequest() (request *LikePhotoRequest) {
	request = &LikePhotoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "LikePhoto", "cloudphoto", "openAPI")
	request.Method = requests.POST
	return
}

func CreateLikePhotoResponse() (response *LikePhotoResponse) {
	response = &LikePhotoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
