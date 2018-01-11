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

func (client *Client) MoveAlbumPhotos(request *MoveAlbumPhotosRequest) (response *MoveAlbumPhotosResponse, err error) {
	response = CreateMoveAlbumPhotosResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) MoveAlbumPhotosWithChan(request *MoveAlbumPhotosRequest) (<-chan *MoveAlbumPhotosResponse, <-chan error) {
	responseChan := make(chan *MoveAlbumPhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MoveAlbumPhotos(request)
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

func (client *Client) MoveAlbumPhotosWithCallback(request *MoveAlbumPhotosRequest, callback func(response *MoveAlbumPhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MoveAlbumPhotosResponse
		var err error
		defer close(result)
		response, err = client.MoveAlbumPhotos(request)
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

type MoveAlbumPhotosRequest struct {
	*requests.RpcRequest
	TargetAlbumId requests.Integer `position:"Query" name:"TargetAlbumId"`
	SourceAlbumId requests.Integer `position:"Query" name:"SourceAlbumId"`
	LibraryId     string           `position:"Query" name:"LibraryId"`
	StoreName     string           `position:"Query" name:"StoreName"`
	PhotoId       *[]string        `position:"Query" name:"PhotoId"  type:"Repeated"`
}

type MoveAlbumPhotosResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Results   []struct {
		Id      int    `json:"Id" xml:"Id"`
		Code    string `json:"Code" xml:"Code"`
		Message string `json:"Message" xml:"Message"`
	} `json:"Results" xml:"Results"`
}

func CreateMoveAlbumPhotosRequest() (request *MoveAlbumPhotosRequest) {
	request = &MoveAlbumPhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "MoveAlbumPhotos", "cloudphoto", "openAPI")
	request.Method = requests.POST
	return
}

func CreateMoveAlbumPhotosResponse() (response *MoveAlbumPhotosResponse) {
	response = &MoveAlbumPhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
