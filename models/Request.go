package models

import "time"

type Request struct {
	Id        int32
	CreatedAt time.Time
	UserUuid  string
	UserIp    string
	Url       string
	Data      string
}
