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

// CreateSearchDataSourceItem 索引一条数据记录。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source-item/create
func (r *SearchService) CreateSearchDataSourceItem(ctx context.Context, request *CreateSearchDataSourceItemReq, options ...MethodOptionFunc) (*CreateSearchDataSourceItemResp, *Response, error) {
	if r.cli.mock.mockSearchCreateSearchDataSourceItem != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#CreateSearchDataSourceItem mock enable")
		return r.cli.mock.mockSearchCreateSearchDataSourceItem(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "CreateSearchDataSourceItem",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/search/v2/data_sources/:data_source_id/items",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createSearchDataSourceItemResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockSearchCreateSearchDataSourceItem mock SearchCreateSearchDataSourceItem method
func (r *Mock) MockSearchCreateSearchDataSourceItem(f func(ctx context.Context, request *CreateSearchDataSourceItemReq, options ...MethodOptionFunc) (*CreateSearchDataSourceItemResp, *Response, error)) {
	r.mockSearchCreateSearchDataSourceItem = f
}

// UnMockSearchCreateSearchDataSourceItem un-mock SearchCreateSearchDataSourceItem method
func (r *Mock) UnMockSearchCreateSearchDataSourceItem() {
	r.mockSearchCreateSearchDataSourceItem = nil
}

// CreateSearchDataSourceItemReq ...
type CreateSearchDataSourceItemReq struct {
	DataSourceID   string                                 `path:"data_source_id" json:"-"`   // 数据源的ID, 示例值: "service_ticket"
	ID             string                                 `json:"id,omitempty"`              // item 在 datasource 中的唯一标识, 示例值: "01010111", 最大长度: `128` 字符
	ACL            []*CreateSearchDataSourceItemReqACL    `json:"acl,omitempty"`             // item 的访问权限控制。 acl 字段为空数组, 则默认数据不可见。如果数据是全员可见, 需要设置 access="allow"; type="user"; value="everyone"
	Metadata       *CreateSearchDataSourceItemReqMetadata `json:"metadata,omitempty"`        // item 的元信息
	StructuredData string                                 `json:"structured_data,omitempty"` // 结构化数据（以 json 字符串传递）, 这些字段是搜索结果的展示字段(特殊字段无须在此另外指定);具体格式可参参考 [通用模块接入指南](/uAjLw4CM/ukTMukTMukTM/search-v2/common-template-intergration-handbook) 请求创建数据项部分, 示例值: "{\"key\":\"value\"}"
	Content        *CreateSearchDataSourceItemReqContent  `json:"content,omitempty"`         // 非结构化数据, 如文档文本, 飞书搜索会用来做召回
}

// CreateSearchDataSourceItemReqACL ...
type CreateSearchDataSourceItemReqACL struct {
	Access *string `json:"access,omitempty"` // 权限类型, 优先级: Deny > Allow, 示例值: "allow", 可选值有: allow: 允许访问, deny: 禁止访问
	Value  *string `json:"value,omitempty"`  // 设置的权限值, 例如 userID, 依赖 type 描述, 注: 在 type 为 user 且 access 为 allow 时, 可填 "everyone" 来表示该数据项对全员可见；, 示例值: "d35e3c23"
	Type   *string `json:"type,omitempty"`   // 权限值类型, 示例值: "user", 可选值有: user: 访问权限控制中指定“用户”可以访问或拒绝访问该条数据, group: (已下线)访问权限控制中指定“用户组”可以访问或拒绝访问该条数据, open_id: 用户的open_id
}

// CreateSearchDataSourceItemReqContent ...
type CreateSearchDataSourceItemReqContent struct {
	Format      *string `json:"format,omitempty"`       // 内容的格式, 示例值: "html", 可选值有: html: html格式, plaintext: 纯文本格式
	ContentData *string `json:"content_data,omitempty"` // 全文数据, 示例值: "这是一个很长的文本"
}

// CreateSearchDataSourceItemReqMetadata ...
type CreateSearchDataSourceItemReqMetadata struct {
	Title           string  `json:"title,omitempty"`             // 该条数据记录对应的标题, 示例值: "工单: 无法创建文章"
	SourceURL       string  `json:"source_url,omitempty"`        // 该条数据记录对应的跳转url, 示例值: "http://www.abc.com.cn"
	CreateTime      *int64  `json:"create_time,omitempty"`       // 数据项的创建时间。Unix 时间, 单位为秒, 示例值: 1618831236
	UpdateTime      *int64  `json:"update_time,omitempty"`       // 数据项的更新时间。Unix 时间, 单位为秒, 示例值: 1618831236
	SourceURLMobile *string `json:"source_url_mobile,omitempty"` // 移动端搜索命中的跳转地址。如果您PC端和移动端有不同的跳转地址, 可以在这里写入移动端专用的url, 我们会在搜索时为您选择合适的地址, 示例值: "https://www.feishu.cn"
}

// CreateSearchDataSourceItemResp ...
type CreateSearchDataSourceItemResp struct {
}

// createSearchDataSourceItemResp ...
type createSearchDataSourceItemResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *CreateSearchDataSourceItemResp `json:"data,omitempty"`
}
