package converter

import (
	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
)

func RequestCreateSpendingTypeToModel(req *domain.RequestCreateSpendingType) *domain.SpendingType {
	return &domain.SpendingType{
		AuditInfo: domain.AuditInfo{
			CreatedAt: 0,
			CreatedBy: "",
			UpdatedAt: 0,
		},
	}
}

func RequestUpdateSpendingTypeToModel(req *domain.RequestUpdateSpendingType) *domain.SpendingType {
	return &domain.SpendingType{
		AuditInfo: domain.AuditInfo{
			CreatedAt: 0,
			CreatedBy: "",
			UpdatedAt: 0,
		},
	}
}

func SpendingTypeModelToResp(i *domain.SpendingType) *domain.ResponseSpendingType {
	return &domain.ResponseSpendingType{
		ID: "",
	}
}
