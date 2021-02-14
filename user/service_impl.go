package user

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

func (s *userService) Register(user *User) error {
	return s.userRepo.Create(user)
}