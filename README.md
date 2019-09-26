# email-center
email center  接入多个email邮箱 防止发送邮件被封  tk可注册免费域名 mail.ru可申请企业邮箱

# 如何使用
初始使用填写配置文件
``` 
email: # email 集合
  - user: "b@a.com"
    password: "password"
    host: "smtp.mail.ru"
    port: 465
  - user: "a@a.com"
    password: "password"
    host: "smtp.mail.ru"
    port: 465
```

### Api
post: http://127.0.0.1/email/seed

参数:
``` 
email string  // 发送目标 (单个邮件: a@a.com)(多个邮件: a@a.com,b@b.com,c@c.com)
head  string  // 发送头
body  string  // 发送体
key   string  // 秘钥 每次都要携带 
```
200 正确返回:
``` 
map[string]interface{}{
    "code":200,
    "msg":"发送成功",
}
```
400 参数错误:
``` 
map[string]interface{}{
    "code":400,
    "msg":"参数错误",
}
```
401 权限错误:
``` 
map[string]interface{}{
    "code":401,
    "msg":"权限错误",
}
```
500 发送失败:
``` 
map[string]interface{}{
    "code":500,
    "msg":"发送失败",
    "data":err.Error(),
}
```
