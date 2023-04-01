package handler

import (
	"context"
	"encoding/json"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/job"
	"github.com/HaleNing/Ning_JobBoard/src/database"
	"github.com/HaleNing/Ning_JobBoard/src/delegate"
	"github.com/HaleNing/Ning_JobBoard/src/param"
	"github.com/HaleNing/Ning_JobBoard/src/utils"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"time"
)

var jobCtx context.Context

func JobApi(api fiber.Router, ctx context.Context) {
	jobCtx = ctx
	api.Post("/job/create", createNewJobHandler)
	api.Get("/job/getByCompany", getJobsByCompanyNameHandler)
	api.Get("/job/allList", getAllJobList)
	api.Get("/job/getList", getListByLimitHandler)
	api.Get("/job/del", delJobByIdHandler)
	api.Post("/job/update", upJobByIdHandler)
	api.Get("/job/doBusy", busyHandler)

}

func upJobByIdHandler(ctx *fiber.Ctx) error {
	if !utils.CheckLogin(ctx) {
		return fiber.NewError(fiber.StatusBadRequest, "you need login")
	}
	upJob := new(param.JobUpParam)
	if err := ctx.BodyParser(upJob); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "job update param is nil")
	}
	log.Println(upJob) // john
	err := delegate.UpJob(upJob)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "job update failure ")
	}
	return ctx.JSON("update job  success")
}

func delJobByIdHandler(ctx *fiber.Ctx) error {
	if !utils.CheckLogin(ctx) {
		return fiber.NewError(fiber.StatusBadRequest, "you need login")
	}
	id := ctx.Query("id")
	delId, _ := strconv.Atoi(id)
	err := delegate.DelJobById(delId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "del have error")
	}
	return ctx.SendString("ok")
}

func getListByLimitHandler(ctx *fiber.Ctx) error {
	p := new(param.JobSelectParam)
	parseErr := ctx.QueryParser(p)
	if parseErr != nil {
		log.Printf("parseErr: %v", parseErr)
		return fiber.NewError(fiber.StatusBadRequest, "query job list fail,please contact administrator")
	}
	jobs, err := delegate.GetJobListByCondition(p.Name, p.Tech, p.IsRemote > 0, p.Area)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "query job list inside fail,please contact administrator")
	}
	return ctx.Send(jobs)
}

func createNewJobHandler(ctx *fiber.Ctx) error {
	if !utils.CheckLogin(ctx) {
		return fiber.NewError(fiber.StatusBadRequest, "you need login")
	}
	newJob := new(param.JobCreateParam)
	if err := ctx.BodyParser(newJob); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "job create param is nil")
	}
	log.Println(newJob) // john
	_, err := database.DBConn.Job.Create().
		SetCompanyName(newJob.CompanyName).
		SetIsRemote(newJob.IsRemote).
		SetArea(newJob.JobArea).
		SetExp(newJob.JobExp).
		SetJobName(newJob.JobName).
		SetDescription(newJob.JobDesc).
		Save(jobCtx)
	if err != nil {
		log.Printf("insert failure %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "job insert failure ")
	}
	if err != nil {
		log.Printf("err:[%v]", err)
		return fiber.NewError(fiber.StatusBadRequest, "create new job error!")
	} else {
		return ctx.JSON("create new job  success")
	}
}

func getJobsByCompanyNameHandler(ctx *fiber.Ctx) error {
	bn := ctx.Query("company")
	log.Println(jobCtx)
	log.Println(bn)
	allJobs, err := database.DBConn.Job.Query().Where(job.IsExist(true)).Where(job.CompanyNameEQ(bn)).Order(ent.Desc("create_time")).All(context.Background())
	if err != nil {
		log.Printf("getJobsByCompanyNameHandler err:[%v]", err)
		return fiber.NewError(fiber.StatusBadRequest, "get job msg error!")
	}
	log.Println(allJobs)
	res, err := json.Marshal(allJobs)
	log.Println(res)
	return ctx.Send(res)
}

func getAllJobList(ctx *fiber.Ctx) error {
	all, err := database.DBConn.Job.Query().Where(job.IsExist(true)).Limit(100).Order(ent.Desc("create_time")).All(context.Background())
	if err != nil {
		log.Printf("getAllJobList have a error :[%v]", err)
	}
	successRes := param.CreateSuccessRes(all)
	res, _ := json.Marshal(successRes)
	return ctx.Send(res)
}

func busyHandler(ctx *fiber.Ctx) error {
	busyChannel := make(chan string, 1)
	timeoutCtx, cancelFunc := context.WithTimeout(jobCtx, 10*time.Second)
	defer cancelFunc()
	go busy(timeoutCtx, busyChannel)
	for {
		select {
		case <-timeoutCtx.Done():
			log.Printf("Context cancelled: %v\n", timeoutCtx.Err())
			return fiber.NewError(fiber.StatusBadRequest, "do busy error!")
		case result := <-busyChannel:
			log.Printf("channel Received: %s\n", result)
			return ctx.JSON(result)
		default:
		}
	}
}
func busy(ctx context.Context, ch chan string) {
	log.Println(ctx)
	log.Println("do busy")
	time.Sleep(5 * time.Second)
	ch <- "end busy"
}
