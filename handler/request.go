package handler

import (
	"fmt"
	"regexp"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type %s) is required", name, typ)
}

func errDocumentNotFound(id string) error {
	return fmt.Errorf("document %s not found", id)
}

func errInvalidUrl(id string) error {
	return fmt.Errorf("invalid url. %s must start with: http|https:// and end with /", id)
}

type CreateUrlRequest struct {
	OriginalPath string `json:"original_path"`
}

func (r *CreateUrlRequest) Validate() error {
	if r.OriginalPath == "" {
		return errParamIsRequired("original_path", "string")
	}
	if !isValidURL((r.OriginalPath)) {
		return errInvalidUrl(r.OriginalPath)
	}
	return nil
}

func isValidURL(url string) bool {
	urlRegex := regexp.MustCompile(`^(http|https):\/\/[a-zA-Z0-9\-_]+(\.[a-zA-Z]{2,})+(\/[a-zA-Z0-9\-._~:/?#[\]@!$&'()*+,;=]*)?$`)

	return urlRegex.MatchString(url)
}
