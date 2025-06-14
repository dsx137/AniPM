package service

import "github.com/dsx137/anipm/anipm-server/pkg/app/repository"

type ServiceProject struct {
	*repository.RepositoryProject
}

func NewServiceProject(repo *repository.RepositoryProject) *ServiceProject {
	return &ServiceProject{RepositoryProject: repo}
}
