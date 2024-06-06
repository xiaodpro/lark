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
	"io"
)

// RecognizeAIVehicleInvoice 机动车发票识别接口, 支持JPG/JPEG/PNG/BMP四种文件类型的一次性的识别。
//
// 单租户限流: 10QPS, 同租户下的应用没有限流, 共享本租户的 10QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/document_ai-v1/vehicle_invoice/recognize
func (r *AIService) RecognizeAIVehicleInvoice(ctx context.Context, request *RecognizeAIVehicleInvoiceReq, options ...MethodOptionFunc) (*RecognizeAIVehicleInvoiceResp, *Response, error) {
	if r.cli.mock.mockAIRecognizeAIVehicleInvoice != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] AI#RecognizeAIVehicleInvoice mock enable")
		return r.cli.mock.mockAIRecognizeAIVehicleInvoice(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "RecognizeAIVehicleInvoice",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/document_ai/v1/vehicle_invoice/recognize",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
		IsFile:                true,
	}
	resp := new(recognizeAIVehicleInvoiceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAIRecognizeAIVehicleInvoice mock AIRecognizeAIVehicleInvoice method
func (r *Mock) MockAIRecognizeAIVehicleInvoice(f func(ctx context.Context, request *RecognizeAIVehicleInvoiceReq, options ...MethodOptionFunc) (*RecognizeAIVehicleInvoiceResp, *Response, error)) {
	r.mockAIRecognizeAIVehicleInvoice = f
}

// UnMockAIRecognizeAIVehicleInvoice un-mock AIRecognizeAIVehicleInvoice method
func (r *Mock) UnMockAIRecognizeAIVehicleInvoice() {
	r.mockAIRecognizeAIVehicleInvoice = nil
}

// RecognizeAIVehicleInvoiceReq ...
type RecognizeAIVehicleInvoiceReq struct {
	File io.Reader `json:"file,omitempty"` // 识别的机动车发票源文件, 示例值: file binary
}

// RecognizeAIVehicleInvoiceResp ...
type RecognizeAIVehicleInvoiceResp struct {
	VehicleInvoice *RecognizeAIVehicleInvoiceRespVehicleInvoice `json:"vehicle_invoice,omitempty"` // 机动车发票信息
}

// RecognizeAIVehicleInvoiceRespVehicleInvoice ...
type RecognizeAIVehicleInvoiceRespVehicleInvoice struct {
	Entities []*RecognizeAIVehicleInvoiceRespVehicleInvoiceEntitie `json:"entities,omitempty"` // 识别出的实体类型
}

// RecognizeAIVehicleInvoiceRespVehicleInvoiceEntitie ...
type RecognizeAIVehicleInvoiceRespVehicleInvoiceEntitie struct {
	Type  string `json:"type,omitempty"`  // 识别的字段种类, 可选值有: invoice_code: 发票代码, invoice_num: 发票号码, date: 开票日期, print_code: 机打代码, print_num: 机打号码, machine_num: 机器编码, buyer_name: 购买方名称, buyer_id: 购买方纳税人识别号, vehicle_type: 车辆类型, product_model: 厂牌型号, certificate_num: 合格证号, engine_num: 发动机号码, vin: 车架号, total_price: 价税合计, total_price_little: 小写金额, saler_name: 销货单位名称, saler_id: 销售方纳税人识别号, saler_addr: 地址, tax_rate: 税率, tax: 税额, price: 不含税价格
	Value string `json:"value,omitempty"` // 识别出字段的文本信息
}

// recognizeAIVehicleInvoiceResp ...
type recognizeAIVehicleInvoiceResp struct {
	Code  int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                         `json:"msg,omitempty"`  // 错误描述
	Data  *RecognizeAIVehicleInvoiceResp `json:"data,omitempty"`
	Error *ErrorDetail                   `json:"error,omitempty"`
}
