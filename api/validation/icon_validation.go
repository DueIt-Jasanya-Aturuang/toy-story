package validation

import (
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
)

func CreateIcon(req *domain.RequestCreateIcon) error {
	errBadRequest := map[string][]string{}

	if len(errBadRequest) != 0 {
		return _error.HttpErrMapOfSlices(errBadRequest, response.CM06)
	}
	return nil
}

func UpdateIcon(req *domain.RequestCreateIcon) error {
	errBadRequest := map[string][]string{}

	if len(errBadRequest) != 0 {
		return _error.HttpErrMapOfSlices(errBadRequest, response.CM06)
	}
	return nil
}
