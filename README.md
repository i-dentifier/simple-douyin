# simple-douyin


具体功能内容参考飞书说明文档

## 环境搭建

- ### 编译运行

工程无其他依赖，直接编译运行即可

```shell
go run main.go
```

- ### 客户端配置

进入抖音客户端后双击右下角-"我"可以打开高级设置

配置BaseUrl为 http://x.x.x.x:8080

ip需要填写本机局域网ip，查看方法：

Windows下可以使用
```shell
ipconfig
```
MacOS下可以使用
```shell
ifconfig en0
```
**重启抖音后如果出现熊的视频说明配置成功**

- ### 数据库配置

    - host: 180.76.52.150
    - user: douyin
    - password: douyin100@

本机可以使用cmd连接访问
```shell
mysql -h 180.76.52.150 -u douyin -p
```

## 已实现功能
### 1. 用户
* 用户注册 `/douyin/user/register/`
* 用户登录 `/douyin/user/login/`
* 用户信息查询 `/douyin/user/`

用户在注册和登录后会由服务器颁发token用于鉴权，token有效期为2小时

需要用户登录前置操作的接口都会接入中间件进行token验证，非法token和过期token将无法通过验证
### 2.Publish
* 投稿操作 `/douyin/publish/action/`
* 投稿列表 `/douyin/publish/list/`

用户投稿和获取投稿列表都需要token登录的鉴权，目前还未对接点赞和评论的操作
注：查看视频会闪退


其余接口建议进行分类，例如user相关的接口在controller, service, dao层分别建立user目录，其中所有文件都统一package为usercontroller, userservice, userdao

## 数据表

库名为simple_douyin, 表暂时只设计了用户相关的user, user_auths, relationships, 具体建表信息在`create_tables.sql`中，建议使用mysql8.0及以上版本，登录数据库后使用以下命令可以将库、表一次性导入

```sql
-- path: create.tables.sql文件所在路径
source path/create_tables.sql
```

## 版本控制

https://github.com/i-dentifier/simple-douyin

考虑使用git进行合作，代码已经上传github并为每个人建立了以姓名首字母为名的分支，大家只在自己的分支上工作就可以了。建议大家尽量只添加新文件而不对已有文件内容进行改动，这样可以避免分支merge操作。如需改动可以在群里沟通。