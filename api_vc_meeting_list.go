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

// GetVCMeetingList 查询会议明细, 具体权限要求请参考[资源介绍](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-room-data/resource-introduction)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting_list/get
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/meeting-room-data/get
func (r *VCService) GetVCMeetingList(ctx context.Context, request *GetVCMeetingListReq, options ...MethodOptionFunc) (*GetVCMeetingListResp, *Response, error) {
	if r.cli.mock.mockVCGetVCMeetingList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] VC#GetVCMeetingList mock enable")
		return r.cli.mock.mockVCGetVCMeetingList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "GetVCMeetingList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/meeting_list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getVCMeetingListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCGetVCMeetingList mock VCGetVCMeetingList method
func (r *Mock) MockVCGetVCMeetingList(f func(ctx context.Context, request *GetVCMeetingListReq, options ...MethodOptionFunc) (*GetVCMeetingListResp, *Response, error)) {
	r.mockVCGetVCMeetingList = f
}

// UnMockVCGetVCMeetingList un-mock VCGetVCMeetingList method
func (r *Mock) UnMockVCGetVCMeetingList() {
	r.mockVCGetVCMeetingList = nil
}

// GetVCMeetingListReq ...
type GetVCMeetingListReq struct {
	StartTime     string  `query:"start_time" json:"-"`     // 查询开始时间（unix时间, 单位sec）, 示例值: 1655276858
	EndTime       string  `query:"end_time" json:"-"`       // 查询结束时间（unix时间, 单位sec）, 示例值: 1655276858
	MeetingStatus *int64  `query:"meeting_status" json:"-"` // 会议状态（不传默认为已结束会议）, 示例值: 2, 可选值有: 1: 进行中, 2: 已结束, 3: 待召开
	MeetingNo     *string `query:"meeting_no" json:"-"`     // 按9位会议号筛选（最多一个筛选条件）, 示例值: 123456789
	UserID        *string `query:"user_id" json:"-"`        // 按参会Lark用户筛选（最多一个筛选条件）, 示例值: ou_3ec3f6a28a0d08c45d895276e8e5e19b
	RoomID        *string `query:"room_id" json:"-"`        // 按参会Rooms筛选（最多一个筛选条件）, 示例值: omm_eada1d61a550955240c28757e7dec3af
	MeetingType   *int64  `query:"meeting_type" json:"-"`   // 按会议类型筛选（最多一个筛选条件）, 示例值: 2, 可选值有: 1: 全部类型（默认）, 2: 视频会议, 3: 本地投屏
	PageSize      *int64  `query:"page_size" json:"-"`      // 分页尺寸大小, 示例值: 20, 默认值: `20`, 取值范围: `20` ～ `100`
	PageToken     *string `query:"page_token" json:"-"`     // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 20
	UserIDType    *IDType `query:"user_id_type" json:"-"`   // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetVCMeetingListResp ...
type GetVCMeetingListResp struct {
	MeetingList []*GetVCMeetingListRespMeeting `json:"meeting_list,omitempty"` // 会议列表
	PageToken   string                         `json:"page_token,omitempty"`   // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore     bool                           `json:"has_more,omitempty"`     // 是否还有更多项
}

// GetVCMeetingListRespMeeting ...
type GetVCMeetingListRespMeeting struct {
	MeetingID            string                                     `json:"meeting_id,omitempty"`             // 9位会议号
	MeetingTopic         string                                     `json:"meeting_topic,omitempty"`          // 会议主题
	MeetingType          int64                                      `json:"meeting_type,omitempty"`           // 会议类型, 可选值有: 1: 全部类型（默认）, 2: 视频会议, 3: 本地投屏
	Organizer            string                                     `json:"organizer,omitempty"`              // 组织者
	Department           string                                     `json:"department,omitempty"`             // 部门
	UserID               string                                     `json:"user_id,omitempty"`                // 用户ID
	EmployeeID           string                                     `json:"employee_id,omitempty"`            // 工号
	Email                string                                     `json:"email,omitempty"`                  // 邮箱
	Mobile               string                                     `json:"mobile,omitempty"`                 // 手机
	MeetingStartTime     string                                     `json:"meeting_start_time,omitempty"`     // 会议开始时间
	MeetingEndTime       string                                     `json:"meeting_end_time,omitempty"`       // 会议结束时间
	MeetingDuration      string                                     `json:"meeting_duration,omitempty"`       // 会议持续时间
	NumberOfParticipants string                                     `json:"number_of_participants,omitempty"` // 参会人数
	NumberOfDevices      string                                     `json:"number_of_devices,omitempty"`      // 累计入会设备数
	Audio                bool                                       `json:"audio,omitempty"`                  // 音频
	Video                bool                                       `json:"video,omitempty"`                  // 视频
	Sharing              bool                                       `json:"sharing,omitempty"`                // 共享
	Recording            bool                                       `json:"recording,omitempty"`              // 录制
	Telephone            bool                                       `json:"telephone,omitempty"`              // 电话
	ReservedRooms        []*GetVCMeetingListRespMeetingReservedRoom `json:"reserved_rooms,omitempty"`         // 关联会议室列表
	HasRelatedDocument   bool                                       `json:"has_related_document,omitempty"`   // 是否有关联文档和纪要
}

// GetVCMeetingListRespMeetingReservedRoom ...
type GetVCMeetingListRespMeetingReservedRoom struct {
	RoomID   string `json:"room_id,omitempty"`   // 会议室ID
	RoomName string `json:"room_name,omitempty"` // 会议室名称
}

// getVCMeetingListResp ...
type getVCMeetingListResp struct {
	Code  int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                `json:"msg,omitempty"`  // 错误描述
	Data  *GetVCMeetingListResp `json:"data,omitempty"`
	Error *ErrorDetail          `json:"error,omitempty"`
}
