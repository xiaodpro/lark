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

// GetApplicationUserVisibleApp
//
// 该接口用于查询用户可用的应用列表, 只能被企业自建应用调用。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMjM3UjLzIzN14yMycTN
func (r *ApplicationService) GetApplicationUserVisibleApp(ctx context.Context, request *GetApplicationUserVisibleAppReq, options ...MethodOptionFunc) (*GetApplicationUserVisibleAppResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationUserVisibleApp != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Application#GetApplicationUserVisibleApp mock enable")
		return r.cli.mock.mockApplicationGetApplicationUserVisibleApp(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationUserVisibleApp",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v1/user/visible_apps",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationUserVisibleAppResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationGetApplicationUserVisibleApp mock ApplicationGetApplicationUserVisibleApp method
func (r *Mock) MockApplicationGetApplicationUserVisibleApp(f func(ctx context.Context, request *GetApplicationUserVisibleAppReq, options ...MethodOptionFunc) (*GetApplicationUserVisibleAppResp, *Response, error)) {
	r.mockApplicationGetApplicationUserVisibleApp = f
}

// UnMockApplicationGetApplicationUserVisibleApp un-mock ApplicationGetApplicationUserVisibleApp method
func (r *Mock) UnMockApplicationGetApplicationUserVisibleApp() {
	r.mockApplicationGetApplicationUserVisibleApp = nil
}

// GetApplicationUserVisibleAppReq ...
type GetApplicationUserVisibleAppReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页起始位置标示, 不填表示从头开始
	PageSize  *int64  `query:"page_size" json:"-"`  // 单页需求最大个数（最大 100）, 0 自动最大个数
	Lang      *string `query:"lang" json:"-"`       // 优先展示的应用信息的语言版本（zh_cn: 中文, en_us: 英文, ja_jp: 日文）
	OpenID    *string `query:"open_id" json:"-"`    // 目标用户 open_id
	UserID    *string `query:"user_id" json:"-"`    // 目标用户 user_id, 与 open_id 至少给其中之一, user_id 优先于 open_id
}

// GetApplicationUserVisibleAppResp ...
type GetApplicationUserVisibleAppResp struct {
	PageToken  string                                   `json:"page_token,omitempty"`  // 下一个请求页应当给的起始位置
	PageSize   int64                                    `json:"page_size,omitempty"`   // 本次请求实际返回的页大小
	TotalCount int64                                    `json:"total_count,omitempty"` // 可用的应用总数
	HasMore    int64                                    `json:"has_more,omitempty"`    // 是否还有更多应用
	Lang       string                                   `json:"lang,omitempty"`        // 当前选择的版本语言
	AppList    *GetApplicationUserVisibleAppRespAppList `json:"app_list,omitempty"`    // 应用列表
}

// GetApplicationUserVisibleAppRespAppList ...
type GetApplicationUserVisibleAppRespAppList struct {
	AppID                string `json:"app_id,omitempty"`                 // 应用 ID
	PrimaryLanguage      string `json:"primary_language,omitempty"`       // 应用首选语言
	AppName              string `json:"app_name,omitempty"`               // 应用名称
	Description          string `json:"description,omitempty"`            // 应用描述
	AvatarURL            string `json:"avatar_url,omitempty"`             // 应用 icon
	AppSceneType         int64  `json:"app_scene_type,omitempty"`         // 应用类型, 0: 企业自建应用；1: 应用商店应用
	Status               int64  `json:"status,omitempty"`                 // 启停状态, 0: 停用；1: 启用
	MobileDefaultAbility int64  `json:"mobile_default_ability,omitempty"` // 移动端默认的应用功能, 0: 未开启；1: 小程序；2: H5；8: 机器人
	PcDefaultAbility     int64  `json:"pc_default_ability,omitempty"`     // PC客户端默认的应用功能, 0: 未开启；1: 小程序；2: H5；8: 机器人
}

// getApplicationUserVisibleAppResp ...
type getApplicationUserVisibleAppResp struct {
	Code int64                             `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 返回码的描述
	Data *GetApplicationUserVisibleAppResp `json:"data,omitempty"` // 返回的业务信息, 仅 code = 0 时有效
}
