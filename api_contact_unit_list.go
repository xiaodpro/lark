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

// GetContactUnitList 通过该接口获取企业的单位列表, 需获取单位的权限
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/list
func (r *ContactService) GetContactUnitList(ctx context.Context, request *GetContactUnitListReq, options ...MethodOptionFunc) (*GetContactUnitListResp, *Response, error) {
	if r.cli.mock.mockContactGetContactUnitList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetContactUnitList mock enable")
		return r.cli.mock.mockContactGetContactUnitList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetContactUnitList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/unit",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getContactUnitListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactGetContactUnitList mock ContactGetContactUnitList method
func (r *Mock) MockContactGetContactUnitList(f func(ctx context.Context, request *GetContactUnitListReq, options ...MethodOptionFunc) (*GetContactUnitListResp, *Response, error)) {
	r.mockContactGetContactUnitList = f
}

// UnMockContactGetContactUnitList un-mock ContactGetContactUnitList method
func (r *Mock) UnMockContactGetContactUnitList() {
	r.mockContactGetContactUnitList = nil
}

// GetContactUnitListReq ...
type GetContactUnitListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 默认50, 取值范围 1-100, 示例值: 50, 默认值: `50`, 最大值: `100`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "AQD9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JKiSIkdexPw="
}

// GetContactUnitListResp ...
type GetContactUnitListResp struct {
	Unitlist  []*GetContactUnitListRespUnit `json:"unitlist,omitempty"`   // 单位列表
	HasMore   bool                          `json:"has_more,omitempty"`   // 是否还有分页数据
	PageToken string                        `json:"page_token,omitempty"` // 分页下次调用的page_token值
}

// GetContactUnitListRespUnit ...
type GetContactUnitListRespUnit struct {
	UnitID   string `json:"unit_id,omitempty"`   // 单位的自定义ID
	Name     string `json:"name,omitempty"`      // 单位的名字
	UnitType string `json:"unit_type,omitempty"` // 单位的类型
}

// getContactUnitListResp ...
type getContactUnitListResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetContactUnitListResp `json:"data,omitempty"`
}
