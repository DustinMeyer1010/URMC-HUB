package db

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/google/uuid"
)

// Returns the name of image file
func SaveImage(image multipart.File, header *multipart.FileHeader) string {
	imageName := fmt.Sprintf("%s_%s.%s", global.Username, uuid.New(), filepath.Ext(header.Filename))
	imageLocation := filepath.Join(DBImageLocation, imageName)
	file, _ := os.Create(imageLocation)
	io.Copy(file, image)
	file.Close()

	return imageName
}
