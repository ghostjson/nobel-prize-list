package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

	type Laureate struct{
		Id string `json:"id"`
		Firstname string `json:"firstname"`
		Surname string `json:"surname"`
		Motivation string `json:"motivation"`
		Share string `json:"share"`
	}

	type Prize struct{
		Year string `json:"year"`
		Category string `json:"category"`
		Laureates []Laureate `json:"laureates"`
	}

	type PrizeList struct{
		Prizes []Prize `json:"prizes"`
	}



func main(){



	router := gin.Default()
	router.Static("/assets", "./public")
	router.Static("/images", "./public/images")
	router.Delims("{{{","}}}")
	router.LoadHTMLGlob("public/*.html")

	router.GET("/", uploadPageHandler)
	router.GET("/all-prize-winners", allPrizeWinnersPageHandler)
	router.GET("/recent-winners", recentWinnersPageHandler)

	// router.GET("api/test", testHandler)

	router.POST("api/upload", uploadHandler)

	router.GET("api/read-file/:filename", readFileHandler)

	router.GET("api/save/yaml/:filename/:json_filename", saveAsYAMLHandler)

	if err := router.Run(":5000"); err != nil{
		log.Fatal(err.Error())
	}
}