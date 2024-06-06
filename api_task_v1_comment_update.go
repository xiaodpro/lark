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

// UpdateTaskV1Comment 该接口用于更新评论内容。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/update
// new doc: https://open.feishu.cn/document/server-docs/task-v1/task-comment/update
//
// Deprecated
func (r *TaskV1Service) UpdateTaskV1Comment(ctx context.Context, request *UpdateTaskV1CommentReq, options ...MethodOptionFunc) (*UpdateTaskV1CommentResp, *Response, error) {
	if r.cli.mock.mockTaskV1UpdateTaskV1Comment != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] TaskV1#UpdateTaskV1Comment mock enable")
		return r.cli.mock.mockTaskV1UpdateTaskV1Comment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "TaskV1",
		API:                   "UpdateTaskV1Comment",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/comments/:comment_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateTaskV1CommentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskV1UpdateTaskV1Comment mock TaskV1UpdateTaskV1Comment method
func (r *Mock) MockTaskV1UpdateTaskV1Comment(f func(ctx context.Context, request *UpdateTaskV1CommentReq, options ...MethodOptionFunc) (*UpdateTaskV1CommentResp, *Response, error)) {
	r.mockTaskV1UpdateTaskV1Comment = f
}

// UnMockTaskV1UpdateTaskV1Comment un-mock TaskV1UpdateTaskV1Comment method
func (r *Mock) UnMockTaskV1UpdateTaskV1Comment() {
	r.mockTaskV1UpdateTaskV1Comment = nil
}

// UpdateTaskV1CommentReq ...
type UpdateTaskV1CommentReq struct {
	TaskID      string  `path:"task_id" json:"-"`       // 任务ID, 示例值: "83912691-2e43-47fc-94a4-d512e03984fa"
	CommentID   string  `path:"comment_id" json:"-"`    // 评论 ID, 示例值: "6937231762296684564"
	UserIDType  *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Content     *string `json:"content,omitempty"`      // 新的评论内容, 示例值: "飞流直下三千尺, 疑是银河落九天", 长度范围: `0` ～ `65536` 字符
	RichContent *string `json:"rich_content,omitempty"` // 新的富文本评论内容（优先使用）, 示例值: "飞流直下三千尺, 疑是银河落九天<at id=7058204817822318612></at>", 长度范围: `0` ～ `65536` 字符
}

// UpdateTaskV1CommentResp ...
type UpdateTaskV1CommentResp struct {
	Comment *UpdateTaskV1CommentRespComment `json:"comment,omitempty"` // 返回修改后的任务评论详情
}

// UpdateTaskV1CommentRespComment ...
type UpdateTaskV1CommentRespComment struct {
	Content         string `json:"content,omitempty"`           // 评论内容, 评论内容和富文本评论内容同时存在时只使用富文本评论内容。
	ParentID        string `json:"parent_id,omitempty"`         // 评论的父ID, 创建评论时若不为空则为某条评论的回复, 若为空则不是回复
	ID              string `json:"id,omitempty"`                // 评论ID, 由飞书服务器发号
	CreateMilliTime string `json:"create_milli_time,omitempty"` // 评论创建的时间戳, 单位为毫秒, 用于展示, 创建时不用填写
	RichContent     string `json:"rich_content,omitempty"`      // 富文本评论内容。语法格式参见[Markdown模块](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/markdown-module)
	CreatorID       string `json:"creator_id,omitempty"`        // 评论的创建者 ID。在创建评论时无需填充该字段
}

// updateTaskV1CommentResp ...
type updateTaskV1CommentResp struct {
	Code  int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                   `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateTaskV1CommentResp `json:"data,omitempty"`
	Error *ErrorDetail             `json:"error,omitempty"`
}
