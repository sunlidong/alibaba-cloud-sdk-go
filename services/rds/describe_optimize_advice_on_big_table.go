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

func (client *Client) DescribeOptimizeAdviceOnBigTable(request *DescribeOptimizeAdviceOnBigTableRequest) (response *DescribeOptimizeAdviceOnBigTableResponse, err error) {
	response = CreateDescribeOptimizeAdviceOnBigTableResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeOptimizeAdviceOnBigTableWithChan(request *DescribeOptimizeAdviceOnBigTableRequest) (<-chan *DescribeOptimizeAdviceOnBigTableResponse, <-chan error) {
	responseChan := make(chan *DescribeOptimizeAdviceOnBigTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOptimizeAdviceOnBigTable(request)
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

func (client *Client) DescribeOptimizeAdviceOnBigTableWithCallback(request *DescribeOptimizeAdviceOnBigTableRequest, callback func(response *DescribeOptimizeAdviceOnBigTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOptimizeAdviceOnBigTableResponse
		var err error
		defer close(result)
		response, err = client.DescribeOptimizeAdviceOnBigTable(request)
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

type DescribeOptimizeAdviceOnBigTableRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeOptimizeAdviceOnBigTableResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	TotalRecordsCount int    `json:"TotalRecordsCount" xml:"TotalRecordsCount"`
	PageNumber        int    `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount   int    `json:"PageRecordCount" xml:"PageRecordCount"`
	Items             struct {
		AdviceOnBigTable []struct {
			DBName    string `json:"DBName" xml:"DBName"`
			TableName string `json:"TableName" xml:"TableName"`
			TableSize int    `json:"TableSize" xml:"TableSize"`
			DataSize  int    `json:"DataSize" xml:"DataSize"`
			IndexSize int    `json:"IndexSize" xml:"IndexSize"`
		} `json:"AdviceOnBigTable" xml:"AdviceOnBigTable"`
	} `json:"Items" xml:"Items"`
}

func CreateDescribeOptimizeAdviceOnBigTableRequest() (request *DescribeOptimizeAdviceOnBigTableRequest) {
	request = &DescribeOptimizeAdviceOnBigTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeOptimizeAdviceOnBigTable", "rds", "openAPI")
	request.Method = requests.POST
	return
}

func CreateDescribeOptimizeAdviceOnBigTableResponse() (response *DescribeOptimizeAdviceOnBigTableResponse) {
	response = &DescribeOptimizeAdviceOnBigTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
