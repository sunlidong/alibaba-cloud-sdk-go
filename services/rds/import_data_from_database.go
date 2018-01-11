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

func (client *Client) ImportDataFromDatabase(request *ImportDataFromDatabaseRequest) (response *ImportDataFromDatabaseResponse, err error) {
	response = CreateImportDataFromDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ImportDataFromDatabaseWithChan(request *ImportDataFromDatabaseRequest) (<-chan *ImportDataFromDatabaseResponse, <-chan error) {
	responseChan := make(chan *ImportDataFromDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportDataFromDatabase(request)
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

func (client *Client) ImportDataFromDatabaseWithCallback(request *ImportDataFromDatabaseRequest, callback func(response *ImportDataFromDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportDataFromDatabaseResponse
		var err error
		defer close(result)
		response, err = client.ImportDataFromDatabase(request)
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

type ImportDataFromDatabaseRequest struct {
	*requests.RpcRequest
	DBInstanceId           string           `position:"Query" name:"DBInstanceId"`
	ImportDataType         string           `position:"Query" name:"ImportDataType"`
	SourceDatabaseIp       string           `position:"Query" name:"SourceDatabaseIp"`
	IsLockTable            string           `position:"Query" name:"IsLockTable"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	SourceDatabaseUserName string           `position:"Query" name:"SourceDatabaseUserName"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	SourceDatabaseDBNames  string           `position:"Query" name:"SourceDatabaseDBNames"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceDatabasePassword string           `position:"Query" name:"SourceDatabasePassword"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	SourceDatabasePort     string           `position:"Query" name:"SourceDatabasePort"`
}

type ImportDataFromDatabaseResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ImportId  int    `json:"ImportId" xml:"ImportId"`
}

func CreateImportDataFromDatabaseRequest() (request *ImportDataFromDatabaseRequest) {
	request = &ImportDataFromDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ImportDataFromDatabase", "rds", "openAPI")
	request.Method = requests.POST
	return
}

func CreateImportDataFromDatabaseResponse() (response *ImportDataFromDatabaseResponse) {
	response = &ImportDataFromDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
