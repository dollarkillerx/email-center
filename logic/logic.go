/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:52 2019-09-25
 */
package logic

import (
	"email-center/config"
	"email-center/defs"
	"gopkg.in/gomail.v2"
	"strings"
	"sync"
)

var (
	num = 0
	mu  sync.Mutex
	cs  sync.Mutex
)

func Logic(data *defs.SeedEmail) error {
	email := getEmail()
	if data.Email[0:1] == "," {
		data.Email = data.Email[1:]
	} else if data.Email[len(data.Email)-1:] == "," {
		data.Email = data.Email[:len(data.Email)-1]
	}
	split := strings.Split(data.Email, ",")

	m := gomail.NewMessage()
	m.SetHeader("From", "Email Center"+"<"+email.User+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", split...)                            //发送给多个用户
	m.SetHeader("Subject", data.Head)                      //设置邮件主题
	m.SetBody("text/html", data.Body)                      //设置邮件正文

	d := gomail.NewDialer(email.Host, email.Port, email.User, email.Password)

	err := d.DialAndSend(m)
	return err
}

func getEmail() config.Email {
	le := len(config.MyConfig.Email) - 1
	i := 0
	mu.Lock()
	if num < le {
		i = num
		num += 1
	} else {
		num = 0
	}
	mu.Unlock()

	cs.Lock()
	email := config.MyConfig.Email[i]
	cs.Unlock()
	return email
}
