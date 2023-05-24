package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var sshUser, sshPrivateKeyPath, sshServer, localScriptPath, sshPassword string
	flag.StringVar(&sshUser, "u", "root", "远程服务器的用户名.")
	flag.StringVar(&sshPrivateKeyPath, "k", "", "* SSH私钥文件的路径.")
	flag.StringVar(&sshPassword, "p", "", "* 远程服务器的密码.")
	flag.StringVar(&sshServer, "s", "127.0.0.1:22", "远程服务器的IP地址或主机名加端口.")
	flag.StringVar(&localScriptPath, "f", "", "* 本地脚本文件的路径.")
	flag.Parse()

	// 判断SSh密码还是私钥
	sshConfig := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// 判断SSh密码还是私钥
	if (sshPassword != "" || sshPrivateKeyPath != "") && localScriptPath != "" {
		if sshPassword != "" {
			passwordAuth := ssh.Password(sshPassword)
			sshConfig.Auth = append(sshConfig.Auth, passwordAuth)
		} else {
			privateKeyAuth := publicKeyFile(sshPrivateKeyPath)
			sshConfig.Auth = append(sshConfig.Auth, privateKeyAuth)
		}
	} else {
		fmt.Println("-f 本地脚本文件的路径不存在.")
		fmt.Println("-p 远程服务器的密码和 -k SSH私钥文件的路径不能同时为空.")
		os.Exit(0)
	}

	// 读取本地脚本文件内容
	scriptBytes, err := ioutil.ReadFile(localScriptPath)
	if err != nil {
		log.Fatal(err)
	}

	// 连接远程服务器
	sshClient, err := ssh.Dial("tcp", sshServer, sshConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer sshClient.Close()

	// 创建会话
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// 执行远程脚本
	output, err := session.CombinedOutput(string(scriptBytes))
	if err != nil {
		log.Fatal(err)
	}

	// 输出结果
	fmt.Println(string(output))
}

// 从私钥文件中获取公钥
func publicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		log.Fatal(err)
	}

	return ssh.PublicKeys(key)
}
