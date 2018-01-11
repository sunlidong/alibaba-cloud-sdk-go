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

func (client *Client) DescribePreCheckResults(request *DescribePreCheckResultsRequest) (response *DescribePreCheckResultsResponse, err error) {
	response = CreateDescribePreCheckResultsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribePreCheckResultsWithChan(request *DescribePreCheckResultsRequest) (<-chan *DescribePreCheckResultsResponse, <-chan error) {
	responseChan := make(chan *DescribePreCheckResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePreCheckResults(request)
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

func (client *Client) DescribePreCheckResultsWithCallback(request *DescribePreCheckResultsRequest, callback func(response *DescribePreCheckResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePreCheckResultsResponse
		var err error
		defer close(result)
		response, err = client.DescribePreCheckResults(request)
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

type DescribePreCheckResultsRequest struct {
	*requests.RpcRequest
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PreCheckId           string           `position:"Query" name:"PreCheckId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribePreCheckResultsResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
	Items        struct {
		PreCheckResult []struct {
			PreCheckName   string `json:"PreCheckName" xml:"PreCheckName"`
			PreCheckResult string `json:"PreCheckResult" xml:"PreCheckResult"`
			FailReasion    string `json:"FailReasion" xml:"FailReasion"`
			RepairMethod   string `json:"RepairMethod" xml:"RepairMethod"`
		} `json:"PreCheckResult" xml:"PreCheckResult"`
	} `json:"Items" xml:"Items"`
}

func CreateDescribePreCheckResultsRequest() (request *DescribePreCheckResultsRequest) {
	request = &DescribePreCheckResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribePreCheckResults", "rds", "openAPI")
	request.Method = requests.POST
	return
}

func CreateDescribePreCheckResultsResponse() (response *DescribePreCheckResultsResponse) {
	response = &DescribePreCheckResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
