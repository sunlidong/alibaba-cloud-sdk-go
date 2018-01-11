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

func (client *Client) DescribeRefreshQuota(request *DescribeRefreshQuotaRequest) (response *DescribeRefreshQuotaResponse, err error) {
	response = CreateDescribeRefreshQuotaResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeRefreshQuotaWithChan(request *DescribeRefreshQuotaRequest) (<-chan *DescribeRefreshQuotaResponse, <-chan error) {
	responseChan := make(chan *DescribeRefreshQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRefreshQuota(request)
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

func (client *Client) DescribeRefreshQuotaWithCallback(request *DescribeRefreshQuotaRequest, callback func(response *DescribeRefreshQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRefreshQuotaResponse
		var err error
		defer close(result)
		response, err = client.DescribeRefreshQuota(request)
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

type DescribeRefreshQuotaRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeRefreshQuotaResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	UrlQuota      string `json:"UrlQuota" xml:"UrlQuota"`
	DirQuota      string `json:"DirQuota" xml:"DirQuota"`
	UrlRemain     string `json:"UrlRemain" xml:"UrlRemain"`
	DirRemain     string `json:"DirRemain" xml:"DirRemain"`
	PreloadQuota  string `json:"PreloadQuota" xml:"PreloadQuota"`
	BlockQuota    string `json:"BlockQuota" xml:"BlockQuota"`
	PreloadRemain string `json:"PreloadRemain" xml:"PreloadRemain"`
	BlockRemain   string `json:"blockRemain" xml:"blockRemain"`
}

func CreateDescribeRefreshQuotaRequest() (request *DescribeRefreshQuotaRequest) {
	request = &DescribeRefreshQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeRefreshQuota", "", "openAPI")
	request.Method = requests.POST
	return
}

func CreateDescribeRefreshQuotaResponse() (response *DescribeRefreshQuotaResponse) {
	response = &DescribeRefreshQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
