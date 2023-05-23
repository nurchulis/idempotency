package user

import (
	"encoding/json"
	"github.com/nurchulis/idempotency/internal/presentations"
	"github.com/nurchulis/idempotency/internal/repositories"
	"github.com/nurchulis/idempotency/internal/ucase/contract"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

type userList struct {
	client   *redis.Client
	userRepo repositories.UserRepo
}

const (
	cacheExpireTime     = 10 * time.Minute
	msgSuccess          = "success"
	msgSuccessWithCache = "success with existing data"
	requireRequestID    = "X-Request-ID is required"
)

func NewUserList(
	client *redis.Client,
	userRepo repositories.UserRepo,
) contract.UseCase {
	return &userList{
		client:   client,
		userRepo: userRepo,
	}
}

func (s *userList) Serve(w http.ResponseWriter, r *http.Request) presentations.Response {
	var (
		requestID = r.Header.Get("X-Request-ID")
		users     []presentations.User
	)

	if requestID == "" {
		return presentations.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: requireRequestID,
		}
	}

	exists, err := s.client.Exists(requestID).Result()
	if err != nil {
		log.Fatal(err)
	}

	if exists == 1 {
		data, err := s.client.Get(requestID).Result()
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal([]byte(data), &users)
		if err != nil {
			log.Fatal(err)
		}

		return presentations.Response{
			Code:    http.StatusOK,
			Message: msgSuccessWithCache,
			Data:    users,
		}
	}

	result, err := s.userRepo.FindAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range result {
		users = append(users, presentations.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	data, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	err = s.client.Set(requestID, string(data), cacheExpireTime).Err()
	if err != nil {
		log.Fatal(err)
	}

	return presentations.Response{
		Code:    http.StatusOK,
		Message: msgSuccess,
		Data:    users,
	}
}
