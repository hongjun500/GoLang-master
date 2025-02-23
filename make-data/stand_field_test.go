package main

import (
	"strings"
	"testing"
)

func TestBuild(t *testing.T) {
	str := "按钮 "

	trimSpace := strings.TrimSpace(str)
	t.Log(trimSpace)

	str = " 南"
	space := strings.TrimSpace(str)
	t.Log(space)
}
