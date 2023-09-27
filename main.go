package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/api/rest"
	"github.com/DueIt-Jasanya-Aturuang/toy-story/api/rest/mapi"
	"github.com/DueIt-Jasanya-Aturuang/toy-story/infra"
	"github.com/DueIt-Jasanya-Aturuang/toy-story/pkg/_repository"
	"github.com/DueIt-Jasanya-Aturuang/toy-story/pkg/_usecase"
)

func main() {
	infra.InitEnv()
	infra.InitLog()
	db := infra.NewPostgresConnection()
	minioC := infra.NewMinioClient()
	//
	minioRepo := _repository.NewMinioImpl(minioC)
	// iconRepo := _repository.NewIconRepositoryImpl(db)
	paymentRepo := _repository.NewPaymentRepositoryImpl(db)
	// spendingTypeRepo := _repository.NewSpendingTypeRepository(db)
	//
	// iconUsecase := _usecase.NewIconUsecaseImpl(iconRepo, minioRepo)
	paymentUsecase := _usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)
	// spendingTypeUsecase := _usecase.NewSpendingTypeUsecaseImpl(spendingTypeRepo)

	paymentHandler := rest.NewPaymentHandlerImpl(paymentUsecase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		r.Use(mapi.BasicAuth)

		r.Get("/payment", paymentHandler.GetAll)
		r.Post("/payment", paymentHandler.Create)
		r.Put("/payment/{id}", paymentHandler.Update)
		r.Delete("/payment/{id}", paymentHandler.Delete)
	})

	err := http.ListenAndServe(infra.AppAddr, r)
	if err != nil {
		panic(err)
	}
}
