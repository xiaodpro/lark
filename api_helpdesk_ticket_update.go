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

// UpdateHelpdeskTicket 该接口用于更新服务台工单详情。只会更新数据, 不会触发相关操作。如修改工单状态到关单, 不会关闭聊天页面。仅支持自建应用。要更新的工单字段必须至少输入一项。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/update
func (r *HelpdeskService) UpdateHelpdeskTicket(ctx context.Context, request *UpdateHelpdeskTicketReq, options ...MethodOptionFunc) (*UpdateHelpdeskTicketResp, *Response, error) {
	if r.cli.mock.mockHelpdeskUpdateHelpdeskTicket != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#UpdateHelpdeskTicket mock enable")
		return r.cli.mock.mockHelpdeskUpdateHelpdeskTicket(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "UpdateHelpdeskTicket",
		Method:              "PUT",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/tickets/:ticket_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(updateHelpdeskTicketResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskUpdateHelpdeskTicket mock HelpdeskUpdateHelpdeskTicket method
func (r *Mock) MockHelpdeskUpdateHelpdeskTicket(f func(ctx context.Context, request *UpdateHelpdeskTicketReq, options ...MethodOptionFunc) (*UpdateHelpdeskTicketResp, *Response, error)) {
	r.mockHelpdeskUpdateHelpdeskTicket = f
}

// UnMockHelpdeskUpdateHelpdeskTicket un-mock HelpdeskUpdateHelpdeskTicket method
func (r *Mock) UnMockHelpdeskUpdateHelpdeskTicket() {
	r.mockHelpdeskUpdateHelpdeskTicket = nil
}

// UpdateHelpdeskTicketReq ...
type UpdateHelpdeskTicketReq struct {
	TicketID         string                                    `path:"ticket_id" json:"-"`          // 工单ID, 示例值: "6945345902185807891"
	Status           *int64                                    `json:"status,omitempty"`            // new status, 1: 已创建, 2: 处理中, 3: 排队中, 5: 待定, 50: 机器人关闭工单, 51: 关闭工单, 示例值: 1
	TagNames         []string                                  `json:"tag_names,omitempty"`         // 新标签名
	Comment          *string                                   `json:"comment,omitempty"`           // 新评论, 示例值: "good"
	CustomizedFields []*UpdateHelpdeskTicketReqCustomizedField `json:"customized_fields,omitempty"` // 自定义字段
	TicketType       *int64                                    `json:"ticket_type,omitempty"`       // ticket stage, 示例值: 1
	Solved           *int64                                    `json:"solved,omitempty"`            // 工单是否解决, 1: 未解决, 2: 已解决, 示例值: 1
	Channel          *int64                                    `json:"channel,omitempty"`           // 工单来源渠道ID, 示例值: 1
}

// UpdateHelpdeskTicketReqCustomizedField ...
type UpdateHelpdeskTicketReqCustomizedField struct {
	ID      *string `json:"id,omitempty"`       // 自定义字段ID, 示例值: "123"
	Value   *string `json:"value,omitempty"`    // 自定义字段值, 示例值: "value"
	KeyName *string `json:"key_name,omitempty"` // 键名, 示例值: "key"
}

// UpdateHelpdeskTicketResp ...
type UpdateHelpdeskTicketResp struct {
}

// updateHelpdeskTicketResp ...
type updateHelpdeskTicketResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *UpdateHelpdeskTicketResp `json:"data,omitempty"`
}