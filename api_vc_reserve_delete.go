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

// DeleteVCReserve 删除一个预约。
//
// 只能删除归属于自己的预约；删除后数据不可恢复
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/delete
func (r *VCService) DeleteVCReserve(ctx context.Context, request *DeleteVCReserveReq, options ...MethodOptionFunc) (*DeleteVCReserveResp, *Response, error) {
	if r.cli.mock.mockVCDeleteVCReserve != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#DeleteVCReserve mock enable")
		return r.cli.mock.mockVCDeleteVCReserve(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "VC",
		API:                 "DeleteVCReserve",
		Method:              "DELETE",
		URL:                 r.cli.openBaseURL + "/open-apis/vc/v1/reserves/:reserve_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(deleteVCReserveResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCDeleteVCReserve mock VCDeleteVCReserve method
func (r *Mock) MockVCDeleteVCReserve(f func(ctx context.Context, request *DeleteVCReserveReq, options ...MethodOptionFunc) (*DeleteVCReserveResp, *Response, error)) {
	r.mockVCDeleteVCReserve = f
}

// UnMockVCDeleteVCReserve un-mock VCDeleteVCReserve method
func (r *Mock) UnMockVCDeleteVCReserve() {
	r.mockVCDeleteVCReserve = nil
}

// DeleteVCReserveReq ...
type DeleteVCReserveReq struct {
	ReserveID string `path:"reserve_id" json:"-"` // 预约ID（预约的唯一标识）, 示例值: "6911188411932033028"
}

// DeleteVCReserveResp ...
type DeleteVCReserveResp struct {
}

// deleteVCReserveResp ...
type deleteVCReserveResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *DeleteVCReserveResp `json:"data,omitempty"`
}
