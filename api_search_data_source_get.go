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

// GetSearchDataSource 获取已经创建的数据源
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/get
func (r *SearchService) GetSearchDataSource(ctx context.Context, request *GetSearchDataSourceReq, options ...MethodOptionFunc) (*GetSearchDataSourceResp, *Response, error) {
	if r.cli.mock.mockSearchGetSearchDataSource != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#GetSearchDataSource mock enable")
		return r.cli.mock.mockSearchGetSearchDataSource(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "GetSearchDataSource",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/search/v2/data_sources/:data_source_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getSearchDataSourceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockSearchGetSearchDataSource mock SearchGetSearchDataSource method
func (r *Mock) MockSearchGetSearchDataSource(f func(ctx context.Context, request *GetSearchDataSourceReq, options ...MethodOptionFunc) (*GetSearchDataSourceResp, *Response, error)) {
	r.mockSearchGetSearchDataSource = f
}

// UnMockSearchGetSearchDataSource un-mock SearchGetSearchDataSource method
func (r *Mock) UnMockSearchGetSearchDataSource() {
	r.mockSearchGetSearchDataSource = nil
}

// GetSearchDataSourceReq ...
type GetSearchDataSourceReq struct {
	DataSourceID string `path:"data_source_id" json:"-"` // 数据源的唯一标识, 示例值: "service_ticket"
}

// GetSearchDataSourceResp ...
type GetSearchDataSourceResp struct {
	DataSource *GetSearchDataSourceRespDataSource `json:"data_source,omitempty"` // 数据源实例
}

// GetSearchDataSourceRespDataSource ...
type GetSearchDataSourceRespDataSource struct {
	ID            string `json:"id,omitempty"`              // 数据源的唯一标识
	Name          string `json:"name,omitempty"`            // data_source的展示名称
	State         int64  `json:"state,omitempty"`           // 数据源状态, 0-未上线, 1-已上线, 可选值有: `0`: 未上线, `1`: 已上线
	Description   string `json:"description,omitempty"`     // 对于数据源的描述
	CreateTime    string `json:"create_time,omitempty"`     // 创建时间, 使用Unix时间戳, 单位为“秒”
	UpdateTime    string `json:"update_time,omitempty"`     // 更新时间, 使用Unix时间戳, 单位为“秒”
	IsExceedQuota bool   `json:"is_exceed_quota,omitempty"` // 是否超限
}

// getSearchDataSourceResp ...
type getSearchDataSourceResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetSearchDataSourceResp `json:"data,omitempty"`
}
