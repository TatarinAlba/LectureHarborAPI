package usecase

type UserUsecase struct {
	repository UserRepository
}

func NewUserUsecase(repository UserRepository) *UserUsecase {
	return &UserUsecase{repository: repository}
}

func (service *UserUsecase) GetTranscription() {

}
