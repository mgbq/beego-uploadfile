package controllers

import (
	"fmt"
)

type FileUploadController struct {
	AuthController
}

//文件下载方法
func (c *FileUploadController) Download() {
	filename := c.GetString("filename")
	fmt.Println("filename6666", filename)

	c.Ctx.Output.Download("./upload/demo.txt")
}

func (c *FileUploadController) Upload() {
	f, h, _ := c.GetFile("myfile")   //获取上传的文件
	path := "./upload/" + h.Filename //文件目录
	f.Close()                        //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	c.SaveToFile("myfile", path)     //存文件
	c.Redirect("/index/index", 302)  //上传成功跳转首页
}

func (c *FileUploadController) Index() {
	// c.display("file_upload")
	c.display("file_index")
}
