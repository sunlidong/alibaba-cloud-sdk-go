package mts

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

func (client *Client) QueryAnalysisJobList(request *QueryAnalysisJobListRequest) (response *QueryAnalysisJobListResponse, err error) {
	response = CreateQueryAnalysisJobListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryAnalysisJobListWithChan(request *QueryAnalysisJobListRequest) (<-chan *QueryAnalysisJobListResponse, <-chan error) {
	responseChan := make(chan *QueryAnalysisJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAnalysisJobList(request)
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

func (client *Client) QueryAnalysisJobListWithCallback(request *QueryAnalysisJobListRequest, callback func(response *QueryAnalysisJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAnalysisJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryAnalysisJobList(request)
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

type QueryAnalysisJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AnalysisJobIds       string           `position:"Query" name:"AnalysisJobIds"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type QueryAnalysisJobListResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	NonExistAnalysisJobIds struct {
		String []string `json:"String" xml:"String"`
	} `json:"NonExistAnalysisJobIds" xml:"NonExistAnalysisJobIds"`
	AnalysisJobList struct {
		AnalysisJob []struct {
			Id           string `json:"Id" xml:"Id"`
			UserData     string `json:"UserData" xml:"UserData"`
			State        string `json:"State" xml:"State"`
			Code         string `json:"Code" xml:"Code"`
			Message      string `json:"Message" xml:"Message"`
			Percent      int    `json:"Percent" xml:"Percent"`
			CreationTime string `json:"CreationTime" xml:"CreationTime"`
			PipelineId   string `json:"PipelineId" xml:"PipelineId"`
			Priority     string `json:"Priority" xml:"Priority"`
			InputFile    struct {
				Bucket   string `json:"Bucket" xml:"Bucket"`
				Location string `json:"Location" xml:"Location"`
				Object   string `json:"Object" xml:"Object"`
			} `json:"InputFile" xml:"InputFile"`
			AnalysisConfig struct {
				QualityControl struct {
					RateQuality     string `json:"RateQuality" xml:"RateQuality"`
					MethodStreaming string `json:"MethodStreaming" xml:"MethodStreaming"`
				} `json:"QualityControl" xml:"QualityControl"`
				PropertiesControl struct {
					Deinterlace string `json:"Deinterlace" xml:"Deinterlace"`
					Crop        struct {
						Mode   string `json:"Mode" xml:"Mode"`
						Width  string `json:"Width" xml:"Width"`
						Height string `json:"Height" xml:"Height"`
						Top    string `json:"Top" xml:"Top"`
						Left   string `json:"Left" xml:"Left"`
					} `json:"Crop" xml:"Crop"`
				} `json:"PropertiesControl" xml:"PropertiesControl"`
			} `json:"AnalysisConfig" xml:"AnalysisConfig"`
			MNSMessageResult struct {
				MessageId    string `json:"MessageId" xml:"MessageId"`
				ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
				ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
			} `json:"MNSMessageResult" xml:"MNSMessageResult"`
			TemplateList struct {
				Template []struct {
					Id        string `json:"Id" xml:"Id"`
					Name      string `json:"Name" xml:"Name"`
					State     string `json:"State" xml:"State"`
					Container struct {
						Format string `json:"Format" xml:"Format"`
					} `json:"Container" xml:"Container"`
					Video struct {
						Codec      string `json:"Codec" xml:"Codec"`
						Profile    string `json:"Profile" xml:"Profile"`
						Bitrate    string `json:"Bitrate" xml:"Bitrate"`
						Crf        string `json:"Crf" xml:"Crf"`
						Width      string `json:"Width" xml:"Width"`
						Height     string `json:"Height" xml:"Height"`
						Fps        string `json:"Fps" xml:"Fps"`
						Gop        string `json:"Gop" xml:"Gop"`
						Preset     string `json:"Preset" xml:"Preset"`
						ScanMode   string `json:"ScanMode" xml:"ScanMode"`
						Bufsize    string `json:"Bufsize" xml:"Bufsize"`
						Maxrate    string `json:"Maxrate" xml:"Maxrate"`
						PixFmt     string `json:"PixFmt" xml:"PixFmt"`
						Degrain    string `json:"Degrain" xml:"Degrain"`
						Qscale     string `json:"Qscale" xml:"Qscale"`
						BitrateBnd struct {
							Max string `json:"Max" xml:"Max"`
							Min string `json:"Min" xml:"Min"`
						} `json:"BitrateBnd" xml:"BitrateBnd"`
					} `json:"Video" xml:"Video"`
					Audio struct {
						Codec      string `json:"Codec" xml:"Codec"`
						Profile    string `json:"Profile" xml:"Profile"`
						Samplerate string `json:"Samplerate" xml:"Samplerate"`
						Bitrate    string `json:"Bitrate" xml:"Bitrate"`
						Channels   string `json:"Channels" xml:"Channels"`
						Qscale     string `json:"Qscale" xml:"Qscale"`
					} `json:"Audio" xml:"Audio"`
					TransConfig struct {
						TransMode string `json:"TransMode" xml:"TransMode"`
					} `json:"TransConfig" xml:"TransConfig"`
					MuxConfig struct {
						Segment struct {
							Duration string `json:"Duration" xml:"Duration"`
						} `json:"Segment" xml:"Segment"`
						Gif struct {
							Loop       string `json:"Loop" xml:"Loop"`
							FinalDelay string `json:"FinalDelay" xml:"FinalDelay"`
						} `json:"Gif" xml:"Gif"`
					} `json:"MuxConfig" xml:"MuxConfig"`
				} `json:"Template" xml:"Template"`
			} `json:"TemplateList" xml:"TemplateList"`
		} `json:"AnalysisJob" xml:"AnalysisJob"`
	} `json:"AnalysisJobList" xml:"AnalysisJobList"`
}

func CreateQueryAnalysisJobListRequest() (request *QueryAnalysisJobListRequest) {
	request = &QueryAnalysisJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryAnalysisJobList", "mts", "openAPI")
	request.Method = requests.POST
	return
}

func CreateQueryAnalysisJobListResponse() (response *QueryAnalysisJobListResponse) {
	response = &QueryAnalysisJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
