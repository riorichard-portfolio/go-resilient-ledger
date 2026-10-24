package fiberhttp

type CreateAccountReq struct {
	Name    string `json:"name" validate:"required"`
	Balance int    `json:"balance" validate:"gte=0"`
}

type FindAccountReq struct {
	ID string `json:"id" validate:"required"`
}
