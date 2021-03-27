package user

// Repository contains the definition of methods used to access data storages such as databases of memory data
type Repository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
	GetByID(id string) (*User, error)
}
