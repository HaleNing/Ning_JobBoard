package service

import (
	"context"
	"encoding/json"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/job"
	"github.com/HaleNing/Ning_JobBoard/src/database"
	"github.com/HaleNing/Ning_JobBoard/src/param"
	"log"
)

func GetJobListByCondition(jobName string, tech string, isRemote bool, area string) ([]byte, error) {
	log.Printf("getListByLimit p: area:[%v]-jobName:[%v]-remote:[%v]-tech:[%v]", area, jobName, isRemote, tech)
	allJobsForLimit, err := database.DBConn.Job.Query().Where(job.IsExist(true)).
		Where(job.IsRemote(isRemote)).
		Where(job.Or(job.JobNameContainsFold(jobName))).
		Where(job.Or(job.CompanyNameContainsFold(tech))).
		Where(job.Or(job.AreaContainsFold(area))).
		Order(ent.Desc("create_time")).All(context.Background())
	if err != nil {
		log.Printf("getListByLimit err:[%v]", err)
		return nil, err
	}
	log.Println(allJobsForLimit)
	res, _ := json.Marshal(allJobsForLimit)
	log.Println(res)
	return res, nil
}

func DelJobById(id int) error {
	log.Printf("DelJobById id:[%v]", id)
	err := database.DBConn.Job.DeleteOneID(id).Exec(context.Background())
	log.Printf("getListByLimit err:[%v]", err)
	return err
}

func UpJob(param *param.JobUpParam) error {
	save, err := database.DBConn.Job.UpdateOneID(param.Id).
		SetCompanyName(param.CompanyName).
		SetIsRemote(param.IsRemote).
		SetArea(param.JobArea).
		SetExp(param.JobExp).
		SetJobName(param.JobName).
		SetDescription(param.JobDesc).
		Save(context.Background())
	log.Printf("res:[%v]", save)
	if err != nil {
		log.Printf("UpJob err:[%v]", err)
	}
	return err
}
