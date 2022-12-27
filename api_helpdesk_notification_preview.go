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

// PreviewHelpdeskNotification 在正式执行推送之前是可以调用此接口预览设置的推送内容。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/preview
func (r *HelpdeskService) PreviewHelpdeskNotification(ctx context.Context, request *PreviewHelpdeskNotificationReq, options ...MethodOptionFunc) (*PreviewHelpdeskNotificationResp, *Response, error) {
	if r.cli.mock.mockHelpdeskPreviewHelpdeskNotification != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#PreviewHelpdeskNotification mock enable")
		return r.cli.mock.mockHelpdeskPreviewHelpdeskNotification(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "PreviewHelpdeskNotification",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/notifications/:notification_id/preview",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(previewHelpdeskNotificationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskPreviewHelpdeskNotification mock HelpdeskPreviewHelpdeskNotification method
func (r *Mock) MockHelpdeskPreviewHelpdeskNotification(f func(ctx context.Context, request *PreviewHelpdeskNotificationReq, options ...MethodOptionFunc) (*PreviewHelpdeskNotificationResp, *Response, error)) {
	r.mockHelpdeskPreviewHelpdeskNotification = f
}

// UnMockHelpdeskPreviewHelpdeskNotification un-mock HelpdeskPreviewHelpdeskNotification method
func (r *Mock) UnMockHelpdeskPreviewHelpdeskNotification() {
	r.mockHelpdeskPreviewHelpdeskNotification = nil
}

// PreviewHelpdeskNotificationReq ...
type PreviewHelpdeskNotificationReq struct {
	NotificationID string `path:"notification_id" json:"-"` // 创建推送接口成功后返回的唯一id, 示例值: "6985032626234982420"
}

// PreviewHelpdeskNotificationResp ...
type PreviewHelpdeskNotificationResp struct {
}

// previewHelpdeskNotificationResp ...
type previewHelpdeskNotificationResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *PreviewHelpdeskNotificationResp `json:"data,omitempty"`
}
