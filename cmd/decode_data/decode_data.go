package main

import (
	"betcorgi/utils"
	"log"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) < 3 {
		log.Println("参数错误")
		log.Println("使用方法,解析下注参数: ./decode_data_win.exe bet 0xxxxxx")
		log.Println("使用方法，解析派奖参数: ./decode_data_win.exe draw 0xxxxxx")
		return
	}
	switch arg[1] {
	case "bet":
		utils.DecodeBetInoutData(arg[2])
	case "draw":
		utils.DecodeDrawInoutData(arg[2])
	default:
		log.Println("当前只支持解码bet和draw参数")
	}

}
