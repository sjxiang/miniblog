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