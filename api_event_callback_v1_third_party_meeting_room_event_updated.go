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

// EventV1ThirdPartyMeetingRoomEventUpdated 当添加了第三方会议室的日程发生变动时（创建/更新/删除）触发此事件。
//
// 了解事件订阅的使用场景和配置流程, 可参见[事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/meeting_room-v1/event/third-room-event-changes
// new doc: https://open.feishu.cn/document/server-docs/calendar-v4/meeting-room-event/event/third-room-event-changes
func (r *EventCallbackService) HandlerEventV1ThirdPartyMeetingRoomEventUpdated(f EventV1ThirdPartyMeetingRoomEventUpdatedHandler) {
	r.cli.eventHandler.eventV1ThirdPartyMeetingRoomEventUpdatedHandler = f
}

// EventV1ThirdPartyMeetingRoomEventUpdatedHandler event EventV1ThirdPartyMeetingRoomEventUpdated handler
type EventV1ThirdPartyMeetingRoomEventUpdatedHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1ThirdPartyMeetingRoomEventUpdated) (string, error)

// EventV1ThirdPartyMeetingRoomEventUpdated ...
type EventV1ThirdPartyMeetingRoomEventUpdated struct {
}
