# 🍔 餐饮服务后端

## 💡 项目简介

实现了类似于美团的一些功能，如搜索餐厅，搜索菜谱等api

## 🗼 代码架构

### 项目结构

<details>
<summary>展开查看</summary>
<pre>
<code>
    ├── app ----------------------------- (项目文件)
    	├── api ------------------------- (api层)
    		├── recipe ------------------ (关于食谱的api)
    		├── restaurant -------------- (关于餐厅的api)
    		├── user -------------------- (关于用户的api)
    	├── global ---------------------- (全局组件)
    	├── internal -------------------- (内部包)
    		├── middleware -------------- (中间件)
    		├── model ------------------- (模型)
    		├── service ----------------- (服务层)
    	├── router ---------------------- (路由层)
    ├── boot ---------------------------- (项目启动包)
    ├── manifest ------------------------ (交付清单)
    	├── config ---------------------- (项目配置)
		├── sql ------------------------- (sql文件)
			├── mongodb ----------------- (mongodb数据集)
			├── mysql ------------------- (mysql表结构)
    ├── utils --------------------------- (工具包)
    ├── build.sh ------------------------ (项目启动shell脚本)
    ├── docker-compose.yml -------------- (docker容器)
</code>
</pre>
</details>

### 技术栈

<img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="15%">

- [gin](https://github.com/gin-gonic/gin)

> `Gin`是一个用Go语言编写的web框架。它是一个类似于`martini`但拥有更好性能的API框架, 由于使用了`httprouter`，速度提高了近40倍。 如果你是性能和高效的追求者, 你会爱上`Gin`。

[Gin框架介绍及使用-李文周的博客](https://www.liwenzhou.com/posts/Go/gin/)

[视频教程](https://www.bilibili.com/video/BV1gJ411p7xC/)

<img src="http://jwt.io/img/logo-asset.svg">

- jwt

> SON Web Token (JWT)是一个开放标准(RFC 7519)，它定义了一种紧凑的、自包含的方式，用于作为JSON对象在各方之间安全地传输信息。该信息可以被验证和信任，因为它是数字签名的。

[五分钟带你了解啥是JWT](https://zhuanlan.zhihu.com/p/86937325)

[JSON Web Token 入门教程](https://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html)

[jwt.io](https://jwt.io/) 可以在这个网站校验 jwt

<img src="https://miro.medium.com/max/700/1*YcVBLTidq861sJhIlVby5w.png" width="50%">

- [zap](https://github.com/uber-go/zap)

> `zap`是`Uber`开发的非常快的、结构化的，分日志级别的Go日志库。根据Uber-go Zap的文档，它的性能比类似的结构化日志包更好，也比标准库更快。具体的性能测试可以去`github`上看到。

[使用zap接收gin框架默认的日志并配置日志归档](https://www.liwenzhou.com/posts/Go/use_zap_in_gin/)

[深入浅析golang zap 日志库使用（含文件切割、分级别存储和全局使用等）](https://www.yisu.com/zixun/154695.html)

<img src="https://github.com/spf13/viper/raw/master/.github/logo.png?raw=true" width="50%">

- [viper](https://github.com/spf13/viper)

> Viper是适用于Go应用程序的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。

[Go语言配置管理神器——Viper中文教程](https://www.liwenzhou.com/posts/Go/viper_tutorial/)

<img src="https://upload.wikimedia.org/wikipedia/zh/thumb/6/62/MySQL.svg/1200px-MySQL.svg.png" width="30%">

- [mysql](https://www.mysql.com/)

> 一个关系型数据库管理系统，由瑞典MySQL AB 公司开发，属于 Oracle 旗下产品。MySQL 是最流行的关系型数据库管理系统关系型数据库管理系统之一，在 WEB 应用方面，MySQL是最好的 RDBMS (Relational Database Management System，关系数据库管理系统) 应用软件之一

[Go操作MySQL](https://www.liwenzhou.com/posts/Go/go_mysql/)

[sqlx库使用指南](https://www.liwenzhou.com/posts/Go/sqlx/)

[GORM入门指南](https://www.liwenzhou.com/posts/Go/gorm/)

[GORM中文文档](https://gorm.io/zh_CN/docs/)

<img src="https://upload.wikimedia.org/wikipedia/en/thumb/6/6b/Redis_Logo.svg/1200px-Redis_Logo.svg.png" width="40%">

- [redis](https://redis.io/)

> 一个开源的、使用C语言编写的、支持网络交互的、可基于内存也可持久化的Key-Value数据库

[Go语言操作Redis](https://www.liwenzhou.com/posts/Go/redis/)

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/00/Mongodb.png/1200px-Mongodb.png" width="40%">

- [MongoDB](https://www.mongodb.com/)

> 文档型数据库，有兴趣的可以自己去了解

[Go语言操作mongoDB](https://www.liwenzhou.com/posts/Go/go_mongodb/)

<img src="https://developers.redhat.com/sites/default/files/styles/article_feature/public/blog/2014/05/homepage-docker-logo.png?itok=zx0e-vcP" width="30%">

- [docker](https://www.docker.com/)

> Google 公司推出的 Go 语言 进行开发实现，基于 Linux 内核的 cgroup，namespace，以及 AUFS 类的 Union FS 等技术的一个容器服务

​	容器用docker-compose部署

## 🚀 功能模块

### API文档

[apipost](https://console-docs.apipost.cn/preview/75cf02d1bc40f846/9a5d4e23d3fc1ea7)
