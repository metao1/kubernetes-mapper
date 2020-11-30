package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// This is the root directory of files
var base = "results"

type File struct {
	Name string `uri:"name" binding:"required"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	files := router.Group("/files")
	logger := log.New(os.Stderr, "", 0)

	files.GET("/:name", func(c *gin.Context) {
		var f File
		if err := c.ShouldBindUri(&f); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		if checkFileIsExist(f.Name) {
			m, cn, err := download(f.Name)
			fmt.Print("Requested for download file %s successfull\n", m)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err})
				return
			}
			c.Header("Content-Disposition", "attachment; filename="+f.Name)
			c.Data(http.StatusOK, "image/png", cn)
		} else {
			logger.Println("file not found!")
			c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
			return
		}
	})

	// Listen and Server in https://127.0.0.1:8080
	_ = router.Run(":8080")
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat("results/" + filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func download(n string) (string, []byte, error) {
	dst := fmt.Sprintf("%s/%s", base, n)
	b, err := ioutil.ReadFile(dst)
	if err != nil {
		return "", nil, err
	}
	m := http.DetectContentType(b[:512])

	return m, b, nil
}
