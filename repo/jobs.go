package repo

import (
	"errors"
	"fmt"
	"graphql/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateJobs(nj models.Job) (models.Job, error) {
	err := r.Db.Create(&nj).Error
	if err != nil {
		fmt.Println("failed to create database")
		return models.Job{}, err
	}
	return nj, nil

}
func (r *Repo) FetchJobById(jid string) (models.Job, error) {
	var jobData models.Job
	result := r.Db.Where("id=?", jid).First(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("company not found")
	}
	return jobData, nil

}
func (r *Repo) FetchAllJob() ([]models.Job, error) {
	var alljobData []models.Job
	result := r.Db.Find(&alljobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return []models.Job{}, errors.New("jobs data not found")

	}
	return alljobData, nil
}
func (r *Repo) FetchAllCompany() ([]models.Company, error) {
	var allcmpnyData []models.Company
	result := r.Db.Find(&allcmpnyData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return []models.Company{}, errors.New("company data not found")

	}
	return allcmpnyData, nil
}
