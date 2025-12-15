package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// refactor This needs to be changed to handler the person username and the name of the file
func SaveImage(image multipart.File, header *multipart.FileHeader) error {
	exe, _ := os.Executable()
	base := filepath.Dir(exe)

	folder := filepath.Join(base, "BOOKMARK_IMAGES")

	os.MkdirAll(folder, 0755)
	now := time.Now().Format("2006-01-02 15:04:05")
	now = strings.ReplaceAll(now, ":", "-")

	filename := "test_" + now + filepath.Ext(header.Filename)
	fullpath := filepath.Join(folder, filename)

	fmt.Println("Saving to:", fullpath)

	dst, _ := os.Create(fullpath)
	_, err := io.Copy(dst, image)

	fmt.Println(err)

	return nil

}
