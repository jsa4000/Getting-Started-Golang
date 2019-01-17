package roles

// Role struct to define an Role
type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// New creates new instance
func New(id string, name string) Role {
	return Role{
		ID:   id,
		Name: name,
	}
}
