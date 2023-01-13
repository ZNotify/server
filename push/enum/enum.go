package enum

type Priority string

const (
	PriorityLow    Priority = "low"    // low
	PriorityNormal Priority = "normal" // normal
	PriorityHigh   Priority = "high"   // high
)

type Sender string

const (
	SenderFcm       Sender = "FCM"       // FCM
	SenderWebPush   Sender = "WebPush"   // WebPush
	SenderWns       Sender = "WNS"       // WNS
	SenderTelegram  Sender = "Telegram"  // Telegram
	SenderWebSocket Sender = "WebSocket" // WebSocket
)
