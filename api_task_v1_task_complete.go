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

// CompleteTaskV1 该接口用于将任务状态修改为“已完成”。
//
// 完成任务是指整个任务全部完成, 而不支持执行者分别完成任务, 执行成功后, 任务对所有关联用户都变为完成状态。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/complete
// new doc: https://open.feishu.cn/document/server-docs/task-v1/task/complete
//
// Deprecated
func (r *TaskV1Service) CompleteTaskV1(ctx context.Context, request *CompleteTaskV1Req, options ...MethodOptionFunc) (*CompleteTaskV1Resp, *Response, error) {
	if r.cli.mock.mockTaskV1CompleteTaskV1 != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] TaskV1#CompleteTaskV1 mock enable")
		return r.cli.mock.mockTaskV1CompleteTaskV1(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "TaskV1",
		API:                   "CompleteTaskV1",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/complete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(completeTaskV1Resp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskV1CompleteTaskV1 mock TaskV1CompleteTaskV1 method
func (r *Mock) MockTaskV1CompleteTaskV1(f func(ctx context.Context, request *CompleteTaskV1Req, options ...MethodOptionFunc) (*CompleteTaskV1Resp, *Response, error)) {
	r.mockTaskV1CompleteTaskV1 = f
}

// UnMockTaskV1CompleteTaskV1 un-mock TaskV1CompleteTaskV1 method
func (r *Mock) UnMockTaskV1CompleteTaskV1() {
	r.mockTaskV1CompleteTaskV1 = nil
}

// CompleteTaskV1Req ...
type CompleteTaskV1Req struct {
	TaskID string `path:"task_id" json:"-"` // 任务 ID, 可通过[创建任务](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/create)时响应体中的id字段获取, 示例值: "bb54ab99-d360-434f-bcaa-a4cc4c05840e"
}

// CompleteTaskV1Resp ...
type CompleteTaskV1Resp struct {
}

// completeTaskV1Resp ...
type completeTaskV1Resp struct {
	Code  int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string              `json:"msg,omitempty"`  // 错误描述
	Data  *CompleteTaskV1Resp `json:"data,omitempty"`
	Error *ErrorDetail        `json:"error,omitempty"`
}
