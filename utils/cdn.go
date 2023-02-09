package utils

import (
	"strings"

	"gitlab.id.vin/gami/gami-common/configs"
	"gitlab.id.vin/gami/gami-common/models"
)

const CDNMode = "/cdn-cgi/image/f=auto,fit=scale-down,quality=75/"

// ConvertImageCDN returns redis key of a id
func ConvertImageCDN(url string) string {
	if configs.AppConfig.Env == models.EnvProd {
		return strings.Replace(url, configs.AppConfig.CDN, configs.AppConfig.CDN+CDNMode, -1)
	}
	return url
}
