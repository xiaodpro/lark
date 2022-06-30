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

// GetApplicationUsageOverview 查看应用在某一天/某一周/某一个月的使用数据, 可以查看租户整体对应用的使用情况, 也可以分部门查看。
//
// 1. 仅支持企业版/旗舰版租户使用
// 2. 一般每天早上10点产出前一天的数据
// 3. 已经支持的指标包括: 应用的活跃用户数、累计用户数、新增用户数
// 4. 数据从飞书3.46版本开始统计, 使用飞书版本3.45及以下版本的用户数据不会被统计到
// 5. 按照部门查看数据时, 会展示当前部门以及其子部门的整体使用情况
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_usage/overview
func (r *ApplicationService) GetApplicationUsageOverview(ctx context.Context, request *GetApplicationUsageOverviewReq, options ...MethodOptionFunc) (*GetApplicationUsageOverviewResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationUsageOverview != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Application#GetApplicationUsageOverview mock enable")
		return r.cli.mock.mockApplicationGetApplicationUsageOverview(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationUsageOverview",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v6/applications/:app_id/app_usage/overview",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationUsageOverviewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationGetApplicationUsageOverview mock ApplicationGetApplicationUsageOverview method
func (r *Mock) MockApplicationGetApplicationUsageOverview(f func(ctx context.Context, request *GetApplicationUsageOverviewReq, options ...MethodOptionFunc) (*GetApplicationUsageOverviewResp, *Response, error)) {
	r.mockApplicationGetApplicationUsageOverview = f
}

// UnMockApplicationGetApplicationUsageOverview un-mock ApplicationGetApplicationUsageOverview method
func (r *Mock) UnMockApplicationGetApplicationUsageOverview() {
	r.mockApplicationGetApplicationUsageOverview = nil
}

// GetApplicationUsageOverviewReq ...
type GetApplicationUsageOverviewReq struct {
	AppID            string            `path:"app_id" json:"-"`              // 目标应用 ID, 示例值: "cli_9f115af860f7901b"
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 调用中使用的部门ID的类型, 示例值: "open_department_id", 可选值有: `department_id`: 以自定义department_id来标识部门, `open_department_id`: 以open_department_id来标识部门, 默认值: `open_department_id`
	Date             string            `json:"date,omitempty"`               // 查询日期, 格式为yyyy-mm-dd, 若cycle_type为1, date可以为任何自然日；若cycle_type为2, 则输入的date必须为周一； 若cycle_type为3, 则输入的date必须为每月1号, 示例值: "2021-07-08"
	CycleType        int64             `json:"cycle_type,omitempty"`         // 活跃周期的统计类型, 示例值: 1, 可选值有: `1`: 日活, 指自然日, 返回当前日期所在日的数据, `2`: 周活, 指自然周, 返回当前日期所在周的数据。若到查询时当周还没结束, 则返回周一到当前日期的数值。例如在2021/7/15 查询2021/7/5 这一周的数据, 则代表的是2021/7/5 ~ 2021/7/11。但若是在2021/7/8 查询2021/7/5 这一周的数据, 则返回的是2021/7/5 ~ 2021/7/7 的数据, `3`: 月活, 指自然月, 返回当前日期所在月的数据。若不满一个月则返回当月1日到截止日期前的数据。例如在2021/8/15 查询 7月的数据, 则代表2021/7/1~2021/7/31。 若在2021/8/15 查询8月的数据, 则代表2021/8/1~2021/8/14的数据
	DepartmentID     *string           `json:"department_id,omitempty"`      // 查询的部门id, 获取方法可参考[部门ID概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview), 若部门id为空, 则返回当前租户的使用数据；若填写部门id, 则返回当前部门的使用数据（包含子部门的用户）；, 若路径参数中department_id_type为空或者为open_department_id, 则此处应该填写部门的 open_department_id；若路径参数中department_id_type为department_id, 则此处应该填写部门的 department_id, 示例值: "od-4e6ac4d14bcd5071a37a39de902c7141"
	Ability          string            `json:"ability,omitempty"`            // 能力类型, 按能力类型进行筛选, 返回对应能力的活跃数据, 示例值: "app", 可选值有: `app`: 返回应用整体的数据, `mp`: 返回小程序能力的数据, `h5`: 返回网页能力的数据, `bot`: 返回机器人能力的数据
}

// GetApplicationUsageOverviewResp ...
type GetApplicationUsageOverviewResp struct {
	Items []*GetApplicationUsageOverviewRespItem `json:"items,omitempty"` // 员工使用应用概览数据
}

// GetApplicationUsageOverviewRespItem ...
type GetApplicationUsageOverviewRespItem struct {
	MetricName  string `json:"metric_name,omitempty"`  // 指标名称, uv: 活跃用户数, total_users: 累计用户数, new_users: 新增用户数
	MetricValue int64  `json:"metric_value,omitempty"` // 指标值
}

// getApplicationUsageOverviewResp ...
type getApplicationUsageOverviewResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *GetApplicationUsageOverviewResp `json:"data,omitempty"`
}
