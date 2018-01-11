package ecs

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

func (client *Client) DescribeInstancesFullStatus(request *DescribeInstancesFullStatusRequest) (response *DescribeInstancesFullStatusResponse, err error) {
	response = CreateDescribeInstancesFullStatusResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeInstancesFullStatusWithChan(request *DescribeInstancesFullStatusRequest) (<-chan *DescribeInstancesFullStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeInstancesFullStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstancesFullStatus(request)
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

func (client *Client) DescribeInstancesFullStatusWithCallback(request *DescribeInstancesFullStatusRequest, callback func(response *DescribeInstancesFullStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstancesFullStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstancesFullStatus(request)
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

type DescribeInstancesFullStatusRequest struct {
	*requests.RpcRequest
	PageSize              requests.Integer `position:"Query" name:"PageSize"`
	EventId               *[]string        `position:"Query" name:"EventId"  type:"Repeated"`
	NotBeforeEnd          string           `position:"Query" name:"NotBefore.End"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	NotBeforeStart        string           `position:"Query" name:"NotBefore.Start"`
	HealthStatus          string           `position:"Query" name:"HealthStatus"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	EventPublishTimeStart string           `position:"Query" name:"EventPublishTime.Start"`
	PageNumber            requests.Integer `position:"Query" name:"PageNumber"`
	Status                string           `position:"Query" name:"Status"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	EventPublishTimeEnd   string           `position:"Query" name:"EventPublishTime.End"`
	EventType             string           `position:"Query" name:"EventType"`
	InstanceId            *[]string        `position:"Query" name:"InstanceId"  type:"Repeated"`
}

type DescribeInstancesFullStatusResponse struct {
	*responses.BaseResponse
	RequestId             string `json:"RequestId" xml:"RequestId"`
	TotalCount            int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber            int    `json:"PageNumber" xml:"PageNumber"`
	PageSize              int    `json:"PageSize" xml:"PageSize"`
	InstanceFullStatusSet struct {
		InstanceFullStatusType []struct {
			InstanceId string `json:"InstanceId" xml:"InstanceId"`
			Status     struct {
				Code int    `json:"Code" xml:"Code"`
				Name string `json:"Name" xml:"Name"`
			} `json:"Status" xml:"Status"`
			HealthStatus struct {
				Code int    `json:"Code" xml:"Code"`
				Name string `json:"Name" xml:"Name"`
			} `json:"HealthStatus" xml:"HealthStatus"`
			ScheduledSystemEventSet struct {
				ScheduledSystemEventType []struct {
					EventId          string `json:"EventId" xml:"EventId"`
					EventPublishTime string `json:"EventPublishTime" xml:"EventPublishTime"`
					NotBefore        string `json:"NotBefore" xml:"NotBefore"`
					EventCycleStatus struct {
						Code int    `json:"Code" xml:"Code"`
						Name string `json:"Name" xml:"Name"`
					} `json:"EventCycleStatus" xml:"EventCycleStatus"`
					EventType struct {
						Code int    `json:"Code" xml:"Code"`
						Name string `json:"Name" xml:"Name"`
					} `json:"EventType" xml:"EventType"`
				} `json:"ScheduledSystemEventType" xml:"ScheduledSystemEventType"`
			} `json:"ScheduledSystemEventSet" xml:"ScheduledSystemEventSet"`
		} `json:"InstanceFullStatusType" xml:"InstanceFullStatusType"`
	} `json:"InstanceFullStatusSet" xml:"InstanceFullStatusSet"`
}

func CreateDescribeInstancesFullStatusRequest() (request *DescribeInstancesFullStatusRequest) {
	request = &DescribeInstancesFullStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstancesFullStatus", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

func CreateDescribeInstancesFullStatusResponse() (response *DescribeInstancesFullStatusResponse) {
	response = &DescribeInstancesFullStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
