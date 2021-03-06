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
	"log"
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

	// 组装body
	//body := strings.Replace(html, "{{.Title}}", data.Head, -1)
	//body = strings.Replace(html,"{{.Body}}",data.Body,-1)

	m := gomail.NewMessage()
	m.SetHeader("From", "Email Center"+"<"+email.User+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", split...)                            //发送给多个用户
	m.SetHeader("Subject", data.Head)                      //设置邮件主题
	m.SetBody("text/html", data.Body)                      //设置邮件正文
	//m.SetBody("text/html", body)                      //设置邮件正文

	log.Println(email)
	return nil

	d := gomail.NewDialer(email.Host, email.Port, email.User, email.Password)

	err := d.DialAndSend(m)
	if err != nil {
		log.Println(email)
	}
	return err
}

func getEmail() config.Email {
	le := len(config.MyConfig.Email)
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

var html = `
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>

<head>
    <meta charset="UTF-8">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>Email Center</title>
    <link href="https://fonts.googleapis.com/css?family=Nunito:300,400,600,700&display=swap" rel="stylesheet">
    <style type="text/css">
        html {
            -webkit-text-size-adjust: none;
            -ms-text-size-adjust: none;
        }

        @media only screen and (min-device-width: 750px) {
            .table750 {
                width: 750px !important;
            }
        }

        @media only screen and (max-device-width: 750px), only screen and (max-width: 750px) {
            table[class="table750"] {
                width: 100% !important;
            }

            .mob_b {
                width: 93% !important;
                max-width: 93% !important;
                min-width: 93% !important;
            }

            .mob_b1 {
                width: 100% !important;
                max-width: 100% !important;
                min-width: 100% !important;
            }

            .mob_left {
                text-align: left !important;
            }

            .mob_soc {
                width: 50% !important;
                max-width: 50% !important;
                min-width: 50% !important;
            }

            .mob_menu {
                width: 50% !important;
                max-width: 50% !important;
                min-width: 50% !important;
                box-shadow: inset -1px -1px 0 0 rgba(255, 255, 255, 0.2);
            }

            .mob_center {
                text-align: center !important;
            }

            .top_pad {
                height: 15px !important;
                max-height: 15px !important;
                min-height: 15px !important;
            }

            .mob_pad {
                width: 15px !important;
                max-width: 15px !important;
                min-width: 15px !important;
            }

            .mob_div {
                display: block !important;
            }
        }

        @media only screen and (max-device-width: 550px), only screen and (max-width: 550px) {
            .mod_div {
                display: block !important;
            }
        }

        .table750 {
            width: 750px;
        }
    </style>
</head>

<body style="margin: 0; padding: 0;">

<table cellpadding="0" cellspacing="0" border="0" width="100%"
       style="background: #f3f3f3; min-width: 350px; font-size: 1px; line-height: normal;">
    <tr>
        <td align="center" valign="top">
            <!--[if (gte mso 9)|(IE)]>
            <table border="0" cellspacing="0" cellpadding="0">
                <tr>
                    <td align="center" valign="top" width="750"><![endif]-->
            <table cellpadding="0" cellspacing="0" border="0" width="750" class="table750"
                   style="width: 100%; max-width: 750px; min-width: 350px; background: #f3f3f3;">
                <tr>
                    <td class="mob_pad" width="25" style="width: 25px; max-width: 25px; min-width: 25px;">&nbsp;</td>
                    <td align="center" valign="top" style="background: #ffffff;">

                        <table cellpadding="0" cellspacing="0" border="0" width="100%"
                               style="width: 100% !important; min-width: 100%; max-width: 100%; background: #f3f3f3;">
                            <tr>
                                <td align="right" valign="top">
                                    <div class="top_pad" style="height: 25px; line-height: 25px; font-size: 23px;">
                                        &nbsp;
                                    </div>
                                </td>
                            </tr>
                        </table>

                        <table cellpadding="0" cellspacing="0" border="0" width="88%"
                               style="width: 88% !important; min-width: 88%; max-width: 88%;">
                            <tr>
                                <td align="left" valign="top">
                                    <div style="height: 39px; line-height: 39px; font-size: 37px;">&nbsp;</div>
                                    <font class="mob_title1" face="'Source Sans Pro', sans-serif" color="#1a1a1a"
                                          style="font-size: 52px; line-height: 55px; font-weight: 300; letter-spacing: -1.5px;">
                                        <a href="https://www.500ml.club" style="text-decoration:none"><span
                                                    class="mob_title1"
                                                    style="font-family: 'Nunito', Arial, Tahoma, Geneva, sans-serif; color: #6777ef; font-size: 48px; line-height: 55px; font-weight: 700; letter-spacing: -1.5px;">DK Email</span></a>
                                    </font>
                                    <div style="height: 73px; line-height: 73px; font-size: 71px;">&nbsp;</div>
                                </td>
                            </tr>
                        </table>

                        <table cellpadding="0" cellspacing="0" border="0" width="88%"
                               style="width: 88% !important; min-width: 88%; max-width: 88%;">
                            <tr>
                                <td align="left" valign="top">
                                    <font face="'Nunito', sans-serif" color="#1a1a1a"
                                          style="font-size: 52px; line-height: 60px; font-weight: 300; letter-spacing: -1.5px;">
                                        <span style="font-family: 'Nunito', Arial, Tahoma, Geneva, sans-serif; color: #1a1a1a; font-size: 52px; line-height: 60px; font-weight: 300; letter-spacing: -1.5px;">{{.Title}}</span>
                                    </font>
                                    <div style="height: 33px; line-height: 33px; font-size: 31px;">&nbsp;</div>
                                    <font face="'Nunito', sans-serif" color="#585858"
                                          style="font-size: 24px; line-height: 32px;">
                                        <span style="font-family: 'Nunito', Arial, Tahoma, Geneva, sans-serif; color: #585858; font-size: 24px; line-height: 32px;">{{.Body}}</span>
                                    </font>
                                    <div style="height: 75px; line-height: 75px; font-size: 73px;">&nbsp;</div>
                                </td>
                            </tr>
                        </table>

                        <table cellpadding="0" cellspacing="0" border="0" width="100%"
                               style="width: 100% !important; min-width: 100%; max-width: 100%; background: #f3f3f3;">
                            <tr>
                                <td align="center" valign="top">
                                    <div style="height: 34px; line-height: 34px; font-size: 32px;">&nbsp;</div>
                                    <table cellpadding="0" cellspacing="0" border="0" width="88%"
                                           style="width: 88% !important; min-width: 88%; max-width: 88%;">
                                        <tr>
                                            <td align="center" valign="top">
                                                <div style="height:12px; line-height: 34px; font-size: 32px;">&nbsp;
                                                </div>
                                                <font face="'Nunito', sans-serif" color="#868686"
                                                      style="font-size: 17px; line-height: 20px;">
                                                    <span style="font-family: 'Nunito', Arial, Tahoma, Geneva, sans-serif; color: #868686; font-size: 17px; line-height: 20px;">由 DollarKiller Email center 发送</span>
                                                </font>
                                                <div style="height: 3px; line-height: 3px; font-size: 1px;">&nbsp;</div>
                                                <font face="'Nunito', sans-serif" color="#1a1a1a"
                                                      style="font-size: 17px; line-height: 20px;">
                                                    <span style="font-family: 'Nunito', Arial, Tahoma, Geneva, sans-serif; color: #1a1a1a; font-size: 17px; line-height: 20px;"><a
                                                                href="https://github.com/dollarkillerx" target="_blank"
                                                                style="font-family: 'Nunito', Arial, Tahoma, Geneva, sans-serif; color: #1a1a1a; font-size: 17px; line-height: 20px; text-decoration: none;">作者地址</a> &nbsp;&nbsp;|&nbsp;&nbsp; <a
                                                                href="https://github.com/dollarkillerx/email-center" target="_blank"
                                                                style="font-family: 'Nunito', Arial, Tahoma, Geneva, sans-serif; color: #1a1a1a; font-size: 17px; line-height: 20px; text-decoration: none;">项目地址</a></span>
                                                </font>
                                                <div style="height: 35px; line-height: 35px; font-size: 33px;">&nbsp;
                                                </div>
                                            </td>
                                        </tr>
                                    </table>
                                </td>
                            </tr>
                        </table>

                    </td>
                    <td class="mob_pad" width="25" style="width: 25px; max-width: 25px; min-width: 25px;">&nbsp;</td>
                </tr>
            </table>
            <!--[if (gte mso 9)|(IE)]>
            </td></tr>
            </table><![endif]-->
        </td>
    </tr>
</table>
</body>

</html>
`
