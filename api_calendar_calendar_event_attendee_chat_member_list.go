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

// GetCalendarEventAttendeeChatMemberList 获取日程的群参与人的群成员列表。
//
// - 当前身份必须有权限查看日程的参与人列表。
// - 当前身份必须在群聊中, 或有权限查看群成员列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee-chat_member/list
func (r *CalendarService) GetCalendarEventAttendeeChatMemberList(ctx context.Context, request *GetCalendarEventAttendeeChatMemberListReq, options ...MethodOptionFunc) (*GetCalendarEventAttendeeChatMemberListResp, *Response, error) {
	if r.cli.mock.mockCalendarGetCalendarEventAttendeeChatMemberList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#GetCalendarEventAttendeeChatMemberList mock enable")
		return r.cli.mock.mockCalendarGetCalendarEventAttendeeChatMemberList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "GetCalendarEventAttendeeChatMemberList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/:attendee_id/chat_members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getCalendarEventAttendeeChatMemberListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarGetCalendarEventAttendeeChatMemberList mock CalendarGetCalendarEventAttendeeChatMemberList method
func (r *Mock) MockCalendarGetCalendarEventAttendeeChatMemberList(f func(ctx context.Context, request *GetCalendarEventAttendeeChatMemberListReq, options ...MethodOptionFunc) (*GetCalendarEventAttendeeChatMemberListResp, *Response, error)) {
	r.mockCalendarGetCalendarEventAttendeeChatMemberList = f
}

// UnMockCalendarGetCalendarEventAttendeeChatMemberList un-mock CalendarGetCalendarEventAttendeeChatMemberList method
func (r *Mock) UnMockCalendarGetCalendarEventAttendeeChatMemberList() {
	r.mockCalendarGetCalendarEventAttendeeChatMemberList = nil
}

// GetCalendarEventAttendeeChatMemberListReq ...
type GetCalendarEventAttendeeChatMemberListReq struct {
	CalendarID string  `path:"calendar_id" json:"-"`   // 日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值: "feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	EventID    string  `path:"event_id" json:"-"`      // 日程ID。参见[日程ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/introduction), 示例值: "xxxxxxxxx_0"
	AttendeeID string  `path:"attendee_id" json:"-"`   // 群参与人 ID。参见[参与人ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/introduction#4998889c), 示例值: "chat_xxxxxx"
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "23jhysaxxxxsysy"
	PageSize   *int64  `query:"page_size" json:"-"`    // 分页大小, 示例值: 10, 最大值: `100`
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetCalendarEventAttendeeChatMemberListResp ...
type GetCalendarEventAttendeeChatMemberListResp struct {
	Items     []*GetCalendarEventAttendeeChatMemberListRespItem `json:"items,omitempty"`      // 群中的群成员, 当type为chat时有效；群成员不支持编辑
	HasMore   bool                                              `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                                            `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetCalendarEventAttendeeChatMemberListRespItem ...
type GetCalendarEventAttendeeChatMemberListRespItem struct {
	RsvpStatus  string `json:"rsvp_status,omitempty"`  // 参与人RSVP状态, 可选值有: needs_action: 参与人尚未回复状态, 或表示会议室预约中, accept: 参与人回复接受, 或表示会议室预约成功, tentative: 参与人回复待定, decline: 参与人回复拒绝, 或表示会议室预约失败, removed: 参与人或会议室已经从日程中被移除
	IsOptional  bool   `json:"is_optional,omitempty"`  // 参与人是否为「可选参加」
	DisplayName string `json:"display_name,omitempty"` // 参与人名称
	OpenID      string `json:"open_id,omitempty"`      // 参与人open_id, 参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 示例值: "ou_xxxxxxxx"
	IsOrganizer bool   `json:"is_organizer,omitempty"` // 参与人是否为日程组织者
	IsExternal  bool   `json:"is_external,omitempty"`  // 参与人是否为外部参与人
}

// getCalendarEventAttendeeChatMemberListResp ...
type getCalendarEventAttendeeChatMemberListResp struct {
	Code int64                                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                      `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarEventAttendeeChatMemberListResp `json:"data,omitempty"`
}
