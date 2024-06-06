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
	"github.com/chyroc/lark/internal"
)

// OpenDocs 打开云文档
//
// 打开云文档（docs）。使用外部浏览器打开文档时，提供入口从飞书中打开。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-docs
// new doc: https://open.feishu.cn/document/common-capabilities/applink-protocol/supported-protocol/open-docs
func (r *AppLinkService) OpenDocs(req *OpenDocsReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/docs/open", req)
}

// OpenDocsReq ...
type OpenDocsReq struct {
	URL string `json:"url,omitempty"` // 要打开的云文档URL
}
