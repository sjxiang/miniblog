package smtp

import (
	"fmt"
	"math/rand"
	"time"
	// "net/smtp"
)

/*

发送验证码
校验

redis 测一下不就完了，查重不久完了


gopkg.in/gomail.v2



*/


var (
	USAGE_SIGNUP = "signup"
	USAGE_FORGET = "forgetpwd"
)


type SMTPServer struct {
	Secret string
}


/*

	邮件服务器

*/

func NewSMTPServer() SMTPServer {
	return SMTPServer{
		
	}
}


type Email struct {

} 

func (e *Email) SendVerificationEmail(email, randomNum, usage string) error {
	return nil
}



func (s *SMTPServer) NewVerificationCode(email, usage string) (string, error) {
	

	return "", nil

	
}


func (s *SMTPServer) ValidateVerificationCode(codeToken, vCode, email, usage string) (bool, error) {
	
	// 验证码失效
	// 验证码错误

	return true, nil
}




// RandStringRunes 生成长度为 length 随机字符串
func RandStringRunes(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")  

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	
	return string(b)
}

// RandNum 生成随机数
func RandNum() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

/*

登录页面
	注册 或登录
	更改密码

注册页面




登录 login
	username 
	password


老外
		first name
		last name
		username 
		邮箱
		password
		

		confirm 链接


		cookie





邮箱
手机号 验证码


中国 
用户名
手机号
邮箱


注册 register signup


发送邮箱

修改密码
	邮箱发连链接


package services

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/k3a/html2text"
	"github.com/vanng822/go-premailer/premailer"
	"gopkg.in/gomail.v2"
	"html/template"
	"leizhenpeng/go-email-verification/initialize"
	"os"
	"path/filepath"
)

type EmailData struct {
	URL      string
	UserName string
	Subject  string
}

func GenEmailVerifyURL(info string) string {
	configNow := initialize.GetConfig()
	return fmt.Sprintf("%s/api/verify_email?info=%s", configNow.BaseUrl, info)
}

func GenEmailData(email string, info string) *EmailData {
	return &EmailData{
		URL:      GenEmailVerifyURL(info),
		UserName: email,
		Subject:  "请激活您的账号",
	}
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(email string, data *EmailData) error {

	// Sender data.
	configNow := initialize.GetConfig()
	from := configNow.EmailFrom
	smtpPass := configNow.SmtpPass
	smtpUser := configNow.SmtpUser
	to := email
	smtpHost := configNow.SmtpHost
	smtpPort := configNow.SmtpPort

	var body bytes.Buffer

	template, err := ParseTemplateDir("templates")
	if err != nil {
		return errors.New("could not parse template")
	}

	template.ExecuteTemplate(&body, "email-verify.html", &data)
	htmlString := body.String()
	prem, _ := premailer.NewPremailerFromString(htmlString, nil)
	htmlInline, err := prem.Transform()
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", htmlInline)
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return errors.New("could not send email")
	}
	return nil

}

*/