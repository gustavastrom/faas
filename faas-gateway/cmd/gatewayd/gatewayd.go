package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io/ioutil"
	"net/http"
)

func dynamicGateway(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	path := r.URL.Path
	pathToName := make(map[string]string)
	for _, container := range containers {
		//fmt.Fprint(w, "ID: " + container.ID, " Image: " + container.Image, " Name: " + container.Labels["faas.name"], "\n")
		pathToName["/faas/" + container.Labels["faas.name"]] = container.Labels["faas.name"]
	}

	_, nameExists := pathToName[r.URL.Path]
	if !nameExists {
		fmt.Fprintf(w, "Error: The function " + path + " doesn't exist")
		return
	}

	value, ok := r.URL.Query()["p"]
	if !ok || len(value[0]) < 1 {
		fmt.Fprintf(w, "Error: param p is missing")
		return
	}

	p := value[0]
	name := pathToName[path]
	resp, err := http.Get("http://" + name + "?p=" + p)
	if err != nil {
		fmt.Fprintf(w, "Error: Can't establish connection to function " + name)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}

func main() {
	http.HandleFunc("/", dynamicGateway)
	http.ListenAndServe(":80", nil)
}