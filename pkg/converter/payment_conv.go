package converter

import (
	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
)

func RequestCreatePaymentToModel(req *domain.RequestCreatePayment) *domain.Payment {
	return &domain.Payment{
		AuditInfo: domain.AuditInfo{
			CreatedAt: 0,
			CreatedBy: "",
			UpdatedAt: 0,
		},
	}
}

func RequestUpdatePaymentToModel(req *domain.RequestUpdatePayment) *domain.Payment {
	return &domain.Payment{
		AuditInfo: domain.AuditInfo{
			CreatedAt: 0,
			CreatedBy: "",
			UpdatedAt: 0,
		},
	}
}

func PaymentModelToResp(i *domain.Payment) *domain.ResponsePayment {
	return &domain.ResponsePayment{
		ID: "",
	}
}
