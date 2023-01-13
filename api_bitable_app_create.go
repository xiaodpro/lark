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

// CreateBitableApp 在指定目录下创建多维表格
//
// 该接口支持调用频率上限为 20 QPM（Query Per Minute, 每分钟请求率）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/create
func (r *BitableService) CreateBitableApp(ctx context.Context, request *CreateBitableAppReq, options ...MethodOptionFunc) (*CreateBitableAppResp, *Response, error) {
	if r.cli.mock.mockBitableCreateBitableApp != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#CreateBitableApp mock enable")
		return r.cli.mock.mockBitableCreateBitableApp(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "CreateBitableApp",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createBitableAppResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableCreateBitableApp mock BitableCreateBitableApp method
func (r *Mock) MockBitableCreateBitableApp(f func(ctx context.Context, request *CreateBitableAppReq, options ...MethodOptionFunc) (*CreateBitableAppResp, *Response, error)) {
	r.mockBitableCreateBitableApp = f
}

// UnMockBitableCreateBitableApp un-mock BitableCreateBitableApp method
func (r *Mock) UnMockBitableCreateBitableApp() {
	r.mockBitableCreateBitableApp = nil
}

// CreateBitableAppReq ...
type CreateBitableAppReq struct {
	Name        *string `json:"name,omitempty"`         // 多维表格App名字, 示例值: "一篇新的多维表格"
	FolderToken *string `json:"folder_token,omitempty"` // 多维表格App归属文件夹, 示例值: "fldbco*CIMltVc"
}

// CreateBitableAppResp ...
type CreateBitableAppResp struct {
	App *CreateBitableAppRespApp `json:"app,omitempty"` // 返回响应体
}

// CreateBitableAppRespApp ...
type CreateBitableAppRespApp struct {
	AppToken    string `json:"app_token,omitempty"`    // 多维表格的 app_token
	Name        string `json:"name,omitempty"`         // 多维表格的名字
	FolderToken string `json:"folder_token,omitempty"` // 多维表格 App 归属文件夹
	URL         string `json:"url,omitempty"`          // 多维表格 App URL
}

// createBitableAppResp ...
type createBitableAppResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *CreateBitableAppResp `json:"data,omitempty"`
}
