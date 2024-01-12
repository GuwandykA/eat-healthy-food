package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"saglyk-backend/internal/config"
	"strconv"
	"time"

	"github.com/vansante/go-ffprobe"
)

func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}

		return nil
	}

	return
}

func CreateFile(folderName string, file *multipart.FileHeader, r *gin.Context) (string, error) {
	cfg := config.GetConfig()
	appName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(filepath.Base(file.Filename)))
	appDirPath := cfg.PublicFilePath + "/" + folderName

	err := os.MkdirAll(appDirPath, os.ModePerm)
	if err != nil {
		return "str", err
	}

	pathApp := fmt.Sprintf("%s/%s", appDirPath, appName)
	if err := r.SaveUploadedFile(file, pathApp); err != nil {
		return "", err
	}

	pathApp = "/public/" + folderName + "/" + appName
	return pathApp, err
}

func RemoveFile(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

func AudioVideoDuration(filePath string) (string, error) {

	data, err := ffprobe.GetProbeData(filePath, 120000*time.Millisecond)
	if err != nil {
		log.Printf("Error getting data: %v", err)
		return "", err
	}

	duration := data.Format.Duration().Milliseconds()
	var stringDuration string
	stringDuration = strconv.FormatInt(duration, 10)
	return stringDuration, nil
}
