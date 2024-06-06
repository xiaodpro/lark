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
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_GetApproval(t *testing.T) {
	as := assert.New(t)

	cli := AppAllPermission.Ins()

	resp, _, err := cli.Approval.GetApproval(ctx, &lark.GetApprovalReq{
		ApprovalCode: ApprovalALLField.Code,
		Locale:       nil,
	})
	as.Nil(err)

	as.Equal("sdk-demo", resp.ApprovalName)
	widgets := []*lark.ApprovalWidget{
		{Type: lark.ApprovalWidgetTypeInput, Name: "单行文本"},
		{Type: lark.ApprovalWidgetTypeTextarea, Name: "多行文本"},
		{Type: lark.ApprovalWidgetTypeText, Name: "说明 1"},
		{Type: lark.ApprovalWidgetTypeNumber, Name: "数字"},
		{Type: lark.ApprovalWidgetTypeAmount, Name: "金额"},
		{Type: lark.ApprovalWidgetTypeFormula, Name: "计算公式"},
		{Type: lark.ApprovalWidgetTypeRadioV2, Name: "单选", Option: &lark.ApprovalWidgetOptions{IsList: true, Options: []*lark.ApprovalWidgetOption{{Text: "1"}, {Text: "2"}}}},
		{Type: lark.ApprovalWidgetTypeCheckboxV2, Name: "多选", Option: &lark.ApprovalWidgetOptions{IsList: true, Options: []*lark.ApprovalWidgetOption{{Text: "1"}, {Text: "2"}, {Text: "3"}}}},
		{Type: lark.ApprovalWidgetTypeDate, Name: "日期"},
		{Type: lark.ApprovalWidgetTypeDateInterval, Name: "DateInterval"},
		{Type: lark.ApprovalWidgetTypeFieldList, Name: "明细", Children: []*lark.ApprovalWidget{{Name: "单行文本"}, {Name: "多行文本"}}},
		{Type: lark.ApprovalWidgetTypeImage, Name: "图片"},
		{Type: lark.ApprovalWidgetTypeAttachmentV2, Name: "附件"},
		{Type: lark.ApprovalWidgetTypeDepartment, Name: "部门"},
		{Type: lark.ApprovalWidgetTypeContact, Name: "联系人"},
		{Type: lark.ApprovalWidgetTypeConnect, Name: "关联审批"},
		{Type: lark.ApprovalWidgetTypeAddress, Name: "地址"},
	}
	as.Equal(len(widgets), len(resp.Form))
	for idx, widget := range resp.Form {
		expectWidget := widgets[idx]
		as.Equal(expectWidget.Name, widget.Name)
		as.Equal(expectWidget.Type, widget.Type)
		as.Equal(len(expectWidget.Children), len(widget.Children))
		if expectWidget.Option != nil {
			as.Equal(len(expectWidget.Option.Options), len(widget.Option.Options))
			for optionIdx, option := range widget.Option.Options {
				expectOption := expectWidget.Option.Options[optionIdx]
				as.Equal(expectOption.Text, option.Text)
			}
			for childrenIdx, children := range widget.Children {
				expectChildren := expectWidget.Children[childrenIdx]
				as.Equal(expectChildren.Name, children.Name)
			}
		}
	}
}

func Test_UnmarshalGetApprovalInstance(t *testing.T) {
	as := assert.New(t)

	s := `{
    "approval_code": "96F7DAD7-5F10-416B-A50C-C3331A2B3010",
    "approval_name": "sdk-demo",
    "comment_list": [],
    "department_id": "",
    "end_time": "0",
    "form": "[{\"id\":\"widget3\",\"name\":\"单行文本\",\"type\":\"input\",\"ext\":null,\"value\":\"test\"},{\"id\":\"widget16215623422760001\",\"name\":\"说明 1\",\"type\":\"text\",\"ext\":null,\"value\":\"说明\"}]",
    "open_id": "ou_6e2bbe38e6551b5b2731a409ef50e8ce",
    "serial_number": "202105230010",
    "start_time": "1621742291784",
    "status": "PENDING",
    "task_list": [
      {
        "end_time": "0",
        "id": "6965330108705980419",
        "node_id": "6db614baa6d5cb208decf9efa2e3eee3",
        "node_name": "直接主管",
        "open_id": "ou_6e2bbe38e6551b5b2731a409ef50e8ce",
        "start_time": "1621742292164",
        "status": "PENDING",
        "type": "AND",
        "user_id": "3gg4cf3g"
      }
    ],
    "timeline": [
      {
        "create_time": "1621742291784",
        "ext": "{\"user_id\":\"1\"}",
        "open_id": "ou_6e2bbe38e6551b5b2731a409ef50e8ce",
        "type": "START",
        "user_id": "3gg4cf3g"
      }
    ],
    "user_id": "3gg4cf3g",
    "uuid": ""
  }
`
	res := new(lark.GetApprovalInstanceResp)
	err := json.Unmarshal([]byte(s), res)
	as.Nil(err)
	as.NotNil(res)
	as.Equal("1", ptrValueString(res.Timeline[0].Ext.UserID))
	printData(res)
}

