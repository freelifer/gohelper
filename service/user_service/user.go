package user_service

type UserService struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (s *UserService) Register() {

}

func (s *UserService) Login() {

}
