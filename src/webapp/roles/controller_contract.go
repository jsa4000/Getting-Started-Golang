package roles

// GetAllRequest struct request
type  GetAllRequest struct {
}

// GetAllResponse struct Response
type  GetAllResponse struct {
	Roles []Role
}

// GetByIDRequest struct request
type  GetByIDRequest struct {
	ID string
}

// GetByIDResponse struct Response
type  GetByIDResponse struct {
	Role Role
}

// CreateRequest request
type  CreateRequest struct {
	Name string
}

// CreateResponse Response
type  CreateResponse struct {
	Role Role
}
	
// DeleteByIDRequest struct request
type  DeleteByIDRequest struct {
	ID string
}
	
// DeleteByIDResponse struct Response
type  DeleteByIDResponse struct {
}
	

