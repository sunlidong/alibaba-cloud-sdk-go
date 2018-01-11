package ram

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

func (client *Client) CreatePolicyVersion(request *CreatePolicyVersionRequest) (response *CreatePolicyVersionResponse, err error) {
	response = CreateCreatePolicyVersionResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreatePolicyVersionWithChan(request *CreatePolicyVersionRequest) (<-chan *CreatePolicyVersionResponse, <-chan error) {
	responseChan := make(chan *CreatePolicyVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePolicyVersion(request)
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

func (client *Client) CreatePolicyVersionWithCallback(request *CreatePolicyVersionRequest, callback func(response *CreatePolicyVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePolicyVersionResponse
		var err error
		defer close(result)
		response, err = client.CreatePolicyVersion(request)
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

type CreatePolicyVersionRequest struct {
	*requests.RpcRequest
	SetAsDefault   requests.Boolean `position:"Query" name:"SetAsDefault"`
	PolicyDocument string           `position:"Query" name:"PolicyDocument"`
	PolicyName     string           `position:"Query" name:"PolicyName"`
}

type CreatePolicyVersionResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	PolicyVersion struct {
		VersionId        string `json:"VersionId" xml:"VersionId"`
		IsDefaultVersion bool   `json:"IsDefaultVersion" xml:"IsDefaultVersion"`
		PolicyDocument   string `json:"PolicyDocument" xml:"PolicyDocument"`
		CreateDate       string `json:"CreateDate" xml:"CreateDate"`
	} `json:"PolicyVersion" xml:"PolicyVersion"`
}

func CreateCreatePolicyVersionRequest() (request *CreatePolicyVersionRequest) {
	request = &CreatePolicyVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "CreatePolicyVersion", "", "openAPI")
	request.Method = requests.POST
	return
}

func CreateCreatePolicyVersionResponse() (response *CreatePolicyVersionResponse) {
	response = &CreatePolicyVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
