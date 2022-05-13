# simple-douyin


具体功能内容参考飞书说明文档

### 环境搭建

- #### 编译运行

工程无其他依赖，直接编译运行即可
    
Windows下可以使用
```shell
go run main.go router.go
```

MacOS下可以使用同Windows的命令或者如下命令
```shell
go build && ./simple-douyin
```

- #### 客户端配置
- 
进入抖音客户端后双击右下角-"我"可以打开高级设置

配置BaseUrl为 http://ip:8080

ip需要填写本机局域网ip，查看方法：

Windows下可以使用
```shell
ipconfg
```
MacOS下可以使用
```shell
ifconfig en0
```
**重启抖音后如果出现熊的视频说明配置成功**

### 功能说明

接口功能不完善，仅作为示例

* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可

### 测试数据

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试