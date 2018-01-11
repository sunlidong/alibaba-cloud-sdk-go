package alidns

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

func (client *Client) DescribeSubDomainRecords(request *DescribeSubDomainRecordsRequest) (response *DescribeSubDomainRecordsResponse, err error) {
	response = CreateDescribeSubDomainRecordsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeSubDomainRecordsWithChan(request *DescribeSubDomainRecordsRequest) (<-chan *DescribeSubDomainRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeSubDomainRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSubDomainRecords(request)
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

func (client *Client) DescribeSubDomainRecordsWithCallback(request *DescribeSubDomainRecordsRequest, callback func(response *DescribeSubDomainRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSubDomainRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSubDomainRecords(request)
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

type DescribeSubDomainRecordsRequest struct {
	*requests.RpcRequest
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	SubDomain    string           `position:"Query" name:"SubDomain"`
	Type         string           `position:"Query" name:"Type"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	Lang         string           `position:"Query" name:"Lang"`
}

type DescribeSubDomainRecordsResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	TotalCount    int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber    int    `json:"PageNumber" xml:"PageNumber"`
	PageSize      int    `json:"PageSize" xml:"PageSize"`
	DomainRecords struct {
		Record []struct {
			DomainName string `json:"DomainName" xml:"DomainName"`
			RecordId   string `json:"RecordId" xml:"RecordId"`
			RR         string `json:"RR" xml:"RR"`
			Type       string `json:"Type" xml:"Type"`
			Value      string `json:"Value" xml:"Value"`
			TTL        int    `json:"TTL" xml:"TTL"`
			Priority   int    `json:"Priority" xml:"Priority"`
			Line       string `json:"Line" xml:"Line"`
			Status     string `json:"Status" xml:"Status"`
			Locked     bool   `json:"Locked" xml:"Locked"`
			Weight     int    `json:"Weight" xml:"Weight"`
		} `json:"Record" xml:"Record"`
	} `json:"DomainRecords" xml:"DomainRecords"`
}

func CreateDescribeSubDomainRecordsRequest() (request *DescribeSubDomainRecordsRequest) {
	request = &DescribeSubDomainRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeSubDomainRecords", "", "openAPI")
	request.Method = requests.POST
	return
}

func CreateDescribeSubDomainRecordsResponse() (response *DescribeSubDomainRecordsResponse) {
	response = &DescribeSubDomainRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
