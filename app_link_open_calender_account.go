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

// OpenCalenderAccount 打开第三方日历账户管理页
//
// 打开第三方日历账户管理页，方便用户添加或导入谷歌/Exchange 日历。移动端打开页面，PC端打开弹层。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-calender/open
// new doc: https://open.feishu.cn/document/common-capabilities/applink-protocol/supported-protocol/open-calender/open
func (r *AppLinkService) OpenCalenderAccount(req *OpenCalenderAccountReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/calendar/account", req)
}

// OpenCalenderAccountReq ...
type OpenCalenderAccountReq struct {
}
