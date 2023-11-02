package repo

import (
	"errors"
	"graphql/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateUser(nu models.User) (models.User, error) {

	err := r.Db.Create(&nu).Error
	if err != nil {
		return models.User{}, err
	}
	return nu, nil

}
func (r *Repo) CreateCpmny(nc models.Company) (models.Company, error) {

	err := r.Db.Create(&nc).Error
	if err != nil {
		return models.Company{}, err
	}
	return nc, nil

}

func (r *Repo) FetchCompanyById(cid string) (models.Company, error) {
	var companyData models.Company
	result := r.Db.Where("id=?", cid).First(&companyData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("company not found")
	}
	return companyData, nil

}
