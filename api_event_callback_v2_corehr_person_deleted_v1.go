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

// EventV2CorehrPersonDeletedV1 个人信息删除
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/person/events/deleted
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/employee/person/deleted
func (r *EventCallbackService) HandlerEventV2CorehrPersonDeletedV1(f EventV2CorehrPersonDeletedV1Handler) {
	r.cli.eventHandler.eventV2CorehrPersonDeletedV1Handler = f
}

// EventV2CorehrPersonDeletedV1Handler event EventV2CorehrPersonDeletedV1 handler
type EventV2CorehrPersonDeletedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2CorehrPersonDeletedV1) (string, error)

// EventV2CorehrPersonDeletedV1 ...
type EventV2CorehrPersonDeletedV1 struct {
	PersonID string `json:"person_id,omitempty"` // 人员ID
}
