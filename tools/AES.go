package tools

import (
    "github.com/astaxie/beego"
    "github.com/DeanChina/wefun.org/logs"
    "crypto/aes"
    "crypto/cipher"
)
var commonIV =[]byte{0x00,0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x09,0x0a,0x0b,0x0c,0x0d,0x0e,0x0f}
/**
AES 加密
**/
func AES(plain string)(string){
    key:=[]byte(beego.AppConfig.String("aeskey"))
    c,err:=aes.NewCipher([]byte(key))
    if nil != err{
    	logs.ErrorLog("aes",err)	
    }
    cfb:=cipher.NewCFBEncrypter(c,commonIV)
    ciphertext:=make([]byte,len(plain))
    cfb.XORKeyStream(ciphertext,[]byte(plain))
    return string(ciphertext)
}

/*
AES 解密
*/
func AESDecrypt(plain string)(string){
	plaintext := []byte(plain)
	key:=[]byte(beego.AppConfig.String("aeskey"))
    c,err:=aes.NewCipher([]byte(key))
    if nil != err{
    	logs.ErrorLog("aes",err)	
    }
    cfbdec:=cipher.NewCFBDecrypter(c,commonIV)
    plaintextCopy := make([]byte,len(plain))
    cfbdec.XORKeyStream(plaintextCopy,plaintext)
    return string(plaintextCopy)
}