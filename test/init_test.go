package test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

type App struct {
	AppID     string
	AppSecret string
}

func (r *App) Ins() *lark.Lark {
	return lark.New(
		lark.WithAppCredential(r.AppID, r.AppSecret),
		lark.WithTimeout(time.Second*20),
	)
}

type Helpdesk struct {
	AppID     string
	AppSecret string
	ID        string
	Token     string
}

func (r *Helpdesk) Ins() *lark.Lark {
	return lark.New(
		lark.WithAppCredential(r.AppID, r.AppSecret),
		lark.WithHelpdeskCredential(r.ID, r.Token),
		lark.WithTimeout(time.Second*20),
	)
}

var HelpdeskAllPermission = Helpdesk{
	AppID:     os.Getenv("LARK_APP_ALL_PERMISSION_APP_ID"),
	AppSecret: os.Getenv("LARK_APP_ALL_PERMISSION_APP_SECRET"),
	ID:        os.Getenv("LARK_HELPDESK_ALL_PERMISSION_ID"),
	Token:     os.Getenv("LARK_HELPDESK_ALL_PERMISSION_TOKEN"),
}

type User struct {
	UserID string
	OpenID string
	Name   string
}

type Chat struct {
	ChatID string
	Name   string
}

var AppNoPermission = App{
	AppID:     os.Getenv("LARK_APP_NO_PERMISSION_APP_ID"),
	AppSecret: os.Getenv("LARK_APP_NO_PERMISSION_APP_SECRET"),
}

var AppALLPermission = App{
	AppID:     os.Getenv("LARK_APP_ALL_PERMISSION_APP_ID"),
	AppSecret: os.Getenv("LARK_APP_ALL_PERMISSION_APP_SECRET"),
}

var UserAdmin = User{
	UserID: os.Getenv("LARK_USER_ADMIN_USER_ID"),
	OpenID: os.Getenv("LARK_USER_ADMIN_OPEN_ID"),
	Name:   os.Getenv("LARK_USER_ADMIN_NAME"),
}

// 这个群公共，必须设置「群公共」三个字
var ChatContainALLPermissionApp = Chat{
	ChatID: os.Getenv("LARK_CHAT_CONTAINS_APP_PERMISSION_APP_CHAT_ID"),
	Name:   "包含「lark-sdk」的群",
}

var ChatNotContainALLPermissionApp = Chat{
	ChatID: os.Getenv("LARK_CHAT_NOT_CONTAINS_APP_PERMISSION_APP_CHAT_ID"),
	Name:   "不包含「lark-sdk」的群",
}

var ChatForSendMessage = Chat{
	ChatID: os.Getenv("LARK_CHAT_FOR_SEND_MSG_CHAT_ID"),
	Name:   "for-send-message",
}

type File struct {
	Key string
}

var File1 = File{
	Key: os.Getenv("LARK_FILE_KEY_TEST_FILE_1_PNG"), // this is file of ./test/file_1.png
}

var File2 = File{
	Key: os.Getenv("LARK_FILE_KEY_TEST_FILE_2_DOC"), // ./test/file_2.docx
}

type Message struct {
	MessageID string
	ChatID    string
}

var MessageAdminSendTextInChatContainAllPermissionApp = Message{
	MessageID: os.Getenv("LARK_MESSAGE_ADMIN_SEND_TEST_IN_CHAT_CONTAINS_ALL_PERMISSION_APP"),
	ChatID:    os.Getenv("LARK_CHAT_CONTAINS_APP_PERMISSION_APP_CHAT_ID"),
}

var MessageAdminSendImageInChatContainAllPermissionApp = Message{
	MessageID: os.Getenv("LARK_MESSAGE_ADMIN_SEND_IMAGE_IN_CHAT_CONTAINS_ALL_PERMISSION_APP"),
	ChatID:    os.Getenv("LARK_CHAT_CONTAINS_APP_PERMISSION_APP_CHAT_ID"),
}

var MessageAllPermissionAppSendTextInChatContainAllPermissionApp = Message{
	MessageID: os.Getenv("LARK_MESSAGE_ALL_PERMISSION_APP_SEND_TEXT_IN_CHAT_CONTAINS_ALL_PERMISSION_APP"),
	ChatID:    os.Getenv("LARK_CHAT_CONTAINS_APP_PERMISSION_APP_CHAT_ID"),
}

func Test_Config(t *testing.T) {
	as := assert.New(t)

	as.NotEmpty(AppNoPermission.AppID)
	as.NotEmpty(AppNoPermission.AppSecret)
	as.NotEmpty(AppALLPermission.AppID)
	as.NotEmpty(AppALLPermission.AppSecret)
	as.NotEmpty(UserAdmin.UserID)
	as.NotEmpty(UserAdmin.OpenID)
	as.NotEmpty(ChatContainALLPermissionApp.ChatID)
	as.NotEmpty(ChatNotContainALLPermissionApp.ChatID)
	as.NotEmpty(ChatForSendMessage.ChatID)
	as.NotEmpty(File1.Key)
	as.NotEmpty(File2.Key)
	as.NotEmpty(MessageAdminSendTextInChatContainAllPermissionApp.ChatID)
	as.NotEmpty(MessageAdminSendTextInChatContainAllPermissionApp.MessageID)
	as.NotEmpty(MessageAdminSendImageInChatContainAllPermissionApp.ChatID)
	as.NotEmpty(MessageAdminSendImageInChatContainAllPermissionApp.MessageID)
	as.NotEmpty(MessageAllPermissionAppSendTextInChatContainAllPermissionApp.ChatID)
	as.NotEmpty(MessageAllPermissionAppSendTextInChatContainAllPermissionApp.MessageID)
}
