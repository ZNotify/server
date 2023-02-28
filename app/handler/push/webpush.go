package push

import (
	"notify-api/app/common"
	"notify-api/app/manager/push"
	"notify-api/app/manager/push/enum"
	"notify-api/app/manager/push/provider/webpush"
)

type info struct {
	Enable    bool   `json:"enable"`
	PublicKey string `json:"public_key"`
}

// WebPush
//
//	@Summary      Endpoint for webpush info check
//	@Id           webpush
//	@Tags         Push
//	@Description  Check if this znotify instance support webpush and get public key
//	@Produce      json
//	@Success      200  {object}  common.Response[info]
//	@Router       /webpush [get]
func WebPush(ctx *common.Context) {
	enable := push.IsSenderActive(enum.SenderWebPush)
	if enable {
		sender, _ := push.GetSender(enum.SenderWebPush)
		webpushSender := sender.(*webpush.Provider)

		ctx.JSONResult(info{
			Enable:    true,
			PublicKey: webpushSender.VAPIDPublicKey,
		})
	} else {
		ctx.JSONResult(info{
			Enable:    false,
			PublicKey: "",
		})
	}
}
