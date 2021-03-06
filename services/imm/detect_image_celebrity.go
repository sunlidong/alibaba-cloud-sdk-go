package imm

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

// DetectImageCelebrity invokes the imm.DetectImageCelebrity API synchronously
// api document: https://help.aliyun.com/api/imm/detectimagecelebrity.html
func (client *Client) DetectImageCelebrity(request *DetectImageCelebrityRequest) (response *DetectImageCelebrityResponse, err error) {
	response = CreateDetectImageCelebrityResponse()
	err = client.DoAction(request, response)
	return
}

// DetectImageCelebrityWithChan invokes the imm.DetectImageCelebrity API asynchronously
// api document: https://help.aliyun.com/api/imm/detectimagecelebrity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetectImageCelebrityWithChan(request *DetectImageCelebrityRequest) (<-chan *DetectImageCelebrityResponse, <-chan error) {
	responseChan := make(chan *DetectImageCelebrityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectImageCelebrity(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DetectImageCelebrityWithCallback invokes the imm.DetectImageCelebrity API asynchronously
// api document: https://help.aliyun.com/api/imm/detectimagecelebrity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetectImageCelebrityWithCallback(request *DetectImageCelebrityRequest, callback func(response *DetectImageCelebrityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectImageCelebrityResponse
		var err error
		defer close(result)
		response, err = client.DetectImageCelebrity(request)
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

// DetectImageCelebrityRequest is the request struct for api DetectImageCelebrity
type DetectImageCelebrityRequest struct {
	*requests.RpcRequest
	ImageUri string `position:"Query" name:"ImageUri"`
	Library  string `position:"Query" name:"Library"`
	Project  string `position:"Query" name:"Project"`
	RealUid  string `position:"Query" name:"RealUid"`
}

// DetectImageCelebrityResponse is the response struct for api DetectImageCelebrity
type DetectImageCelebrityResponse struct {
	*responses.BaseResponse
	RequestId string          `json:"RequestId" xml:"RequestId"`
	ImageUri  string          `json:"ImageUri" xml:"ImageUri"`
	Celebrity []CelebrityItem `json:"Celebrity" xml:"Celebrity"`
}

// CreateDetectImageCelebrityRequest creates a request to invoke DetectImageCelebrity API
func CreateDetectImageCelebrityRequest() (request *DetectImageCelebrityRequest) {
	request = &DetectImageCelebrityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "DetectImageCelebrity", "imm", "openAPI")
	return
}

// CreateDetectImageCelebrityResponse creates a response to parse from DetectImageCelebrity response
func CreateDetectImageCelebrityResponse() (response *DetectImageCelebrityResponse) {
	response = &DetectImageCelebrityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
