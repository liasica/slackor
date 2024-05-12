// Copyright (C) slackor. 2024-present.
//
// Created at 2024-05-12, by liasica

package slackor

var (
	cfg = &config{}
)

type config struct {
	token string
}

func SetToken(token string) {
	cfg.token = token
}

func GetToken() string {
	return cfg.token
}
