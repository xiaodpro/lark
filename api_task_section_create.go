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

// CreateTaskSection 为清单或我负责的任务列表创建一个自定义分组。创建时可以需要提供名称和可选的配置。如果不指定位置, 新分组会放到指定resource的自定义分组列表的最后。
//
// 当在清单中创建自定义分组时, 需要设置`resourse_type`为"tasklist", `resource_id`设为清单的GUID。
// 当为我负责任务列表中创建自定义分组时, 需要设置`resource_type`为"my_tasks", 不需要设置`resource_id`。调用身份只能为自己的我负责的任务列表创建自定义分组。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/section/create
func (r *TaskService) CreateTaskSection(ctx context.Context, request *CreateTaskSectionReq, options ...MethodOptionFunc) (*CreateTaskSectionResp, *Response, error) {
	if r.cli.mock.mockTaskCreateTaskSection != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#CreateTaskSection mock enable")
		return r.cli.mock.mockTaskCreateTaskSection(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "CreateTaskSection",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/sections",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createTaskSectionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskCreateTaskSection mock TaskCreateTaskSection method
func (r *Mock) MockTaskCreateTaskSection(f func(ctx context.Context, request *CreateTaskSectionReq, options ...MethodOptionFunc) (*CreateTaskSectionResp, *Response, error)) {
	r.mockTaskCreateTaskSection = f
}

// UnMockTaskCreateTaskSection un-mock TaskCreateTaskSection method
func (r *Mock) UnMockTaskCreateTaskSection() {
	r.mockTaskCreateTaskSection = nil
}

// CreateTaskSectionReq ...
type CreateTaskSectionReq struct {
	UserIDType   *IDType `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值: open_id, 默认值: `open_id`
	Name         string  `json:"name,omitempty"`          // 自定义分组名。不允许为空, 最大100个utf8字符, 示例值: "已经审核过的任务"
	ResourceType string  `json:"resource_type,omitempty"` // 自定义分组的资源类型, 支持"tasklist"（清单）或者"my_tasks"（我负责的）, 示例值: "tasklist", 默认值: `tasklist`
	ResourceID   *string `json:"resource_id,omitempty"`   // 自定义分组要归属的资源id。当`resource_type`为"tasklist"时这里需要填写清单的GUID；当`resource_type`为"my_tasks"时, 无需填写, 示例值: "cc371766-6584-cf50-a222-c22cd9055004"
	InsertBefore *string `json:"insert_before,omitempty"` // 要将新分组插入到自定义分分组的前面的目标分组的guid, `insert_before`和`insert_after`均不设置时表示将新分组放到已有的所有自定义分组之后, 如果同时设置`insert_before`和`insert_after`, 接口会报错, 示例值: "e6e37dcc-f75a-5936-f589-12fb4b5c80c2", 最大长度: `100` 字符
	InsertAfter  *string `json:"insert_after,omitempty"`  // 要将新分组插入到自定义分分组的后面的目标分组的guid, `insert_before`和`insert_after`均不设置时表示将新分组放到已有的所有自定义分组之后, 如果同时设置`insert_before`和`insert_after`, 接口会报错, 示例值: "e6e37dcc-f75a-5936-f589-12fb4b5c80c2", 最大长度: `100` 字符
}

// CreateTaskSectionResp ...
type CreateTaskSectionResp struct {
	Section *CreateTaskSectionRespSection `json:"section,omitempty"` // 创建的自定义分组数据
}

// CreateTaskSectionRespSection ...
type CreateTaskSectionRespSection struct {
	Guid         string                                `json:"guid,omitempty"`          // 自定义分组的guid
	Name         string                                `json:"name,omitempty"`          // 自定义分组的名字
	ResourceType string                                `json:"resource_type,omitempty"` // 资源类型
	IsDefault    bool                                  `json:"is_default,omitempty"`    // 分组是否为默认自定义分组
	Creator      *CreateTaskSectionRespSectionCreator  `json:"creator,omitempty"`       // 自定义分组的创建者
	Tasklist     *CreateTaskSectionRespSectionTasklist `json:"tasklist,omitempty"`      // 如果该分组归属于清单, 展示清单的简要信息
	CreatedAt    string                                `json:"created_at,omitempty"`    // 自定义分组创建时间戳(ms)
	UpdatedAt    string                                `json:"updated_at,omitempty"`    // 自定义分组最近一次更新时间戳(ms)
}

// CreateTaskSectionRespSectionCreator ...
type CreateTaskSectionRespSectionCreator struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// CreateTaskSectionRespSectionTasklist ...
type CreateTaskSectionRespSectionTasklist struct {
	Guid string `json:"guid,omitempty"` // 清单的全局唯一ID
	Name string `json:"name,omitempty"` // 清单名字
}

// createTaskSectionResp ...
type createTaskSectionResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *CreateTaskSectionResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}
