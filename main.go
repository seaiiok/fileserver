package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	hislogs = [10]string{}
)

type fileinfo struct {
	Name      string `json:"name,omitempty"`
	Time      string `json:"time,omitempty"`
	Timestrap int64
}

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("templates/*")
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	os.MkdirAll("./uploads", 0644)

	r.GET("/file", func(c *gin.Context) {
		items := []*fileinfo{}
		if files, err := os.ReadDir("./uploads"); err == nil {
			for _, f := range files {
				finfo, _ := f.Info()
				items = append(items, &fileinfo{
					Name:      f.Name(),
					Time:      finfo.ModTime().Format(time.DateTime),
					Timestrap: finfo.ModTime().Unix(),
				})
			}
		}
		sort.Sort(fileinfos(items))
		c.HTML(http.StatusOK, "/templates/file.tmpl", gin.H{
			"files": items,
		})
	})
	r.GET("/download/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		b, err := os.ReadFile(filepath.Join("./uploads", filename))
		if err != nil {
			c.AbortWithStatus(http.StatusNotExtended)
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", "application/text/plain")
		c.Header("Accept-Length", fmt.Sprintf("%d", len(b)))
		c.Writer.Write(b)
		addHislog("download", filename)
	})

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/templates/upload.tmpl", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}
		files := form.File["files"]
		filenames := []string{}
		for _, file := range files {
			filename := filepath.Join("./uploads", file.Filename)
			filenames = append(filenames, file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
				return
			}
		}

		c.Redirect(http.StatusMovedPermanently, "/file")
		addHislog("upload", filenames...)
	})
	r.GET("/logs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/templates/logs.tmpl", gin.H{
			"logs": hislogs,
		})
	})

	r.Run(":9000")
}

func addHislog(action string, filenames ...string) {
	for i := len(hislogs) - 1; i >= 0; i-- {
		if i == 0 {
			hislogs[i] = fmt.Sprintf("%s %v %s", time.Now().Format(time.DateTime), filenames, action)
			continue
		}

		hislogs[i] = hislogs[i-1]
	}
}

type fileinfos []*fileinfo

func (x fileinfos) Len() int           { return len(x) }
func (x fileinfos) Less(i, j int) bool { return x[i].Timestrap > x[j].Timestrap }
func (x fileinfos) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// loadTemplate 加载由 go-assets-builder 嵌入的模板
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
