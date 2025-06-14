package repository

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dsx137/anipm/anipm-server/pkg/app/config"
	"github.com/dsx137/anipm/anipm-server/pkg/app/entity"
	"github.com/dsx137/anipm/anipm-server/pkg/util"
	"github.com/sirupsen/logrus"
)

type RepositoryProject struct {
	BaseDir string
}

func NewRepositoryProject(baseDir string) *RepositoryProject {
	return &RepositoryProject{BaseDir: baseDir}
}

func (repo *RepositoryProject) Save(project *entity.EntityProject) error {
	projectPath := filepath.Join(config.BaseDir, project.Name)
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return err
	}
	succeeded := false
	defer func() {
		if !succeeded {
			os.Remove(projectPath)
		}
	}()

	err := util.WriteFileWithJson(filepath.Join(projectPath, config.TrailFile), project)
	if err != nil {
		return err
	}
	succeeded = true
	return nil
}

func (repo *RepositoryProject) ExistByName(name string) bool {
	path := filepath.Join(repo.BaseDir, name)

	stat, err := os.Stat(path)
	if err != nil && !stat.IsDir() {
		return false
	}

	initJsonPath := filepath.Join(config.BaseDir, stat.Name(), config.TrailFile)
	_, err = os.Stat(initJsonPath)

	return err == nil
}

func (repo *RepositoryProject) FindByName(name string) (*entity.EntityProject, error) {
	if repo.ExistByName(name) {
		return &entity.EntityProject{Name: name}, nil
	} else {
		return nil, fmt.Errorf("project not exist")
	}
}

func (repo *RepositoryProject) FindAll() ([]*entity.EntityProject, error) {
	projects := []*entity.EntityProject{}
	entries, err := os.ReadDir(config.BaseDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read project directory list: %s", err)
	}

	for _, entry := range entries {
		func() {
			project, err := repo.FindByName(entry.Name())
			if err != nil {
				logrus.Errorf("Failed to find project: %s", entry.Name())
				return
			}
			projects = append(projects, project)
			// f, err := os.Open(initJsonPath)
			// if err != nil {
			// 	if err := os.Remove(filepath.Join(config.BaseDir, entry.Name())); err != nil && !os.IsNotExist(err) {
			// 		logrus.Errorf("Failed to remove project directory: %s", entry.Name())
			// 	}
			// 	return
			// }
			// defer f.Close()

			// var proj pojo.RequestCreateProject
			// dec := json.NewDecoder(f)
			// if err := dec.Decode(&proj); err != nil {
			// 	return
			// }
		}()
	}

	return projects, nil
}
