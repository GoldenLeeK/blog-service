package global

import (
	"github.com/GoldenLeeK/blog-service/pkg/logger"
	"github.com/GoldenLeeK/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
