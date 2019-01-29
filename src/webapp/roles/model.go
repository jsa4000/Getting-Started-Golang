package roles

// Role struct to define an Role
type Role struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name"`
}

// New creates new instance
func New(id, name string) Role {
	return Role{
		ID:   id,
		Name: name,
	}
}
