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

// GetFaceVerifyAuthResult
//
// 无源人脸比对接入需申请白名单, 接入前请联系飞书开放平台工作人员, 邮箱: openplatform@bytedance.com。
// 无源人脸比对流程, 开发者后台通过调用此接口请求飞书后台, 对本次活体比对结果做校验。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/face/query-recognition-result
func (r *HumanAuthService) GetFaceVerifyAuthResult(ctx context.Context, request *GetFaceVerifyAuthResultReq, options ...MethodOptionFunc) (*GetFaceVerifyAuthResultResp, *Response, error) {
	if r.cli.mock.mockHumanAuthGetFaceVerifyAuthResult != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] HumanAuth#GetFaceVerifyAuthResult mock enable")
		return r.cli.mock.mockHumanAuthGetFaceVerifyAuthResult(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "HumanAuth",
		API:                   "GetFaceVerifyAuthResult",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/face_verify/v1/query_auth_result",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getFaceVerifyAuthResultResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHumanAuthGetFaceVerifyAuthResult mock HumanAuthGetFaceVerifyAuthResult method
func (r *Mock) MockHumanAuthGetFaceVerifyAuthResult(f func(ctx context.Context, request *GetFaceVerifyAuthResultReq, options ...MethodOptionFunc) (*GetFaceVerifyAuthResultResp, *Response, error)) {
	r.mockHumanAuthGetFaceVerifyAuthResult = f
}

// UnMockHumanAuthGetFaceVerifyAuthResult un-mock HumanAuthGetFaceVerifyAuthResult method
func (r *Mock) UnMockHumanAuthGetFaceVerifyAuthResult() {
	r.mockHumanAuthGetFaceVerifyAuthResult = nil
}

// GetFaceVerifyAuthResultReq ...
type GetFaceVerifyAuthResultReq struct {
	ReqOrderNo string  `query:"req_order_no" json:"-"` // 人脸识别单次唯一标识, 由`tt.startFaceVerify`接口返回
	OpenID     *string `query:"open_id" json:"-"`      // 用户应用标识, 与employee_id二选其一
	EmployeeID *string `query:"employee_id" json:"-"`  // 用户租户标识, 与open_id二选其一
}

// GetFaceVerifyAuthResultResp ...
type GetFaceVerifyAuthResultResp struct {
	AuthState     int64 `json:"auth_state,omitempty"`     // 认证结果, 0: 认证中, 1: 认证成功, 1: 认证失败
	AuthTimpstamp int64 `json:"auth_timpstamp,omitempty"` // 认证时间, unix 时间戳
}

// getFaceVerifyAuthResultResp ...
type getFaceVerifyAuthResultResp struct {
	Code int64                        `json:"code,omitempty"` // 返回码, 非0为失败
	Msg  string                       `json:"msg,omitempty"`  // 返回信息, 返回码的描述
	Data *GetFaceVerifyAuthResultResp `json:"data,omitempty"` // 业务数据
}
