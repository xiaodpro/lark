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

// EventV2ContactEmployeeTypeEnumCreatedV3 新建人员类型会发出对应事件。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/events/created
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/employee_type_enum/events/created
func (r *EventCallbackService) HandlerEventV2ContactEmployeeTypeEnumCreatedV3(f EventV2ContactEmployeeTypeEnumCreatedV3Handler) {
	r.cli.eventHandler.eventV2ContactEmployeeTypeEnumCreatedV3Handler = f
}

// EventV2ContactEmployeeTypeEnumCreatedV3Handler event EventV2ContactEmployeeTypeEnumCreatedV3 handler
type EventV2ContactEmployeeTypeEnumCreatedV3Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2ContactEmployeeTypeEnumCreatedV3) (string, error)

// EventV2ContactEmployeeTypeEnumCreatedV3 ...
type EventV2ContactEmployeeTypeEnumCreatedV3 struct {
	NewEnum *EventV2ContactEmployeeTypeEnumCreatedV3NewEnum `json:"new_enum,omitempty"` // 新枚举类型
}

// EventV2ContactEmployeeTypeEnumCreatedV3NewEnum ...
type EventV2ContactEmployeeTypeEnumCreatedV3NewEnum struct {
	EnumID      string                                                       `json:"enum_id,omitempty"`      // 枚举值id
	EnumValue   string                                                       `json:"enum_value,omitempty"`   // 枚举的编号值, 创建新的人员类型后, 系统生成对应编号。对应[创建用户接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/create)中用户信息的employee_type字段值
	Content     string                                                       `json:"content,omitempty"`      // 枚举内容, 长度范围: `1` ～ `100` 字符
	EnumType    int64                                                        `json:"enum_type,omitempty"`    // 类型, 可选值有: 1: 内置类型, 2: 自定义
	EnumStatus  int64                                                        `json:"enum_status,omitempty"`  // 使用状态, 可选值有: 1: 激活, 2: 未激活
	I18nContent []*EventV2ContactEmployeeTypeEnumCreatedV3NewEnumI18nContent `json:"i18n_content,omitempty"` // i18n定义
}

// EventV2ContactEmployeeTypeEnumCreatedV3NewEnumI18nContent ...
type EventV2ContactEmployeeTypeEnumCreatedV3NewEnumI18nContent struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}
