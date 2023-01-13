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

// EventV2VCMeetingRecordingReadyV1 发生在录制文件上传完毕时。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=vc&version=v1&resource=meeting&event=recording_ready)
//
// 收到该事件后, 方可进行录制文件获取、授权等操作。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/events/recording_ready
func (r *EventCallbackService) HandlerEventV2VCMeetingRecordingReadyV1(f EventV2VCMeetingRecordingReadyV1Handler) {
	r.cli.eventHandler.eventV2VCMeetingRecordingReadyV1Handler = f
}

// EventV2VCMeetingRecordingReadyV1Handler event EventV2VCMeetingRecordingReadyV1 handler
type EventV2VCMeetingRecordingReadyV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2VCMeetingRecordingReadyV1) (string, error)

// EventV2VCMeetingRecordingReadyV1 ...
type EventV2VCMeetingRecordingReadyV1 struct {
	Meeting  *EventV2VCMeetingRecordingReadyV1Meeting `json:"meeting,omitempty"`  // 会议数据
	URL      string                                   `json:"url,omitempty"`      // 会议录制链接
	Duration string                                   `json:"duration,omitempty"` // 录制总时长（单位msec）
}

// EventV2VCMeetingRecordingReadyV1Meeting ...
type EventV2VCMeetingRecordingReadyV1Meeting struct {
	ID        string                                        `json:"id,omitempty"`         // 会议ID（视频会议的唯一标识, 视频会议开始后才会产生）
	Topic     string                                        `json:"topic,omitempty"`      // 会议主题
	MeetingNo string                                        `json:"meeting_no,omitempty"` // 9位会议号（飞书用户可通过输入9位会议号快捷入会）
	Owner     *EventV2VCMeetingRecordingReadyV1MeetingOwner `json:"owner,omitempty"`      // 会议拥有者
}

// EventV2VCMeetingRecordingReadyV1MeetingOwner ...
type EventV2VCMeetingRecordingReadyV1MeetingOwner struct {
	ID *EventV2VCMeetingRecordingReadyV1MeetingOwnerID `json:"id,omitempty"` // 用户 ID
}

// EventV2VCMeetingRecordingReadyV1MeetingOwnerID ...
type EventV2VCMeetingRecordingReadyV1MeetingOwnerID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
