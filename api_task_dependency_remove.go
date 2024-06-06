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

// RemoveTaskDependency 从一个任务移除一个或者多个依赖。移除时只需要输入要移除的`task_guid`即可。
//
// 注意, 如果要移除的依赖非当前任务的依赖, 会被自动忽略。接口会返回成功。
// 移除任务依赖时, 需要当前任务的编辑权限。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/remove_dependencies
func (r *TaskService) RemoveTaskDependency(ctx context.Context, request *RemoveTaskDependencyReq, options ...MethodOptionFunc) (*RemoveTaskDependencyResp, *Response, error) {
	if r.cli.mock.mockTaskRemoveTaskDependency != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#RemoveTaskDependency mock enable")
		return r.cli.mock.mockTaskRemoveTaskDependency(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "RemoveTaskDependency",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/tasks/:task_guid/remove_dependencies",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(removeTaskDependencyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskRemoveTaskDependency mock TaskRemoveTaskDependency method
func (r *Mock) MockTaskRemoveTaskDependency(f func(ctx context.Context, request *RemoveTaskDependencyReq, options ...MethodOptionFunc) (*RemoveTaskDependencyResp, *Response, error)) {
	r.mockTaskRemoveTaskDependency = f
}

// UnMockTaskRemoveTaskDependency un-mock TaskRemoveTaskDependency method
func (r *Mock) UnMockTaskRemoveTaskDependency() {
	r.mockTaskRemoveTaskDependency = nil
}

// RemoveTaskDependencyReq ...
type RemoveTaskDependencyReq struct {
	TaskGuid     string                                `path:"task_guid" json:"-"`     // 要移除依赖的任务GUID, 示例值: "93b7bd05-35e6-4371-b3c9-6b7cbd7100c0"
	Dependencies []*RemoveTaskDependencyReqDependencie `json:"dependencies,omitempty"` // 要移除的依赖, 长度范围: `1` ～ `50`
}

// RemoveTaskDependencyReqDependencie ...
type RemoveTaskDependencyReqDependencie struct {
	TaskGuid string `json:"task_guid,omitempty"` // 依赖任务的GUID, 示例值: "93b7bd05-35e6-4371-b3c9-6b7cbd7100c0"
}

// RemoveTaskDependencyResp ...
type RemoveTaskDependencyResp struct {
	Dependencies []*RemoveTaskDependencyRespDependencie `json:"dependencies,omitempty"` // 移除之后的任务GUID
}

// RemoveTaskDependencyRespDependencie ...
type RemoveTaskDependencyRespDependencie struct {
	Type     string `json:"type,omitempty"`      // 依赖类型, 可选值有: prev: 前置依赖, next: 后置依赖
	TaskGuid string `json:"task_guid,omitempty"` // 依赖任务的GUID
}

// removeTaskDependencyResp ...
type removeTaskDependencyResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *RemoveTaskDependencyResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
