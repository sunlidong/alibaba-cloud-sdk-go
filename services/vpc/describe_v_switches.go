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

func (client *Client) DescribeVSwitches(request *DescribeVSwitchesRequest) (response *DescribeVSwitchesResponse, err error) {
	response = CreateDescribeVSwitchesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeVSwitchesWithChan(request *DescribeVSwitchesRequest) (<-chan *DescribeVSwitchesResponse, <-chan error) {
	responseChan := make(chan *DescribeVSwitchesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVSwitches(request)
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

func (client *Client) DescribeVSwitchesWithCallback(request *DescribeVSwitchesRequest, callback func(response *DescribeVSwitchesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVSwitchesResponse
		var err error
		defer close(result)
		response, err = client.DescribeVSwitches(request)
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

type DescribeVSwitchesRequest struct {
	*requests.RpcRequest
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	VSwitchName          string           `position:"Query" name:"VSwitchName"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	IsDefault            requests.Boolean `position:"Query" name:"IsDefault"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	VpcId                string           `position:"Query" name:"VpcId"`
}

type DescribeVSwitchesResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	VSwitches  struct {
		VSwitch []struct {
			VSwitchId               string `json:"VSwitchId" xml:"VSwitchId"`
			VpcId                   string `json:"VpcId" xml:"VpcId"`
			Status                  string `json:"Status" xml:"Status"`
			CidrBlock               string `json:"CidrBlock" xml:"CidrBlock"`
			ZoneId                  string `json:"ZoneId" xml:"ZoneId"`
			AvailableIpAddressCount int    `json:"AvailableIpAddressCount" xml:"AvailableIpAddressCount"`
			Description             string `json:"Description" xml:"Description"`
			VSwitchName             string `json:"VSwitchName" xml:"VSwitchName"`
			CreationTime            string `json:"CreationTime" xml:"CreationTime"`
			IsDefault               bool   `json:"IsDefault" xml:"IsDefault"`
		} `json:"VSwitch" xml:"VSwitch"`
	} `json:"VSwitches" xml:"VSwitches"`
}

func CreateDescribeVSwitchesRequest() (request *DescribeVSwitchesRequest) {
	request = &DescribeVSwitchesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVSwitches", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

func CreateDescribeVSwitchesResponse() (response *DescribeVSwitchesResponse) {
	response = &DescribeVSwitchesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
