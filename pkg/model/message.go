package model

type CreateUserRequest struct {
	Name PersonName `json:"name"`
}

type CreateUserResponse struct {
	ID UserID `json:"id"`
}
