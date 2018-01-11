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

func (client *Client) DescribeRouterInterfaces(request *DescribeRouterInterfacesRequest) (response *DescribeRouterInterfacesResponse, err error) {
	response = CreateDescribeRouterInterfacesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeRouterInterfacesWithChan(request *DescribeRouterInterfacesRequest) (<-chan *DescribeRouterInterfacesResponse, <-chan error) {
	responseChan := make(chan *DescribeRouterInterfacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRouterInterfaces(request)
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

func (client *Client) DescribeRouterInterfacesWithCallback(request *DescribeRouterInterfacesRequest, callback func(response *DescribeRouterInterfacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRouterInterfacesResponse
		var err error
		defer close(result)
		response, err = client.DescribeRouterInterfaces(request)
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

type DescribeRouterInterfacesRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer                  `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string                            `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer                  `position:"Query" name:"PageNumber"`
	ResourceOwnerId      requests.Integer                  `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer                  `position:"Query" name:"OwnerId"`
	Filter               *[]DescribeRouterInterfacesFilter `position:"Query" name:"Filter"  type:"Repeated"`
}

type DescribeRouterInterfacesFilter struct {
	Key   string    `name:"Key"`
	Value *[]string `name:"Value" type:"Repeated"`
}

type DescribeRouterInterfacesResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	PageNumber         int    `json:"PageNumber" xml:"PageNumber"`
	PageSize           int    `json:"PageSize" xml:"PageSize"`
	TotalCount         int    `json:"TotalCount" xml:"TotalCount"`
	RouterInterfaceSet struct {
		RouterInterfaceType []struct {
			RouterInterfaceId               string `json:"RouterInterfaceId" xml:"RouterInterfaceId"`
			OppositeRegionId                string `json:"OppositeRegionId" xml:"OppositeRegionId"`
			Role                            string `json:"Role" xml:"Role"`
			Spec                            string `json:"Spec" xml:"Spec"`
			Name                            string `json:"Name" xml:"Name"`
			Description                     string `json:"Description" xml:"Description"`
			RouterId                        string `json:"RouterId" xml:"RouterId"`
			RouterType                      string `json:"RouterType" xml:"RouterType"`
			CreationTime                    string `json:"CreationTime" xml:"CreationTime"`
			EndTime                         string `json:"EndTime" xml:"EndTime"`
			ChargeType                      string `json:"ChargeType" xml:"ChargeType"`
			Status                          string `json:"Status" xml:"Status"`
			BusinessStatus                  string `json:"BusinessStatus" xml:"BusinessStatus"`
			ConnectedTime                   string `json:"ConnectedTime" xml:"ConnectedTime"`
			OppositeInterfaceId             string `json:"OppositeInterfaceId" xml:"OppositeInterfaceId"`
			OppositeInterfaceSpec           string `json:"OppositeInterfaceSpec" xml:"OppositeInterfaceSpec"`
			OppositeInterfaceStatus         string `json:"OppositeInterfaceStatus" xml:"OppositeInterfaceStatus"`
			OppositeInterfaceBusinessStatus string `json:"OppositeInterfaceBusinessStatus" xml:"OppositeInterfaceBusinessStatus"`
			OppositeRouterId                string `json:"OppositeRouterId" xml:"OppositeRouterId"`
			OppositeRouterType              string `json:"OppositeRouterType" xml:"OppositeRouterType"`
			OppositeInterfaceOwnerId        string `json:"OppositeInterfaceOwnerId" xml:"OppositeInterfaceOwnerId"`
			AccessPointId                   string `json:"AccessPointId" xml:"AccessPointId"`
			OppositeAccessPointId           string `json:"OppositeAccessPointId" xml:"OppositeAccessPointId"`
			HealthCheckSourceIp             string `json:"HealthCheckSourceIp" xml:"HealthCheckSourceIp"`
			HealthCheckTargetIp             string `json:"HealthCheckTargetIp" xml:"HealthCheckTargetIp"`
			OppositeVpcInstanceId           string `json:"OppositeVpcInstanceId" xml:"OppositeVpcInstanceId"`
			VpcInstanceId                   string `json:"VpcInstanceId" xml:"VpcInstanceId"`
		} `json:"RouterInterfaceType" xml:"RouterInterfaceType"`
	} `json:"RouterInterfaceSet" xml:"RouterInterfaceSet"`
}

func CreateDescribeRouterInterfacesRequest() (request *DescribeRouterInterfacesRequest) {
	request = &DescribeRouterInterfacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeRouterInterfaces", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

func CreateDescribeRouterInterfacesResponse() (response *DescribeRouterInterfacesResponse) {
	response = &DescribeRouterInterfacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
