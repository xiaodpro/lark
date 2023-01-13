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

// EventV2VCRoomLevelCreatedV1 当创建会议室层级时, 会触发该事件。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=vc&version=v1&resource=room_level&event=created)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_level/events/created
func (r *EventCallbackService) HandlerEventV2VCRoomLevelCreatedV1(f EventV2VCRoomLevelCreatedV1Handler) {
	r.cli.eventHandler.eventV2VCRoomLevelCreatedV1Handler = f
}

// EventV2VCRoomLevelCreatedV1Handler event EventV2VCRoomLevelCreatedV1 handler
type EventV2VCRoomLevelCreatedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2VCRoomLevelCreatedV1) (string, error)

// EventV2VCRoomLevelCreatedV1 ...
type EventV2VCRoomLevelCreatedV1 struct {
	RoomLevel *EventV2VCRoomLevelCreatedV1RoomLevel `json:"room_level,omitempty"` // 层级信息
}

// EventV2VCRoomLevelCreatedV1RoomLevel ...
type EventV2VCRoomLevelCreatedV1RoomLevel struct {
	RoomLevelID   string   `json:"room_level_id,omitempty"`   // 层级ID
	Name          string   `json:"name,omitempty"`            // 层级名称
	ParentID      string   `json:"parent_id,omitempty"`       // 父层级ID
	Path          []string `json:"path,omitempty"`            // 层级路径
	HasChild      bool     `json:"has_child,omitempty"`       // 是否有子层级
	CustomGroupID string   `json:"custom_group_id,omitempty"` // 自定义层级ID
}
