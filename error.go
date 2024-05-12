// Copyright (C) slackor. 2024-present.
//
// Created at 2024-05-12, by liasica

package slackor

import (
	"fmt"
	"os"
)

func Exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
