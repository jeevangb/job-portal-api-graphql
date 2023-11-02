package service

import (
	"errors"
	"fmt"
	"graphql/graph/model"
	"graphql/models"
	"strconv"
)

func (s *Service) CreateJob(input model.NewJob) (*model.Job, error) {
	cid, err := strconv.ParseUint(input.Cid, 10, 64)
	if err != nil {
		fmt.Println("please provicde id in proper format")
		return &model.Job{}, err
	}
	jobDetails := models.Job{
		Title: input.Title,
		Desc:  input.Desc,
		Cid:   uint(cid),
	}
	jobData, err := s.R.CreateJobs(jobDetails)
	if err != nil {
		fmt.Println("error at create job table")
		return &model.Job{}, nil
	}
	jid := strconv.FormatUint(uint64(jobData.ID), 10)
	companyId := strconv.FormatUint(uint64(jobDetails.Cid), 10)
	return &model.Job{
		ID:    jid,
		Title: jobData.Title,
		Desc:  jobData.Desc,
		Cid:   companyId,
	}, nil

}

func (s *Service) ViewJobByID(cid string) (*model.Job, error) {

	jobDetails, err := s.R.FetchJobById(cid)
	if err != nil {
		return nil, errors.New("job data not found")

	}
	return &model.Job{
		ID:    cid,
		Title: jobDetails.Title,
		Desc:  jobDetails.Desc,
	}, nil

}
func (s *Service) ViewAllJob() ([]*model.Job, error) {

	alljobdata, err := s.R.FetchAllJob()
	if err != nil {
		return nil, errors.New("job data not found")

	}

	jobDetails := []*model.Job{}

	for _, value := range alljobdata {
		jobData := model.Job{
			ID:    strconv.FormatUint(uint64(value.ID), 10),
			Title: value.Title,
			Desc:  value.Desc,
			Cid:   strconv.FormatUint(uint64(value.Cid), 10),
		}
		jobDetails = append(jobDetails, &jobData)

	}
	return jobDetails, nil

}
