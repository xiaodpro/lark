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

// ReplyCalendarEvent 调用该接口以当前身份（应用或用户）回复日程。
//
// 当前身份由 Header Authorization 的 Token 类型决定。tenant_access_token 指应用身份, user_access_token 指用户身份。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/reply
func (r *CalendarService) ReplyCalendarEvent(ctx context.Context, request *ReplyCalendarEventReq, options ...MethodOptionFunc) (*ReplyCalendarEventResp, *Response, error) {
	if r.cli.mock.mockCalendarReplyCalendarEvent != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Calendar#ReplyCalendarEvent mock enable")
		return r.cli.mock.mockCalendarReplyCalendarEvent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "ReplyCalendarEvent",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/reply",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(replyCalendarEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarReplyCalendarEvent mock CalendarReplyCalendarEvent method
func (r *Mock) MockCalendarReplyCalendarEvent(f func(ctx context.Context, request *ReplyCalendarEventReq, options ...MethodOptionFunc) (*ReplyCalendarEventResp, *Response, error)) {
	r.mockCalendarReplyCalendarEvent = f
}

// UnMockCalendarReplyCalendarEvent un-mock CalendarReplyCalendarEvent method
func (r *Mock) UnMockCalendarReplyCalendarEvent() {
	r.mockCalendarReplyCalendarEvent = nil
}

// ReplyCalendarEventReq ...
type ReplyCalendarEventReq struct {
	CalendarID string `path:"calendar_id" json:"-"`  // 日程所在的日历 ID。了解更多, 参见[日历 ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值: "feishu.cn_HF9U2MbibE8PPpjro6xjqa@group.calendar.feishu.cn"
	EventID    string `path:"event_id" json:"-"`     // 日程 ID, 创建日程时会返回日程 ID。你也可以调用以下接口获取某一日历的 ID, [获取日程列表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/list), [搜索日程](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/search), 示例值: "75d28f9b-e35c-4230-8a83-4a661497db54_0"
	RsvpStatus string `json:"rsvp_status,omitempty"` // 日程参与人 RSVP 状态, 即日程回复状态, 示例值: "accept", 可选值有: accept: 接受, decline: 拒绝, tentative: 待定
}

// ReplyCalendarEventResp ...
type ReplyCalendarEventResp struct {
}

// replyCalendarEventResp ...
type replyCalendarEventResp struct {
	Code  int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                  `json:"msg,omitempty"`  // 错误描述
	Data  *ReplyCalendarEventResp `json:"data,omitempty"`
	Error *ErrorDetail            `json:"error,omitempty"`
}