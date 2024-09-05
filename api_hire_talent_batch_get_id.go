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

// BatchGetHireTalent 通过手机号或邮箱获取人才 ID。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/talent/batch_get_id
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/candidate-management/talent/batch_get_id
func (r *HireService) BatchGetHireTalent(ctx context.Context, request *BatchGetHireTalentReq, options ...MethodOptionFunc) (*BatchGetHireTalentResp, *Response, error) {
	if r.cli.mock.mockHireBatchGetHireTalent != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Hire#BatchGetHireTalent mock enable")
		return r.cli.mock.mockHireBatchGetHireTalent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "BatchGetHireTalent",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/talents/batch_get_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetHireTalentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireBatchGetHireTalent mock HireBatchGetHireTalent method
func (r *Mock) MockHireBatchGetHireTalent(f func(ctx context.Context, request *BatchGetHireTalentReq, options ...MethodOptionFunc) (*BatchGetHireTalentResp, *Response, error)) {
	r.mockHireBatchGetHireTalent = f
}

// UnMockHireBatchGetHireTalent un-mock HireBatchGetHireTalent method
func (r *Mock) UnMockHireBatchGetHireTalent() {
	r.mockHireBatchGetHireTalent = nil
}

// BatchGetHireTalentReq ...
type BatchGetHireTalentReq struct {
	MobileCode               *string  `json:"mobile_code,omitempty"`                // 手机国家区号, 默认值: 86, 即中国大陆地区, 示例值: "86"
	MobileNumberList         []string `json:"mobile_number_list,omitempty"`         // 手机号, 区号均采用 mobile_code 参数的值, 最多 100 个, 示例值: ["182900291190"]
	EmailList                []string `json:"email_list,omitempty"`                 // 邮箱信息列表, 最多 100 个, 示例值: ["foo@bytedance.com"]
	IdentificationType       *int64   `json:"identification_type,omitempty"`        // 证件类型, 可参考招聘枚举常量文档下的 IdentificationType 枚举定义, 示例值: 1
	IdentificationNumberList []string `json:"identification_number_list,omitempty"` // 证件号, 示例值: ["130xxxxxxx"]
}

// BatchGetHireTalentResp ...
type BatchGetHireTalentResp struct {
	TalentList []*BatchGetHireTalentRespTalent `json:"talent_list,omitempty"` // 人才信息列表
}

// BatchGetHireTalentRespTalent ...
type BatchGetHireTalentRespTalent struct {
	TalentID             string `json:"talent_id,omitempty"`             // 人才 ID
	MobileCode           string `json:"mobile_code,omitempty"`           // 手机国家区号
	MobileNumber         string `json:"mobile_number,omitempty"`         // 手机号
	Email                string `json:"email,omitempty"`                 // 邮箱
	IdentificationType   int64  `json:"identification_type,omitempty"`   // 证件类型, 可参考招聘枚举常量 IdentificationType 枚举定义
	IdentificationNumber string `json:"identification_number,omitempty"` // 证件号
}

// batchGetHireTalentResp ...
type batchGetHireTalentResp struct {
	Code  int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                  `json:"msg,omitempty"`  // 错误描述
	Data  *BatchGetHireTalentResp `json:"data,omitempty"`
	Error *ErrorDetail            `json:"error,omitempty"`
}