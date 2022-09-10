package account

import (
	"fita-backend-test/app/account/data/repository"
	"fita-backend-test/app/account/domain/usecase"
	"fita-backend-test/app/account/infrastructure"
)

type AccountService interface {
	usecase.UserUsecase
}

func NewAccountService() AccountService {
	memory := infrastructure.NewInMemory()
	repository := repository.NewUserRepository(memory)

	return usecase.NewUserUsecase(repository)
}