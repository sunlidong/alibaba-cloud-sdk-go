package ess

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

func (client *Client) AttachInstances(request *AttachInstancesRequest) (response *AttachInstancesResponse, err error) {
	response = CreateAttachInstancesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AttachInstancesWithChan(request *AttachInstancesRequest) (<-chan *AttachInstancesResponse, <-chan error) {
	responseChan := make(chan *AttachInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachInstances(request)
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

func (client *Client) AttachInstancesWithCallback(request *AttachInstancesRequest, callback func(response *AttachInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachInstancesResponse
		var err error
		defer close(result)
		response, err = client.AttachInstances(request)
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

type AttachInstancesRequest struct {
	*requests.RpcRequest
	LoadBalancerWeight11 requests.Integer `position:"Query" name:"LoadBalancerWeight.11"`
	LoadBalancerWeight12 requests.Integer `position:"Query" name:"LoadBalancerWeight.12"`
	LoadBalancerWeight10 requests.Integer `position:"Query" name:"LoadBalancerWeight.10"`
	LoadBalancerWeight15 requests.Integer `position:"Query" name:"LoadBalancerWeight.15"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	LoadBalancerWeight16 requests.Integer `position:"Query" name:"LoadBalancerWeight.16"`
	LoadBalancerWeight13 requests.Integer `position:"Query" name:"LoadBalancerWeight.13"`
	LoadBalancerWeight14 requests.Integer `position:"Query" name:"LoadBalancerWeight.14"`
	InstanceId10         string           `position:"Query" name:"InstanceId.10"`
	InstanceId12         string           `position:"Query" name:"InstanceId.12"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId11         string           `position:"Query" name:"InstanceId.11"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	InstanceId19         string           `position:"Query" name:"InstanceId.19"`
	InstanceId17         string           `position:"Query" name:"InstanceId.17"`
	InstanceId18         string           `position:"Query" name:"InstanceId.18"`
	InstanceId6          string           `position:"Query" name:"InstanceId.6"`
	InstanceId15         string           `position:"Query" name:"InstanceId.15"`
	LoadBalancerWeight1  requests.Integer `position:"Query" name:"LoadBalancerWeight.1"`
	InstanceId7          string           `position:"Query" name:"InstanceId.7"`
	InstanceId16         string           `position:"Query" name:"InstanceId.16"`
	InstanceId8          string           `position:"Query" name:"InstanceId.8"`
	InstanceId13         string           `position:"Query" name:"InstanceId.13"`
	InstanceId9          string           `position:"Query" name:"InstanceId.9"`
	InstanceId14         string           `position:"Query" name:"InstanceId.14"`
	LoadBalancerWeight18 requests.Integer `position:"Query" name:"LoadBalancerWeight.18"`
	LoadBalancerWeight4  requests.Integer `position:"Query" name:"LoadBalancerWeight.4"`
	InstanceId2          string           `position:"Query" name:"InstanceId.2"`
	LoadBalancerWeight17 requests.Integer `position:"Query" name:"LoadBalancerWeight.17"`
	LoadBalancerWeight5  requests.Integer `position:"Query" name:"LoadBalancerWeight.5"`
	InstanceId3          string           `position:"Query" name:"InstanceId.3"`
	LoadBalancerWeight2  requests.Integer `position:"Query" name:"LoadBalancerWeight.2"`
	InstanceId4          string           `position:"Query" name:"InstanceId.4"`
	LoadBalancerWeight19 requests.Integer `position:"Query" name:"LoadBalancerWeight.19"`
	LoadBalancerWeight3  requests.Integer `position:"Query" name:"LoadBalancerWeight.3"`
	InstanceId5          string           `position:"Query" name:"InstanceId.5"`
	LoadBalancerWeight8  requests.Integer `position:"Query" name:"LoadBalancerWeight.8"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	LoadBalancerWeight9  requests.Integer `position:"Query" name:"LoadBalancerWeight.9"`
	LoadBalancerWeight6  requests.Integer `position:"Query" name:"LoadBalancerWeight.6"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	LoadBalancerWeight7  requests.Integer `position:"Query" name:"LoadBalancerWeight.7"`
	InstanceId1          string           `position:"Query" name:"InstanceId.1"`
	LoadBalancerWeight20 requests.Integer `position:"Query" name:"LoadBalancerWeight.20"`
	InstanceId20         string           `position:"Query" name:"InstanceId.20"`
}

type AttachInstancesResponse struct {
	*responses.BaseResponse
	ScalingActivityId string `json:"ScalingActivityId" xml:"ScalingActivityId"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
}

func CreateAttachInstancesRequest() (request *AttachInstancesRequest) {
	request = &AttachInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "AttachInstances", "ess", "openAPI")
	request.Method = requests.POST
	return
}

func CreateAttachInstancesResponse() (response *AttachInstancesResponse) {
	response = &AttachInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
