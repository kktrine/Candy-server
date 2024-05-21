package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Order struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

func main() {
	k := flag.String("k", "", "Candy type")
	c := flag.Int("c", 0, "Count of candy to buy")
	m := flag.Int("m", 0, "Amount of money")
	port := flag.Int("port", 0, "port")
	flag.Parse()
	order := Order{
		Money:      *m,
		CandyType:  *k,
		CandyCount: *c,
	}
	caCert, err := os.ReadFile("minica.pem")
	if err != nil {
		log.Fatalf("could not read CA certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	clientCert, err := tls.LoadX509KeyPair("candy.tld/cert.pem", "candy.tld/key.pem")
	if err != nil {
		log.Fatalf("could not load client key pair: %s", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      caCertPool,
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("could not marshal order: %s", err)
	}
	addr := "https://candy.tld:" + strconv.Itoa(*port) + "/buy_candy"
	resp, err := client.Post(addr, "application/json", strings.NewReader(string(orderJSON)))
	if err != nil {
		log.Fatalf("could not send request: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("could not read response: %s", err)
	}

	fmt.Println(string(body))
}
