package push

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

func (client *Client) QueryDeviceStat(request *QueryDeviceStatRequest) (response *QueryDeviceStatResponse, err error) {
	response = CreateQueryDeviceStatResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryDeviceStatWithChan(request *QueryDeviceStatRequest) (<-chan *QueryDeviceStatResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceStatResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceStat(request)
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

func (client *Client) QueryDeviceStatWithCallback(request *QueryDeviceStatRequest, callback func(response *QueryDeviceStatResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceStatResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceStat(request)
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

type QueryDeviceStatRequest struct {
	*requests.RpcRequest
	EndTime    string           `position:"Query" name:"EndTime"`
	StartTime  string           `position:"Query" name:"StartTime"`
	AppKey     requests.Integer `position:"Query" name:"AppKey"`
	DeviceType string           `position:"Query" name:"DeviceType"`
	QueryType  string           `position:"Query" name:"QueryType"`
}

type QueryDeviceStatResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	AppDeviceStats struct {
		AppDeviceStat []struct {
			Time       string `json:"Time" xml:"Time"`
			Count      int    `json:"Count" xml:"Count"`
			DeviceType string `json:"DeviceType" xml:"DeviceType"`
		} `json:"AppDeviceStat" xml:"AppDeviceStat"`
	} `json:"AppDeviceStats" xml:"AppDeviceStats"`
}

func CreateQueryDeviceStatRequest() (request *QueryDeviceStatRequest) {
	request = &QueryDeviceStatRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "QueryDeviceStat", "", "openAPI")
	request.Method = requests.POST
	return
}

func CreateQueryDeviceStatResponse() (response *QueryDeviceStatResponse) {
	response = &QueryDeviceStatResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
