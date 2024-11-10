package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/nanovms/ops/lepton"
	"github.com/nanovms/ops/provider"
)

func rword() string {
	f, err := os.Open("/words")
	if err != nil {
		fmt.Println(err)
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	words := strings.Split(string(bytes), "\n")

	return words[rand.Int()%len(words)]
}

func createWorker() {
	c := lepton.NewConfig()

	name := rword()

	c.CloudConfig.ProjectID = "prod-1033"
	c.CloudConfig.Zone = "us-west2-a"
	c.CloudConfig.ImageName = "worker"
	c.RunConfig.InstanceName = name

	p, err := provider.CloudProvider("gcp", &c.CloudConfig)
	if err != nil {
		fmt.Println(err)
	}

	ctx := lepton.NewContext(c)

	err = p.CreateInstance(ctx)
	if err != nil {
		fmt.Println(err)
	}

}

func createHandler(w http.ResponseWriter, r *http.Request) {
	go func() {
		createWorker()
	}()

	fmt.Fprintf(w, "hello!")
}

type Worker struct {
	Ip      string
	Payload string
}

var workers = []Worker{}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	a, err := json.Marshal(workers)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(a))
}

func getIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func reportHandler(w http.ResponseWriter, r *http.Request) {

	payload := r.URL.Query().Get("payload")
	wrk := Worker{
		Payload: payload,
		Ip:      getIP(r),
	}
	workers = append(workers, wrk)

	fmt.Fprintf(w, "hello!")
}

func main() {
	fmt.Printf("listening on port 8080 ...\n")
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/report", reportHandler)
	http.HandleFunc("/create", createHandler)

	http.ListenAndServe("0.0.0.0:8080", nil)
}
