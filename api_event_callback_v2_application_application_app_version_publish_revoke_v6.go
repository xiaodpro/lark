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

// EventV2ApplicationApplicationAppVersionPublishRevokeV6
//
// 通过订阅该事件, 可接收应用撤回发布申请事件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/events/publish_revoke
func (r *EventCallbackService) HandlerEventV2ApplicationApplicationAppVersionPublishRevokeV6(f EventV2ApplicationApplicationAppVersionPublishRevokeV6Handler) {
	r.cli.eventHandler.eventV2ApplicationApplicationAppVersionPublishRevokeV6Handler = f
}

// EventV2ApplicationApplicationAppVersionPublishRevokeV6Handler event EventV2ApplicationApplicationAppVersionPublishRevokeV6 handler
type EventV2ApplicationApplicationAppVersionPublishRevokeV6Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2ApplicationApplicationAppVersionPublishRevokeV6) (string, error)

// EventV2ApplicationApplicationAppVersionPublishRevokeV6 ...
type EventV2ApplicationApplicationAppVersionPublishRevokeV6 struct {
	OperatorID *EventV2ApplicationApplicationAppVersionPublishRevokeV6OperatorID `json:"operator_id,omitempty"` // 用户 ID
	CreatorID  *EventV2ApplicationApplicationAppVersionPublishRevokeV6CreatorID  `json:"creator_id,omitempty"`  // 用户 ID
	AppID      string                                                            `json:"app_id,omitempty"`      // 撤回应用的 id
	VersionID  string                                                            `json:"version_id,omitempty"`  // 撤回应用的版本 id
}

// EventV2ApplicationApplicationAppVersionPublishRevokeV6OperatorID ...
type EventV2ApplicationApplicationAppVersionPublishRevokeV6OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

// EventV2ApplicationApplicationAppVersionPublishRevokeV6CreatorID ...
type EventV2ApplicationApplicationAppVersionPublishRevokeV6CreatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
