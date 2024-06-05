package img

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
)

func Create(g *gin.Context) (string, error) {
	uid := g.Request.PostFormValue("uid")
	if uid == "" {
		return "", errors.New("uid is null")
	}
	file, header, err := g.Request.FormFile("file")
	if err != nil {
		return "", errors.New("Failed to retrieve file")
	}

	ext := filepath.Ext(header.Filename) //获取文件后缀
	if ext != ".jpg" && ext != ".webp" && ext != ".png" {
		return "", errors.New("文件格式不符合")
	}

	header.Filename = fmt.Sprintf("%v%s", uid, ext)

	// 指定上传目录，比如 "uploads"
	uploadDir := "./public/images"

	// 确保上传目录存在
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", errors.New("Failed to create upload directory")
	}

	// 构建文件的完整路径
	filePath := filepath.Join(uploadDir, header.Filename)

	// 创建文件用于写入
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", errors.New("Failed to create file")
	}
	defer outFile.Close()

	// 复制文件内容
	if _, err := io.Copy(outFile, file); err != nil {
		return "", errors.New("Failed to copy file")
	}
	return "http://localhost:8080/static/images/" + header.Filename, nil
}
