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

// BatchGetMeetingRoomRoomID 该接口用于根据租户自定义会议室ID查询会议室ID。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYzMxYjL2MTM24iNzEjN
func (r *MeetingRoomService) BatchGetMeetingRoomRoomID(ctx context.Context, request *BatchGetMeetingRoomRoomIDReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomRoomIDResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomBatchGetMeetingRoomRoomID != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MeetingRoom#BatchGetMeetingRoomRoomID mock enable")
		return r.cli.mock.mockMeetingRoomBatchGetMeetingRoomRoomID(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "MeetingRoom",
		API:                   "BatchGetMeetingRoomRoomID",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/meeting_room/room/batch_get_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetMeetingRoomRoomIDResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMeetingRoomBatchGetMeetingRoomRoomID mock MeetingRoomBatchGetMeetingRoomRoomID method
func (r *Mock) MockMeetingRoomBatchGetMeetingRoomRoomID(f func(ctx context.Context, request *BatchGetMeetingRoomRoomIDReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomRoomIDResp, *Response, error)) {
	r.mockMeetingRoomBatchGetMeetingRoomRoomID = f
}

// UnMockMeetingRoomBatchGetMeetingRoomRoomID un-mock MeetingRoomBatchGetMeetingRoomRoomID method
func (r *Mock) UnMockMeetingRoomBatchGetMeetingRoomRoomID() {
	r.mockMeetingRoomBatchGetMeetingRoomRoomID = nil
}

// BatchGetMeetingRoomRoomIDReq ...
type BatchGetMeetingRoomRoomIDReq struct {
	CustomRoomIDs string `query:"custom_room_ids" json:"-"` // 用于查询指定会议室的租户自定义会议室ID
}

// BatchGetMeetingRoomRoomIDResp ...
type BatchGetMeetingRoomRoomIDResp struct {
	Rooms []*BatchGetMeetingRoomRoomIDRespRoom `json:"rooms,omitempty"` // 会议室列表
}

// BatchGetMeetingRoomRoomIDRespRoom ...
type BatchGetMeetingRoomRoomIDRespRoom struct {
	RoomID       string `json:"room_id,omitempty"`        // 会议室 ID
	CustomRoomID string `json:"custom_room_id,omitempty"` // 租户自定义会议室 ID
}

// batchGetMeetingRoomRoomIDResp ...
type batchGetMeetingRoomRoomIDResp struct {
	Code int64                          `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 返回码的描述, "success" 表示成功, 其他为错误提示信息
	Data *BatchGetMeetingRoomRoomIDResp `json:"data,omitempty"` // 返回业务信息
}
