package main

import (
	"lcq/controller"
	"lcq/services"
	"net/http"
	"time"
	"lcq/db"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	db.MysqlInit()
	services.ServicesInit()
	router := controller.InitRouter()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}


func B( ) int { //-1代表没有局部极小值
	temp := []int{1, 2, 4, 5, 3, 6, 7}
	 var result int

	for {
		result = C(temp)
		if result != -1 {
			break
		}
		if temp[len(temp)/2] > temp[len(temp)/2 -1] {
			temp = temp[:len(temp)/2]
		}else {
			temp = temp[len(temp)/2:len(temp)-1]
		}
	}
	return result
}
func C (temp []int) int {

	if temp[1] > temp[0] {
		return temp[0]
	}
	if temp[len(temp)-1] < temp[len(temp)-2] {
		return temp[len(temp)-1]
	}

	if temp[len(temp)/2] < temp[len(temp)/2 -1] && temp[len(temp)/2] > temp[len(temp)/2 + 1] {
		return temp[len(temp)/2]
	}

	return -1
}