package main

import (
	"github.com/DueIt-Jasanya-Aturuang/toy-story/infra"
)

func main() {
	infra.InitEnv()
	infra.InitLog()
	// db := infra.NewPostgresConnection()
	// minioC := infra.NewMinioClient()
	//
	// minioRepo := _repository.NewMinioImpl(minioC)
	// iconRepo := _repository.NewIconRepositoryImpl(db)
	// paymentRepo := _repository.NewPaymentRepositoryImpl(db)
	// spendingTypeRepo := _repository.NewSpendingTypeRepository(db)
	//
	// iconUsecase := _usecase.NewIconUsecaseImpl(iconRepo, minioRepo)
	// paymentUsecase := _usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)
	// spendingTypeUsecase := _usecase.NewSpendingTypeUsecaseImpl(spendingTypeRepo)
}
