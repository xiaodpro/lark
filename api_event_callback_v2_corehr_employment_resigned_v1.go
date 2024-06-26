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

// EventV2CorehrEmploymentResignedV1 员工完成离职, 即离职日期的次日凌晨时, 员工雇佣状态更改为“离职”后触发该事件。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/employment/events/resigned
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/offboarding/resigned
func (r *EventCallbackService) HandlerEventV2CorehrEmploymentResignedV1(f EventV2CorehrEmploymentResignedV1Handler) {
	r.cli.eventHandler.eventV2CorehrEmploymentResignedV1Handler = f
}

// EventV2CorehrEmploymentResignedV1Handler event EventV2CorehrEmploymentResignedV1 handler
type EventV2CorehrEmploymentResignedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2CorehrEmploymentResignedV1) (string, error)

// EventV2CorehrEmploymentResignedV1 ...
type EventV2CorehrEmploymentResignedV1 struct {
	EmploymentID string `json:"employment_id,omitempty"` // 主对象ID
}
