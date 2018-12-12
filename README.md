# Baas 系统

本系统是一个基于HTTP协议的restful的API网关，可灵活自由的操作任意**mysql**数据库，作为后端数据API服务。

## 解决问题

几乎每次做一个APP或者应用，后端都需要给前端开发一套restful API接口，在不使用第三方平台的情况下，目前最快的方式就是通过一些语言框架生成CURD接口，但生成代码，需要进行二次加工才可以使用，操作复杂，也容易出错，出现问题，相对于新手来说，可能还不如他写代码来的快。



这里为了解决这个问题，开发出一套程序，数据表自动支持restful API操作。 为了保证大家对速度的最求，这里使用了Golang作为开发语言。



## 使用场景

1. 提供APP开发、小程序开发一切的数据自动API服务。




# 需要开发的操作
## 各类服务

- [ ] 权限验证
- [ ] 日志服务
- [ ] 项目管理（一个项目一个库，第二期）



## 数据库操作

- [ ] 增加
- [ ] 修改
- [ ] 条件修改
- [x] ID删除
- [ ] 条件删除（第二期）
- [x] 查询一条
- [ ] 查询列表
- [ ] 条件查询

## 表结构操作

- [ ] 创建表
- [ ] 删除表
- [ ] 修改表结构

## 索引操作

- [ ] 表索引列表
- [ ] 创建索引
- [ ] 删除索引
- [ ] 修改索引

## 用户表操作

- [ ]  登陆
- [ ] 注册 
- [ ] 找回密码 
- [ ] 第三方登陆 
- [ ] 重置密码



## 客户端SDK

- [ ] js
- [ ] nodejs



## WEBUI

- [ ] UI界面管理数据库（第二期）



## 开发文档

- [ ] 源码部署文档

- [ ] API Server使用说明
- [ ] 目录结构说明



# 目录结构说明

还未修改

```
.
├── gorm.db
├── hello.go
├── common
│   ├── utils.go        //small tools function
│   └── database.go     //DB connect manager
└── users
    ├── models.go       //data models define & DB operation
    ├── serializers.go  //response computing & format
    ├── routers.go      //business logic & router binding
    ├── middlewares.go  //put the before & after logic of handle request
    └── validators.go   //form/json checker
```


