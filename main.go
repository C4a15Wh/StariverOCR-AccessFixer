package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	simpchinese "golang.org/x/text/encoding/simplifiedchinese"

	"main.go/common"
)

func main() {
	common.Logger(0, "欢迎使用在线OCR访问修复工具，本工具将引导您修复大部分在线OCR无法访问、速度慢的问题。")
	common.Logger(0, "在开始修复前，请先关闭您的计算机上正在运行的杀毒软件以保证修复工作正常进行。")

	SystemTime := time.Now()
	common.Logger(0, "当前系统时间戳: "+fmt.Sprint(SystemTime.Unix()))

	common.Logger(0, "正在解析本地环境信息...")

	StageStatus := make([]bool, 5)
	Response, err := common.HttpGet("https://api.stariver.org/ip/")

	if err != nil {
		common.Logger(2, "连接到星河服务器...失败")
		common.Logger(2, err.Error())
		StageStatus[0] = false
	} else {
		StageStatus[0] = true
	}

	if StageStatus[0] {
		common.Logger(0, "互联网接入地址为: "+string(Response))
	}

	Response, err = common.HttpGet("http://whois.pconline.com.cn/ip.jsp")
	if err != nil {
		common.Logger(2, "获取网络区域...失败")
		common.Logger(2, err.Error())
		StageStatus[1] = false
	} else {
		StageStatus[1] = true
		b, _ := simpchinese.GBK.NewDecoder().Bytes(Response)
		common.Logger(0, "您当前位于: "+string(b))
	}

	common.Logger(0, "正在修改Hosts...")
	SystemEnv := os.Getenv("windir")
	if len(SystemEnv) <= 2 {
		SystemEnv = "C:\\Windows" // 兜底
		common.Logger(1, "获取环境变量'windir'失败，将使用默认值"+SystemEnv)
	}
	common.Logger(0, "Windows Dir: "+os.Getenv("windir"))

	HostsFile, err := ioutil.ReadFile(SystemEnv + "\\System32\\drivers\\etc\\hosts")
	if err != nil {
		StageStatus[2] = false
		common.Logger(2, "无法读取Hosts文件，请以管理员权限运行本程序并关闭杀毒软件。")
		common.Logger(2, err.Error())
		common.Logger(2, "修复失败!")
		time.Sleep(200000)
		return
	}

	HostsString := string(HostsFile)
	HostsString += "\n59.56.100.180  access.c4a15wh.cn"

	err = ioutil.WriteFile(SystemEnv+"\\System32\\drivers\\etc\\hosts", []byte(HostsString), 0777)
	if err != nil {
		common.Logger(2, "无法写入Hosts文件，请以管理员权限运行本程序并关闭杀毒软件。")
		common.Logger(2, err.Error())
		common.Logger(2, "修复失败!")
		time.Sleep(200000)
		return
	}

	common.Logger(0, "写入Hosts文件成功!")
	common.Logger(0, "正在刷新系统DNS...")
	exec.Command("ipconfig", "/flushdns")
	common.Logger(0, "刷新完成! ")
	exec.Command("ping", "access.c4a15wh.cn")
	common.Logger(0, "修复完成，后续如有问题请联系QQ: 3039504176或胖次团子!")
	time.Sleep(200000)
	return
}
