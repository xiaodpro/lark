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

// EventV1OrderPaid
//
// 了解事件订阅的使用场景和配置流程, 请点击查看 [事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
// 用户购买应用商店付费应用成功后发送给应用ISV的通知事件。
// - 订阅条件: 只有应用商店应用才能订阅此事件。自建应用无此事件。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/event/public-app-purchase
func (r *EventCallbackService) HandlerEventV1OrderPaid(f EventV1OrderPaidHandler) {
	r.cli.eventHandler.eventV1OrderPaidHandler = f
}

// EventV1OrderPaidHandler event EventV1OrderPaid handler
type EventV1OrderPaidHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1OrderPaid) (string, error)

// EventV1OrderPaid ...
type EventV1OrderPaid struct {
	Type          string `json:"type,omitempty"`            // 事件类型. 如: order_paid
	AppID         string `json:"app_id,omitempty"`          // 应用ID. 如: cli_9daeceab98721136
	OrderID       string `json:"order_id,omitempty"`        // 用户购买付费方案时对订单ID 可作为唯一标识. 如: 6704894492631105539
	PricePlanID   string `json:"price_plan_id,omitempty"`   // 付费方案ID. 如: price_9d86fa1333b8110c
	PricePlanType string `json:"price_plan_type,omitempty"` // 用户购买方案类型 "trial" -试用；"permanent"-免费；"per_year"-企业年付费；"per_month"-企业月付费；"per_seat_per_year"-按人按年付费；"per_seat_per_month"-按人按月付费；"permanent_count"-按次付费. 如: per_seat_per_month
	Seats         int64  `json:"seats,omitempty"`           // 表示购买了多少人份. 如: 20
	BuyCount      int64  `json:"buy_count,omitempty"`       // 套餐购买数量 目前都为1. 如: 1
	CreateTime    string `json:"create_time,omitempty"`     // 如: 1502199207
	PayTime       string `json:"pay_time,omitempty"`        // 如: 1502199209
	BuyType       string `json:"buy_type,omitempty"`        // 购买类型 buy普通购买 upgrade为升级购买 renew为续费购买. 如: buy
	SrcOrderID    string `json:"src_order_id,omitempty"`    // 当前为升级购买时(buy_type 为upgrade), 该字段表示原订单ID, 升级后原订单失效, 状态变为已升级(业务方需要处理). 如: 6704894492631105539
	OrderPayPrice int64  `json:"order_pay_price,omitempty"` // 订单支付价格 单位分, . 如: 10000
	TenantKey     string `json:"tenant_key,omitempty"`      // 购买应用的企业标示. 如: 2f98c01bc23f6847
}
