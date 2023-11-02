package repo

import (
	"errors"
	"graphql/models"

	"gorm.io/gorm"
)

type Repo struct {
	Db *gorm.DB
}
type UserRepo interface {
	CreateUser(nu models.User) (models.User, error)
	CreateCpmny(nu models.Company) (models.Company, error)
	FetchCompanyById(cid string) (models.Company, error)
	CreateJobs(nj models.Job) (models.Job, error)
	FetchJobById(jid string) (models.Job, error)
	FetchAllJob() ([]models.Job, error)
	FetchAllCompany() ([]models.Company, error)
}

func NewRepo(g *gorm.DB) (UserRepo, error) {
	if g == nil {
		return nil, errors.New("db cannot be nil")

	}
	return &Repo{
		Db: g,
	}, nil
}
