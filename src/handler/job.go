package handler

import (
	"context"
	"encoding/json"
	"github.com/HaleNing/bustrack/src/Model/ent/job"
	"github.com/HaleNing/bustrack/src/database"
	"github.com/HaleNing/bustrack/src/param"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

var (
	jobCtx context.Context
)

func JobApi(api fiber.Router, ctx context.Context) {
	jobCtx = ctx
	api.Post("/job/create", createNewJobHandler)
	api.Get("/job/getByCompany", getJobsByCompanyNameHandler)
	api.Get("/job/doBusy", busyHandler)

}

func createNewJobHandler(ctx *fiber.Ctx) error {
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
		return err
	}
	if err != nil {
		log.Printf("err:[%v]", err)
		return fiber.NewError(fiber.StatusBadRequest, "create new job error!")
	} else {
		return ctx.SendString("create new job  success")
	}
}

func updateBusHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
	//param := new(param.BusUpdateParam)
	//if err := ctx.BodyParser(param); err != nil {
	//	log.Println(err)
	//	return fiber.NewError(fiber.StatusBadRequest, "param is nil")
	//}
	//res, err := delegate.UpBusName(busCtx, param)
	//log.Println(res)
	//log.Println(err)
	//if err != nil {
	//	return fiber.NewError(fiber.StatusBadRequest, "update have error")
	//} else {
	//	return ctx.SendString("update bus success")
	//}
}

func getJobsByCompanyNameHandler(ctx *fiber.Ctx) error {
	bn := ctx.Query("company")
	log.Println(jobCtx)
	log.Println(bn)
	allJobs, err := database.DBConn.Job.Query().Where(job.IsExist(true)).Where(job.CompanyNameEQ(bn)).All(jobCtx)
	if err != nil {
		log.Printf("getJobsByCompanyNameHandler err:[%v]", err)
		return fiber.NewError(fiber.StatusBadRequest, "get job msg error!")
	}
	log.Println(allJobs)
	res, err := json.Marshal(allJobs)
	log.Println(res)
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
			return ctx.SendString(result)
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
