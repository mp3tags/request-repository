package service

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/mp3tags/request-repository-proto"
	"log"
	"request-repository/models"
)

func (s *Service) ListRequests(ctx context.Context, params *request_repository.ListRequestsParams) (*request_repository.ListRequestsResponse, error) {
	response := new(request_repository.ListRequestsResponse)

	requests, total, err := s.listRequests(params)
	if err != nil {
		log.Println(err)
	}

	response.Total = total
	for _, request := range requests {
		createdAt, err := ptypes.TimestampProto(request.CreatedAt)
		if err != nil {
			return response, err
		}

		response.Requests = append(response.Requests, &request_repository.Request{
			Id:        request.Id,
			UserUuid:  request.UserUuid,
			UserIp:    request.UserIp,
			Url:       request.Url,
			Data:      request.Data,
			CreatedAt: createdAt,
		})
	}
	return response, err
}

func (s *Service) listRequests(params *request_repository.ListRequestsParams) ([]models.Request, int32, error) {
	var requests []models.Request
	var err error
	var total int32

	results, err := s.Db.Query("SELECT id, created_at, user_uuid, user_ip, url, data FROM request ORDER BY id desc LIMIT ? OFFSET ?", params.Limit, params.Offset)
	if err != nil {
		return requests, total, err
	}

	for results.Next() {
		var request models.Request
		// for each row, scan the result into our tag composite object
		err = results.Scan(&request.Id, &request.CreatedAt, &request.UserUuid, &request.UserIp, &request.Url, &request.Data)
		if err != nil {
			return requests, total, err
		}
		requests = append(requests, request)
	}

	err = s.Db.QueryRow("SELECT count(*) FROM request").Scan(&total)
	if err != nil {
		return requests, total, err
	}

	return requests, total, err
}
