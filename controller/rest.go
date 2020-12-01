/**@author Mehrdad Karami**/

package controller

import (
	"fmt"
	k8mapper "github.com/metao1/kubernetes-mapper/service"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// This is the root directory of files
var base = "results"

type File struct {
	Name string `uri:"name" binding:"required"`
}

func StartServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	files := router.Group("/files")
	logger := log.New(os.Stderr, "visualizer", 0)
	counter := 0
	files.GET("/:name", func(c *gin.Context) {
		var f File
		if err := c.ShouldBindUri(&f); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		if checkFileIsExist(f.Name) {
			m, cn, err := download(f.Name)
			fmt.Printf("Requested for download file %s successfull\n", m)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err})
				return
			}
			c.Header("Content-Disposition", "attachment; filename="+f.Name)
			c.Data(http.StatusOK, "image/png", cn)
		} else {
			logger.Println("file not found!")
			c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		}
	})

	files.POST("/virtualize", func(c *gin.Context) {
		fmt.Print("creating plot requested.\n")
		var f interface{}
		if err := c.ShouldBindUri(&f); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		namespace := c.Request.FormValue("namespace")
		outputFileName := c.Request.FormValue("output_file_name")
		fileType := c.Request.FormValue("file_type")
		if namespace != "" && fileType != "" && outputFileName != "" {
			fmt.Printf("creating plot initiating started for %s, %s, %s %d times.\n", namespace, fileType, outputFileName, counter)
			if counter == 0 {
				counter = k8mapper.Initialize(namespace)
			}
			fmt.Print("creating plot initiated for namespace%.\n", namespace)
			err := k8mapper.CreatePlot(namespace, outputFileName, fileType)
			fmt.Print("creating plot process done.\n")
			if err == nil {
				fmt.Printf("plot %s created successfully.\n", namespace)
				c.JSON(http.StatusBadRequest, gin.H{"successful": namespace})
				return
			}
		}
		fmt.Printf("unsuccessful creating plot.\n")
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameters missed"})
	})

	fmt.Printf("Listen and Serve in https://127.0.0.1:8080\n")
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
