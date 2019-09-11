package service

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/mp3tags/request-repository-proto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService_CreateRequest(t *testing.T) {
	asrt := assert.New(t)

	srv := Service{}
	srv.Db = ConnectToDb()

	_, err := srv.CreateRequest(context.Background(), &request_repository.CreateRequestParams{
		UserUuid:  "12",
		UserIp:    "14",
		Url:       "asad",
		CreatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
		Data:      "asdadad",
	})
	asrt.NoError(err)
	_, err = srv.Db.Exec("DELETE from request")
	asrt.NoError(err)
}
