package ram

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

func (client *Client) SetSecurityPreference(request *SetSecurityPreferenceRequest) (response *SetSecurityPreferenceResponse, err error) {
	response = CreateSetSecurityPreferenceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetSecurityPreferenceWithChan(request *SetSecurityPreferenceRequest) (<-chan *SetSecurityPreferenceResponse, <-chan error) {
	responseChan := make(chan *SetSecurityPreferenceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetSecurityPreference(request)
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

func (client *Client) SetSecurityPreferenceWithCallback(request *SetSecurityPreferenceRequest, callback func(response *SetSecurityPreferenceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetSecurityPreferenceResponse
		var err error
		defer close(result)
		response, err = client.SetSecurityPreference(request)
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

type SetSecurityPreferenceRequest struct {
	*requests.RpcRequest
	LoginSessionDuration        requests.Integer `position:"Query" name:"LoginSessionDuration"`
	AllowUserToManageAccessKeys requests.Boolean `position:"Query" name:"AllowUserToManageAccessKeys"`
	LoginNetworkMasks           string           `position:"Query" name:"LoginNetworkMasks"`
	AllowUserToChangePassword   requests.Boolean `position:"Query" name:"AllowUserToChangePassword"`
	AllowUserToManagePublicKeys requests.Boolean `position:"Query" name:"AllowUserToManagePublicKeys"`
	AllowUserToManageMFADevices requests.Boolean `position:"Query" name:"AllowUserToManageMFADevices"`
	EnableSaveMFATicket         requests.Boolean `position:"Query" name:"EnableSaveMFATicket"`
}

type SetSecurityPreferenceResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	SecurityPreference struct {
		LoginProfilePreference struct {
			EnableSaveMFATicket       bool   `json:"EnableSaveMFATicket" xml:"EnableSaveMFATicket"`
			AllowUserToChangePassword bool   `json:"AllowUserToChangePassword" xml:"AllowUserToChangePassword"`
			LoginSessionDuration      int    `json:"LoginSessionDuration" xml:"LoginSessionDuration"`
			LoginNetworkMasks         string `json:"LoginNetworkMasks" xml:"LoginNetworkMasks"`
		} `json:"LoginProfilePreference" xml:"LoginProfilePreference"`
		AccessKeyPreference struct {
			AllowUserToManageAccessKeys bool `json:"AllowUserToManageAccessKeys" xml:"AllowUserToManageAccessKeys"`
		} `json:"AccessKeyPreference" xml:"AccessKeyPreference"`
		PublicKeyPreference struct {
			AllowUserToManagePublicKeys bool `json:"AllowUserToManagePublicKeys" xml:"AllowUserToManagePublicKeys"`
		} `json:"PublicKeyPreference" xml:"PublicKeyPreference"`
		MFAPreference struct {
			AllowUserToManageMFADevices bool `json:"AllowUserToManageMFADevices" xml:"AllowUserToManageMFADevices"`
		} `json:"MFAPreference" xml:"MFAPreference"`
	} `json:"SecurityPreference" xml:"SecurityPreference"`
}

func CreateSetSecurityPreferenceRequest() (request *SetSecurityPreferenceRequest) {
	request = &SetSecurityPreferenceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "SetSecurityPreference", "", "openAPI")
	request.Method = requests.POST
	return
}

func CreateSetSecurityPreferenceResponse() (response *SetSecurityPreferenceResponse) {
	response = &SetSecurityPreferenceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
