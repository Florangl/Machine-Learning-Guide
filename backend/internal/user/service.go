package user

func GetUserByIDService(id int) (User, error) {
	return FindUserByID(id)
}
