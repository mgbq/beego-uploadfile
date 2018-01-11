package controllers

import (
	"fmt"
)

type FileUploadController struct {
	BaseController
}

//文件下载方法
func (this *FileUploadController) Download() {
	filename := this.GetString("filename")
	fmt.Println("filename6666", filename)

	this.Ctx.Output.Download("./upload/demo.txt")
}

func (this *FileUploadController) Upload() {
	f, h, _ := this.GetFile("myfile")  //获取上传的文件
	path := "./upload/" + h.Filename   //文件目录
	f.Close()                          //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	this.SaveToFile("myfile", path)    //存文件
	this.Redirect("/index/index", 302) //上传成功跳转首页
}
