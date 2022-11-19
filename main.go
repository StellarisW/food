package main

import "main/boot"

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.MysqlDBSetup()
	boot.MongoDBSetup()
	boot.RedisSetup()
	boot.ServerSetup()
}
