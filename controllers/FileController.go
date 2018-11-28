package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"os"
	"fmt"
	"crypto/md5"
	"math/rand"
	"beeblog/models"
)

type FileController struct {
	beego.Controller
}

func (this *FileController) Upload() {
	filename := ""
	for filename, _ = range this.Ctx.Request.MultipartForm.File {}
	f, h, filerr := this.GetFile(filename) //获取上传的文件
	if filerr != nil {
		fmt.Println("err:", filerr)
		this.Data["json"] = models.ReurnError(500, "后缀名不符合上传要求")
		this.ServeJSON()
		return
	}
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		this.Data["json"] = models.ReurnError(500, "后缀名不符合上传要求")
		this.ServeJSON()
		return
	}
	//创建目录
	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
	err := os.MkdirAll(uploadDir, 777)
	if err != nil {
		this.Data["json"] = models.ReurnError(500, "")
		this.ServeJSON()
		return
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ))
	fileName := fmt.Sprintf("%x", hashName) + ext
	fpath := uploadDir + fileName
	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = this.SaveToFile(filename, fpath)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = models.ReurnError(500, "")
		this.ServeJSON()
		return
	}
	this.Data["json"] = models.ReurnData("",fpath)
	this.ServeJSON()
	return
}
