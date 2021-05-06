package usecase

import (
	"reflect"
	"testing"

	"github.com/taisa831/sample-go-project/domain/model"
)

func TestUserUsecase_List(t *testing.T) {
	tests := []struct {
		name string
		u    *UserUsecase
		want func() []model.User
	}{
		{
			name: "List",
			want: func() []model.User {
				user := model.NewUser()
				return []model.User{user}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserUsecase{}
			if got := u.List(); !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("UserUsecase.List() = %v, want %v", got, tt.want())
			}
		})
	}
}
