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

// BatchDeleteContactGroupMember 从普通用户组中批量移除成员 (目前仅支持移除用户, 暂不支持移除部门）。如果应用的通讯录权限范围是“全部员工”, 则可将任何成员移出任何用户组。如果应用的通讯录权限范围不是“全部员工”, 则仅可将通讯录权限范围中的成员从通讯录权限范围的用户组中移除, [点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。
//
// 请求体中的member_type, 目前仅支持user, 未来将支持department。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/batch_remove
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/group-member/batch_remove
func (r *ContactService) BatchDeleteContactGroupMember(ctx context.Context, request *BatchDeleteContactGroupMemberReq, options ...MethodOptionFunc) (*BatchDeleteContactGroupMemberResp, *Response, error) {
	if r.cli.mock.mockContactBatchDeleteContactGroupMember != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Contact#BatchDeleteContactGroupMember mock enable")
		return r.cli.mock.mockContactBatchDeleteContactGroupMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "BatchDeleteContactGroupMember",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/group/:group_id/member/batch_remove",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchDeleteContactGroupMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactBatchDeleteContactGroupMember mock ContactBatchDeleteContactGroupMember method
func (r *Mock) MockContactBatchDeleteContactGroupMember(f func(ctx context.Context, request *BatchDeleteContactGroupMemberReq, options ...MethodOptionFunc) (*BatchDeleteContactGroupMemberResp, *Response, error)) {
	r.mockContactBatchDeleteContactGroupMember = f
}

// UnMockContactBatchDeleteContactGroupMember un-mock ContactBatchDeleteContactGroupMember method
func (r *Mock) UnMockContactBatchDeleteContactGroupMember() {
	r.mockContactBatchDeleteContactGroupMember = nil
}

// BatchDeleteContactGroupMemberReq ...
type BatchDeleteContactGroupMemberReq struct {
	GroupID string                                    `path:"group_id" json:"-"` // 用户组ID, 示例值: "test_group"
	Members []*BatchDeleteContactGroupMemberReqMember `json:"members,omitempty"` // 待移除成员, 长度范围: `1` ～ `100`
}

// BatchDeleteContactGroupMemberReqMember ...
type BatchDeleteContactGroupMemberReqMember struct {
	MemberID     string  `json:"member_id,omitempty"`      // 成员ID, 示例值: "u287xj12"
	MemberType   string  `json:"member_type,omitempty"`    // 用户组成员的类型, 取值为 user或department, 示例值: "user"
	MemberIDType *IDType `json:"member_id_type,omitempty"` // 当member_type为user时, member_id_type表示user_id_type, 可选值为open_id, union_id, user_id。仅在请求参数中有效, 响应体中不会返回此参数, 示例值: "user_id"
}

// BatchDeleteContactGroupMemberResp ...
type BatchDeleteContactGroupMemberResp struct {
}

// batchDeleteContactGroupMemberResp ...
type batchDeleteContactGroupMemberResp struct {
	Code  int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                             `json:"msg,omitempty"`  // 错误描述
	Data  *BatchDeleteContactGroupMemberResp `json:"data,omitempty"`
	Error *ErrorDetail                       `json:"error,omitempty"`
}