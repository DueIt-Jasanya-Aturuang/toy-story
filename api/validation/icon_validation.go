package validation

import (
	"fmt"
	"strings"

	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
)

func CreateIcon(req *domain.RequestCreateIcon) error {
	err := map[string][]string{}

	if req.Icon == nil {
		err["icon"] = append(err["icon"], required)
	}
	if req.Icon != nil && req.Icon.Size > 0 {
		if req.Icon.Size > 2097152 {
			err["icon"] = append(err["icon"], fmt.Sprintf(fileSize, 2048, 2))
		}
		if !checkContentType(req.Icon.Header.Get("Content-Type"), image) {
			err["icon"] = append(err["icon"], fmt.Sprintf(fileContent, strings.Join(contentType(image), " or ")))
		}
	}

	if req.Title == "" {
		err["title"] = append(err["title"], required)
	}
	title := maxMinString(req.Title, 3, 32)
	if title != "" {
		err["title"] = append(err["title"], title)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}
	return nil
}

func UpdateIcon(req *domain.RequestUpdateIcon) error {
	err := map[string][]string{}

	if req.Icon == nil {
		err["icon"] = append(err["icon"], required)
	}
	if req.Icon != nil && req.Icon.Size > 0 {
		if req.Icon.Size > 2097152 {
			err["icon"] = append(err["icon"], fmt.Sprintf(fileSize, 2048, 2))
		}
		if !checkContentType(req.Icon.Header.Get("Content-Type"), image) {
			err["icon"] = append(err["icon"], fmt.Sprintf(fileContent, strings.Join(contentType(image), " or ")))
		}
	}

	if req.Title == "" {
		err["title"] = append(err["title"], required)
	}
	title := maxMinString(req.Title, 3, 32)
	if title != "" {
		err["title"] = append(err["title"], title)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}

	return nil
}
