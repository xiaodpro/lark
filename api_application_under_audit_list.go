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

// GetApplicationUnderAuditList 查看本企业下所有待审核的自建应用列表
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/underauditlist
func (r *ApplicationService) GetApplicationUnderAuditList(ctx context.Context, request *GetApplicationUnderAuditListReq, options ...MethodOptionFunc) (*GetApplicationUnderAuditListResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationUnderAuditList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Application#GetApplicationUnderAuditList mock enable")
		return r.cli.mock.mockApplicationGetApplicationUnderAuditList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationUnderAuditList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v6/applications/underauditlist",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationUnderAuditListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationGetApplicationUnderAuditList mock ApplicationGetApplicationUnderAuditList method
func (r *Mock) MockApplicationGetApplicationUnderAuditList(f func(ctx context.Context, request *GetApplicationUnderAuditListReq, options ...MethodOptionFunc) (*GetApplicationUnderAuditListResp, *Response, error)) {
	r.mockApplicationGetApplicationUnderAuditList = f
}

// UnMockApplicationGetApplicationUnderAuditList un-mock ApplicationGetApplicationUnderAuditList method
func (r *Mock) UnMockApplicationGetApplicationUnderAuditList() {
	r.mockApplicationGetApplicationUnderAuditList = nil
}

// GetApplicationUnderAuditListReq ...
type GetApplicationUnderAuditListReq struct {
	Lang       string  `query:"lang" json:"-"`         // 指定返回的语言, 示例值: "zh_cn", 可选值有: zh_cn: 中文, en_us: 英文, ja_jp: 日文, 最小长度: `1` 字符
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "new-e3c5a0627cdf0c2e057da7257b90376a"
	PageSize   *int64  `query:"page_size" json:"-"`    // 分页大小, 示例值: 10, 默认值: `20`, 最大值: `50`
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetApplicationUnderAuditListResp ...
type GetApplicationUnderAuditListResp struct {
	Items     []*GetApplicationUnderAuditListRespItem `json:"items,omitempty"`      // 待审核应用列表
	HasMore   bool                                    `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                                  `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetApplicationUnderAuditListRespItem ...
type GetApplicationUnderAuditListRespItem struct {
	AppID            string                                       `json:"app_id,omitempty"`             // 应用的 app_id
	CreatorID        string                                       `json:"creator_id,omitempty"`         // 应用创建者（所有者）
	Status           int64                                        `json:"status,omitempty"`             // 应用状态, 可选值有: 0: 停用状态, 1: 启用状态, 2: 未启用状态, 3: 未知状态
	SceneType        int64                                        `json:"scene_type,omitempty"`         // 应用类型, 可选值有: 0: 自建应用, 1: 应用商店应用, 2: 个人应用商店应用, 3: 未知应用类型
	PaymentType      int64                                        `json:"payment_type,omitempty"`       // 付费类型, 可选值有: 0: 免费, 1: 付费
	RedirectURLs     []string                                     `json:"redirect_urls,omitempty"`      // 安全设置中的重定向 URL
	OnlineVersionID  string                                       `json:"online_version_id,omitempty"`  // 发布在线上的应用版本 ID, 若没有则为空
	UnauditVersionID string                                       `json:"unaudit_version_id,omitempty"` // 在审核中的版本 ID, 若没有则为空
	AppName          string                                       `json:"app_name,omitempty"`           // 应用名称
	AvatarURL        string                                       `json:"avatar_url,omitempty"`         // 应用图标 url
	Description      string                                       `json:"description,omitempty"`        // 应用默认描述
	Scopes           []*GetApplicationUnderAuditListRespItemScope `json:"scopes,omitempty"`             // 应用权限列表
	BackHomeURL      string                                       `json:"back_home_url,omitempty"`      // 后台主页地址
	I18n             []*GetApplicationUnderAuditListRespItemI18n  `json:"i18n,omitempty"`               // 应用的国际化信息列表
	PrimaryLanguage  string                                       `json:"primary_language,omitempty"`   // 应用主语言, 可选值有: zh_cn: 中文, en_us: 英文, ja_jp: 日文
	CommonCategories []string                                     `json:"common_categories,omitempty"`  // 应用分类的国际化描述
	Owner            *GetApplicationUnderAuditListRespItemOwner   `json:"owner,omitempty"`              // 应用的所有者信息
}

// GetApplicationUnderAuditListRespItemI18n ...
type GetApplicationUnderAuditListRespItemI18n struct {
	I18nKey     string `json:"i18n_key,omitempty"`    // 国际化语言的 key, 可选值有: zh_cn: 中文, en_us: 英文, ja_jp: 日文
	Name        string `json:"name,omitempty"`        // 应用国际化名称
	Description string `json:"description,omitempty"` // 应用国际化描述（副标题）
	HelpUse     string `json:"help_use,omitempty"`    // 帮助国际化文档链接
}

// GetApplicationUnderAuditListRespItemOwner ...
type GetApplicationUnderAuditListRespItemOwner struct {
	Type     int64  `json:"type,omitempty"`      // 应用所有者类型, 可选值有: 0: 飞书科技, 1: 飞书合作伙伴, 2: 企业内成员
	OwnerID  string `json:"owner_id,omitempty"`  // 应用所有者ID
	Name     string `json:"name,omitempty"`      // 应用开发商名称(仅商店应用返回)
	HelpDesk string `json:"help_desk,omitempty"` // 应用开发商服务台链接(仅商店应用返回)
	Email    string `json:"email,omitempty"`     // 应用开发商的邮箱(仅商店应用返回)
	Phone    string `json:"phone,omitempty"`     // 应用开发商的手机号(仅商店应用返回)
}

// GetApplicationUnderAuditListRespItemScope ...
type GetApplicationUnderAuditListRespItemScope struct {
	Scope       string `json:"scope,omitempty"`       // 应用权限
	Description string `json:"description,omitempty"` // 应用权限的国际化描述
	Level       int64  `json:"level,omitempty"`       // 权限等级描述, 可选值有: 1: 普通权限, 2: 高级权限, 3: 超敏感权限, 0: 未知等级
}

// getApplicationUnderAuditListResp ...
type getApplicationUnderAuditListResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *GetApplicationUnderAuditListResp `json:"data,omitempty"`
}
