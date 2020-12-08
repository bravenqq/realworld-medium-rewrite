// Package usecases provides ...
package usecases

import "iohttps.com/live/realworld-medium-rewrite/domain"

type Token struct {
	ID     SUUID
	UserID domain.NUUID
}

type SUUID string

type TokenRepository interface {
}

type TokeInteractor struct {
	TokenRepos TokenRepository
}

//Login 判断用户登录
func (itor Token) Login()
