package main

import (
	"bluebell/config"
	"bluebell/global"
	"bluebell/pkg/snowflake"
	"bluebell/pkg/validator"
	"bluebell/router"
	"fmt"
	"go.uber.org/zap"
)

//-----代码分层介绍
//controller： 服务的入口，负责处理路由，参数校验，请求转发，返回结果
//service：服务层，负责处理业务逻辑
//models：负责数据与存储相关功能，就是数据库查询和插入这种操作 都放在models里， models里定义的结构体是是请求req 和 返回rsp的结构体

//-----错误返回介绍
//1.controller 层返回json格式和接受service层 传过来的错误 并记录日志
//2.service层接受models层的错误，并继续往上一层返回 到controller层
//3.models层 定义错误信息，并return.

func main() {
	//1.加载配置
	if err := config.Init(); err != nil {
		fmt.Printf("init config failed err: %v\n", err)
		return
	}

	//2.初始化日志
	if err := global.InitLogger(config.Conf, config.Conf.App.Mode); err != nil {
		fmt.Printf("init logger failed err: %v\n", err)
		return
	}
	zap.L().Debug("init logger success...")
	defer zap.L().Sync()

	//3.初始化MySQL链接
	err := global.InitMysql(config.Conf.DataBase)
	if err != nil {
		zap.L().Error("初始化mysql失败 %v", zap.Error(err))
		return
	}

	defer func() {
		if err = global.MysqlClose(); err != nil {
			zap.L().Error("%v", zap.Error(err))
			return
		}
	}()
	zap.L().Debug("mysql init success")

	//4. 初始化Redis链接
	//if err := redis.Init(); err != nil {
	//	fmt.Printf("init redis failed err: %v\n", err)
	//	return
	//}

	//4. 雪花算法
	if err = snowflake.Init(config.Conf.StartTime, config.Conf.MachineID); err != nil {
		zap.L().Error("初始化snowflake 失败 %v", zap.Error(err))
	}

	//5.校验器 中文翻译
	if err = validator.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	//6.初始化路由
	r := router.InitRouter()
	r.Run(fmt.Sprintf(":%s", config.Conf.App.Port))
}
