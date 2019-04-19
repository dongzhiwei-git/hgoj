package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/yinrenxin/hgoj/models"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"

	//"strconv"
	_ "github.com/yinrenxin/hgoj/models"
	_ "github.com/yinrenxin/hgoj/routers"
)

func main() {
	initTemplate()
	initSession()
	initStatic()
	initLogs()
	beego.Run()
}


func initLogs() {
	logs.SetLogger(logs.AdapterFile,`{"filename":"logs/app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func initSession() {
	gob.Register(models.Users{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "hgoj"
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./data/session"
}


func initStatic() {
	beego.SetStaticPath("/","static")
	//beego.SetStaticPath("/problem", "static")
}

func initTemplate() {
	_ = beego.AddFuncMap("get_res", func(result int16) (string) {
		//res, _ := strconv.Atoi(result)
		return models.JUDGERES[int(result)]
	})


	_ = beego.AddFuncMap("get_res_class", func(result int16)(string){

		return models.JUDGERESCLSAA[int(result)]
	})

	_ = beego.AddFuncMap("menu_eq", func(a,b string)(string){
		if a == b {
			return "active"
		} else {
			return ""
		}
	})
	_ = beego.AddFuncMap("cal_rate", func(a,b time.Time)(float64){
		startTime := a
		endTime := b
		totalTime := endTime.Sub(startTime).Minutes()
		t := time.Now().Sub(startTime).Minutes()
		percentage := (t/totalTime)*100
		if t > totalTime {
			percentage = 100
		}

		if t < 0 {
			percentage = 0
		}
		return percentage
	})


	_ = beego.AddFuncMap("is_finish", func(a,b time.Time)(bool){
		startTime := a
		endTime := b
		totalTime := endTime.Sub(startTime).Minutes()
		t := time.Now().Sub(startTime).Minutes()
		if t > totalTime {
			return true
		}
		return false
	})


	_ = beego.AddFuncMap("ParseTime", func(temp float64)(string){
		s := int(temp)
		h := s/(60*60)
		s = s - h*60*60
		m := s/60
		s = s - m*60
		th := strconv.Itoa(h)
		tm := strconv.Itoa(m)
		ts := strconv.Itoa(s)
		return th+":"+tm+":"+ts
	})
}