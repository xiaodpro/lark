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

// EventV1RemoveBot
//
// 为了更好地提升该事件的安全性, 我们对其进行了升级, 请尽快迁移至[新版本>>](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-bot/events/deleted)
// 机器人被从群聊中移除时触发此事件。
// - 依赖条件: 应用必须开启了[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ugzMugzMugzM/event/bot-removed-from-group
func (r *EventCallbackService) HandlerEventV1RemoveBot(f EventV1RemoveBotHandler) {
	r.cli.eventHandler.eventV1RemoveBotHandler = f
}

// EventV1RemoveBotHandler event EventV1RemoveBot handler
type EventV1RemoveBotHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1RemoveBot) (string, error)

// EventV1RemoveBot ...
type EventV1RemoveBot struct {
	AppID               string                              `json:"app_id,omitempty"`                 // 如: cli_9c8609450f78d102
	ChatI18nNames       *EventV1RemoveBotEventChatI18nNames `json:"chat_i18n_names,omitempty"`        // 群名称国际化字段
	ChatName            string                              `json:"chat_name,omitempty"`              // 如: 群名称
	ChatOwnerEmployeeID string                              `json:"chat_owner_employee_id,omitempty"` // 群主的employee_id（即“用户ID”。如果群主是机器人则没有这个字段, 仅企业自建应用返回）. 如: ca51d83b
	ChatOwnerName       string                              `json:"chat_owner_name,omitempty"`        // 群主姓名. 如: xxx
	ChatOwnerOpenID     string                              `json:"chat_owner_open_id,omitempty"`     // 群主的open_id. 如: ou_18eac85d35a26f989317ad4f02e8bbbb
	OpenChatID          string                              `json:"open_chat_id,omitempty"`           // 群聊的id. 如: oc_e520983d3e4f5ec7306069bffe672aa3
	OperatorEmployeeID  string                              `json:"operator_employee_id,omitempty"`   // 操作者的emplolyee_id, 仅企业自建应用返回. 如: ca51d83b
	OperatorName        string                              `json:"operator_name,omitempty"`          // 操作者姓名. 如: yyy
	OperatorOpenID      string                              `json:"operator_open_id,omitempty"`       // 操作者的open_id. 如: ou_18eac85d35a26f989317ad4f02e8bbbb
	OwnerIsBot          bool                                `json:"owner_is_bot,omitempty"`           // 群主是否是机器人. 如: false
	TenantKey           string                              `json:"tenant_key,omitempty"`             // 企业标识. 如: 736588c9260f175d
	Type                string                              `json:"type,omitempty"`                   // 移除机器人: remove_bot. 如: remove_bot
}

// EventV1RemoveBotEventChatI18nNames ...
type EventV1RemoveBotEventChatI18nNames struct {
	EnUs string `json:"en_us,omitempty"` // 如: 英文标题
	ZhCn string `json:"zh_cn,omitempty"` // 如: 中文标题
}
