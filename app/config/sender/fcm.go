package senderConfig

type FCMConfig struct {
	// Should be the content of service account or refresh token JSON credentials you got from Google.
	Credential string `json:"credential" jsonschema:"required,title=Firebase messaging credential json"`
}
