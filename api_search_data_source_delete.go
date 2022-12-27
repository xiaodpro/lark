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

// DeleteSearchDataSource 删除一个已存在的数据源。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/delete
func (r *SearchService) DeleteSearchDataSource(ctx context.Context, request *DeleteSearchDataSourceReq, options ...MethodOptionFunc) (*DeleteSearchDataSourceResp, *Response, error) {
	if r.cli.mock.mockSearchDeleteSearchDataSource != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#DeleteSearchDataSource mock enable")
		return r.cli.mock.mockSearchDeleteSearchDataSource(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "DeleteSearchDataSource",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/search/v2/data_sources/:data_source_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteSearchDataSourceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockSearchDeleteSearchDataSource mock SearchDeleteSearchDataSource method
func (r *Mock) MockSearchDeleteSearchDataSource(f func(ctx context.Context, request *DeleteSearchDataSourceReq, options ...MethodOptionFunc) (*DeleteSearchDataSourceResp, *Response, error)) {
	r.mockSearchDeleteSearchDataSource = f
}

// UnMockSearchDeleteSearchDataSource un-mock SearchDeleteSearchDataSource method
func (r *Mock) UnMockSearchDeleteSearchDataSource() {
	r.mockSearchDeleteSearchDataSource = nil
}

// DeleteSearchDataSourceReq ...
type DeleteSearchDataSourceReq struct {
	DataSourceID string `path:"data_source_id" json:"-"` // 数据源的唯一标识, 示例值: "6953903108179099667"
}

// DeleteSearchDataSourceResp ...
type DeleteSearchDataSourceResp struct {
}

// deleteSearchDataSourceResp ...
type deleteSearchDataSourceResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *DeleteSearchDataSourceResp `json:"data,omitempty"`
}
