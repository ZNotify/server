package enum

type Sender string

const (
	SenderFcm       Sender = "FCM"       // FCM
	SenderWebPush   Sender = "WebPush"   // WebPush
	SenderWns       Sender = "WNS"       // WNS
	SenderTelegram  Sender = "Telegram"  // Telegram
	SenderWebSocket Sender = "WebSocket" // WebSocket
)
