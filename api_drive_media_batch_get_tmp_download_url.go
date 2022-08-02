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

// BatchGetDriveMediaTmpDownloadURL 通过file_token获取素材临时下载链接, 链接时效性是24小时, 过期失效。
//
// 该接口不支持太高的并发, 且调用频率上限为5QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/batch_get_tmp_download_url
func (r *DriveService) BatchGetDriveMediaTmpDownloadURL(ctx context.Context, request *BatchGetDriveMediaTmpDownloadURLReq, options ...MethodOptionFunc) (*BatchGetDriveMediaTmpDownloadURLResp, *Response, error) {
	if r.cli.mock.mockDriveBatchGetDriveMediaTmpDownloadURL != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#BatchGetDriveMediaTmpDownloadURL mock enable")
		return r.cli.mock.mockDriveBatchGetDriveMediaTmpDownloadURL(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "BatchGetDriveMediaTmpDownloadURL",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/medias/batch_get_tmp_download_url",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchGetDriveMediaTmpDownloadURLResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveBatchGetDriveMediaTmpDownloadURL mock DriveBatchGetDriveMediaTmpDownloadURL method
func (r *Mock) MockDriveBatchGetDriveMediaTmpDownloadURL(f func(ctx context.Context, request *BatchGetDriveMediaTmpDownloadURLReq, options ...MethodOptionFunc) (*BatchGetDriveMediaTmpDownloadURLResp, *Response, error)) {
	r.mockDriveBatchGetDriveMediaTmpDownloadURL = f
}

// UnMockDriveBatchGetDriveMediaTmpDownloadURL un-mock DriveBatchGetDriveMediaTmpDownloadURL method
func (r *Mock) UnMockDriveBatchGetDriveMediaTmpDownloadURL() {
	r.mockDriveBatchGetDriveMediaTmpDownloadURL = nil
}

// BatchGetDriveMediaTmpDownloadURLReq ...
type BatchGetDriveMediaTmpDownloadURLReq struct {
	FileTokens []string `query:"file_tokens" json:"-"` // 文件标识符列表, 示例值: boxcnrHpsg1QDqXAAAyachabcef
	Extra      *string  `query:"extra" json:"-"`       // 拓展信息(可选), 示例值: "[请参考-上传点类型及对应Extra说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/introduction)"
}

// BatchGetDriveMediaTmpDownloadURLResp ...
type BatchGetDriveMediaTmpDownloadURLResp struct {
	TmpDownloadURLs []*BatchGetDriveMediaTmpDownloadURLRespTmpDownloadURL `json:"tmp_download_urls,omitempty"` // 临时下载列表
}

// BatchGetDriveMediaTmpDownloadURLRespTmpDownloadURL ...
type BatchGetDriveMediaTmpDownloadURLRespTmpDownloadURL struct {
	FileToken      string `json:"file_token,omitempty"`       // 文件标识符
	TmpDownloadURL string `json:"tmp_download_url,omitempty"` // 文件临时下载链接
}

// batchGetDriveMediaTmpDownloadURLResp ...
type batchGetDriveMediaTmpDownloadURLResp struct {
	Code int64                                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                `json:"msg,omitempty"`  // 错误描述
	Data *BatchGetDriveMediaTmpDownloadURLResp `json:"data,omitempty"`
}