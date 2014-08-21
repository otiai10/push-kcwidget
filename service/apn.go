package service

import "fmt"
import apns "github.com/anachronistic/apns"

type ApnClient struct {
	// implements PushClient
	set      PushSet
	err      error
	host     string
	certPath string
	keyPath  string
}

func init() {
	//
}

func (c *ApnClient) Send() error {
	return c.sendToApn(c.set)
}

// もしPersistent Connetctionを使うことになったら
// initとかでvarにchanつくって流し込むようにしよう
// とりあえずひとつひとつリクエストする
func (c *ApnClient) sendToApn(set PushSet) (e error) {
	payload := apns.NewPayload()
	// {{{ 雑
	var message string = "デフォルトメッセージ"
	if len(set.Events()) > 0 {
		message = set.Events()[0].Message
	}
	payload.Alert = message
	// }}}
	payload.Sound = "default"

	pushNotification := apns.NewPushNotification()
	pushNotification.DeviceToken = c.set.Token()
	pushNotification.AddPayload(payload)

	client := apns.NewClient(
		c.host,
		c.certPath,
		c.keyPath,
	)

	if resp := client.Send(pushNotification); resp.Error != nil {
		return fmt.Errorf("[APN Error] %v", resp.Error)
	}
	return
}
