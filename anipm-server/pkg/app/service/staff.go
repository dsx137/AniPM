package service

import "github.com/dsx137/anipm/anipm-server/pkg/app/repository"

type ServiceStaff struct {
	*repository.RepositoryStaff
}

func NewServiceStaff(repo *repository.RepositoryStaff) *ServiceStaff {
	return &ServiceStaff{RepositoryStaff: repo}
}
