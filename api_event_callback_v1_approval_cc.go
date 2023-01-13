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

// EventV1ApprovalCc 创建抄送或者被抄送人已读抄送后, 会向开发者推送消息。
//
// 1. 创建抄送, 推送 "operate" 为 "CREATE" 的事件。
// 1. 被抄送人已读抄送, 推送 "operate" 为 "READ" 的事件。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/common-event/approval-cc-event
func (r *EventCallbackService) HandlerEventV1ApprovalCc(f EventV1ApprovalCcHandler) {
	r.cli.eventHandler.eventV1ApprovalCcHandler = f
}

// EventV1ApprovalCcHandler event EventV1ApprovalCc handler
type EventV1ApprovalCcHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1ApprovalCc) (string, error)

// EventV1ApprovalCc ...
type EventV1ApprovalCc struct {
	AppID        string `json:"app_id,omitempty"` // 如: cli_xxx
	TenantKey    string `json:"tenant_key,omitempty"`
	Type         string `json:"type,omitempty"`          // approval_cc 固定字段
	ApprovalCode string `json:"approval_code,omitempty"` // 审批定义 Code
	InstanceCode string `json:"instance_code,omitempty"` // 审批实例 Code
	ID           string `json:"id,omitempty"`            // 抄送 ID
	UserID       string `json:"user_id,omitempty"`       // 被抄送人
	CreateTime   int64  `json:"create_time,omitempty"`   // 抄送时间
	Operate      string `json:"operate,omitempty"`       // 操作类型  CREATE: 抄送  READ: 已读
	From         string `json:"from,omitempty"`          // 抄送人, 可能为空
}
