package publishservice

import (
	"fmt"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"simple-douyin/dao/publish"
	"simple-douyin/model"
	"simple-douyin/util"
	"time"
)

type PublishActionFlow struct {
	pubActionDao *publishdao.PubActionDao
	video        *model.Video
	data         *multipart.FileHeader
	curTime      string
}

func Publish(data *multipart.FileHeader, title string, userId uint32) error {
	return newPublishActionFlow(data, title, userId).Do()
}

func newPublishActionFlow(data *multipart.FileHeader, title string, userId uint32) *PublishActionFlow {
	return &PublishActionFlow{
		pubActionDao: publishdao.NewPubActionDaoInstance(),
		video:        &model.Video{Title: title, UserId: userId},
		data:         data,
		curTime:      time.Now().Format("20060102_150405"),
	}
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
	/* ---------- Update by xhy ----------*/
	// 创建需要上传文件的句柄
	src, err := p.data.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 生成文件名和playUrl
	ext := filepath.Ext(p.data.Filename)
	// 文件名uid_YYYYMMDD_hhmmss.(ext)
	finalName := fmt.Sprintf("%d_%v%v", p.video.UserId, p.curTime, ext)
	// 目录名douyin/video/
	p.video.PlayUrl = "http://180.76.52.150/douyin/video/" + finalName

	// 上传到服务器
	return util.UploadFile(src, finalName, "video")

	/* ---------- Author: xhy ---------- */
	/* ---------- upload to local ---------- */
	// src, err := p.data.Open()
	// if err != nil {
	// 	return err
	// }
	// defer src.Close()
	// // 待修改
	// // 1.路径问题
	// //
	// //
	// // 1.路径检查
	// if _, err := util.MkDir("./public/video/"); err != nil {
	// 	return err
	// }
	// // 2.进行copy操作
	// filename := filepath.Base(p.data.Filename)
	// finalName := fmt.Sprintf("%d_%s", p.video.UserId, filename)
	// p.video.PlayUrl = filepath.Join("./public/video/", finalName)
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

	/* ---------- Author: ny ---------- */
	/* ---------- remote cover ---------- */
	/* ---------- Update by xhy ----------*/
	// 2.执行ffmpeg命令

	// 文件名uid_YYYYMMDD_hhmmss.jpg
	finalName := fmt.Sprintf("%d_%v.jpg", p.video.UserId, p.curTime)
	// 目录名 douyin/cover/
	p.video.CoverUrl = "http://180.76.52.150/douyin/cover/" + finalName
	tmpCoverPath := filepath.Join("./public/", finalName)

	cmd := exec.Command("ffmpeg", "-i", p.video.PlayUrl,
		"-y", "-ss", "00:00:00", "-frames", "1", tmpCoverPath)
	if err := cmd.Run(); err != nil {
		return err
	}
	src, err := os.Open(tmpCoverPath)
	if err != nil {
		return err
	}
	err = util.UploadFile(src, finalName, "cover")
	// 记得close否则会因为占用无法删除
	src.Close()
	os.Remove(tmpCoverPath)
	return err

	/* ---------- Author: xhy ---------- */
	/* ---------- local cover ---------- */
	// 1.路径检查
	// if _, err := util.MkDir("./public/cover/"); err != nil {
	// 	return err
	// }
	// 2.执行ffmpeg命令
	// filename := strings.Split(filepath.Base(p.data.Filename), ".")[0]
	// finalName := fmt.Sprintf("%d_%s_cover.jpg", p.video.UserId, filename)
	// p.video.CoverUrl = filepath.Join("./public/cover/", finalName)
	// cmd := exec.Command("ffmpeg", "-i", p.video.PlayUrl,
	// 	"-y", "-ss", "00:00:00", "-frames", "1", p.video.CoverUrl)
	// if err := cmd.Run(); err != nil {
	// 	return err
	// }
	// return nil
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
