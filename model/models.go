package model

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identify      string
	ClientIp      string
	ClientPort    string
	LoginTime     *time.Time "default:null"
	HeartbeatTime *time.Time "default:null"
	LogOutTime    *time.Time "default:null"
	IsLogout      bool
	DeviceInfo    string
}
