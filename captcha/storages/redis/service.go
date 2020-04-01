// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package redis

import (
	"github.com/go-redis/redis"
	"github.com/orivil/service"
	redis2 "github.com/orivil/services/memory/redis"
)

type Service struct {
	redisService service.Provider
}

func (s *Service) New(ctn *service.Container) (value interface{}, err error) {
	client := ctn.MustGet(&s.redisService).(*redis.Client)
	return NewStorage(client), nil
}

func NewService(redisService *redis2.Service) *Service {
	return &Service{redisService: redisService}
}