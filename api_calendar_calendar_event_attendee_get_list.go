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

// GetCalendarEventAttendeeList 获取日程的参与人列表, 若参与者列表中有群组, 请使用 [获取参与人群成员列表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee-chat_member/list) 。
//
// - 当前身份必须对日历有reader、writer或owner权限（调用[获取日历](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get)接口, role字段可查看权限）。
// - 当前身份必须有权限查看日程的参与人列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/list
func (r *CalendarService) GetCalendarEventAttendeeList(ctx context.Context, request *GetCalendarEventAttendeeListReq, options ...MethodOptionFunc) (*GetCalendarEventAttendeeListResp, *Response, error) {
	if r.cli.mock.mockCalendarGetCalendarEventAttendeeList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#GetCalendarEventAttendeeList mock enable")
		return r.cli.mock.mockCalendarGetCalendarEventAttendeeList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "GetCalendarEventAttendeeList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getCalendarEventAttendeeListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarGetCalendarEventAttendeeList mock CalendarGetCalendarEventAttendeeList method
func (r *Mock) MockCalendarGetCalendarEventAttendeeList(f func(ctx context.Context, request *GetCalendarEventAttendeeListReq, options ...MethodOptionFunc) (*GetCalendarEventAttendeeListResp, *Response, error)) {
	r.mockCalendarGetCalendarEventAttendeeList = f
}

// UnMockCalendarGetCalendarEventAttendeeList un-mock CalendarGetCalendarEventAttendeeList method
func (r *Mock) UnMockCalendarGetCalendarEventAttendeeList() {
	r.mockCalendarGetCalendarEventAttendeeList = nil
}

// GetCalendarEventAttendeeListReq ...
type GetCalendarEventAttendeeListReq struct {
	CalendarID string  `path:"calendar_id" json:"-"`   // 日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值: "feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	EventID    string  `path:"event_id" json:"-"`      // 日程ID。参见[日程ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/introduction), 示例值: "xxxxxxxxx_0"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "780TRhwXXXXX"
	PageSize   *int64  `query:"page_size" json:"-"`    // 分页大小, 示例值: 10, 最大值: `100`
}

// GetCalendarEventAttendeeListResp ...
type GetCalendarEventAttendeeListResp struct {
	Items     []*GetCalendarEventAttendeeListRespItem `json:"items,omitempty"`      // 日程的参与者列表
	HasMore   bool                                    `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                                  `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetCalendarEventAttendeeListRespItem ...
type GetCalendarEventAttendeeListRespItem struct {
	Type                  CalendarEventAttendeeType                                    `json:"type,omitempty"`                   // 参与人类型, 可选值有: user: 用户, chat: 群组, resource: 会议室, third_party: 邮箱
	AttendeeID            string                                                       `json:"attendee_id,omitempty"`            // 参与人ID。参见[参与人ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/introduction#4998889c)
	RsvpStatus            string                                                       `json:"rsvp_status,omitempty"`            // 参与人RSVP状态, 可选值有: needs_action: 参与人尚未回复状态, 或表示会议室预约中, accept: 参与人回复接受, 或表示会议室预约成功, tentative: 参与人回复待定, decline: 参与人回复拒绝, 或表示会议室预约失败, removed: 参与人或会议室已经从日程中被移除
	IsOptional            bool                                                         `json:"is_optional,omitempty"`            // 参与人是否为「可选参加」, 无法编辑群参与人的此字段
	IsOrganizer           bool                                                         `json:"is_organizer,omitempty"`           // 参与人是否为日程组织者
	IsExternal            bool                                                         `json:"is_external,omitempty"`            // 参与人是否为外部参与人；外部参与人不支持编辑
	DisplayName           string                                                       `json:"display_name,omitempty"`           // 参与人名称
	ChatMembers           []*GetCalendarEventAttendeeListRespItemChatMember            `json:"chat_members,omitempty"`           // 群中的群成员, 当type为Chat时有效；群成员不支持编辑
	UserID                string                                                       `json:"user_id,omitempty"`                // 参与人的用户id, 依赖于user_id_type返回对应的取值, 当is_external为true时, 此字段只会返回open_id或者union_id, 参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	ChatID                string                                                       `json:"chat_id,omitempty"`                // chat类型参与人的群组chat_id, 参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description)
	RoomID                string                                                       `json:"room_id,omitempty"`                // resource类型参与人的会议室room_id
	ThirdPartyEmail       string                                                       `json:"third_party_email,omitempty"`      // third_party类型参与人的邮箱
	OperateID             string                                                       `json:"operate_id,omitempty"`             // 如果日程是使用应用身份创建的, 在添加会议室的时候, 用来指定会议室的联系人, 在会议室视图展示。参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	ResourceCustomization []*GetCalendarEventAttendeeListRespItemResourceCustomization `json:"resource_customization,omitempty"` // 会议室的个性化配置
}

// GetCalendarEventAttendeeListRespItemChatMember ...
type GetCalendarEventAttendeeListRespItemChatMember struct {
	RsvpStatus  string `json:"rsvp_status,omitempty"`  // 参与人RSVP状态, 可选值有: needs_action: 参与人尚未回复状态, 或表示会议室预约中, accept: 参与人回复接受, 或表示会议室预约成功, tentative: 参与人回复待定, decline: 参与人回复拒绝, 或表示会议室预约失败, removed: 参与人或会议室已经从日程中被移除
	IsOptional  bool   `json:"is_optional,omitempty"`  // 参与人是否为「可选参加」
	DisplayName string `json:"display_name,omitempty"` // 参与人名称
	IsOrganizer bool   `json:"is_organizer,omitempty"` // 参与人是否为日程组织者
	IsExternal  bool   `json:"is_external,omitempty"`  // 参与人是否为外部参与人
}

// GetCalendarEventAttendeeListRespItemResourceCustomization ...
type GetCalendarEventAttendeeListRespItemResourceCustomization struct {
	IndexKey     string                                                             `json:"index_key,omitempty"`     // 每个配置的唯一ID
	InputContent string                                                             `json:"input_content,omitempty"` // 当type类型为填空时, 该参数需要填入
	Options      []*GetCalendarEventAttendeeListRespItemResourceCustomizationOption `json:"options,omitempty"`       // 每个配置的选项
}

// GetCalendarEventAttendeeListRespItemResourceCustomizationOption ...
type GetCalendarEventAttendeeListRespItemResourceCustomizationOption struct {
	OptionKey     string `json:"option_key,omitempty"`     // 每个选项的唯一ID
	OthersContent string `json:"others_content,omitempty"` // 当type类型为其它选项时, 该参数需要填入
}

// getCalendarEventAttendeeListResp ...
type getCalendarEventAttendeeListResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarEventAttendeeListResp `json:"data,omitempty"`
}