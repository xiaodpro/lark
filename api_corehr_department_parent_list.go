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

// GetCoreHRDepartmentParentList 该接口用来递归获取部门的父部门信息, 并按照由子到父的顺序返回有权限的父部门信息列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/department/parents
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/department/parents
func (r *CoreHRService) GetCoreHRDepartmentParentList(ctx context.Context, request *GetCoreHRDepartmentParentListReq, options ...MethodOptionFunc) (*GetCoreHRDepartmentParentListResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRDepartmentParentList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRDepartmentParentList mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRDepartmentParentList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRDepartmentParentList",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/departments/parents",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRDepartmentParentListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRDepartmentParentList mock CoreHRGetCoreHRDepartmentParentList method
func (r *Mock) MockCoreHRGetCoreHRDepartmentParentList(f func(ctx context.Context, request *GetCoreHRDepartmentParentListReq, options ...MethodOptionFunc) (*GetCoreHRDepartmentParentListResp, *Response, error)) {
	r.mockCoreHRGetCoreHRDepartmentParentList = f
}

// UnMockCoreHRGetCoreHRDepartmentParentList un-mock CoreHRGetCoreHRDepartmentParentList method
func (r *Mock) UnMockCoreHRGetCoreHRDepartmentParentList() {
	r.mockCoreHRGetCoreHRDepartmentParentList = nil
}

// GetCoreHRDepartmentParentListReq ...
type GetCoreHRDepartmentParentListReq struct {
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门 ID 类型, 示例值: open_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_corehr_department_id: 以 people_corehr_department_id 来标识部门, 默认值: `open_department_id`
	PageToken        *string           `query:"page_token" json:"-"`         // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: eVQrYzJBNDNONlk4VFZBZVlSdzlKdFJ4bVVHVExENDNKVHoxaVdiVnViQT0=
	PageSize         *int64            `query:"page_size" json:"-"`          // 分页大小, 示例值: 10, 默认值: `20`
	DepartmentIDList []string          `json:"department_id_list,omitempty"` // 部门 ID 列表, 一次性最多传入 100 个部门 ID, 示例值: ["6893014062142064111"], 长度范围: `1` ～ `100`
}

// GetCoreHRDepartmentParentListResp ...
type GetCoreHRDepartmentParentListResp struct {
	Items []*GetCoreHRDepartmentParentListRespItem `json:"items,omitempty"` // 父部门查询结果
}

// GetCoreHRDepartmentParentListRespItem ...
type GetCoreHRDepartmentParentListRespItem struct {
	DepartmentID         string                                                   `json:"department_id,omitempty"`          // 部门 ID
	ParentDepartmentList []*GetCoreHRDepartmentParentListRespItemParentDepartment `json:"parent_department_list,omitempty"` // 父部门列表, 部门按照至底向上的顺序返回
}

// GetCoreHRDepartmentParentListRespItemParentDepartment ...
type GetCoreHRDepartmentParentListRespItemParentDepartment struct {
	DepartmentID       string                                                                 `json:"department_id,omitempty"`        // 部门 ID
	DepartmentName     []*GetCoreHRDepartmentParentListRespItemParentDepartmentDepartmentName `json:"department_name,omitempty"`      // 部门名称
	ParentDepartmentID string                                                                 `json:"parent_department_id,omitempty"` // 上级部门 ID
	Active             bool                                                                   `json:"active,omitempty"`               // 是否启用
	IsRoot             bool                                                                   `json:"is_root,omitempty"`              // 是否根部门
}

// GetCoreHRDepartmentParentListRespItemParentDepartmentDepartmentName ...
type GetCoreHRDepartmentParentListRespItemParentDepartmentDepartmentName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// getCoreHRDepartmentParentListResp ...
type getCoreHRDepartmentParentListResp struct {
	Code int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHRDepartmentParentListResp `json:"data,omitempty"`
}
