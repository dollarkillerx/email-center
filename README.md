# email-center
email center  接入多个email邮箱 防止发送邮件被封  tk可注册免费域名 mail.ru可申请企业邮箱

# 如何使用
初始使用填写配置文件
``` 
email: # email 集合
  - user: "notice@dollarkiller.com"
    password: "%Y4I4qjlqKAy"
    host: "smtp.mail.ru"
    port: 465
```

### Api
post: http://127.0.0.1/email/seed

value:
``` 
email string  // 发送目标 (单个邮件: a@a.com)(多个邮件: a@a.com,b@b.com,c@c.com)
head  string  // 发送头
body  string  // 发送体 
```
