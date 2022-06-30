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

// EventV2ContactCustomAttrEventUpdatedV3
//
// 通过该事件订阅成员字段变更。old_object 展示更新字段的原始值。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=contact&version=v3&resource=custom_attr_event&event=updated)
// 触发事件的动作有「打开/关闭」开关、「增加/删除」成员字段。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/custom_attr_event/events/updated
func (r *EventCallbackService) HandlerEventV2ContactCustomAttrEventUpdatedV3(f EventV2ContactCustomAttrEventUpdatedV3Handler) {
	r.cli.eventHandler.eventV2ContactCustomAttrEventUpdatedV3Handler = f
}

// EventV2ContactCustomAttrEventUpdatedV3Handler event EventV2ContactCustomAttrEventUpdatedV3 handler
type EventV2ContactCustomAttrEventUpdatedV3Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2ContactCustomAttrEventUpdatedV3) (string, error)

// EventV2ContactCustomAttrEventUpdatedV3 ...
type EventV2ContactCustomAttrEventUpdatedV3 struct {
	Object    *EventV2ContactCustomAttrEventUpdatedV3Object    `json:"object,omitempty"`     // 变更后信息
	OldObject *EventV2ContactCustomAttrEventUpdatedV3OldObject `json:"old_object,omitempty"` // 变更前信息
}

// EventV2ContactCustomAttrEventUpdatedV3Object ...
type EventV2ContactCustomAttrEventUpdatedV3Object struct {
	ContactFieldKey []string `json:"contact_field_key,omitempty"` // 通讯录字段键值
	AllowOpenQuery  bool     `json:"allow_open_query,omitempty"`  // 开关是否打开
}

// EventV2ContactCustomAttrEventUpdatedV3OldObject ...
type EventV2ContactCustomAttrEventUpdatedV3OldObject struct {
	ContactFieldKey []string `json:"contact_field_key,omitempty"` // 通讯录字段键值
	AllowOpenQuery  bool     `json:"allow_open_query,omitempty"`  // 开关是否打开
}
