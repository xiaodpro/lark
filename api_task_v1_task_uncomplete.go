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

// UncompleteTaskV1 该接口用于取消任务的已完成状态。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/uncomplete
// new doc: https://open.feishu.cn/document/server-docs/task-v1/task/uncomplete
//
// Deprecated
func (r *TaskV1Service) UncompleteTaskV1(ctx context.Context, request *UncompleteTaskV1Req, options ...MethodOptionFunc) (*UncompleteTaskV1Resp, *Response, error) {
	if r.cli.mock.mockTaskV1UncompleteTaskV1 != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] TaskV1#UncompleteTaskV1 mock enable")
		return r.cli.mock.mockTaskV1UncompleteTaskV1(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "TaskV1",
		API:                   "UncompleteTaskV1",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/uncomplete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(uncompleteTaskV1Resp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskV1UncompleteTaskV1 mock TaskV1UncompleteTaskV1 method
func (r *Mock) MockTaskV1UncompleteTaskV1(f func(ctx context.Context, request *UncompleteTaskV1Req, options ...MethodOptionFunc) (*UncompleteTaskV1Resp, *Response, error)) {
	r.mockTaskV1UncompleteTaskV1 = f
}

// UnMockTaskV1UncompleteTaskV1 un-mock TaskV1UncompleteTaskV1 method
func (r *Mock) UnMockTaskV1UncompleteTaskV1() {
	r.mockTaskV1UncompleteTaskV1 = nil
}

// UncompleteTaskV1Req ...
type UncompleteTaskV1Req struct {
	TaskID string `path:"task_id" json:"-"` // 任务 ID, 示例值: "bb54ab99-d360-434f-bcaa-a4cc4c05840e"
}

// UncompleteTaskV1Resp ...
type UncompleteTaskV1Resp struct {
}

// uncompleteTaskV1Resp ...
type uncompleteTaskV1Resp struct {
	Code  int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                `json:"msg,omitempty"`  // 错误描述
	Data  *UncompleteTaskV1Resp `json:"data,omitempty"`
	Error *ErrorDetail          `json:"error,omitempty"`
}
