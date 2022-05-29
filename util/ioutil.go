package util

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"mime/multipart"
	"os"
	"time"
)

// MkDir 判断并创建目录
// path 路径
func MkDir(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsExist(err) {
		return true, nil
	}
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return false, err
	}
	return true, nil
}

// UploadFile 上传文件到服务器
// src 需要上传文件的句柄指针
// filename 上传后的文件名
// path 上传的相对目录, cover为封面, video为视频
func UploadFile(src multipart.File, filename string, path string) error {
	// ssh连接配置
	sshConfig := &ssh.ClientConfig{
		User: "qxy",
		Auth: []ssh.AuthMethod{
			ssh.Password("qxy100@"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		ClientVersion:   "",
		Timeout:         10 * time.Second,
	}

	// 建立与ssh服务器的连接
	sshClient, err := ssh.Dial("tcp", "180.76.52.150:22", sshConfig)
	if err != nil {
		return err
	}
	defer sshClient.Close()

	// 建立sftp连接
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return err
	}
	defer sftpClient.Close()

	// 建立远端数据流
	remoteFile, err := sftpClient.Create(sftp.Join("/var/www/html/douyin", path, filename))
	if err != nil {
		return err
	}
	defer remoteFile.Close()

	// 本地文件流拷贝到上传文件流
	_, err = io.Copy(remoteFile, src)

	return err
}
