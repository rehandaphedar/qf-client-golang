package main

//go:generate oapi-codegen -config config.yaml -package content -o pkg/content/main.gen.go schemas/content-apis/v4.json
//go:generate oapi-codegen -config config.yaml -package oauth2 -o pkg/oauth2/main.gen.go schemas/oauth2-apis/v1.json
//go:generate oapi-codegen -config config.yaml -package user -o pkg/user/main.gen.go schemas/user-apis/v1.json
