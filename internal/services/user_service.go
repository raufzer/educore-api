package services


type UserService struct {
	createUser(*models.User) error
	getUser(*string) (*models.User, error)
	getAllUsers() ([]*models.User, error)
	updateUser(*model.User) error
	deleteUser(*string) error
}


