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

// UpdateDepartment 该接口用于更新当前部门所有信息。
//
// - 调用该接口需要具有该部门与更新部门信息涉及的通讯录权限。
// - 没有填写的字段会被置为空值（order字段除外）。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/update
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/department/update
func (r *ContactService) UpdateDepartment(ctx context.Context, request *UpdateDepartmentReq, options ...MethodOptionFunc) (*UpdateDepartmentResp, *Response, error) {
	if r.cli.mock.mockContactUpdateDepartment != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Contact#UpdateDepartment mock enable")
		return r.cli.mock.mockContactUpdateDepartment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "UpdateDepartment",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/departments/:department_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateDepartmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactUpdateDepartment mock ContactUpdateDepartment method
func (r *Mock) MockContactUpdateDepartment(f func(ctx context.Context, request *UpdateDepartmentReq, options ...MethodOptionFunc) (*UpdateDepartmentResp, *Response, error)) {
	r.mockContactUpdateDepartment = f
}

// UnMockContactUpdateDepartment un-mock ContactUpdateDepartment method
func (r *Mock) UnMockContactUpdateDepartment() {
	r.mockContactUpdateDepartment = nil
}

// UpdateDepartmentReq ...
type UpdateDepartmentReq struct {
	DepartmentID           string                       `path:"department_id" json:"-"`              // 部门ID, 需要与查询参数中传入的department_id_type类型保持一致, 注意: 除需要满足正则规则外, 同时不能以od-开头, 示例值: "D096", 最大长度: `64` 字符, 正则校验: `^[a-zA-Z0-9][a-zA-Z0-9_\-@.]{0, 63}$`
	UserIDType             *IDType                      `query:"user_id_type" json:"-"`              // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType       *DepartmentIDType            `query:"department_id_type" json:"-"`        // 此次调用中使用的部门ID的类型, 示例值: open_department_id, 可选值有: department_id: 用来标识租户内一个唯一的部门, open_department_id: 用来在具体某个应用中标识一个部门, 同一个部门 在不同应用中的 open_department_id 相同。, 默认值: `open_department_id`
	Name                   string                       `json:"name,omitempty"`                      // 部门名称, 注意: 不可包含斜杠, 示例值: "DemoName", 最小长度: `1` 字符
	I18nName               *UpdateDepartmentReqI18nName `json:"i18n_name,omitempty"`                 // 国际化的部门名称, 注意: 不可包含斜杠
	ParentDepartmentID     string                       `json:"parent_department_id,omitempty"`      // 父部门的ID, * 在根部门下创建新部门, 该参数值为 “0”, 示例值: "D067"
	LeaderUserID           *string                      `json:"leader_user_id,omitempty"`            // 部门主管用户ID, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	Order                  *string                      `json:"order,omitempty"`                     // 部门的排序, 即部门在其同级部门的展示顺序, 示例值: "100"
	UnitIDs                []string                     `json:"unit_ids,omitempty"`                  // 部门单位自定义ID列表, 当前只支持一个, 示例值: ["custom_unit_id"]
	CreateGroupChat        *bool                        `json:"create_group_chat,omitempty"`         // 是否创建部门群, 默认不创建, 创建部门群时, 默认群名为部门名, 默认群主为部门主负责人, 示例值: false
	Leaders                []*UpdateDepartmentReqLeader `json:"leaders,omitempty"`                   // 部门负责人
	GroupChatEmployeeTypes []int64                      `json:"group_chat_employee_types,omitempty"` // 部门群雇员类型限制。[]空列表时, 表示为无任何雇员类型。类型字段可包含以下值, 支持多个类型值；若有多个, 用英文', '分隔: 1、正式员工, 2、实习生, 3、外包, 4、劳务, 5、顾问, 6、其他自定义类型字段, 可通过下方接口获取到该租户的自定义员工类型的名称, 参见[获取人员类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list), 示例值: [1]
}

// UpdateDepartmentReqI18nName ...
type UpdateDepartmentReqI18nName struct {
	ZhCn *string `json:"zh_cn,omitempty"` // 部门的中文名, 示例值: "Demo名称"
	JaJp *string `json:"ja_jp,omitempty"` // 部门的日文名, 示例值: "デモ名"
	EnUs *string `json:"en_us,omitempty"` // 部门的英文名, 示例值: "Demo Name"
}

// UpdateDepartmentReqLeader ...
type UpdateDepartmentReqLeader struct {
	LeaderType int64  `json:"leaderType,omitempty"` // 负责人类型, 示例值: 1, 可选值有: 1: 主负责人, 2: 副负责人
	LeaderID   string `json:"leaderID,omitempty"`   // 负责人ID, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
}

// UpdateDepartmentResp ...
type UpdateDepartmentResp struct {
	Department *UpdateDepartmentRespDepartment `json:"department,omitempty"` // 部门信息
}

// UpdateDepartmentRespDepartment ...
type UpdateDepartmentRespDepartment struct {
	Name                   string                                  `json:"name,omitempty"`                      // 部门名称, 注意: 不可包含斜杠, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	I18nName               *UpdateDepartmentRespDepartmentI18nName `json:"i18n_name,omitempty"`                 // 国际化的部门名称, 注意: 不可包含斜杠, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	ParentDepartmentID     string                                  `json:"parent_department_id,omitempty"`      // 父部门的ID, * 在根部门下创建新部门, 该参数值为 “0”, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	DepartmentID           string                                  `json:"department_id,omitempty"`             // 本部门的自定义部门ID, 注意: 除需要满足正则规则外, 同时不能以`od-`开头, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	OpenDepartmentID       string                                  `json:"open_department_id,omitempty"`        // 部门的open_id, 类型与通过请求的查询参数传入的department_id_type相同
	LeaderUserID           string                                  `json:"leader_user_id,omitempty"`            // 部门主管用户ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	ChatID                 string                                  `json:"chat_id,omitempty"`                   // 部门群ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	Order                  string                                  `json:"order,omitempty"`                     // 部门的排序, 即部门在其同级部门的展示顺序, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	UnitIDs                []string                                `json:"unit_ids,omitempty"`                  // 部门单位自定义ID列表, 当前只支持一个, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	MemberCount            int64                                   `json:"member_count,omitempty"`              // 当前部门及其下属部门下用户（包含部门负责人）的个数, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	Status                 *UpdateDepartmentRespDepartmentStatus   `json:"status,omitempty"`                    // 部门状态, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	Leaders                []*UpdateDepartmentRespDepartmentLeader `json:"leaders,omitempty"`                   // 部门负责人
	GroupChatEmployeeTypes []int64                                 `json:"group_chat_employee_types,omitempty"` // 部门群雇员类型限制。[]空列表时, 表示为无任何雇员类型。类型字段可包含以下值, 支持多个类型值；若有多个, 用英文', '分隔: 1、正式员工, 2、实习生, 3、外包, 4、劳务, 5、顾问, 6、其他自定义类型字段, 可通过下方接口获取到该租户的自定义员工类型的名称, 参见[获取人员类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list)。
	PrimaryMemberCount     int64                                   `json:"primary_member_count,omitempty"`      // 当前部门及其下属的主属成员（即成员的主部门为当前部门）的数量, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
}

// UpdateDepartmentRespDepartmentI18nName ...
type UpdateDepartmentRespDepartmentI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 部门的中文名
	JaJp string `json:"ja_jp,omitempty"` // 部门的日文名
	EnUs string `json:"en_us,omitempty"` // 部门的英文名
}

// UpdateDepartmentRespDepartmentLeader ...
type UpdateDepartmentRespDepartmentLeader struct {
	LeaderType int64  `json:"leaderType,omitempty"` // 负责人类型, 可选值有: 1: 主负责人, 2: 副负责人
	LeaderID   string `json:"leaderID,omitempty"`   // 负责人ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
}

// UpdateDepartmentRespDepartmentStatus ...
type UpdateDepartmentRespDepartmentStatus struct {
	IsDeleted bool `json:"is_deleted,omitempty"` // 是否被删除
}

// updateDepartmentResp ...
type updateDepartmentResp struct {
	Code  int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateDepartmentResp `json:"data,omitempty"`
	Error *ErrorDetail          `json:"error,omitempty"`
}
