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

// DeleteHelpdeskCategory 该接口用于删除知识库分类详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/delete
func (r *HelpdeskService) DeleteHelpdeskCategory(ctx context.Context, request *DeleteHelpdeskCategoryReq, options ...MethodOptionFunc) (*DeleteHelpdeskCategoryResp, *Response, error) {
	if r.cli.mock.mockHelpdeskDeleteHelpdeskCategory != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#DeleteHelpdeskCategory mock enable")
		return r.cli.mock.mockHelpdeskDeleteHelpdeskCategory(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "DeleteHelpdeskCategory",
		Method:              "DELETE",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/categories/:id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(deleteHelpdeskCategoryResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskDeleteHelpdeskCategory mock HelpdeskDeleteHelpdeskCategory method
func (r *Mock) MockHelpdeskDeleteHelpdeskCategory(f func(ctx context.Context, request *DeleteHelpdeskCategoryReq, options ...MethodOptionFunc) (*DeleteHelpdeskCategoryResp, *Response, error)) {
	r.mockHelpdeskDeleteHelpdeskCategory = f
}

// UnMockHelpdeskDeleteHelpdeskCategory un-mock HelpdeskDeleteHelpdeskCategory method
func (r *Mock) UnMockHelpdeskDeleteHelpdeskCategory() {
	r.mockHelpdeskDeleteHelpdeskCategory = nil
}

// DeleteHelpdeskCategoryReq ...
type DeleteHelpdeskCategoryReq struct {
	ID string `path:"id" json:"-"` // 知识库分类ID, 示例值: "6948728206392295444"
}

// DeleteHelpdeskCategoryResp ...
type DeleteHelpdeskCategoryResp struct {
}

// deleteHelpdeskCategoryResp ...
type deleteHelpdeskCategoryResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *DeleteHelpdeskCategoryResp `json:"data,omitempty"`
}
