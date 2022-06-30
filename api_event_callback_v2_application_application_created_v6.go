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

// EventV2ApplicationApplicationCreatedV6
//
// 当企业内有新的应用被创建时推送此事件{使用示例}(url=/api/tools/api_explore/api_explore_config?project=application&version=v6&resource=application&event=created)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/events/created
func (r *EventCallbackService) HandlerEventV2ApplicationApplicationCreatedV6(f EventV2ApplicationApplicationCreatedV6Handler) {
	r.cli.eventHandler.eventV2ApplicationApplicationCreatedV6Handler = f
}

// EventV2ApplicationApplicationCreatedV6Handler event EventV2ApplicationApplicationCreatedV6 handler
type EventV2ApplicationApplicationCreatedV6Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2ApplicationApplicationCreatedV6) (string, error)

// EventV2ApplicationApplicationCreatedV6 ...
type EventV2ApplicationApplicationCreatedV6 struct {
	OperatorID      *EventV2ApplicationApplicationCreatedV6OperatorID `json:"operator_id,omitempty"`      // 用户 ID
	AppID           string                                            `json:"app_id,omitempty"`           // 应用 ID
	Name            string                                            `json:"name,omitempty"`             // 应用名称
	Description     string                                            `json:"description,omitempty"`      // 应用描述
	Avatar          string                                            `json:"avatar,omitempty"`           // 应用图标链接
	AppSceneType    int64                                             `json:"app_scene_type,omitempty"`   // 应用类型, 0: 自建应用, 1: 应用商店应用
	PrimaryLanguage string                                            `json:"primary_language,omitempty"` // 应用主语言
}

// EventV2ApplicationApplicationCreatedV6OperatorID ...
type EventV2ApplicationApplicationCreatedV6OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
