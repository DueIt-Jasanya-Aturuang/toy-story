package converter

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
)

func RequestCreateIconToModel(req *domain.RequestCreateIcon, icon string) *domain.Icon {
	id := uuid.NewV4().String()
	return &domain.Icon{
		ID:    id,
		Title: req.Title,
		Icon:  icon,
		AuditInfo: domain.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: "",
			UpdatedAt: time.Now().Unix(),
		},
	}
}

func RequestUpdateIconToModel(req *domain.RequestUpdateIcon, icon string) *domain.Icon {
	return &domain.Icon{
		ID:    req.ID,
		Title: req.Title,
		Icon:  icon,
		AuditInfo: domain.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: "",
			UpdatedAt: time.Now().Unix(),
		},
	}
}

func IconModelToResp(i *domain.Icon) *domain.ResponseIcon {
	return &domain.ResponseIcon{
		ID:    i.ID,
		Title: i.Title,
		Icon:  i.Icon,
	}
}
