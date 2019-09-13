package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/mp3tags/request-repository-proto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService_ListRequests(t *testing.T) {
	asrt := assert.New(t)

	srv := Service{}
	srv.Db = ConnectToDb()

	var err error

	_, err = srv.CreateRequest(context.Background(), &request_repository.CreateRequestParams{
		UserUuid:  "1",
		UserIp:    "11",
		Url:       "asad",
		CreatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
		Data:      "asdadad",
	})
	asrt.NoError(err)

	_, err = srv.CreateRequest(context.Background(), &request_repository.CreateRequestParams{
		UserUuid:  "2",
		UserIp:    "22",
		Url:       "dsa",
		CreatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
		Data:      "dad",
	})
	asrt.NoError(err)

	list, err := srv.ListRequests(context.Background(), &request_repository.ListRequestsParams{
		Limit:  10,
		Offset: 0,
	})
	asrt.NoError(err)
	asrt.Equal(2, len(list.Requests))
	asrt.Equal(int32(2), list.Total)
	if len(list.Requests) == 2 {
		asrt.Equal("1", list.Requests[1].UserUuid)
		asrt.Equal("11", list.Requests[1].UserIp)
		asrt.Equal("asad", list.Requests[1].Url)
		asrt.Equal("asdadad", list.Requests[1].Data)
		asrt.NotNil(list.Requests[1].CreatedAt)

		asrt.Equal("2", list.Requests[0].UserUuid)
		asrt.Equal("22", list.Requests[0].UserIp)
		asrt.Equal("dsa", list.Requests[0].Url)
		asrt.Equal("dad", list.Requests[0].Data)
		asrt.NotNil(list.Requests[0].CreatedAt)
	}

	_, err = srv.Db.Exec("DELETE FROM request")
	asrt.NoError(err)
}
