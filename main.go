package main

import (
	"github.com/ZJGSU-ACM/GoOnlineJudge/controller"
	"github.com/ZJGSU-ACM/GoOnlineJudge/controller/admin"
	"github.com/ZJGSU-ACM/GoOnlineJudge/controller/contest"
	_ "github.com/ZJGSU-ACM/GoOnlineJudge/schedule"
	"github.com/ZJGSU-ACM/restweb"
	"log"
)

func main() {

	restweb.RegisterController(&controller.ContestController{})
	restweb.RegisterController(&controller.RanklistController{})
	restweb.RegisterController(&controller.ProblemController{})
	restweb.RegisterController(&controller.SessController{})
	restweb.RegisterController(&controller.StatusController{})
	restweb.RegisterController(&controller.NewsController{})
	restweb.RegisterController(&controller.FAQController{})
	restweb.RegisterController(&controller.OSCController{})
	restweb.RegisterController(&controller.HomeController{})
	restweb.RegisterController(&controller.UserController{})
	restweb.RegisterController(&admin.AdminNotice{})
	restweb.RegisterController(&admin.AdminNews{})
	restweb.RegisterController(&admin.AdminTestdata{})
	restweb.RegisterController(&admin.AdminRejudge{})
	restweb.RegisterController(&admin.AdminContest{})
	restweb.RegisterController(&admin.AdminProblem{})
	restweb.RegisterController(&admin.AdminUser{})
	restweb.RegisterController(&admin.AdminHome{})
	restweb.RegisterController(&admin.AdminImage{})
	restweb.RegisterController(&contest.Contest{})
	restweb.RegisterController(&contest.ContestRanklist{})
	restweb.RegisterController(&contest.ContestProblem{})
	restweb.RegisterController(&contest.ContestStatus{})

	restweb.AddFile("/static/", ".")
	log.Fatal(restweb.Run())
}
