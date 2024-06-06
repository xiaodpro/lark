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

// QueryCoreHRDepartmentTimeline 查询指定生效的部门基本信息, 含部门名称、上级、编码、负责人、是否启用、描述等信息
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/department/query_timeline
func (r *CoreHRService) QueryCoreHRDepartmentTimeline(ctx context.Context, request *QueryCoreHRDepartmentTimelineReq, options ...MethodOptionFunc) (*QueryCoreHRDepartmentTimelineResp, *Response, error) {
	if r.cli.mock.mockCoreHRQueryCoreHRDepartmentTimeline != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#QueryCoreHRDepartmentTimeline mock enable")
		return r.cli.mock.mockCoreHRQueryCoreHRDepartmentTimeline(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "QueryCoreHRDepartmentTimeline",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/departments/query_timeline",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(queryCoreHRDepartmentTimelineResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRQueryCoreHRDepartmentTimeline mock CoreHRQueryCoreHRDepartmentTimeline method
func (r *Mock) MockCoreHRQueryCoreHRDepartmentTimeline(f func(ctx context.Context, request *QueryCoreHRDepartmentTimelineReq, options ...MethodOptionFunc) (*QueryCoreHRDepartmentTimelineResp, *Response, error)) {
	r.mockCoreHRQueryCoreHRDepartmentTimeline = f
}

// UnMockCoreHRQueryCoreHRDepartmentTimeline un-mock CoreHRQueryCoreHRDepartmentTimeline method
func (r *Mock) UnMockCoreHRQueryCoreHRDepartmentTimeline() {
	r.mockCoreHRQueryCoreHRDepartmentTimeline = nil
}

// QueryCoreHRDepartmentTimelineReq ...
type QueryCoreHRDepartmentTimelineReq struct {
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: people_corehr_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `people_corehr_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门 ID 类型, 示例值: people_corehr_department_id, 可选值有: open_department_id: [飞书]用来在具体某个应用中标识一个部门, 同一个department_id 在不同应用中的 open_department_id 相同。, department_id: [飞书]用来标识租户内一个唯一的部门。, people_corehr_department_id: [飞书人事]用来标识「飞书人事」中的部门。, 默认值: `people_corehr_department_id`
	DepartmentIDs    []string          `json:"department_ids,omitempty"`     // 部门 ID 列表, 可请求[搜索部门信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/department/search)获取, 示例值: ["7094136522860922111"], 长度范围: `0` ～ `100`
	EffectiveDate    string            `json:"effective_date,omitempty"`     // 生效日期, 示例值: "2020-01-01", 长度范围: `10` ～ `10` 字符, 正则校验: `^((([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8]))))|((([0-9]{2})(0[48]|[2468][048]|[13579][26])|((0[48]|[2468][048]|[3579][26])00))-02-29))$`
	Fields           []string          `json:"fields,omitempty"`             // 返回数据的字段列表, 可选["department_name", "code", "active", "parent_department_id", "manager", "description", "effective_date"], 示例值: ["department_name"], 长度范围: `0` ～ `100`
}

// QueryCoreHRDepartmentTimelineResp ...
type QueryCoreHRDepartmentTimelineResp struct {
	Items []*QueryCoreHRDepartmentTimelineRespItem `json:"items,omitempty"` // 部门信息
}

// QueryCoreHRDepartmentTimelineRespItem ...
type QueryCoreHRDepartmentTimelineRespItem struct {
	ID                 string                                              `json:"id,omitempty"`                   // 部门 ID
	VersionID          string                                              `json:"version_id,omitempty"`           // 部门版本 ID
	Names              []*QueryCoreHRDepartmentTimelineRespItemName        `json:"names,omitempty"`                // 部门名称
	ParentDepartmentID string                                              `json:"parent_department_id,omitempty"` // 上级部门 ID, 字段权限要求: 获取部门组织架构信息
	Manager            string                                              `json:"manager,omitempty"`              // 部门负责人雇佣 ID, 枚举值及详细信息可通过[查询员工信息]接口查询获得, 字段权限要求: 获取部门负责人信息
	Code               string                                              `json:"code,omitempty"`                 // 编码
	EffectiveDate      string                                              `json:"effective_date,omitempty"`       // 生效日期
	Active             bool                                                `json:"active,omitempty"`               // 是否启用
	Descriptions       []*QueryCoreHRDepartmentTimelineRespItemDescription `json:"descriptions,omitempty"`         // 描述
	CustomFields       []*QueryCoreHRDepartmentTimelineRespItemCustomField `json:"custom_fields,omitempty"`        // 自定义字段, 字段权限要求: 获取部门自定义字段
}

// QueryCoreHRDepartmentTimelineRespItemCustomField ...
type QueryCoreHRDepartmentTimelineRespItemCustomField struct {
	CustomApiName string                                                `json:"custom_api_name,omitempty"` // 自定义字段 apiname, 即自定义字段的唯一标识
	Name          *QueryCoreHRDepartmentTimelineRespItemCustomFieldName `json:"name,omitempty"`            // 自定义字段名称
	Type          int64                                                 `json:"type,omitempty"`            // 自定义字段类型
	Value         string                                                `json:"value,omitempty"`           // 字段值, 是 json 转义后的字符串, 根据元数据定义不同, 字段格式不同（如 123, 123.23, "true", ["id1", "id2"], "2006-01-02 15:04:05"）
}

// QueryCoreHRDepartmentTimelineRespItemCustomFieldName ...
type QueryCoreHRDepartmentTimelineRespItemCustomFieldName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// QueryCoreHRDepartmentTimelineRespItemDescription ...
type QueryCoreHRDepartmentTimelineRespItemDescription struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// QueryCoreHRDepartmentTimelineRespItemName ...
type QueryCoreHRDepartmentTimelineRespItemName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// queryCoreHRDepartmentTimelineResp ...
type queryCoreHRDepartmentTimelineResp struct {
	Code  int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                             `json:"msg,omitempty"`  // 错误描述
	Data  *QueryCoreHRDepartmentTimelineResp `json:"data,omitempty"`
	Error *ErrorDetail                       `json:"error,omitempty"`
}
