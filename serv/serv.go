package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"
)

func main() {
	// Используем свободные порты
	httpPort := ":8082"  // Новый HTTP порт
	httpsPort := ":9443" // HTTPS порт

	// Генерация сертификата
	cert, err := generateSelfSignedCert()
	if err != nil {
		log.Fatal("Ошибка генерации сертификата:", err)
	}

	// Обработчик для всех запросов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Server work on port: %s\n", httpsPort)
		fmt.Fprintf(w, "Server time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Fprintf(w, "Your IP: %s\n", r.RemoteAddr)
		fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
	})

	// Запускаем HTTP сервер для редиректа
	go func() {
		log.Printf("HTTP сервер запущен на %s", httpPort)
		http.ListenAndServe(httpPort, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			target := "https://" + r.Host + httpsPort + r.URL.String()
			http.Redirect(w, r, target, http.StatusMovedPermanently)
		}))
	}()

	// Настраиваем HTTPS сервер
	server := &http.Server{
		Addr: httpsPort,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	log.Printf("HTTPS сервер запущен на %s", httpsPort)
	log.Printf("Откройте в браузере: https://localhost%s", httpsPort)
	log.Fatal(server.ListenAndServeTLS("", ""))
}

func generateSelfSignedCert() (tls.Certificate, error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return tls.Certificate{}, err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Local Test"},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:    []string{"localhost"},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return tls.Certificate{}, err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	})
	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	})

	return tls.X509KeyPair(certPEM, keyPEM)
}
