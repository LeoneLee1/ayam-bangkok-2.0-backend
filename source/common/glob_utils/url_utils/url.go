package urlutils

import "github.com/gin-gonic/gin"

func BuildStorageURL(c *gin.Context, path string) string {
	if path == "" {
		return ""
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	return scheme + "://" + c.Request.Host + "/storage/" + path
}