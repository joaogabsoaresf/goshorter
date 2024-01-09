package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type %s) is required", name, typ)
}

// func errParamAlreadyExist(name, typ string) error {
// 	return fmt.Errorf("param %s (type %s) already exist", name, typ)
// }

type CreateUrlRequest struct {
	OriginalPath string `json:"original_path"`
}

func (r *CreateUrlRequest) Validate() error {
	if r.OriginalPath == "" {
		return errParamIsRequired("original_path", "string")
	}
	return nil
}

// TO-DO:
// - Criar validacao no banco de dados para domain
