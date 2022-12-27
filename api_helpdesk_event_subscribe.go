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

// SubscribeHelpdeskEvent 本接口用于订阅服务台事件。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/event/subscribe
func (r *HelpdeskService) SubscribeHelpdeskEvent(ctx context.Context, request *SubscribeHelpdeskEventReq, options ...MethodOptionFunc) (*SubscribeHelpdeskEventResp, *Response, error) {
	if r.cli.mock.mockHelpdeskSubscribeHelpdeskEvent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#SubscribeHelpdeskEvent mock enable")
		return r.cli.mock.mockHelpdeskSubscribeHelpdeskEvent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "SubscribeHelpdeskEvent",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/events/subscribe",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(subscribeHelpdeskEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskSubscribeHelpdeskEvent mock HelpdeskSubscribeHelpdeskEvent method
func (r *Mock) MockHelpdeskSubscribeHelpdeskEvent(f func(ctx context.Context, request *SubscribeHelpdeskEventReq, options ...MethodOptionFunc) (*SubscribeHelpdeskEventResp, *Response, error)) {
	r.mockHelpdeskSubscribeHelpdeskEvent = f
}

// UnMockHelpdeskSubscribeHelpdeskEvent un-mock HelpdeskSubscribeHelpdeskEvent method
func (r *Mock) UnMockHelpdeskSubscribeHelpdeskEvent() {
	r.mockHelpdeskSubscribeHelpdeskEvent = nil
}

// SubscribeHelpdeskEventReq ...
type SubscribeHelpdeskEventReq struct {
	Events []*SubscribeHelpdeskEventReqEvent `json:"events,omitempty"` // 可订阅的事件列表
}

// SubscribeHelpdeskEventReqEvent ...
type SubscribeHelpdeskEventReqEvent struct {
	Type    string `json:"type,omitempty"`    // 事件类型, 示例值: "helpdesk.ticket_message"
	Subtype string `json:"subtype,omitempty"` // 事件子类型, 示例值: "ticket_message.created_v1"
}

// SubscribeHelpdeskEventResp ...
type SubscribeHelpdeskEventResp struct {
}

// subscribeHelpdeskEventResp ...
type subscribeHelpdeskEventResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *SubscribeHelpdeskEventResp `json:"data,omitempty"`
}
