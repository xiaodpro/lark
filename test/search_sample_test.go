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

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Search_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Search

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateSearchDataSource(ctx, &lark.CreateSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Search

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchSearchMessage(func(ctx context.Context, request *lark.SearchMessageReq, options ...lark.MethodOptionFunc) (*lark.SearchMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchSearchMessage()

			_, _, err := moduleCli.SearchMessage(ctx, &lark.SearchMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchSearchApp(func(ctx context.Context, request *lark.SearchAppReq, options ...lark.MethodOptionFunc) (*lark.SearchAppResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchSearchApp()

			_, _, err := moduleCli.SearchApp(ctx, &lark.SearchAppReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchCreateSearchDataSource(func(ctx context.Context, request *lark.CreateSearchDataSourceReq, options ...lark.MethodOptionFunc) (*lark.CreateSearchDataSourceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchCreateSearchDataSource()

			_, _, err := moduleCli.CreateSearchDataSource(ctx, &lark.CreateSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchGetSearchDataSource(func(ctx context.Context, request *lark.GetSearchDataSourceReq, options ...lark.MethodOptionFunc) (*lark.GetSearchDataSourceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchGetSearchDataSource()

			_, _, err := moduleCli.GetSearchDataSource(ctx, &lark.GetSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchUpdateSearchDataSource(func(ctx context.Context, request *lark.UpdateSearchDataSourceReq, options ...lark.MethodOptionFunc) (*lark.UpdateSearchDataSourceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchUpdateSearchDataSource()

			_, _, err := moduleCli.UpdateSearchDataSource(ctx, &lark.UpdateSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchGetSearchDataSourceList(func(ctx context.Context, request *lark.GetSearchDataSourceListReq, options ...lark.MethodOptionFunc) (*lark.GetSearchDataSourceListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchGetSearchDataSourceList()

			_, _, err := moduleCli.GetSearchDataSourceList(ctx, &lark.GetSearchDataSourceListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchDeleteSearchDataSource(func(ctx context.Context, request *lark.DeleteSearchDataSourceReq, options ...lark.MethodOptionFunc) (*lark.DeleteSearchDataSourceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchDeleteSearchDataSource()

			_, _, err := moduleCli.DeleteSearchDataSource(ctx, &lark.DeleteSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchBatchCreateSearchDataSourceItem(func(ctx context.Context, request *lark.BatchCreateSearchDataSourceItemReq, options ...lark.MethodOptionFunc) (*lark.BatchCreateSearchDataSourceItemResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchBatchCreateSearchDataSourceItem()

			_, _, err := moduleCli.BatchCreateSearchDataSourceItem(ctx, &lark.BatchCreateSearchDataSourceItemReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchCreateSearchDataSourceItem(func(ctx context.Context, request *lark.CreateSearchDataSourceItemReq, options ...lark.MethodOptionFunc) (*lark.CreateSearchDataSourceItemResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchCreateSearchDataSourceItem()

			_, _, err := moduleCli.CreateSearchDataSourceItem(ctx, &lark.CreateSearchDataSourceItemReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchGetSearchDataSourceItem(func(ctx context.Context, request *lark.GetSearchDataSourceItemReq, options ...lark.MethodOptionFunc) (*lark.GetSearchDataSourceItemResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchGetSearchDataSourceItem()

			_, _, err := moduleCli.GetSearchDataSourceItem(ctx, &lark.GetSearchDataSourceItemReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchDeleteSearchDataSourceItem(func(ctx context.Context, request *lark.DeleteSearchDataSourceItemReq, options ...lark.MethodOptionFunc) (*lark.DeleteSearchDataSourceItemResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchDeleteSearchDataSourceItem()

			_, _, err := moduleCli.DeleteSearchDataSourceItem(ctx, &lark.DeleteSearchDataSourceItemReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchUpdateSearchSchema(func(ctx context.Context, request *lark.UpdateSearchSchemaReq, options ...lark.MethodOptionFunc) (*lark.UpdateSearchSchemaResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchUpdateSearchSchema()

			_, _, err := moduleCli.UpdateSearchSchema(ctx, &lark.UpdateSearchSchemaReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchDeleteSearchSchema(func(ctx context.Context, request *lark.DeleteSearchSchemaReq, options ...lark.MethodOptionFunc) (*lark.DeleteSearchSchemaResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchDeleteSearchSchema()

			_, _, err := moduleCli.DeleteSearchSchema(ctx, &lark.DeleteSearchSchemaReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchGetSearchSchema(func(ctx context.Context, request *lark.GetSearchSchemaReq, options ...lark.MethodOptionFunc) (*lark.GetSearchSchemaResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchGetSearchSchema()

			_, _, err := moduleCli.GetSearchSchema(ctx, &lark.GetSearchSchemaReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockSearchCreateSearchSchema(func(ctx context.Context, request *lark.CreateSearchSchemaReq, options ...lark.MethodOptionFunc) (*lark.CreateSearchSchemaResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchCreateSearchSchema()

			_, _, err := moduleCli.CreateSearchSchema(ctx, &lark.CreateSearchSchemaReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Search

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchMessage(ctx, &lark.SearchMessageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchApp(ctx, &lark.SearchAppReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateSearchDataSource(ctx, &lark.CreateSearchDataSourceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetSearchDataSource(ctx, &lark.GetSearchDataSourceReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateSearchDataSource(ctx, &lark.UpdateSearchDataSourceReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetSearchDataSourceList(ctx, &lark.GetSearchDataSourceListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteSearchDataSource(ctx, &lark.DeleteSearchDataSourceReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchCreateSearchDataSourceItem(ctx, &lark.BatchCreateSearchDataSourceItemReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateSearchDataSourceItem(ctx, &lark.CreateSearchDataSourceItemReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetSearchDataSourceItem(ctx, &lark.GetSearchDataSourceItemReq{
				DataSourceID: "x",
				ItemID:       "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteSearchDataSourceItem(ctx, &lark.DeleteSearchDataSourceItemReq{
				DataSourceID: "x",
				ItemID:       "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateSearchSchema(ctx, &lark.UpdateSearchSchemaReq{
				SchemaID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteSearchSchema(ctx, &lark.DeleteSearchSchemaReq{
				SchemaID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetSearchSchema(ctx, &lark.GetSearchSchemaReq{
				SchemaID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateSearchSchema(ctx, &lark.CreateSearchSchemaReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Search
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchMessage(ctx, &lark.SearchMessageReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchApp(ctx, &lark.SearchAppReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateSearchDataSource(ctx, &lark.CreateSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetSearchDataSource(ctx, &lark.GetSearchDataSourceReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateSearchDataSource(ctx, &lark.UpdateSearchDataSourceReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetSearchDataSourceList(ctx, &lark.GetSearchDataSourceListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteSearchDataSource(ctx, &lark.DeleteSearchDataSourceReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchCreateSearchDataSourceItem(ctx, &lark.BatchCreateSearchDataSourceItemReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateSearchDataSourceItem(ctx, &lark.CreateSearchDataSourceItemReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetSearchDataSourceItem(ctx, &lark.GetSearchDataSourceItemReq{
				DataSourceID: "x",
				ItemID:       "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteSearchDataSourceItem(ctx, &lark.DeleteSearchDataSourceItemReq{
				DataSourceID: "x",
				ItemID:       "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateSearchSchema(ctx, &lark.UpdateSearchSchemaReq{
				SchemaID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteSearchSchema(ctx, &lark.DeleteSearchSchemaReq{
				SchemaID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetSearchSchema(ctx, &lark.GetSearchSchemaReq{
				SchemaID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateSearchSchema(ctx, &lark.CreateSearchSchemaReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
