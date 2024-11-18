package services


type UserService struct {
	CreateUser(*models.User
	GetUser(*string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(*model.User) error
	DeleteUser(*string) erro
}