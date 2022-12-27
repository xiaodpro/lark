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

// EventV2IMMessageReactionDeletedV1 消息被删除某一个表情回复后触发此事件{使用示例}(url=/api/tools/api_explore/api_explore_config?project=im&version=v1&resource=message.reaction&event=deleted)
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 具备[获取单聊、群组消息] 或 [获取与发送单聊、群组消息]权限, 并订阅 [消息与群组] 分类下的 [消息被取消reaction] 事件才可接收推送
// - 机器人只能收到所在群聊内的消息被删除表情回复事件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/events/deleted
func (r *EventCallbackService) HandlerEventV2IMMessageReactionDeletedV1(f EventV2IMMessageReactionDeletedV1Handler) {
	r.cli.eventHandler.eventV2IMMessageReactionDeletedV1Handler = f
}

// EventV2IMMessageReactionDeletedV1Handler event EventV2IMMessageReactionDeletedV1 handler
type EventV2IMMessageReactionDeletedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2IMMessageReactionDeletedV1) (string, error)

// EventV2IMMessageReactionDeletedV1 ...
type EventV2IMMessageReactionDeletedV1 struct {
	MessageID    string                                         `json:"message_id,omitempty"`    // 消息的 open_message_id
	ReactionType *EventV2IMMessageReactionDeletedV1ReactionType `json:"reaction_type,omitempty"` // 表情回复的资源类型
	OperatorType string                                         `json:"operator_type,omitempty"` // 操作人类型, 注意事项: 如果操作人类型是"user", 则会返回 [user_id], 如果操作人类型是"app", 则会返回 [app_id]
	UserID       *EventV2IMMessageReactionDeletedV1UserID       `json:"user_id,omitempty"`       // 用户 ID
	AppID        string                                         `json:"app_id,omitempty"`        // 应用 ID
	ActionTime   string                                         `json:"action_time,omitempty"`   // 表情回复被添加时的时间戳（单位: ms）
}

// EventV2IMMessageReactionDeletedV1ReactionType ...
type EventV2IMMessageReactionDeletedV1ReactionType struct {
	EmojiType string `json:"emoji_type,omitempty"` // emoji类型 [emoji类型列举](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)
}

// EventV2IMMessageReactionDeletedV1UserID ...
type EventV2IMMessageReactionDeletedV1UserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
