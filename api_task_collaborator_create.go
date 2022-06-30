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

// CreateTaskCollaborator 该接口用于新增任务执行者, 一次性可以添加多个执行者。新增的执行者必须是表示是用户的ID。
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
	TaskID     string   `path:"task_id" json:"-"`       // 任务 ID, 示例值: "83912691-2e43-47fc-94a4-d512e03984fa"
	UserIDType *IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: <md-enum>, <md-enum-item key="open_id" >用户的 open id</md-enum-item>, <md-enum-item key="union_id" >用户的 union id</md-enum-item>, <md-enum-item key="user_id" >用户的 user id</md-enum-item>, </md-enum>, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ID         *string  `json:"id,omitempty"`           // 任务执行者的 ID, 示例值: "ou_99e1a581b36ecc4862cbfbce473f1234"
	IDList     []string `json:"id_list,omitempty"`      // 执行者的用户ID列表, 示例值: ["ou_550cc75233d8b7b9fcbdad65f34433f4", "ou_d1e9d27cf3235b40ca9a67c67ef088b0"]
}

// CreateTaskCollaboratorResp ...
type CreateTaskCollaboratorResp struct {
	Collaborator *CreateTaskCollaboratorRespCollaborator `json:"collaborator,omitempty"` // 返回创建成功后的任务执行者列表
}

// CreateTaskCollaboratorRespCollaborator ...
type CreateTaskCollaboratorRespCollaborator struct {
	ID     string   `json:"id,omitempty"`      // 任务执行者的 ID
	IDList []string `json:"id_list,omitempty"` // 执行者的用户ID列表。
}

// createTaskCollaboratorResp ...
type createTaskCollaboratorResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *CreateTaskCollaboratorResp `json:"data,omitempty"`
}
