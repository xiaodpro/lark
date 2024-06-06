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

// ListIMTagRelation 查询实体与标签的绑定关系
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/group/im-v2/biz_entity_tag_relation/get
func (r *MessageService) ListIMTagRelation(ctx context.Context, request *ListIMTagRelationReq, options ...MethodOptionFunc) (*ListIMTagRelationResp, *Response, error) {
	if r.cli.mock.mockMessageListIMTagRelation != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Message#ListIMTagRelation mock enable")
		return r.cli.mock.mockMessageListIMTagRelation(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "ListIMTagRelation",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v2/biz_entity_tag_relation",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(listIMTagRelationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageListIMTagRelation mock MessageListIMTagRelation method
func (r *Mock) MockMessageListIMTagRelation(f func(ctx context.Context, request *ListIMTagRelationReq, options ...MethodOptionFunc) (*ListIMTagRelationResp, *Response, error)) {
	r.mockMessageListIMTagRelation = f
}

// UnMockMessageListIMTagRelation un-mock MessageListIMTagRelation method
func (r *Mock) UnMockMessageListIMTagRelation() {
	r.mockMessageListIMTagRelation = nil
}

// ListIMTagRelationReq ...
type ListIMTagRelationReq struct {
	TagBizType  string `query:"tag_biz_type" json:"-"`  // 业务类型, 示例值: chat, 可选值有: chat: chat类型
	BizEntityID string `query:"biz_entity_id" json:"-"` // 业务实体id, 示例值: 71616xxxx
}

// ListIMTagRelationResp ...
type ListIMTagRelationResp struct {
	TagInfoWithBindVersions []*ListIMTagRelationRespTagInfoWithBindVersion `json:"tag_info_with_bind_versions,omitempty"` // 标签内容及绑定时间
}

// ListIMTagRelationRespTagInfoWithBindVersion ...
type ListIMTagRelationRespTagInfoWithBindVersion struct {
	TagInfo     *ListIMTagRelationRespTagInfoWithBindVersionTagInfo `json:"tag_info,omitempty"`     // 标签内容
	BindVersion string                                              `json:"bind_version,omitempty"` // 绑定时间
}

// ListIMTagRelationRespTagInfoWithBindVersionTagInfo ...
type ListIMTagRelationRespTagInfoWithBindVersionTagInfo struct {
	ID         string                                                        `json:"id,omitempty"`          // id
	TagType    string                                                        `json:"tag_type,omitempty"`    // 标签类型
	Name       string                                                        `json:"name,omitempty"`        // name
	I18nNames  []*ListIMTagRelationRespTagInfoWithBindVersionTagInfoI18nName `json:"i18n_names,omitempty"`  // i18n name
	CreateTime string                                                        `json:"create_time,omitempty"` // 创建时间
	UpdateTime string                                                        `json:"update_time,omitempty"` // 更新时间
}

// ListIMTagRelationRespTagInfoWithBindVersionTagInfoI18nName ...
type ListIMTagRelationRespTagInfoWithBindVersionTagInfoI18nName struct {
	Locale string `json:"locale,omitempty"` // 语言
	Name   string `json:"name,omitempty"`   // 名称
}

// listIMTagRelationResp ...
type listIMTagRelationResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *ListIMTagRelationResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}
