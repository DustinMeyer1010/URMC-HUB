package service

import (
	"net/http"
	"os/exec"

	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func OpenDirectory(directory models.Directory) error {
	err := exec.Command("explorer", directory.Path).Start()

	if err != nil {
		return &errs.ApiError{StatusCode: http.StatusBadRequest, Type: "INVALID_PATH", Msg: "FAILED TO OPEN DIRECTORY"}
	}

	return nil
}
