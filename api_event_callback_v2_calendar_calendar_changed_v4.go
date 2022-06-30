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

// EventV2CalendarCalendarChangedV4
//
// 当订阅用户的日历列表有日历变动时触发此事件。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar&event=changed)
// 应用首先需要调用上述接口建立订阅关系。应用收到该事件后, 使用事件的 user_list 字段中的用户对应的 user_access_token 调用[获取日历列表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/list)接口拉取增量的变更数据
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/events/changed
func (r *EventCallbackService) HandlerEventV2CalendarCalendarChangedV4(f EventV2CalendarCalendarChangedV4Handler) {
	r.cli.eventHandler.eventV2CalendarCalendarChangedV4Handler = f
}

// EventV2CalendarCalendarChangedV4Handler event EventV2CalendarCalendarChangedV4 handler
type EventV2CalendarCalendarChangedV4Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2CalendarCalendarChangedV4) (string, error)

// EventV2CalendarCalendarChangedV4 ...
type EventV2CalendarCalendarChangedV4 struct {
	UserIDList []*EventV2CalendarCalendarChangedV4UserID `json:"user_id_list,omitempty"` // 需要推送事件的用户列表
}

// EventV2CalendarCalendarChangedV4UserID ...
type EventV2CalendarCalendarChangedV4UserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
