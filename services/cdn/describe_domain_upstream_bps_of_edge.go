package cdn

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

func (client *Client) DescribeDomainUpstreamBpsOfEdge(request *DescribeDomainUpstreamBpsOfEdgeRequest) (response *DescribeDomainUpstreamBpsOfEdgeResponse, err error) {
	response = CreateDescribeDomainUpstreamBpsOfEdgeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainUpstreamBpsOfEdgeWithChan(request *DescribeDomainUpstreamBpsOfEdgeRequest) (<-chan *DescribeDomainUpstreamBpsOfEdgeResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainUpstreamBpsOfEdgeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainUpstreamBpsOfEdge(request)
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

func (client *Client) DescribeDomainUpstreamBpsOfEdgeWithCallback(request *DescribeDomainUpstreamBpsOfEdgeRequest, callback func(response *DescribeDomainUpstreamBpsOfEdgeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainUpstreamBpsOfEdgeResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainUpstreamBpsOfEdge(request)
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

type DescribeDomainUpstreamBpsOfEdgeRequest struct {
	*requests.RpcRequest
	EndTime       string           `position:"Query" name:"EndTime"`
	StartTime     string           `position:"Query" name:"StartTime"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeDomainUpstreamBpsOfEdgeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	BpsDatas  struct {
		DomainBpsModel []struct {
			Time string  `json:"Time" xml:"Time"`
			Bps  float64 `json:"Bps" xml:"Bps"`
		} `json:"DomainBpsModel" xml:"DomainBpsModel"`
	} `json:"BpsDatas" xml:"BpsDatas"`
}

func CreateDescribeDomainUpstreamBpsOfEdgeRequest() (request *DescribeDomainUpstreamBpsOfEdgeRequest) {
	request = &DescribeDomainUpstreamBpsOfEdgeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainUpstreamBpsOfEdge", "", "openAPI")
	request.Method = requests.POST
	return
}

func CreateDescribeDomainUpstreamBpsOfEdgeResponse() (response *DescribeDomainUpstreamBpsOfEdgeResponse) {
	response = &DescribeDomainUpstreamBpsOfEdgeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
