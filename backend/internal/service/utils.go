package service

import (
	"net/http"
	"os/exec"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func OpenDirectory(directory models.Directory) *customError.Error {
	err := exec.Command("explorer", directory.Path).Start()

	if err != nil {
		return &customError.Error{StatusCode: http.StatusBadRequest, Type: "INVALID_PATH", Msg: "FAILED TO OPEN DIRECTORY"}
	}

	return nil
}
