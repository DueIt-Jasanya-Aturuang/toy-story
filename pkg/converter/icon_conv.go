package converter

import (
	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
)

func RequestCreateIconToModel(req *domain.RequestCreateIcon) *domain.Icon {
	return &domain.Icon{
		AuditInfo: domain.AuditInfo{
			CreatedAt: 0,
			CreatedBy: "",
			UpdatedAt: 0,
		},
	}
}

func RequestUpdateIconToModel(req *domain.RequestUpdateIcon) *domain.Icon {
	return &domain.Icon{
		AuditInfo: domain.AuditInfo{
			CreatedAt: 0,
			CreatedBy: "",
			UpdatedAt: 0,
		},
	}
}

func IconModelToResp(i *domain.Icon) *domain.ResponseIcon {
	return &domain.ResponseIcon{
		ID: "",
	}
}
