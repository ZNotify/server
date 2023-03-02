package senderConfig

type WebPushConfig struct {
	// The public key of the VAPID key pair. Read webpush documentation for more information.
	VAPIDPublicKey string `json:"vapid_public_key" jsonschema:"required,title=VAPID Public Key"`
	// The private key of the VAPID key pair. Read webpush documentation for more information.
	VAPIDPrivateKey string `json:"vapid_private_key" jsonschema:"required,title=VAPID Private Key"`
	// The mail address of the sender.
	Mailto string `json:"mail_to" jsonschema:"required,title=Mail to,format=email"`
}
