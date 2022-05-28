package publishservice

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"simple-douyin/dao/publish"
	"simple-douyin/model"
	"strings"
	"time"
)

type PublishActionFlow struct {
	pubActionDao *publishdao.PubActionDao
	video        *model.Video
	data         *multipart.FileHeader
}

func Publish(data *multipart.FileHeader, title string, userId uint32) error {
	return newPublishActionFlow(data, title, userId).Do()
}

func newPublishActionFlow(data *multipart.FileHeader, title string, userId uint32) *PublishActionFlow {
	return &PublishActionFlow{
		pubActionDao: publishdao.NewPubActionDaoInstance(),
		video:        &model.Video{Title: title, UserId: userId},
		data:         data,
	}
}

// uploadFile 上传文件到服务器
// src 需要上传文件的句柄指针
// filename 上传后的文件名
// path 上传的相对目录, cover为封面, video为视频
func uploadFile(src multipart.File, filename string, path string) error {
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

func (p *PublishActionFlow) Do() error {
	// 保存文件
	if err := p.saveUploadedFile(); err != nil {
		return err
	}
	// 生成封面信息
	if err := p.saveCover(); err != nil {
		return err
	}
	// 组装创建videoInfo
	if err := p.CreateVideo(); err != nil {
		return err
	}
	return nil
}

func (p *PublishActionFlow) saveUploadedFile() error {
	/* ---------- Author: ny  ---------- */
	/* ---------- upload to remote  ---------- */
	// 创建需要上传文件的句柄
	src, err := p.data.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 生成文件名和playUrl
	filename := filepath.Base(p.data.Filename)
	finalName := fmt.Sprintf("%d_%s", p.video.UserId, filename)
	p.video.PlayUrl = "http://180.76.52.150/douyin/video/" + finalName

	// 上传到服务器
	return uploadFile(src, finalName, "video")

	/* ---------- Author: xhy ---------- */
	/* ---------- upload to local ---------- */
	// src, err := p.data.Open()
	// if err != nil {
	// 	return err
	// }
	// defer src.Close()
	// 待修改
	// 1.路径问题
	//

	// 1.路径检查
	// if _, err := util.MkDir("./public/video/"); err != nil {
	// 	return err
	// }
	// 2.进行copy操作
	// filename := filepath.Base(p.data.Filename)
	// finalName := fmt.Sprintf("%d_%s", p.video.UserId, filename)
	// p.video.PlayUrl = filepath.Join("./public/video/", finalName)
	// p.video.PlayUrl = "http://180.76.52.150/douyin/video/" + finalName
	// out, err := os.Create(p.video.PlayUrl)
	// if err != nil {
	// 	return err
	// }
	// defer out.Close()
	// _, err = io.Copy(out, src)
	// return err
}

func (p *PublishActionFlow) saveCover() error {
	// 待修改
	// 1.路径问题
	// 2.ffmpeg的绝对路径

	// 1.路径检查
	// if _, err := util.MkDir("./public/cover/"); err != nil {
	// 	return err
	// }
	// 2.执行ffmpeg命令
	filename := strings.Split(filepath.Base(p.data.Filename), ".")[0]
	finalName := fmt.Sprintf("%d_%s_cover.jpg", p.video.UserId, filename)
	p.video.CoverUrl = "http://180.76.52.150/douyin/cover/" + finalName

	tmpCoverPath := filepath.Join("./public", finalName)
	cmd := exec.Command("ffmpeg", "-i", p.video.PlayUrl,
		"-y", "-ss", "00:00:00", "-frames", "1", tmpCoverPath)
	if err := cmd.Run(); err != nil {
		return err
	}
	src, err := os.Open(tmpCoverPath)
	if err != nil {
		return err
	}
	err = uploadFile(src, finalName, "cover")
	os.Remove(tmpCoverPath)
	return err
}

func (p *PublishActionFlow) CreateVideo() error {
	return p.pubActionDao.CreateVideo(p.video)
}

// func (p *PublishActionFlow) rollBackFile() {
//
// }
//
// func (p *PublishActionFlow) rollBackCover() {
//
// }
