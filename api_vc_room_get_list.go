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

// GetVCRoomList 该接口可以用来查询某个会议室层级下会议室列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room/list
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/room/list
func (r *VCService) GetVCRoomList(ctx context.Context, request *GetVCRoomListReq, options ...MethodOptionFunc) (*GetVCRoomListResp, *Response, error) {
	if r.cli.mock.mockVCGetVCRoomList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] VC#GetVCRoomList mock enable")
		return r.cli.mock.mockVCGetVCRoomList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "GetVCRoomList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/rooms",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getVCRoomListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCGetVCRoomList mock VCGetVCRoomList method
func (r *Mock) MockVCGetVCRoomList(f func(ctx context.Context, request *GetVCRoomListReq, options ...MethodOptionFunc) (*GetVCRoomListResp, *Response, error)) {
	r.mockVCGetVCRoomList = f
}

// UnMockVCGetVCRoomList un-mock VCGetVCRoomList method
func (r *Mock) UnMockVCGetVCRoomList() {
	r.mockVCGetVCRoomList = nil
}

// GetVCRoomListReq ...
type GetVCRoomListReq struct {
	PageSize    *int64  `query:"page_size" json:"-"`     // 分页大小, 示例值: 10, 最大值: `100`
	PageToken   *string `query:"page_token" json:"-"`    // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 10
	RoomLevelID *string `query:"room_level_id" json:"-"` // 层级ID, 当需要获取租户下会议室列表时, room_level_id可传空, 示例值: omb_4ad1a2c7a2fbc5fc9570f38456931293
	UserIDType  *IDType `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetVCRoomListResp ...
type GetVCRoomListResp struct {
	Rooms     []*GetVCRoomListRespRoom `json:"rooms,omitempty"`      // 会议室列表
	PageToken string                   `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                     `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetVCRoomListRespRoom ...
type GetVCRoomListRespRoom struct {
	RoomID       string                           `json:"room_id,omitempty"`        // 会议室ID
	Name         string                           `json:"name,omitempty"`           // 会议室名称
	Capacity     int64                            `json:"capacity,omitempty"`       // 会议室能容纳的人数
	Description  string                           `json:"description,omitempty"`    // 会议室的相关描述
	DisplayID    string                           `json:"display_id,omitempty"`     // 会议室的展示ID
	CustomRoomID string                           `json:"custom_room_id,omitempty"` // 自定义的会议室ID
	RoomLevelID  string                           `json:"room_level_id,omitempty"`  // 层级ID
	Path         []string                         `json:"path,omitempty"`           // 层级路径
	RoomStatus   *GetVCRoomListRespRoomRoomStatus `json:"room_status,omitempty"`    // 会议室状态
	Device       []*GetVCRoomListRespRoomDevice   `json:"device,omitempty"`         // 设施信息列表
}

// GetVCRoomListRespRoomDevice ...
type GetVCRoomListRespRoomDevice struct {
	Name string `json:"name,omitempty"` // 设施名称
}

// GetVCRoomListRespRoomRoomStatus ...
type GetVCRoomListRespRoomRoomStatus struct {
	Status           bool     `json:"status,omitempty"`             // 是否启用会议室
	ScheduleStatus   bool     `json:"schedule_status,omitempty"`    // 会议室未来状态为启用或禁用
	DisableStartTime string   `json:"disable_start_time,omitempty"` // 禁用开始时间（unix时间, 单位sec）
	DisableEndTime   string   `json:"disable_end_time,omitempty"`   // 禁用结束时间（unix时间, 单位sec, 数值0表示永久禁用）
	DisableReason    string   `json:"disable_reason,omitempty"`     // 禁用原因
	ContactIDs       []string `json:"contact_ids,omitempty"`        // 联系人列表, id类型由user_id_type参数决定
	DisableNotice    bool     `json:"disable_notice,omitempty"`     // 是否在禁用时发送通知给预定了该会议室的员工
	ResumeNotice     bool     `json:"resume_notice,omitempty"`      // 是否在恢复启用时发送通知给联系人
}

// getVCRoomListResp ...
type getVCRoomListResp struct {
	Code  int64              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string             `json:"msg,omitempty"`  // 错误描述
	Data  *GetVCRoomListResp `json:"data,omitempty"`
	Error *ErrorDetail       `json:"error,omitempty"`
}
