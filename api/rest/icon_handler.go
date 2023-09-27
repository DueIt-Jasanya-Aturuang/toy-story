package rest

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/api/rest/helper"
	"github.com/DueIt-Jasanya-Aturuang/toy-story/api/validation"
	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
	"github.com/DueIt-Jasanya-Aturuang/toy-story/pkg/_usecase"
)

type IconHandlerImpl struct {
	iconUsecase domain.IconUsecase
}

func NewIconHandlerImpl(
	iconUsecase domain.IconUsecase,
) *IconHandlerImpl {
	return &IconHandlerImpl{
		iconUsecase: iconUsecase,
	}
}

func (h *IconHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestCreateIcon)

	err := helper.ParseForm(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	_, fileHeader, _ := r.FormFile("icon")
	req.Icon = fileHeader

	err = validation.CreateIcon(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	icon, err := h.iconUsecase.Create(r.Context(), req)
	if err != nil {
		if errors.Is(err, _usecase.TitleIconExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"title": {
					err.Error(),
				},
			}, response.CM06)
		}

		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, icon, "created icon berhasil")
}

func (h *IconHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestUpdateIcon)

	err := helper.ParseForm(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01))
		return
	}
	req.ID = id

	_, fileHeader, _ := r.FormFile("icon")
	req.Icon = fileHeader

	err = validation.UpdateIcon(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	icon, err := h.iconUsecase.Update(r.Context(), req)
	if err != nil {
		if errors.Is(err, _usecase.TitleIconExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"title": {
					err.Error(),
				},
			}, response.CM06)
		} else if errors.Is(err, _usecase.IconNotExist) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}

		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, icon, "update icon berhasil")
}

func (h *IconHandlerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	payments, err := h.iconUsecase.GetAll(r.Context())
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, payments, "data icon")
}

func (h *IconHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01))
		return
	}

	err := h.iconUsecase.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, _usecase.IconNotExist) {
			helper.SuccessResponseEncode(w, nil, "delete icon berhasil, icon gak ada")
			return
		}

		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, nil, "delete icon berhasil")
}
