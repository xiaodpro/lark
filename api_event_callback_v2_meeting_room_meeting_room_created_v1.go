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

// EventV2MeetingRoomMeetingRoomCreatedV1 会议室被创建将触发此事件。
//
// 了解事件订阅的使用场景和配置流程, 请点击查看 [事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/meeting_room-v1/meeting_room/events/created
// new doc: https://open.feishu.cn/document/server-docs/historic-version/meeting_room-v1/event/created
func (r *EventCallbackService) HandlerEventV2MeetingRoomMeetingRoomCreatedV1(f EventV2MeetingRoomMeetingRoomCreatedV1Handler) {
	r.cli.eventHandler.eventV2MeetingRoomMeetingRoomCreatedV1Handler = f
}

// EventV2MeetingRoomMeetingRoomCreatedV1Handler event EventV2MeetingRoomMeetingRoomCreatedV1 handler
type EventV2MeetingRoomMeetingRoomCreatedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2MeetingRoomMeetingRoomCreatedV1) (string, error)

// EventV2MeetingRoomMeetingRoomCreatedV1 ...
type EventV2MeetingRoomMeetingRoomCreatedV1 struct {
	RoomName string `json:"room_name,omitempty"` // 会议室名称
	RoomID   string `json:"room_id,omitempty"`   // 会议室 ID
}
