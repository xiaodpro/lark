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

// GetOKRReview 根据周期和用户查询复盘信息。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/review/query
// new doc: https://open.feishu.cn/document/server-docs/okr-v1/review/query
func (r *OKRService) GetOKRReview(ctx context.Context, request *GetOKRReviewReq, options ...MethodOptionFunc) (*GetOKRReviewResp, *Response, error) {
	if r.cli.mock.mockOKRGetOKRReview != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] OKR#GetOKRReview mock enable")
		return r.cli.mock.mockOKRGetOKRReview(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "OKR",
		API:                   "GetOKRReview",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/okr/v1/reviews/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getOKRReviewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockOKRGetOKRReview mock OKRGetOKRReview method
func (r *Mock) MockOKRGetOKRReview(f func(ctx context.Context, request *GetOKRReviewReq, options ...MethodOptionFunc) (*GetOKRReviewResp, *Response, error)) {
	r.mockOKRGetOKRReview = f
}

// UnMockOKRGetOKRReview un-mock OKRGetOKRReview method
func (r *Mock) UnMockOKRGetOKRReview() {
	r.mockOKRGetOKRReview = nil
}

// GetOKRReviewReq ...
type GetOKRReviewReq struct {
	UserIDType *IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_admin_id: 以people_admin_id来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	UserIDs    []string `query:"user_ids" json:"-"`     // 目标用户id列表, 最多5个, 示例值: ou-asdasdasdasdasd, 最大长度: `5`
	PeriodIDs  []string `query:"period_ids" json:"-"`   // period_id列表, 最多5个, 示例值: 6951461264858777132, 最大长度: `5`
}

// GetOKRReviewResp ...
type GetOKRReviewResp struct {
	ReviewList []*GetOKRReviewRespReview `json:"review_list,omitempty"` // OKR复盘 列表
}

// GetOKRReviewRespReview ...
type GetOKRReviewRespReview struct {
	UserID           *GetOKRReviewRespReviewUserID         `json:"user_id,omitempty"`            // 复盘的用户
	ReviewPeriodList []*GetOKRReviewRespReviewReviewPeriod `json:"review_period_list,omitempty"` // 用户对应的OKR复盘列表
}

// GetOKRReviewRespReviewReviewPeriod ...
type GetOKRReviewRespReviewReviewPeriod struct {
	PeriodID           string                                              `json:"period_id,omitempty"`            // 周期ID
	CycleReviewList    []*GetOKRReviewRespReviewReviewPeriodCycleReview    `json:"cycle_review_list,omitempty"`    // 复盘文档
	ProgressReportList []*GetOKRReviewRespReviewReviewPeriodProgressReport `json:"progress_report_list,omitempty"` // 进展报告
}

// GetOKRReviewRespReviewReviewPeriodCycleReview ...
type GetOKRReviewRespReviewReviewPeriodCycleReview struct {
	URL        string `json:"url,omitempty"`         // 文档链接
	CreateTime string `json:"create_time,omitempty"` // 创建时间 毫秒
}

// GetOKRReviewRespReviewReviewPeriodProgressReport ...
type GetOKRReviewRespReviewReviewPeriodProgressReport struct {
	URL        string `json:"url,omitempty"`         // 文档链接
	CreateTime string `json:"create_time,omitempty"` // 创建时间 毫秒
}

// GetOKRReviewRespReviewUserID ...
type GetOKRReviewRespReviewUserID struct {
	OpenID string `json:"open_id,omitempty"` // 用户的 open_id
	UserID string `json:"user_id,omitempty"` // 用户的 user_id
}

// getOKRReviewResp ...
type getOKRReviewResp struct {
	Code  int64             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string            `json:"msg,omitempty"`  // 错误描述
	Data  *GetOKRReviewResp `json:"data,omitempty"`
	Error *ErrorDetail      `json:"error,omitempty"`
}
