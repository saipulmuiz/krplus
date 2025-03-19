package config

import (
	"github.com/saipulmuiz/krplus/pkg/serror"
	"github.com/saipulmuiz/krplus/service/handler/rest"
	"github.com/saipulmuiz/krplus/service/repository"
	"github.com/saipulmuiz/krplus/service/usecase"
)

func (cfg *Config) InitService() (errx serror.SError) {
	userRepo := repository.NewUserRepository(cfg.DB)
	userUsecase := usecase.NewUserUsecase(userRepo)

	creditRepo := repository.NewCreditRepo(cfg.DB)
	creditUsecase := usecase.NewCreditUsecase(creditRepo, userRepo)

	transactionRepo := repository.NewTransactionRepo(cfg.DB)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo, creditRepo, userRepo)

	route := rest.CreateHandler(
		userUsecase,
		creditUsecase,
		transactionUsecase,
	)

	cfg.Server = route

	return nil
}
