package publishservice

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	publishdao "simple-douyin/dao/publish"
	"simple-douyin/model"
	"simple-douyin/util"
	"strings"
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

func (p *PublishActionFlow) Do() error {
	//保存文件
	if err := p.saveUploadedFile(); err != nil {
		return err
	}
	//生成封面信息
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
	src, err := p.data.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//待修改
	//1.路径问题
	//

	//1.路径检查
	if _, err := util.MkDir("./public/video/"); err != nil {
		return err
	}
	//2.进行copy操作
	filename := filepath.Base(p.data.Filename)
	finalName := fmt.Sprintf("%d_%s", p.video.UserId, filename)
	p.video.PlayUrl = filepath.Join(`./public/video/`, finalName)
	out, err := os.Create(p.video.PlayUrl)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}

func (p *PublishActionFlow) saveCover() error {
	//待修改
	//1.路径问题
	//2.ffmpeg的绝对路径

	//1.路径检查
	if _, err := util.MkDir("./public/cover/"); err != nil {
		return err
	}
	//2.执行ffmpeg命令
	filename := strings.Split(filepath.Base(p.data.Filename), ".")[0]
	finalName := fmt.Sprintf("%d_%s_cover.jpg", p.video.UserId, filename)
	p.video.CoverUrl = filepath.Join(`./public/cover/`, finalName)
	//s = "ffmpeg", "-i " + p.videoInfo.PlayUrl,
	//	"-ss", "00:00:00", "-frames:v 1 ", p.videoInfo.CoverUrl
	cmd := exec.Command("ffmpeg", "-i", p.video.PlayUrl,
		"-ss", "00:00:00", "-frames", "1", p.video.CoverUrl)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (p *PublishActionFlow) CreateVideo() error {
	return p.pubActionDao.CreateVideo(p.video)
}

//func (p *PublishActionFlow) rollBackFile() {
//
//}
//
//func (p *PublishActionFlow) rollBackCover() {
//
//}
