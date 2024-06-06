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

// CreateHelpdeskFAQ 该接口用于创建知识库。
//
// 注意事项:
// user_access_token 访问, 需要操作者是当前服务台的客服、管理员或所有者
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/create
// new doc: https://open.feishu.cn/document/server-docs/helpdesk-v1/faq-management/faq/create
func (r *HelpdeskService) CreateHelpdeskFAQ(ctx context.Context, request *CreateHelpdeskFAQReq, options ...MethodOptionFunc) (*CreateHelpdeskFAQResp, *Response, error) {
	if r.cli.mock.mockHelpdeskCreateHelpdeskFAQ != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Helpdesk#CreateHelpdeskFAQ mock enable")
		return r.cli.mock.mockHelpdeskCreateHelpdeskFAQ(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "CreateHelpdeskFAQ",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/faqs",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(createHelpdeskFAQResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskCreateHelpdeskFAQ mock HelpdeskCreateHelpdeskFAQ method
func (r *Mock) MockHelpdeskCreateHelpdeskFAQ(f func(ctx context.Context, request *CreateHelpdeskFAQReq, options ...MethodOptionFunc) (*CreateHelpdeskFAQResp, *Response, error)) {
	r.mockHelpdeskCreateHelpdeskFAQ = f
}

// UnMockHelpdeskCreateHelpdeskFAQ un-mock HelpdeskCreateHelpdeskFAQ method
func (r *Mock) UnMockHelpdeskCreateHelpdeskFAQ() {
	r.mockHelpdeskCreateHelpdeskFAQ = nil
}

// CreateHelpdeskFAQReq ...
type CreateHelpdeskFAQReq struct {
	FAQ *CreateHelpdeskFAQReqFAQ `json:"faq,omitempty"` // 知识库详情
}

// CreateHelpdeskFAQReqFAQ ...
type CreateHelpdeskFAQReqFAQ struct {
	CategoryID     *string  `json:"category_id,omitempty"`     // 知识库分类ID, 示例值: "6836004780707807251"
	Question       string   `json:"question,omitempty"`        // 问题, 示例值: "问题"
	Answer         *string  `json:"answer,omitempty"`          // 答案, 示例值: "答案"
	AnswerRichText *string  `json:"answer_richtext,omitempty"` // 富文本答案和答案必须有一个必填。Json Array格式, 富文本结构请见[了解更多: 富文本](https://open.feishu.cn/document/ukTMukTMukTM/uITM0YjLyEDN24iMxQjN), 示例值: "{\"content\":\"这只是一个测试, 医保问题\", \"type\":\"text\"}"
	Tags           []string `json:"tags,omitempty"`            // 相似问题, 示例值: ["问", "题"]
}

// CreateHelpdeskFAQResp ...
type CreateHelpdeskFAQResp struct {
	FAQ *CreateHelpdeskFAQRespFAQ `json:"faq,omitempty"` // 知识库详情
}

// CreateHelpdeskFAQRespFAQ ...
type CreateHelpdeskFAQRespFAQ struct {
	FAQID          string                                    `json:"faq_id,omitempty"`          // 知识库ID
	ID             string                                    `json:"id,omitempty"`              // 知识库旧版ID, 请使用faq_id
	HelpdeskID     string                                    `json:"helpdesk_id,omitempty"`     // 服务台ID
	Question       string                                    `json:"question,omitempty"`        // 问题
	Answer         string                                    `json:"answer,omitempty"`          // 答案
	AnswerRichText []*CreateHelpdeskFAQRespFAQAnswerRichText `json:"answer_richtext,omitempty"` // 富文本答案
	CreateTime     int64                                     `json:"create_time,omitempty"`     // 创建时间
	UpdateTime     int64                                     `json:"update_time,omitempty"`     // 修改时间
	Categories     []*HelpdeskCategory                       `json:"categories,omitempty"`      // 分类
	Tags           []string                                  `json:"tags,omitempty"`            // 相似问题列表
	ExpireTime     int64                                     `json:"expire_time,omitempty"`     // 失效时间
	UpdateUser     *CreateHelpdeskFAQRespFAQUpdateUser       `json:"update_user,omitempty"`     // 更新用户
	CreateUser     *CreateHelpdeskFAQRespFAQCreateUser       `json:"create_user,omitempty"`     // 创建用户
}

// CreateHelpdeskFAQRespFAQAnswerRichText ...
type CreateHelpdeskFAQRespFAQAnswerRichText struct {
	Content string `json:"content,omitempty"` // 内容
	Type    string `json:"type,omitempty"`    // 类型
}

// CreateHelpdeskFAQRespFAQCreateUser ...
type CreateHelpdeskFAQRespFAQCreateUser struct {
	ID         string `json:"id,omitempty"`         // 用户ID
	AvatarURL  string `json:"avatar_url,omitempty"` // 用户头像url
	Name       string `json:"name,omitempty"`       // 用户名
	Department string `json:"department,omitempty"` // 所在部门名称
	City       string `json:"city,omitempty"`       // 城市
	Country    string `json:"country,omitempty"`    // 国家代号(CountryCode), 参考: http://www.mamicode.com/info-detail-2186501.html
}

// CreateHelpdeskFAQRespFAQUpdateUser ...
type CreateHelpdeskFAQRespFAQUpdateUser struct {
	ID         string `json:"id,omitempty"`         // 用户ID
	AvatarURL  string `json:"avatar_url,omitempty"` // 用户头像url
	Name       string `json:"name,omitempty"`       // 用户名
	Department string `json:"department,omitempty"` // 所在部门名称
	City       string `json:"city,omitempty"`       // 城市
	Country    string `json:"country,omitempty"`    // 国家代号(CountryCode), 参考: http://www.mamicode.com/info-detail-2186501.html
}

// createHelpdeskFAQResp ...
type createHelpdeskFAQResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *CreateHelpdeskFAQResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}
