package vpc

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

func (client *Client) DescribeNqas(request *DescribeNqasRequest) (response *DescribeNqasResponse, err error) {
	response = CreateDescribeNqasResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeNqasWithChan(request *DescribeNqasRequest) (<-chan *DescribeNqasResponse, <-chan error) {
	responseChan := make(chan *DescribeNqasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNqas(request)
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

func (client *Client) DescribeNqasWithCallback(request *DescribeNqasRequest, callback func(response *DescribeNqasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNqasResponse
		var err error
		defer close(result)
		response, err = client.DescribeNqas(request)
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

type DescribeNqasRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	IsDefault            requests.Boolean `position:"Query" name:"IsDefault"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	NqaId                string           `position:"Query" name:"NqaId"`
	RouterId             string           `position:"Query" name:"RouterId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeNqasResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Nqas       struct {
		Nqa []struct {
			NqaId         string `json:"NqaId" xml:"NqaId"`
			RegionId      string `json:"RegionId" xml:"RegionId"`
			Status        string `json:"Status" xml:"Status"`
			RouterId      string `json:"RouterId" xml:"RouterId"`
			DestinationIp string `json:"DestinationIp" xml:"DestinationIp"`
		} `json:"Nqa" xml:"Nqa"`
	} `json:"Nqas" xml:"Nqas"`
}

func CreateDescribeNqasRequest() (request *DescribeNqasRequest) {
	request = &DescribeNqasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeNqas", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

func CreateDescribeNqasResponse() (response *DescribeNqasResponse) {
	response = &DescribeNqasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
