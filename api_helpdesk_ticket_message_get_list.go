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

// GetHelpdeskTicketMessageList 该接口用于获取服务台工单消息详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket-message/list
// new doc: https://open.feishu.cn/document/server-docs/helpdesk-v1/ticket-management/ticket-message/list
func (r *HelpdeskService) GetHelpdeskTicketMessageList(ctx context.Context, request *GetHelpdeskTicketMessageListReq, options ...MethodOptionFunc) (*GetHelpdeskTicketMessageListResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetHelpdeskTicketMessageList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Helpdesk#GetHelpdeskTicketMessageList mock enable")
		return r.cli.mock.mockHelpdeskGetHelpdeskTicketMessageList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetHelpdeskTicketMessageList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/tickets/:ticket_id/messages",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getHelpdeskTicketMessageListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskGetHelpdeskTicketMessageList mock HelpdeskGetHelpdeskTicketMessageList method
func (r *Mock) MockHelpdeskGetHelpdeskTicketMessageList(f func(ctx context.Context, request *GetHelpdeskTicketMessageListReq, options ...MethodOptionFunc) (*GetHelpdeskTicketMessageListResp, *Response, error)) {
	r.mockHelpdeskGetHelpdeskTicketMessageList = f
}

// UnMockHelpdeskGetHelpdeskTicketMessageList un-mock HelpdeskGetHelpdeskTicketMessageList method
func (r *Mock) UnMockHelpdeskGetHelpdeskTicketMessageList() {
	r.mockHelpdeskGetHelpdeskTicketMessageList = nil
}

// GetHelpdeskTicketMessageListReq ...
type GetHelpdeskTicketMessageListReq struct {
	TicketID  string `path:"ticket_id" json:"-"`   // 工单ID, 示例值: "6948728206392295444"
	TimeStart *int64 `query:"time_start" json:"-"` // 起始时间, 示例值: 1617960686
	TimeEnd   *int64 `query:"time_end" json:"-"`   // 结束时间, 示例值: 1617960687
	Page      *int64 `query:"page" json:"-"`       // 页数ID, 示例值: 1
	PageSize  *int64 `query:"page_size" json:"-"`  // 消息数量, 最大200, 默认20, 示例值: 10
}

// GetHelpdeskTicketMessageListResp ...
type GetHelpdeskTicketMessageListResp struct {
	Messages []*GetHelpdeskTicketMessageListRespMessage `json:"messages,omitempty"` // 工单消息列表
	Total    int64                                      `json:"total,omitempty"`    // 消息总数
}

// GetHelpdeskTicketMessageListRespMessage ...
type GetHelpdeskTicketMessageListRespMessage struct {
	ID          string  `json:"id,omitempty"`           // 工单消息ID
	MessageID   string  `json:"message_id,omitempty"`   // chat消息ID
	MessageType MsgType `json:"message_type,omitempty"` // 消息类型；text: 纯文本；post: 富文本；image: 图像；file: 文件；media: 视频
	CreatedAt   int64   `json:"created_at,omitempty"`   // 创建时间
	Content     string  `json:"content,omitempty"`      // 内容
	UserName    string  `json:"user_name,omitempty"`    // 用户名
	AvatarURL   string  `json:"avatar_url,omitempty"`   // 用户图片url
	UserID      string  `json:"user_id,omitempty"`      // 用户open ID
}

// getHelpdeskTicketMessageListResp ...
type getHelpdeskTicketMessageListResp struct {
	Code  int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                            `json:"msg,omitempty"`  // 错误描述
	Data  *GetHelpdeskTicketMessageListResp `json:"data,omitempty"`
	Error *ErrorDetail                      `json:"error,omitempty"`
}
