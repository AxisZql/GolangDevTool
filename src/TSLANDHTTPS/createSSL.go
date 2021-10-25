package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func CreateSsl() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)//随机生成证书的序列号（生产环境使用的是CA分发的唯一号码）
	subject := pkix.Name{//设置证书的标题
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}
//KeyUsage和ExtKeyUsage字段表示的值表明这个X.509证书实用于进行服务器是否验证操作的
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageDataEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")}, //设置证书只能在设定ip使用
	}
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	cerOut, _ := os.Create("cert.pem") //创建证书文件
	pem.Encode(cerOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	keyOut, _ := os.Create("key.pem") //创建服务器私钥文件
	defer keyOut.Close()
	defer cerOut.Close()
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})

}
