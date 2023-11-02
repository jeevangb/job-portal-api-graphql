package service

import (
	"errors"
	"fmt"
	"graphql/graph/model"
	"graphql/models"
	"strconv"
)

func (s *Service) CreateCompany(input model.NewCompany) (*model.Company, error) {
	cmpnyDetails := models.Company{
		Name:     input.Name,
		Location: input.Location,
	}
	cmpnyDetails, err := s.R.CreateCpmny(cmpnyDetails)
	if err != nil {
		return nil, errors.New("company creation failed")

	}
	cid := strconv.FormatUint(uint64(cmpnyDetails.ID), 10)
	return &model.Company{
		ID:       cid,
		Name:     cmpnyDetails.Name,
		Location: cmpnyDetails.Location,
	}, nil

}

func (s *Service) ViewCompanyByID(cid string) (*model.Company, error) {

	cmpnyDetails, err := s.R.FetchCompanyById(cid)
	if err != nil {
		return nil, errors.New("company data not found")

	}
	return &model.Company{
		ID:       cid,
		Name:     cmpnyDetails.Name,
		Location: cmpnyDetails.Location,
	}, nil

}

func (s *Service) ViewAllCompany() ([]*model.Company, error) {

	allcmpnydata, err := s.R.FetchAllCompany()
	if err != nil {
		fmt.Println("companies data are not found")
		return nil, err

	}

	cmpnyDetails := []*model.Company{}

	for _, value := range allcmpnydata {
		cmpnyData := model.Company{
			ID:       strconv.FormatUint(uint64(value.ID), 10),
			Name:     value.Name,
			Location: value.Location,
		}
		cmpnyDetails = append(cmpnyDetails, &cmpnyData)

	}
	return cmpnyDetails, nil

}
