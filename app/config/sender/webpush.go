package senderConfig

type WebPushConfig struct {
	VAPIDPublicKey  string `yaml:"vapid_public_key"`
	VAPIDPrivateKey string `yaml:"vapid_private_key"`
	Mailto          string `yaml:"mailto"`
}
