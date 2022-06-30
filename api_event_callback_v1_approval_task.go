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

// EventV1ApprovalTask
//
//
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ugDNyUjL4QjM14CO0ITN
func (r *EventCallbackService) HandlerEventV1ApprovalTask(f EventV1ApprovalTaskHandler) {
	r.cli.eventHandler.eventV1ApprovalTaskHandler = f
}

// EventV1ApprovalTaskHandler event EventV1ApprovalTask handler
type EventV1ApprovalTaskHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1ApprovalTask) (string, error)

// EventV1ApprovalTask ...
type EventV1ApprovalTask struct {
	AppID        string `json:"app_id,omitempty"`  // 如: cli_xxx
	OpenID       string `json:"open_id,omitempty"` // 如: xxx
	TenantKey    string `json:"tenant_key,omitempty"`
	Type         string `json:"type,omitempty"`          // 固定 approval_task
	ApprovalCode string `json:"approval_code,omitempty"` // 审批定义 Code
	InstanceCode string `json:"instance_code,omitempty"` // 审批实例 Code
	TaskID       string `json:"task_id,omitempty"`       // 审批任务 ID
	UserID       string `json:"user_id,omitempty"`       // 操作人 ID（当 task 为自动通过类型时，user_id 为空）
	Status       string `json:"status,omitempty"`        // 任务状态, REVERTED - 已还原, PENDING - 进行中, APPROVED - 已通过, REJECTED - 已拒绝, TRANSFERRED - 已转交, DONE - 已完成
	OperateTime  string `json:"operate_time,omitempty"`  // 事件发生时间
}
