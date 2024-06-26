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

// DeleteTask 删除一个任务。
//
// 删除后任务无法再被获取到。
// 删除任务需要任务的可编辑权限。详情见[任务功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/overview)中的“任务是如何鉴权的？”章节。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/delete
func (r *TaskService) DeleteTask(ctx context.Context, request *DeleteTaskReq, options ...MethodOptionFunc) (*DeleteTaskResp, *Response, error) {
	if r.cli.mock.mockTaskDeleteTask != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#DeleteTask mock enable")
		return r.cli.mock.mockTaskDeleteTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "DeleteTask",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/tasks/:task_guid",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskDeleteTask mock TaskDeleteTask method
func (r *Mock) MockTaskDeleteTask(f func(ctx context.Context, request *DeleteTaskReq, options ...MethodOptionFunc) (*DeleteTaskResp, *Response, error)) {
	r.mockTaskDeleteTask = f
}

// UnMockTaskDeleteTask un-mock TaskDeleteTask method
func (r *Mock) UnMockTaskDeleteTask() {
	r.mockTaskDeleteTask = nil
}

// DeleteTaskReq ...
type DeleteTaskReq struct {
	TaskGuid string `path:"task_guid" json:"-"` // 要删除的任务guid, 示例值: "e297ddff-06ca-4166-b917-4ce57cd3a7a0", 最大长度: `100` 字符
}

// DeleteTaskResp ...
type DeleteTaskResp struct {
}

// deleteTaskResp ...
type deleteTaskResp struct {
	Code  int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string          `json:"msg,omitempty"`  // 错误描述
	Data  *DeleteTaskResp `json:"data,omitempty"`
	Error *ErrorDetail    `json:"error,omitempty"`
}
