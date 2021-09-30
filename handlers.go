package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

// route /
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// route api/test/
func testHandler(c *gin.Context)  {

		type Test struct {
			Name string `yaml:"name"`
		}

		var test = Test{Name: "Hello"}
		testString, err := yaml.Marshal(&test)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error",
			})
		}

		file, err2 := os.Create("data.yaml")
		if err2 != nil {
			log.Fatal(err)
		}

		defer file.Close()

		_, err3 := file.WriteString(string(testString))
		if err3 != nil {
			log.Fatal(err3)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": string(testString),
		})
}

// route api/upload
func uploadHandler(c *gin.Context)  {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error with file upload",
			})
			return
		}

		filename := header.Filename

		out, err := os.Create("storage/" + filename)
		if err != nil {
			log.Fatal(err)
		}

		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil{
			log.Fatal(err)
		}

		var prizesList PrizeList

		content,err := ioutil.ReadFile("storage/" + filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error occurred during file read",
			})
			return
		}

		
		if err2 := json.Unmarshal(content, &prizesList); err2 != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error occurred during parsing",
			})
			return
		}
		

		prizeListYaml, err3 := yaml.Marshal(&prizesList)
		if err3 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error while yaml parsing",
			})
			return
		}

		file1, err4 := os.Create("storage/" + filename +".yaml")
		if err4 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error while creating yaml file"})
			return
		}

		defer file1.Close()

		_, err5 := file1.WriteString(string(prizeListYaml))
		if err5 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error while writing to yaml file",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully finished.",
			"content": prizesList,
		})
}

// route api/read-file/:filename
func readFileHandler(c *gin.Context){
	fileName := c.Param("filename")

	var prizesList PrizeList

	content, err := ioutil.ReadFile("storage/" + fileName)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occurred during file read",
		})
	}

	if strings.HasSuffix(fileName, ".json"){
		if err2 := json.Unmarshal(content, &prizesList); err2 != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error occurred during parsing",
			})
			return
		}
	}else if strings.HasSuffix(fileName, ".yaml"){
		if err3 := yaml.Unmarshal(content, &prizesList); err3 != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error occurred during parsing",
			})
			return
		}
	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid filename. Filename should end with .json or .yaml",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": prizesList,
	})
}


// route /api/save/yaml/:filename/:json_filename

func saveAsYAMLHandler(c *gin.Context)  {
	filename := c.Param("filename")
	jsonFilename := c.Param("json_filename")

	fmt.Println(jsonFilename)
	var prizesList PrizeList

	content,err := ioutil.ReadFile("storage/" + jsonFilename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occurred during file read",
		})
		return
	}

	
	if err2 := json.Unmarshal(content, &prizesList); err2 != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occurred during parsing",
		})
		return
	}
	

	prizeListYaml, err3 := yaml.Marshal(&prizesList)
	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while yaml parsing",
		})
		return
	}

	file1, err4 := os.Create("storage/" + filename)
	if err4 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while creating yaml file"})
		return
	}

	defer file1.Close()

	_, err5 := file1.WriteString(string(prizeListYaml))
	if err5 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while writing to yaml file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully finished.",
		"content": prizesList,
	})
}
