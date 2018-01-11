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

func (client *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
	response = CreateCreateUserResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateUserWithChan(request *CreateUserRequest) (<-chan *CreateUserResponse, <-chan error) {
	responseChan := make(chan *CreateUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUser(request)
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

func (client *Client) CreateUserWithCallback(request *CreateUserRequest, callback func(response *CreateUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUserResponse
		var err error
		defer close(result)
		response, err = client.CreateUser(request)
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

type CreateUserRequest struct {
	*requests.RpcRequest
	MobilePhone string `position:"Query" name:"MobilePhone"`
	UserName    string `position:"Query" name:"UserName"`
	Email       string `position:"Query" name:"Email"`
	Comments    string `position:"Query" name:"Comments"`
	DisplayName string `position:"Query" name:"DisplayName"`
}

type CreateUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	User      struct {
		UserId      string `json:"UserId" xml:"UserId"`
		UserName    string `json:"UserName" xml:"UserName"`
		DisplayName string `json:"DisplayName" xml:"DisplayName"`
		MobilePhone string `json:"MobilePhone" xml:"MobilePhone"`
		Email       string `json:"Email" xml:"Email"`
		Comments    string `json:"Comments" xml:"Comments"`
		CreateDate  string `json:"CreateDate" xml:"CreateDate"`
	} `json:"User" xml:"User"`
}

func CreateCreateUserRequest() (request *CreateUserRequest) {
	request = &CreateUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "CreateUser", "", "openAPI")
	request.Method = requests.POST
	return
}

func CreateCreateUserResponse() (response *CreateUserResponse) {
	response = &CreateUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
