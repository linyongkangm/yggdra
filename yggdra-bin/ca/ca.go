package ca

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"net"
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

func GenCertificate() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Country:            []string{"CN"},
		Province:           []string{"BeiJing"},
		Organization:       []string{"yggdra"},
		OrganizationalUnit: []string{"certYggdra"},
		CommonName:         "yggdra",
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certOut, _ := os.Create(config.ROOT_CA_CRT)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	keyOut, _ := os.Create(config.ROOT_CA_KEY)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}

func GetTLSConfig() (cert tls.Certificate, err error) {
	rawCert, _ := ioutil.ReadFile(config.ROOT_CA_CRT)
	rawKey, _ := ioutil.ReadFile(config.ROOT_CA_KEY)
	return tls.X509KeyPair(rawCert, rawKey)
}
