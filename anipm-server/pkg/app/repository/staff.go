package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dsx137/anipm/anipm-server/pkg/app/config"
	"github.com/dsx137/anipm/anipm-server/pkg/app/entity"
	"github.com/dsx137/anipm/anipm-server/pkg/util"
)

type RepositoryStaff struct {
	BaseDir string
}

func NewRepositoryStaff(baseDir string) *RepositoryStaff { return &RepositoryStaff{BaseDir: baseDir} }

func (repo *RepositoryStaff) FindAll() ([]*entity.EntityStaff, error) {
	staffs := []*entity.EntityStaff{}

	jsonPath := filepath.Join(config.BaseDir, config.StaffFile)
	fi, err := os.Open(jsonPath)
	if os.IsNotExist(err) {
		_, err = os.Create(jsonPath)
		if err != nil {
			return nil, fmt.Errorf("failed to create staff file: %w", err)
		}
		err = util.WriteFileWithJson(jsonPath, staffs)
		if err != nil {
			return nil, fmt.Errorf("failed to write staff file: %w", err)
		}
		return staffs, nil
	} else if err != nil {
		return nil, err
	}
	defer fi.Close()

	err = json.NewDecoder(fi).Decode(&staffs)
	if err != nil {
		return nil, err
	}
	return staffs, nil
}

func (repo *RepositoryStaff) FindByName(name string) (*entity.EntityStaff, error) {
	staffs, err := repo.FindAll()
	if err != nil {
		return nil, err
	}

	for _, staff := range staffs {
		if staff.Name == name {
			return staff, nil
		}
	}
	return nil, nil
}

func (repo *RepositoryStaff) Save(staff *entity.EntityStaff) error {
	staffs, err := repo.FindAll()
	if err != nil {
		return err
	}
	for _, s := range staffs {
		if s.Name == staff.Name {
			return nil
		}
	}

	return util.WriteFileWithJson(filepath.Join(config.StaffFile), staffs)
}

func (repo *RepositoryStaff) Delete(name string) error {
	staffs, err := repo.FindAll()
	if err != nil {
		return err
	}
	for i, s := range staffs {
		if s.Name == name {
			staffs = append(staffs[:i], staffs[i+1:]...)
		}
	}
	return util.WriteFileWithJson(filepath.Join(repo.BaseDir, config.StaffFile), staffs)
}
