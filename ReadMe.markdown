# 重要提示

1.本项目从起步到开发到部署均有文档以及详细教程

2.本项目需要您有一定的golang redis mysql 基础

3.您完全可以通过我们的教程和文档完成一切操作，因此我们不再提供免费的技术服务

4.如果您将此项目用于商业用途，请遵守Apache2.0协议并保留作者技术支持声明。您需保留如下版权声明信息，其余信息功能不做任何限制。如需剔除请

## 1. 基本介绍

### 1.1 项目介绍

> Gin-hong-api是一个基于  [gin](https://gin-gonic.com) 开发的适合PHP程序员转行的开发基础平台，集成jwt鉴权，动态路由，redis remember，redis tag动态管理，简易redis队列处理，异常多途径报警，任务调度。并提供多种示例文件，让您把更多时间专注在业务开发上。

### 1.2 贡献指南
Hi! 首先感谢你使用 gin-hone-api。

gin-hone-api 是一套为PHP程序员快速转go搭建gin项目准备的一整套架构式的开源框架，旨在快速搭建中小型项目。

gin-hone-api 的成长离不开大家的支持，如果你愿意为 gin-hone-api 贡献代码或提供建议，请阅读以下内容。

#### 1.2.1 Issue 规范
- issue 仅用于提交 Bug 或 Feature 以及设计相关的内容，其它内容可能会被直接关闭。如果你在使用时产生了疑问，请到 Slack 里咨询。

- 在提交 issue 之前，请搜索相关内容是否已被提出。

#### 1.2.2 Pull Request 规范
- 请先 fork 一份到自己的项目下，不要直接在仓库下建分支。

- commit 信息要以`[文件名]: 描述信息` 的形式填写，例如 `README.md: fix xxx bug`。

- 如果是修复 bug，请在 PR 中给出描述信息。

- 合并代码需要两名维护人员参与：一人进行 review 后 approve，另一人再次 review，通过后即可合并。

## 2. 使用说明

```
- golang版本 >= v1.18
- IDE推荐：Goland
```

### 2.1 server项目

使用 `Goland` 等编辑工具，打开server目录，不可以打开 gin-vue-admin 根目录

```bash
# 克隆项目
git clone https://github.com/justfu/gin-hong-api.git
# 进入主文件夹
cd gin-hong-api
# 脚本编辑且运行项目 具体可以自己行在sh里面去加
./install.sh
```

## 3. 技术选型

- 后端：用 [Gin](https://gin-gonic.com/) 快速搭建基础restful风格API，[Gin](https://gin-gonic.com/) 是一个go语言编写的Web框架。
- 数据库：采用`MySql` > (5.7) 版本 数据库引擎 InnoDB，使用 [gorm](http://gorm.cn) 实现对数据库的基本操作。
- 缓存：使用`Redis`
- 配置文件：使用 [configor](github.com/jinzhu/configor) 实现`yaml`格式的配置文件。
- 日志：使用 [zap](https://github.com/uber-go/zap) 实现日志记录。
- 定时任务：使用 [quartz](github.com/reugn/go-quartz/quartz) 实现定时任务调度。

## 4. 项目架构

### 4.1 目录结构

```├─gin-hong-api
      ├─bin 编译后的文件夹
      │  ├─api //web项目编译文件
      │  └─cmd //定时任务编译文件
      ├─cmd
      │  ├─queue //队列自动执行
      │  │  └─job
      ├─common
      │  ├─alarm//错误预警
      │  ├─env //环境配置
      │  ├─function //公用方法
      │  ├─queue //公共队列推送方法
      │  ├─service
      │  └─time
      ├─config//配置文件
      │  └─extra
      ├─controller//控制器
      │  ├─app
      │  └─smallapp
      ├─core //系统核心文件
      │  └─redis
      ├─entity //返回实体类文件
      ├─event //队列事件执行方法
      │  ├─addLog
      │  └─exeWords
      ├─extend //扩展文件
      ├─fonts  //字体文件
      ├─handler //基于gin自定义handler
      ├─imgs //输出图片
      │  └─out
      ├─lib //队列接口
      ├─model //数据库结构体 请求结构体 返回数据结构体
      ├─routers //路由文件
      ├─service //服务
      └─test //基本文件测试
```

## 5. 主要功能

- 文件上传下载：实现基于 `阿里云` 的文件上传操作(请开发自己去各个平台的申请对应 `token` 或者对应`key`)。
- 用户管理：系统管理员分配用户角色和角色权限。
- 实现基于redis的简单队列任务处理(支持可选并发处理任务)
- 实现redis remember
- 实现redis cacheTag统一管理
- 实现异常多途径报警
- 实现指定json结构返回
- 任务调度
- 基于yaml多环境配置
- JWT
- 词云生成
- JIEBA分词
- 更多功能待开发

## 8. 贡献者

感谢您对gin-hong-api的贡献!

## 9. 捐赠

如果你觉得这个项目对你有帮助，你可以请作者喝饮料 

## 10. 商用注意事项

如果您将此项目用于商业用途，请遵守Apache2.0协议并保留作者技术支持声明。

Docker:构建命令  ```docker build -t gin:v1.0 .```
Docker:启动命令: ```docker run -d -p80:9999 --name="gin" gin:v1.0```