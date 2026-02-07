package uploadutils

import (
	"ayam_bangkok/source/pkg/logger"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var allowedImageExt = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

func isValidImage(file *multipart.FileHeader) bool {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	return allowedImageExt[ext]
}

func UploadImage(c *gin.Context, fieldName, folderPath string) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", err
	}

	if !isValidImage(file) {
		return "", err
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	fileName := uuid.New().String() + ext

	folderPath = strings.TrimSuffix(folderPath, "/")

	fullPath := folderPath + "/" + fileName

	if err := c.SaveUploadedFile(file, fullPath); err != nil {
		logger.Error().Err(err).Msg("Failed to upload file")
		return "", err
	}

	return fileName, nil
}