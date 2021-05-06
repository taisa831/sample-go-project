package usecase

import "github.com/taisa831/sample-go-project/domain/model"

type UserUsecase struct{}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}

func (u *UserUsecase) List() []model.User {
	user := model.NewUser()
	return []model.User{user}
}

func (u *UserUsecase) Edit() model.User {
	return model.User{
		ID:        0,
		Email:     "",
		FirstName: "",
		LastName:  "",
	}
}
