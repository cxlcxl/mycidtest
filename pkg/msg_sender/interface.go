package msg_sender

import "xiaoniuds.com/cid/pkg/errs"

type TextMsg struct {
	MsgFmt string
	Params []interface{}
}

type MarkdownMsg struct {
	Title   string
	Content string
}

type NotifyMsg struct {
}

type MessageSenderInterface interface {
	MarkdownMessage(markdown *MarkdownMsg) *NotifyMsg
	PostMessage(rows interface{})
	TextMessage(msg *TextMsg) *NotifyMsg
	Send(message *NotifyMsg) *errs.MyErr
}

var (
	messageSenderMap = map[string]MessageSenderInterface{}
)

func NewMessageSender(notifyMethod string, webhooks []string) MessageSenderInterface {
	return nil
}
