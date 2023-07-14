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

// EventV2ContactDepartmentCreatedV3 创建通讯录部门时发送该事件给订阅应用。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=contact&version=v3&resource=department&event=created)
//
// 只有当应用拥有被改动字段的数据权限时, 才会接收到事件。具体的数据权限与字段的关系请参考[应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN), 或查看事件体参数列表的字段描述。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/events/created
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/department/events/created
func (r *EventCallbackService) HandlerEventV2ContactDepartmentCreatedV3(f EventV2ContactDepartmentCreatedV3Handler) {
	r.cli.eventHandler.eventV2ContactDepartmentCreatedV3Handler = f
}

// EventV2ContactDepartmentCreatedV3Handler event EventV2ContactDepartmentCreatedV3 handler
type EventV2ContactDepartmentCreatedV3Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2ContactDepartmentCreatedV3) (string, error)

// EventV2ContactDepartmentCreatedV3 ...
type EventV2ContactDepartmentCreatedV3 struct {
	Object *EventV2ContactDepartmentCreatedV3Object `json:"object,omitempty"` // 部门信息
}

// EventV2ContactDepartmentCreatedV3Object ...
type EventV2ContactDepartmentCreatedV3Object struct {
	Name               string                                                   `json:"name,omitempty"`                 // 部门名称, 最小长度: `1` 字符, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	ParentDepartmentID string                                                   `json:"parent_department_id,omitempty"` // 父部门的部门open_department_id [部门相关ID概念](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview#23857fe0), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	DepartmentID       string                                                   `json:"department_id,omitempty"`        // 本部门的department_id [部门相关ID概念](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview#23857fe0), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	OpenDepartmentID   string                                                   `json:"open_department_id,omitempty"`   // 部门的open_department_id [部门相关ID概念](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview#23857fe0)
	LeaderUserID       string                                                   `json:"leader_user_id,omitempty"`       // 部门主管用户open_id [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	ChatID             string                                                   `json:"chat_id,omitempty"`              // 部门群ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	Order              int64                                                    `json:"order,omitempty"`                // 部门的排序, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	Status             *EventV2ContactDepartmentCreatedV3ObjectStatus           `json:"status,omitempty"`               // 部门状态, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	Leaders            []*EventV2ContactDepartmentCreatedV3ObjectLeader         `json:"leaders,omitempty"`              // 部门负责人
	DepartmentHrbps    []*EventV2ContactDepartmentCreatedV3ObjectDepartmentHrbp `json:"department_hrbps,omitempty"`     // 部门HRBP, 字段权限要求: 查询部门 HRBP 信息
}

// EventV2ContactDepartmentCreatedV3ObjectDepartmentHrbp ...
type EventV2ContactDepartmentCreatedV3ObjectDepartmentHrbp struct {
	UnionID string `json:"union_id,omitempty"` // Union ID
	UserID  string `json:"user_id,omitempty"`  // User ID
	OpenID  string `json:"open_id,omitempty"`  // Open ID
}

// EventV2ContactDepartmentCreatedV3ObjectLeader ...
type EventV2ContactDepartmentCreatedV3ObjectLeader struct {
	LeaderType int64  `json:"leaderType,omitempty"` // 负责人类型, 可选值有: 1: 主负责人, 2: 副负责人
	LeaderID   string `json:"leaderID,omitempty"`   // 负责人ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
}

// EventV2ContactDepartmentCreatedV3ObjectStatus ...
type EventV2ContactDepartmentCreatedV3ObjectStatus struct {
	IsDeleted bool `json:"is_deleted,omitempty"` // 是否被删除
}
