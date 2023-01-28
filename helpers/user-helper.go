package helpers

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/folklore13/task-5-vix-btpns-ario/database"
	"github.com/folklore13/task-5-vix-btpns-ario/models"
)

type UserHelper interface {
	Update(user database.UserUpdateData) models.User
	Profile(userID string) models.User
}

type userHelper struct {
	userRepo models.UserRepo
}

func NewUserHelper(userHelp models.UserRepo) UserHelper {
	return &userHelper{
		userRepo: userHelp,
	}
}

func (helper *userHelper) Update(user database.UserUpdateData) models.User {
	userUpdate := models.User{}
	err := smapping.FillStruct(&userUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed mao %v:", err)
	}
	updatedUser := helper.userRepo.UpdateUser(userUpdate)
	return updatedUser
}

func (helper *userHelper) Profile(userID string) models.User {
	return helper.userRepo.ProfileUser(userID)
}