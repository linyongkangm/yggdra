package ca

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"os"
	"time"
	"yggdra/config"
)

// 生成根证书签名
func GenRootCertificateSign() (rootCsr *x509.Certificate) {
	rootCsr = &x509.Certificate{
		Version:      3,
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			Country:            []string{"CN"},
			Province:           []string{"GuangZhou"},
			Locality:           []string{"GuangZhou"},
			Organization:       []string{"Yggdra"},
			OrganizationalUnit: []string{"YggdraProxy"},
			CommonName:         "Yggdra Root CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            1,
		MaxPathLenZero:        false,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}
	return
}

// 生成私钥
func GenPrivateKey() (key *rsa.PrivateKey) {
	key, _ = rsa.GenerateKey(rand.Reader, 2048)
	return
}

// 生成根证书
func GenRootCertificate() {
	rootCsr := GenRootCertificateSign()
	rootKey := GenPrivateKey()
	rootDer, _ := x509.CreateCertificate(rand.Reader, rootCsr, rootCsr, rootKey.Public(), rootKey)
	//4.将得到的证书放入pem.Block结构体中
	block := pem.Block{
		Type:    "CERTIFICATE",
		Headers: nil,
		Bytes:   rootDer,
	}
	crtFile, _ := os.Create(config.ROOT_CA_CRT)
	defer crtFile.Close()
	pem.Encode(crtFile, &block)

	//6.将私钥中的密钥对放入pem.Block结构体中
	block = pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(rootKey),
	}
	keyFile, _ := os.Create(config.ROOT_CA_KEY)
	defer keyFile.Close()
	pem.Encode(keyFile, &block)
}

func GenmMitMCertificateSign(commonName string) (csr *x509.Certificate) {
	csr = &x509.Certificate{
		Version:      3,
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			Country:            []string{"CN"},
			Province:           []string{"GuangZhou"},
			Locality:           []string{"GuangZhou"},
			Organization:       []string{"Yggdra"},
			OrganizationalUnit: []string{"YggdraProxy"},
			CommonName:         commonName,
		},
		DNSNames:              []string{commonName},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		BasicConstraintsValid: true,
		IsCA:                  false,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	return
}

func GenMitMCertificate(commonName string) {
	rootCsrBytes, _ := ioutil.ReadFile(config.ROOT_CA_CRT)
	rootCsr, _ := x509.ParseCertificate(rootCsrBytes)
	rootKeyBytes, _ := ioutil.ReadFile(config.ROOT_CA_KEY)
	rootKey, _ := x509.ParsePKCS1PrivateKey(rootKeyBytes)
	csr := GenmMitMCertificateSign(commonName)
	key := GenPrivateKey()
	der, _ := x509.CreateCertificate(rand.Reader, csr, rootCsr, key.Public(), rootKey)
	//4.将得到的证书放入pem.Block结构体中
	block := pem.Block{
		Type:    "CERTIFICATE",
		Headers: nil,
		Bytes:   der,
	}
	crtFile, _ := os.Create(config.ROOT_CA_CRT)
	defer crtFile.Close()
	pem.Encode(crtFile, &block)

	//6.将私钥中的密钥对放入pem.Block结构体中
	block = pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(key),
	}
	keyFile, _ := os.Create(config.ROOT_CA_KEY)
	defer keyFile.Close()
	pem.Encode(keyFile, &block)
}
