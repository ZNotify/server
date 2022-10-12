package types

type Response[T any] struct {
	Code int `json:"code" default:"200"`
	Body T   `json:"body"`
}
