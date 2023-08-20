# meg-backup-gozero
## 项目介绍
****
*此项目使用go-zero框架，go-zero是一个云原生微服务框架，内建级联超时控制、限流、自适应熔断、自适应降级等微服务治理能力，无需配置额外代码；集成了各种工程实践的web和rpc框架。go-zero虽然是一个微服务框架，但是它也支持写单体应用，详细见go-zero官网*
***
*目录树*
```
├─api               // api文件
│  │  ├─preidct
│  │  └─test
│  ├─logic          // logic文件，编写接口逻辑
│  │  ├─admin
│  │  ├─doctor
│  │  ├─download
│  │  ├─preidct
│  │  └─test
│  ├─svc            // 服务发现
│  └─types          // 根据api文件生成
├─models            // 数据库相关，本项目使用postgresql
│  ├─doctor
│  ├─patient
│  └─quota  
└─response          // 返回信息模板
```
*这里并不是全部的目录，我认为用go-zero写单体应用只需要将api文件和etc目录下的配置文件写好，然后使用goctl脚手架，而后补全logic文件就完成70%~80%的工作了*

## goctl相关使用
***
*由于此项目使用的数据库为postgresql，goctl不能像生成mysql相关的model文件那样直接使用sql建表语句生成，而是需要从postgresql数据库中生成*
```api
goctl model pg datasource -url="postgres://postgres:password@address:5432/database?sslmode=disable" -table="table" -c -dir . 
```
*使用goctl生成api服务*
```api
goctl api go -api .\api\main.api -dir .   
```
*运行项目*
```api
go run .\main.go -f .\etc\main-api.yaml 
```
## 配置依赖项
***
```api
go mod tidy
```