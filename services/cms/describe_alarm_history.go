package cms

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

func (client *Client) DescribeAlarmHistory(request *DescribeAlarmHistoryRequest) (response *DescribeAlarmHistoryResponse, err error) {
	response = CreateDescribeAlarmHistoryResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeAlarmHistoryWithChan(request *DescribeAlarmHistoryRequest) (<-chan *DescribeAlarmHistoryResponse, <-chan error) {
	responseChan := make(chan *DescribeAlarmHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAlarmHistory(request)
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

func (client *Client) DescribeAlarmHistoryWithCallback(request *DescribeAlarmHistoryRequest, callback func(response *DescribeAlarmHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAlarmHistoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeAlarmHistory(request)
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

type DescribeAlarmHistoryRequest struct {
	*requests.RpcRequest
	EndTime    string           `position:"Query" name:"EndTime"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	StartTime  string           `position:"Query" name:"StartTime"`
	Status     string           `position:"Query" name:"Status"`
	RuleName   string           `position:"Query" name:"RuleName"`
	State      string           `position:"Query" name:"State"`
	Namespace  string           `position:"Query" name:"Namespace"`
	MetricName string           `position:"Query" name:"MetricName"`
	AlertName  string           `position:"Query" name:"AlertName"`
	Page       requests.Integer `position:"Query" name:"Page"`
	Ascending  requests.Boolean `position:"Query" name:"Ascending"`
	GroupId    string           `position:"Query" name:"GroupId"`
	OnlyCount  requests.Boolean `position:"Query" name:"OnlyCount"`
}

type DescribeAlarmHistoryResponse struct {
	*responses.BaseResponse
	Success          bool   `json:"Success" xml:"Success"`
	Code             string `json:"Code" xml:"Code"`
	Message          string `json:"Message" xml:"Message"`
	Total            string `json:"Total" xml:"Total"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	AlarmHistoryList struct {
		AlarmHistory []struct {
			Id              string `json:"Id" xml:"Id"`
			AlertName       string `json:"AlertName" xml:"AlertName"`
			GroupId         string `json:"GroupId" xml:"GroupId"`
			Namespace       string `json:"Namespace" xml:"Namespace"`
			MetricName      string `json:"MetricName" xml:"MetricName"`
			Dimensions      string `json:"Dimensions" xml:"Dimensions"`
			Expression      string `json:"Expression" xml:"Expression"`
			EvaluationCount int    `json:"EvaluationCount" xml:"EvaluationCount"`
			Value           string `json:"Value" xml:"Value"`
			AlertTime       int    `json:"AlertTime" xml:"AlertTime"`
			LastTime        int    `json:"LastTime" xml:"LastTime"`
			Level           string `json:"Level" xml:"Level"`
			PreLevel        string `json:"PreLevel" xml:"PreLevel"`
			RuleName        string `json:"ruleName" xml:"ruleName"`
			State           string `json:"State" xml:"State"`
			Status          int    `json:"Status" xml:"Status"`
			UserId          string `json:"UserId" xml:"UserId"`
			Webhooks        string `json:"Webhooks" xml:"Webhooks"`
			ContactGroups   struct {
				ContactGroup []string `json:"ContactGroup" xml:"ContactGroup"`
			} `json:"ContactGroups" xml:"ContactGroups"`
			Contacts struct {
				Contact []string `json:"Contact" xml:"Contact"`
			} `json:"Contacts" xml:"Contacts"`
			ContactALIIMs struct {
				ContactALIIM []string `json:"ContactALIIM" xml:"ContactALIIM"`
			} `json:"ContactALIIMs" xml:"ContactALIIMs"`
			ContactSmses struct {
				ContactSms []string `json:"ContactSms" xml:"ContactSms"`
			} `json:"ContactSmses" xml:"ContactSmses"`
			ContactMails struct {
				ContactMail []string `json:"ContactMail" xml:"ContactMail"`
			} `json:"ContactMails" xml:"ContactMails"`
		} `json:"AlarmHistory" xml:"AlarmHistory"`
	} `json:"AlarmHistoryList" xml:"AlarmHistoryList"`
}

func CreateDescribeAlarmHistoryRequest() (request *DescribeAlarmHistoryRequest) {
	request = &DescribeAlarmHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "DescribeAlarmHistory", "cms", "openAPI")
	request.Method = requests.POST
	return
}

func CreateDescribeAlarmHistoryResponse() (response *DescribeAlarmHistoryResponse) {
	response = &DescribeAlarmHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
