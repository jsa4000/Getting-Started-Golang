package users

// GetAllRequest struct request
type GetAllRequest struct {
}

// GetAllResponse struct Response
type GetAllResponse struct {
	Users []User
}

// GetByIDRequest struct request
type GetByIDRequest struct {
	ID string
}

// GetByIDResponse struct Response
type GetByIDResponse struct {
	User User
}

// CreateRequest request
type CreateRequest struct {
	Name     string `json:"name" validate:"min=0,max=255,required"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"min=0,max=255,required"`
}

// CreateResponse Response
type CreateResponse struct {
	User User
}

// DeleteByIDRequest struct request
type DeleteByIDRequest struct {
	ID string
}

// DeleteByIDResponse struct Response
type DeleteByIDResponse struct {
}
