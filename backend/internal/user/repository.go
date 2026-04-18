package user

import "meditrack/configs"

func FindUserByID(id int) (User, error) {
	var user User
	err := configs.DB.First(&user, id).Error
	return user, err
}
