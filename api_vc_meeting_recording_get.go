// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// GetVCMeetingRecording 获取一个会议的录制文件。
//
// 会议结束后并且收到了"录制完成"的事件方可获取录制文件；只有会议owner（通过开放平台预约的会议即为预约人）有权限获取；录制时间太短(<5s)有可能无法生成录制文件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/get
func (r *VCService) GetVCMeetingRecording(ctx context.Context, request *GetVCMeetingRecordingReq, options ...MethodOptionFunc) (*GetVCMeetingRecordingResp, *Response, error) {
	if r.cli.mock.mockVCGetVCMeetingRecording != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#GetVCMeetingRecording mock enable")
		return r.cli.mock.mockVCGetVCMeetingRecording(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "GetVCMeetingRecording",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/meetings/:meeting_id/recording",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getVCMeetingRecordingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCGetVCMeetingRecording mock VCGetVCMeetingRecording method
func (r *Mock) MockVCGetVCMeetingRecording(f func(ctx context.Context, request *GetVCMeetingRecordingReq, options ...MethodOptionFunc) (*GetVCMeetingRecordingResp, *Response, error)) {
	r.mockVCGetVCMeetingRecording = f
}

// UnMockVCGetVCMeetingRecording un-mock VCGetVCMeetingRecording method
func (r *Mock) UnMockVCGetVCMeetingRecording() {
	r.mockVCGetVCMeetingRecording = nil
}

// GetVCMeetingRecordingReq ...
type GetVCMeetingRecordingReq struct {
	MeetingID string `path:"meeting_id" json:"-"` // 会议ID（视频会议的唯一标识, 视频会议开始后才会产生）, 示例值: "6911188411932033028"
}

// GetVCMeetingRecordingResp ...
type GetVCMeetingRecordingResp struct {
	Recording *GetVCMeetingRecordingRespRecording `json:"recording,omitempty"` // 录制文件数据
}

// GetVCMeetingRecordingRespRecording ...
type GetVCMeetingRecordingRespRecording struct {
	URL      string `json:"url,omitempty"`      // 录制文件URL
	Duration string `json:"duration,omitempty"` // 录制总时长（单位msec）
}

// getVCMeetingRecordingResp ...
type getVCMeetingRecordingResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *GetVCMeetingRecordingResp `json:"data,omitempty"`
}
