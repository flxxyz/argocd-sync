package main

import (
	"log"
	"os"
	"strconv"

	"github.com/safe2008/argocd-actions/pkg/argocd"
)

func main() {
	insecure, error := strconv.ParseBool(os.Getenv("INPUT_INSECURE"))

	if error != nil {
		log.Fatalf("error: %v", error)
	}

	options := argocd.APIOptions{
		Address:  os.Getenv("INPUT_ADDRESS"),
		Token:    os.Getenv("INPUT_TOKEN"),
		Insecure: insecure,
	}

	api := argocd.NewAPI(options)

	err := api.Sync(os.Getenv("INPUT_APPNAME"))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
