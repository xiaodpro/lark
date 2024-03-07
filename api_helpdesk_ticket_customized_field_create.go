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

// CreateHelpdeskTicketCustomizedField 该接口用于创建自定义字段。
//
// 注意事项:
// user_access_token 访问, 需要操作者是当前服务台的管理员或所有者
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/create-ticket-customized-field
// new doc: https://open.feishu.cn/document/server-docs/helpdesk-v1/ticket-management/ticket_customized_field/create-ticket-customized-field
func (r *HelpdeskService) CreateHelpdeskTicketCustomizedField(ctx context.Context, request *CreateHelpdeskTicketCustomizedFieldReq, options ...MethodOptionFunc) (*CreateHelpdeskTicketCustomizedFieldResp, *Response, error) {
	if r.cli.mock.mockHelpdeskCreateHelpdeskTicketCustomizedField != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Helpdesk#CreateHelpdeskTicketCustomizedField mock enable")
		return r.cli.mock.mockHelpdeskCreateHelpdeskTicketCustomizedField(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "CreateHelpdeskTicketCustomizedField",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/ticket_customized_fields",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(createHelpdeskTicketCustomizedFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskCreateHelpdeskTicketCustomizedField mock HelpdeskCreateHelpdeskTicketCustomizedField method
func (r *Mock) MockHelpdeskCreateHelpdeskTicketCustomizedField(f func(ctx context.Context, request *CreateHelpdeskTicketCustomizedFieldReq, options ...MethodOptionFunc) (*CreateHelpdeskTicketCustomizedFieldResp, *Response, error)) {
	r.mockHelpdeskCreateHelpdeskTicketCustomizedField = f
}

// UnMockHelpdeskCreateHelpdeskTicketCustomizedField un-mock HelpdeskCreateHelpdeskTicketCustomizedField method
func (r *Mock) UnMockHelpdeskCreateHelpdeskTicketCustomizedField() {
	r.mockHelpdeskCreateHelpdeskTicketCustomizedField = nil
}

// CreateHelpdeskTicketCustomizedFieldReq ...
type CreateHelpdeskTicketCustomizedFieldReq struct {
	HelpdeskID            string                  `json:"helpdesk_id,omitempty"`             // 服务台ID, 示例值: "1542164574896126"
	KeyName               string                  `json:"key_name,omitempty"`                // 键名, 示例值: "test dropdown"
	DisplayName           string                  `json:"display_name,omitempty"`            // 名称, 示例值: "test dropdown"
	Position              string                  `json:"position,omitempty"`                // 字段在列表后台管理列表中的位置, 示例值: "3"
	FieldType             string                  `json:"field_type,omitempty"`              // 类型, string - 单行文本, multiline - 多行文本, dropdown - 下拉列表, dropdown_nested - 级联下拉, 示例值: "dropdown"
	Description           string                  `json:"description,omitempty"`             // 描述, 示例值: "下拉示例"
	Visible               bool                    `json:"visible,omitempty"`                 // 是否可见, 示例值: true
	Editable              bool                    `json:"editable,omitempty"`                // 是否可以修改, 示例值: true
	Required              bool                    `json:"required,omitempty"`                // 是否必填, 示例值: false
	DropdownOptions       *HelpdeskDropdownOption `json:"dropdown_options,omitempty"`        // 下拉列表选项
	DropdownAllowMultiple *bool                   `json:"dropdown_allow_multiple,omitempty"` // 是否支持多选, 仅在字段类型是dropdown的时候有效, 示例值: true
}

// CreateHelpdeskTicketCustomizedFieldResp ...
type CreateHelpdeskTicketCustomizedFieldResp struct {
}

// createHelpdeskTicketCustomizedFieldResp ...
type createHelpdeskTicketCustomizedFieldResp struct {
	Code int64                                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                   `json:"msg,omitempty"`  // 错误描述
	Data *CreateHelpdeskTicketCustomizedFieldResp `json:"data,omitempty"`
}
