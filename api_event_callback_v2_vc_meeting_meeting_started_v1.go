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

// EventV2VCMeetingMeetingStartedV1 发生在会议开始时, 目前仅提供预约会议的相关事件。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=vc&version=v1&resource=meeting&event=meeting_started)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/events/meeting_started
func (r *EventCallbackService) HandlerEventV2VCMeetingMeetingStartedV1(f EventV2VCMeetingMeetingStartedV1Handler) {
	r.cli.eventHandler.eventV2VCMeetingMeetingStartedV1Handler = f
}

// EventV2VCMeetingMeetingStartedV1Handler event EventV2VCMeetingMeetingStartedV1 handler
type EventV2VCMeetingMeetingStartedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2VCMeetingMeetingStartedV1) (string, error)

// EventV2VCMeetingMeetingStartedV1 ...
type EventV2VCMeetingMeetingStartedV1 struct {
	Meeting  *EventV2VCMeetingMeetingStartedV1Meeting  `json:"meeting,omitempty"`  // 会议数据
	Operator *EventV2VCMeetingMeetingStartedV1Operator `json:"operator,omitempty"` // 事件操作人
}

// EventV2VCMeetingMeetingStartedV1Meeting ...
type EventV2VCMeetingMeetingStartedV1Meeting struct {
	ID            string                                           `json:"id,omitempty"`             // 会议ID（视频会议的唯一标识, 视频会议开始后才会产生）
	Topic         string                                           `json:"topic,omitempty"`          // 会议主题
	MeetingNo     string                                           `json:"meeting_no,omitempty"`     // 9位会议号（飞书用户可通过输入9位会议号快捷入会）
	MeetingSource int64                                            `json:"meeting_source,omitempty"` // 会议创建源, 可选值有: 1: 日程会议, 2: 即时会议, 3: 面试会议, 4: 开放平台会议, 100: 其他会议类型
	StartTime     string                                           `json:"start_time,omitempty"`     // 会议结束时间（unix时间, 单位: 秒）
	EndTime       string                                           `json:"end_time,omitempty"`       // 会议结束时间（unix时间, 单位: 秒）
	HostUser      *EventV2VCMeetingMeetingStartedV1MeetingHostUser `json:"host_user,omitempty"`      // 会议主持人
	Owner         *EventV2VCMeetingMeetingStartedV1MeetingOwner    `json:"owner,omitempty"`          // 会议拥有者
}

// EventV2VCMeetingMeetingStartedV1MeetingHostUser ...
type EventV2VCMeetingMeetingStartedV1MeetingHostUser struct {
	ID       *EventV2VCMeetingMeetingStartedV1MeetingHostUserID `json:"id,omitempty"`        // 用户 ID
	UserRole int64                                              `json:"user_role,omitempty"` // 用户会中角色, 可选值有: 1: 普通参会人, 2: 主持人, 3: 联席主持人
	UserType int64                                              `json:"user_type,omitempty"` // 用户类型, 可选值有: 1: lark用户, 2: rooms用户（建议使用open_id作为user_id_type用于获取此类用户）, 3: 文档用户, 4: neo单品用户, 5: neo单品游客用户, 6: pstn用户, 7: sip用户
}

// EventV2VCMeetingMeetingStartedV1MeetingHostUserID ...
type EventV2VCMeetingMeetingStartedV1MeetingHostUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

// EventV2VCMeetingMeetingStartedV1MeetingOwner ...
type EventV2VCMeetingMeetingStartedV1MeetingOwner struct {
	ID       *EventV2VCMeetingMeetingStartedV1MeetingOwnerID `json:"id,omitempty"`        // 用户 ID
	UserRole int64                                           `json:"user_role,omitempty"` // 用户会中角色, 可选值有: 1: 普通参会人, 2: 主持人, 3: 联席主持人
	UserType int64                                           `json:"user_type,omitempty"` // 用户类型, 可选值有: 1: lark用户, 2: rooms用户（建议使用open_id作为user_id_type用于获取此类用户）, 3: 文档用户, 4: neo单品用户, 5: neo单品游客用户, 6: pstn用户, 7: sip用户
}

// EventV2VCMeetingMeetingStartedV1MeetingOwnerID ...
type EventV2VCMeetingMeetingStartedV1MeetingOwnerID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

// EventV2VCMeetingMeetingStartedV1Operator ...
type EventV2VCMeetingMeetingStartedV1Operator struct {
	ID       *EventV2VCMeetingMeetingStartedV1OperatorID `json:"id,omitempty"`        // 用户 ID
	UserRole int64                                       `json:"user_role,omitempty"` // 用户会中角色, 可选值有: 1: 普通参会人, 2: 主持人, 3: 联席主持人
	UserType int64                                       `json:"user_type,omitempty"` // 用户类型, 可选值有: 1: lark用户, 2: rooms用户（建议使用open_id作为user_id_type用于获取此类用户）, 3: 文档用户, 4: neo单品用户, 5: neo单品游客用户, 6: pstn用户, 7: sip用户
}

// EventV2VCMeetingMeetingStartedV1OperatorID ...
type EventV2VCMeetingMeetingStartedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
