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

// DeleteApprovalComment 逻辑删除某审批实例下的一条评论或评论回复（不包含审批同意、拒绝、转交等附加的理由或意见）。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance-comment/delete
func (r *ApprovalService) DeleteApprovalComment(ctx context.Context, request *DeleteApprovalCommentReq, options ...MethodOptionFunc) (*DeleteApprovalCommentResp, *Response, error) {
	if r.cli.mock.mockApprovalDeleteApprovalComment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#DeleteApprovalComment mock enable")
		return r.cli.mock.mockApprovalDeleteApprovalComment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "DeleteApprovalComment",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/instances/:instance_id/comments/:comment_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteApprovalCommentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalDeleteApprovalComment mock ApprovalDeleteApprovalComment method
func (r *Mock) MockApprovalDeleteApprovalComment(f func(ctx context.Context, request *DeleteApprovalCommentReq, options ...MethodOptionFunc) (*DeleteApprovalCommentResp, *Response, error)) {
	r.mockApprovalDeleteApprovalComment = f
}

// UnMockApprovalDeleteApprovalComment un-mock ApprovalDeleteApprovalComment method
func (r *Mock) UnMockApprovalDeleteApprovalComment() {
	r.mockApprovalDeleteApprovalComment = nil
}

// DeleteApprovalCommentReq ...
type DeleteApprovalCommentReq struct {
	InstanceID string  `path:"instance_id" json:"-"`   // 审批实例code（或者租户自定义审批实例ID）, 示例值: "6A123516-FB88-470D-A428-9AF58B71B3C0"
	CommentID  string  `path:"comment_id" json:"-"`    // 评论ID, 示例值: "7081516627711606803"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	UserID     string  `query:"user_id" json:"-"`      // 根据user_id_type填写用户ID, 示例值: "ou_806a18fb5bdf525e38ba219733bdbd73"
}

// DeleteApprovalCommentResp ...
type DeleteApprovalCommentResp struct {
	CommentID string `json:"comment_id,omitempty"` // 删除的评论ID
}

// deleteApprovalCommentResp ...
type deleteApprovalCommentResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *DeleteApprovalCommentResp `json:"data,omitempty"`
}
