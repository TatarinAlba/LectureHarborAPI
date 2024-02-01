package usecase

import "github.com/TatarinAlba/LectureHarborAPI/internal/entity"

type (
	UserRepository interface {
		CreateUser(user entity.User) error
		UpdateUser(user entity.User) error
		RemoveUser(user entity.User) error
		GetUserById(id int64) (entity.User, error)
	}
)
