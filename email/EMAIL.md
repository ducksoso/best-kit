
## golang中邮件相关的包

```
net/mail
net/smtp
net/textproto
mime
mime/multipart
mime/quotedprintable

```

### 基于smtp包的邮件发送示例，基于官方文档：

```
// Set up authentication information.
  auth := smtp.PlainAuth("", "xxxx@qq.com", "xxxxx", "smtp.qq.com")

  // Connect to the server, authenticate, set the sender and recipient,
  // and send the email all in one step.
  to := []string{"xxxx@xx.com"} //收件地址
  msg := []byte("To: xxx@xxx.com\r\n" +
    "Subject: discount Gophers!\r\n" +
    "\r\n" +
    "This is the email body.\r\n")
  err := smtp.SendMail("smtp.qq.com:25", auth, "xxx@xx.com", to, msg)
  if err != nil {
    log.Fatal(err)
  }
```

代码解析：

```
func PlainAuth(identity, username, password, host string) Auth

// identity 参数一般为空
// username 一般为发送者邮件地址
// password 通过邮箱账号得到的登录口令，在通过代理发送邮件时候需要。
// host 设置使用的那个平台的邮箱


func SendMail(addr string, a Auth, from string, to []string, msg []byte) error 
// addr 参数是平台地址参数，各大平台都有自己的地址，不同的平台去官网查询即可，该地址参数要带一个端口，默认一般是25
// a 是上面生成的 安全认证(也可以说是一个账户信息，如同使用邮件之前需要登陆邮箱一样)
// from 参数是发送者地址，此参数要保持和auth 中的参数用户名一致
// to 存放一个地址收件邮箱信息。
// msg 是一个消息体，要满足标准协议

```


## 目前用的最多的包是gomail

https://github.com/go-gomail/gomail