func Test_UnmarshalCreateApproval(t *testing.T) {
	as := assert.New(t)

	s := `{
    "approval_code":"4202AD96-9EC1-4284-9C48-B923CDC4F30B",
    "user_id":"59a92c4a",
    "open_id":"ou_806a18fb5bdf525e38ba219733bdbd73",
    "form":"[{\"id\":\"111\",\"type\":\"input\",\"value\":\"11111\"},{\"id\":\"222\",\"required\":true,\"type\":\"dateInterval\",\"value\":{\"start\":\"2019-10-01T08:12:01+08:00\",\"end\":\"2019-10-02T08:12:01+08:00\",\"interval\": 2.0}},{\"id\":\"333\",\"type\":\"radioV2\",\"value\":\"k2b8mkx0-h71x5gljn3i-1\"},{\"id\":\"444\",\"type\":\"number\", \"value\":\"4\"},{\"id\":\"555\",\"type\":\"textarea\",\"value\":\"fsafs\"}]",
    "node_approver_user_id_list":[
        {"key": "46e6d96cfa756980907209209ec03b64","value":["59a92c4a"]},
        {"key": "manager_node_id","value":["59a92c4a"]}
    ],
    "node_approver_open_id_list":[
        {"key": "46e6d96cfa756980907209209ec03b64","value":["ou_806a18fb5bdf525e38ba219733bdbd73","ou_806a18f"]},
        {"key": "manager_node_id","value":["ou_806a18fb5bdf525e38ba219733bdbd73"]}
    ],
    "node_cc_user_id_list":[
        {"key": "46e6d96cfa756980907209209ec03b64","value":["59a92c4a"]},
        {"key": "manager_node_id","value":["59a92c4a"]}
    ],
    "node_cc_open_id_list":[
        {"key": "46e6d96cfa756980907209209ec03b64","value":["ou_806a18fb5bdf525e38ba219733bdbd73"]},
        {"key": "manager_node_id","value":["ou_806a18fb5bdf525e38ba219733bdbd73"]}
    ]
}
`
	res := new(lark.CreateApprovalInstanceReq)
	err := json.Unmarshal([]byte(s), res)
	as.Nil(err)
	as.NotNil(res)
	as.Equal("46e6d96cfa756980907209209ec03b64", *res.NodeApproverOpenIDList[0].Key)
	as.Equal("ou_806a18f", res.NodeApproverOpenIDList[0].Value[1])
	printData(res)
}

func Test_Create_CancelApproval(t *testing.T) {
	as := assert.New(t)

	cli := AppAllPermission.Ins()

	t.Run("cancel", func(t *testing.T) {
		instanceCode, _ := testCreateApproval(t, cli, ApprovalALLField.Code, UserAdmin.UserID, UserAdmin.OpenID)
		_, _, err := cli.Approval.CancelApprovalInstance(ctx, &lark.CancelApprovalInstanceReq{
			ApprovalCode: ApprovalALLField.Code,
			InstanceCode: instanceCode,
			UserID:       UserAdmin.OpenID,
		})
		as.Nil(err)
	})

	t.Run("approve-reject", func(t *testing.T) {
		t.SkipNow()
		taskDone := map[string]bool{}
		instanceCode, instance := testCreateApproval(t, cli, ApprovalALLField.Code, UserAdmin.UserID, UserAdmin.OpenID)
		for taskIdx, task := range instance.TaskList {
			if taskDone[task.ID] {
				continue
			}
			taskDone[task.ID] = true
			fmt.Println("task", task.ID, task.NodeID)
			_, _, err := cli.Approval.ApproveApprovalInstance(ctx, &lark.ApproveApprovalInstanceReq{
				ApprovalCode: ApprovalALLField.Code,
				InstanceCode: instanceCode,
				UserID:       UserAdmin.UserID,
				TaskID:       task.ID,
				Comment:      ptrString(fmt.Sprintf("task: %d, approve", taskIdx)),
			})
			as.Nil(err)
		}

		resp, _, err := cli.Approval.GetApprovalInstance(ctx, &lark.GetApprovalInstanceReq{
			InstanceID: instanceCode,
			Locale:     nil,
		})
		as.Nil(err)
		for taskIdx, task := range resp.TaskList {
			if taskDone[task.ID] {
				continue
			}
			taskDone[task.ID] = true
			fmt.Println("task", task.ID, task.NodeID)
			_, _, err := cli.Approval.RejectApprovalInstance(ctx, &lark.RejectApprovalInstanceReq{
				ApprovalCode: ApprovalALLField.Code,
				InstanceCode: instanceCode,
				UserID:       UserAdmin.UserID,
				TaskID:       task.ID,
				Comment:      ptrString(fmt.Sprintf("task: %d, approve", taskIdx)),
			})
			as.Nil(err)
		}
	})
}

func testCreateApproval(t *testing.T, cli *lark.Lark, approvalCode, userID, openID string) (string, *lark.GetApprovalInstanceResp) {
	as := assert.New(t)

	var widgetDefine lark.ApprovalWidgetList
	{
		resp, _, err := cli.Approval.GetApproval(ctx, &lark.GetApprovalReq{
			ApprovalCode: approvalCode,
			Locale:       nil,
		})
		as.Nil(err)
		widgetDefine = resp.Form
	}

	v := lark.ApprovalWidgetList{
		{
			ID:    widgetDefine[0].ID,
			Type:  lark.ApprovalWidgetTypeInput,
			Value: "test",
		},
	}

	instanceCode := ""
	{
		resp, _, err := cli.Approval.CreateApprovalInstance(ctx, &lark.CreateApprovalInstanceReq{
			ApprovalCode:           approvalCode,
			UserID:                 &userID,
			OpenID:                 &openID,
			DepartmentID:           nil,
			Form:                   v,
			NodeApproverUserIDList: nil,
			NodeApproverOpenIDList: nil,
			UUID:                   nil,
		})
		as.Nil(err)
		instanceCode = resp.InstanceCode
	}

	resp, _, err := cli.Approval.GetApprovalInstance(ctx, &lark.GetApprovalInstanceReq{
		InstanceID: instanceCode,
		Locale:     nil,
	})
	as.Nil(err)

	return instanceCode, resp
}
