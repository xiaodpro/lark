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

// GetHireApplicationList 根据限定条件获取投递列表信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/list
func (r *HireService) GetHireApplicationList(ctx context.Context, request *GetHireApplicationListReq, options ...MethodOptionFunc) (*GetHireApplicationListResp, *Response, error) {
	if r.cli.mock.mockHireGetHireApplicationList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireApplicationList mock enable")
		return r.cli.mock.mockHireGetHireApplicationList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireApplicationList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/applications",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireApplicationListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireApplicationList mock HireGetHireApplicationList method
func (r *Mock) MockHireGetHireApplicationList(f func(ctx context.Context, request *GetHireApplicationListReq, options ...MethodOptionFunc) (*GetHireApplicationListResp, *Response, error)) {
	r.mockHireGetHireApplicationList = f
}

// UnMockHireGetHireApplicationList un-mock HireGetHireApplicationList method
func (r *Mock) UnMockHireGetHireApplicationList() {
	r.mockHireGetHireApplicationList = nil
}

// GetHireApplicationListReq ...
type GetHireApplicationListReq struct {
	ProcessID       *string `query:"process_id" json:"-"`        // 按流程过滤, 招聘流程 ID, 枚举值通过接口「获取招聘流程信息」接口获取, 示例值: "6960663240925956554"
	StageID         *string `query:"stage_id" json:"-"`          // 按招聘阶段过滤, 招聘阶段 ID, 枚举值通过「获取招聘流程信息」接口获取, 示例值: "614218419274131"
	TalentID        *string `query:"talent_id" json:"-"`         // 按人才过滤, 示例值: "6891560630172518670"
	ActiveStatus    *string `query:"active_status" json:"-"`     // 按活跃状态筛选 1=活跃投递, 2=非活跃投递, 3=全部, 示例值: "1"
	JobID           *string `query:"job_id" json:"-"`            // 职位 ID, 示例值: "7334134355464633"
	PageToken       *string `query:"page_token" json:"-"`        // 查询游标, 由上一页结果返回, 第一页不传, 示例值: "1"
	PageSize        *int64  `query:"page_size" json:"-"`         // 每页限制, 每页最大不超过100, 示例值: 100
	UpdateStartTime *string `query:"update_start_time" json:"-"` // 最早更新时间, 毫秒级时间戳, 示例值: "1618500278663"
	UpdateEndTime   *string `query:"update_end_time" json:"-"`   // 最晚更新时间, 毫秒级时间戳, 示例值: "1618500278663"
}

// GetHireApplicationListResp ...
type GetHireApplicationListResp struct {
	Items     []string `json:"items,omitempty"`      // 投递数据列表
	PageToken string   `json:"page_token,omitempty"` // 游标, 翻下一页数据时使用
	HasMore   bool     `json:"has_more,omitempty"`   // 是否还有下一页数据
}

// getHireApplicationListResp ...
type getHireApplicationListResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *GetHireApplicationListResp `json:"data,omitempty"`
}
