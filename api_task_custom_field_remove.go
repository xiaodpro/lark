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

// RemoveTaskCustomField 将自定义字段从资源中移出。移除后, 该资源将无法再使用该字段。目前资源的类型支持"tasklist"。
//
// 如果要移除自定义字段本来就不存在于资源, 本接口将正常返回。
// 注意自定义字段是通过清单来实现授权的, 如果将自定义字段从所有关联的清单中移除, 就意味着任何调用身份都无法再访问改自定义字段。
// 需要资源的可编辑权限。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field/remove
func (r *TaskService) RemoveTaskCustomField(ctx context.Context, request *RemoveTaskCustomFieldReq, options ...MethodOptionFunc) (*RemoveTaskCustomFieldResp, *Response, error) {
	if r.cli.mock.mockTaskRemoveTaskCustomField != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#RemoveTaskCustomField mock enable")
		return r.cli.mock.mockTaskRemoveTaskCustomField(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "RemoveTaskCustomField",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/custom_fields/:custom_field_guid/remove",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(removeTaskCustomFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskRemoveTaskCustomField mock TaskRemoveTaskCustomField method
func (r *Mock) MockTaskRemoveTaskCustomField(f func(ctx context.Context, request *RemoveTaskCustomFieldReq, options ...MethodOptionFunc) (*RemoveTaskCustomFieldResp, *Response, error)) {
	r.mockTaskRemoveTaskCustomField = f
}

// UnMockTaskRemoveTaskCustomField un-mock TaskRemoveTaskCustomField method
func (r *Mock) UnMockTaskRemoveTaskCustomField() {
	r.mockTaskRemoveTaskCustomField = nil
}

// RemoveTaskCustomFieldReq ...
type RemoveTaskCustomFieldReq struct {
	CustomFieldGuid string `path:"custom_field_guid" json:"-"` // 自定义字段GUID。自定义字段GUID。可以通过[创建自定义字段](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field/create)接口创建, 或者通过[列取自定义字段](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field/list)接口查询得到, 示例值: "0110a4bd-f24b-4a93-8c1a-1732b94f9593"
	ResourceType    string `json:"resource_type,omitempty"`    // 要从某个资源移除自定义字段的资源类型, 目前只支持清单, 示例值: "tasklist"
	ResourceID      string `json:"resource_id,omitempty"`      // 要从某个资源移除自定义字段的资源id, `resource_type`为"tasklist"时, 需填写清单的GUID, 示例值: "0110a4bd-f24b-4a93-8c1a-1732b94f9593"
}

// RemoveTaskCustomFieldResp ...
type RemoveTaskCustomFieldResp struct {
}

// removeTaskCustomFieldResp ...
type removeTaskCustomFieldResp struct {
	Code  int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                     `json:"msg,omitempty"`  // 错误描述
	Data  *RemoveTaskCustomFieldResp `json:"data,omitempty"`
	Error *ErrorDetail               `json:"error,omitempty"`
}
