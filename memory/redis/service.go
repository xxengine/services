// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package redis

import (
	"github.com/go-redis/redis"
	"github.com/orivil/service"
	"github.com/orivil/services/cfg"
	"github.com/orivil/xcfg"
)

type Service struct {
	configService service.Provider
	db            int
}

func (s *Service) New(ctn *service.Container) (value interface{}, err error) {
	envs := ctn.MustGet(&s.configService).(xcfg.Env)
	env := &Env{}
	err = envs.UnmarshalSub("redis", env)
	if err != nil {
		panic(err)
	}
	var client *redis.Client
	client, err = env.Init(s.db)
	if err != nil {
		panic(err)
	}
	return client, nil
}

func NewService(configService *cfg.Service, DB int) *Service {
	return &Service{
		configService: configService,
		db:            DB,
	}
}
