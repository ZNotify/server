package common

type Response[T any] struct {
	Code int `json:"code" default:"200"`
	Body T   `json:"body"`
}

type BadRequestResponse struct {
	Code int    `json:"code" default:"400"`
	Body string `json:"body"`
}

type UnauthorizedResponse struct {
	Code int    `json:"code" default:"401"`
	Body string `json:"body"`
}

type NotFoundResponse struct {
	Code int    `json:"code" default:"404"`
	Body string `json:"body"`
}

type InternalServerErrorResponse struct {
	Code int    `json:"code" default:"500"`
	Body string `json:"body"`
}
