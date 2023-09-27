package _usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
	"github.com/DueIt-Jasanya-Aturuang/toy-story/infra"
	"github.com/DueIt-Jasanya-Aturuang/toy-story/pkg/converter"
)

type IconUsecaseImpl struct {
	iconRepo  domain.IconRepository
	minioRepo domain.MinioRepo
}

func NewIconUsecaseImpl(
	iconRepo domain.IconRepository,
	minioRepo domain.MinioRepo,
) domain.IconUsecase {
	return &IconUsecaseImpl{
		iconRepo:  iconRepo,
		minioRepo: minioRepo,
	}
}

func (i *IconUsecaseImpl) Create(ctx context.Context, req *domain.RequestCreateIcon) (*domain.ResponseIcon, error) {
	icon, err := i.iconRepo.GetByTitle(ctx, req.Title)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}
	if icon != nil {
		return nil, TitleIconExist
	}

	fileExt := filepath.Ext(req.Icon.Filename)
	fileName := i.minioRepo.GenerateFileName(fileExt, "icon-images/public/")
	iconConv := converter.RequestCreateIconToModel(req, fmt.Sprintf("/%s/%s", infra.MinIoBucket, fileName))

	err = i.minioRepo.UploadFile(ctx, req.Icon, fileName)
	if err != nil {
		return nil, err
	}

	err = i.iconRepo.Create(ctx, iconConv)
	if err != nil {
		return nil, err
	}

	iconResp := converter.IconModelToResp(iconConv)
	return iconResp, nil
}

func (i *IconUsecaseImpl) Update(ctx context.Context, req *domain.RequestUpdateIcon) (*domain.ResponseIcon, error) {
	icon, err := i.iconRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, IconNotExist
		}
		return nil, err
	}

	if icon.Title != req.Title {
		iconTitle, err := i.iconRepo.GetByTitle(ctx, req.Title)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return nil, err
			}
		}
		if iconTitle != nil {
			return nil, TitleIconExist
		}
	}

	fileName := icon.Icon
	reqImageCondition := req.Icon != nil && req.Icon.Size > 0

	if reqImageCondition {
		fileExt := filepath.Ext(req.Icon.Filename)
		fileName = i.minioRepo.GenerateFileName(fileExt, "icon-images/public/")

		err = i.minioRepo.UploadFile(ctx, req.Icon, fileName)
		if err != nil {
			return nil, err
		}

		imageDelArr := strings.Split(icon.Icon, "/")
		imageDel := fmt.Sprintf("/%s/%s/%s", imageDelArr[2], imageDelArr[3], imageDelArr[4])
		err = i.minioRepo.DeleteFile(ctx, imageDel)
		if err != nil {
			return nil, err
		}
	}

	iconConv := converter.RequestUpdateIconToModel(req, fmt.Sprintf("/%s/%s", infra.MinIoBucket, fileName))

	err = i.iconRepo.Update(ctx, iconConv)
	if err != nil {
		return nil, err
	}

	iconResp := converter.IconModelToResp(iconConv)
	return iconResp, nil
}

func (i *IconUsecaseImpl) GetAll(ctx context.Context) (*[]domain.ResponseIcon, error) {
	icons, err := i.iconRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var iconResps []domain.ResponseIcon
	for _, icon := range *icons {
		iconResp := *converter.IconModelToResp(&icon)
		iconResps = append(iconResps, iconResp)
	}

	return &iconResps, nil
}

func (i *IconUsecaseImpl) Delete(ctx context.Context, id string) error {
	icon, err := i.iconRepo.GetByID(ctx, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		return IconNotExist
	}

	imageDelArr := strings.Split(icon.Icon, "/")
	imageDel := fmt.Sprintf("/%s/%s/%s", imageDelArr[2], imageDelArr[3], imageDelArr[4])
	if err = i.minioRepo.DeleteFile(ctx, imageDel); err != nil {
		return err
	}

	err = i.iconRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
