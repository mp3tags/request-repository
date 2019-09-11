package service

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mp3tags/request-repository-proto"
	"log"
)

func (s *Service) CreateRequest(ctx context.Context, params *request_repository.CreateRequestParams) (*empty.Empty, error) {
	err := s.createRequest(params)
	if err != nil {
		log.Println(err)
	}
	return &empty.Empty{}, err
}

func (s *Service) createRequest(params *request_repository.CreateRequestParams) error {
	insForm, err := s.Db.Prepare("INSERT INTO request (created_at, user_uuid, user_ip, url, data) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}

	createdAt, err := ptypes.Timestamp(params.CreatedAt)
	if err != nil {
		return err
	}

	_, err = insForm.Exec(createdAt, params.UserUuid, params.UserIp, params.Url, params.Data)
	if err != nil {
		return err
	}

	return nil
}
