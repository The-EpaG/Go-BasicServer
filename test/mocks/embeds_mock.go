package mocks

import (
	_ "embed"
)

var (
	//go:embed user/user-external.res.json
	UserExternalRes []byte
)
