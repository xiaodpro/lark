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

// CreateTaskCollaborator 该接口用于新增任务执行者, 一次性可以添加多个执行者。
//
// 只有任务的创建者和执行者才能添加执行者, 关注人无权限添加。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-collaborator/create
func (r *TaskService) CreateTaskCollaborator(ctx context.Context, request *CreateTaskCollaboratorReq, options ...MethodOptionFunc) (*CreateTaskCollaboratorResp, *Response, error) {
	if r.cli.mock.mockTaskCreateTaskCollaborator != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#CreateTaskCollaborator mock enable")
		return r.cli.mock.mockTaskCreateTaskCollaborator(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "CreateTaskCollaborator",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/collaborators",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createTaskCollaboratorResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskCreateTaskCollaborator mock TaskCreateTaskCollaborator method
func (r *Mock) MockTaskCreateTaskCollaborator(f func(ctx context.Context, request *CreateTaskCollaboratorReq, options ...MethodOptionFunc) (*CreateTaskCollaboratorResp, *Response, error)) {
	r.mockTaskCreateTaskCollaborator = f
}

// UnMockTaskCreateTaskCollaborator un-mock TaskCreateTaskCollaborator method
func (r *Mock) UnMockTaskCreateTaskCollaborator() {
	r.mockTaskCreateTaskCollaborator = nil
}

// CreateTaskCollaboratorReq ...
type CreateTaskCollaboratorReq struct {
	TaskID     string   `path:"task_id" json:"-"`       // 任务 ID, 可通过[创建任务](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/create)时响应体中的id字段获取, 示例值: "83912691-2e43-47fc-94a4-d512e03984fa"
	UserIDType *IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ID         *string  `json:"id,omitempty"`           // 任务执行者的 ID, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档[如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 已经废弃, 为了向前兼容早期只支持单次添加一个人的情况而保留, 但不再推荐使用, 建议使用id_list字段, 示例值: "ou_99e1a581b36ecc4862cbfbce473f1234"
	IDList     []string `json:"id_list,omitempty"`      // 执行者的用户ID列表, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档[如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 示例值: ["ou_550cc75233d8b7b9fcbdad65f34433f4", "ou_d1e9d27cf3235b40ca9a67c67ef088b0"]
}

// CreateTaskCollaboratorResp ...
type CreateTaskCollaboratorResp struct {
	Collaborator *CreateTaskCollaboratorRespCollaborator `json:"collaborator,omitempty"` // 返回创建成功后的任务执行者列表
}

// CreateTaskCollaboratorRespCollaborator ...
type CreateTaskCollaboratorRespCollaborator struct {
	ID     string   `json:"id,omitempty"`      // 任务执行者的 ID, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档[如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 已经废弃, 为了向前兼容早期只支持单次添加一个人的情况而保留, 但不再推荐使用, 建议使用id_list字段
	IDList []string `json:"id_list,omitempty"` // 执行者的用户ID列表, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档[如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get)。
}

// createTaskCollaboratorResp ...
type createTaskCollaboratorResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *CreateTaskCollaboratorResp `json:"data,omitempty"`
}
