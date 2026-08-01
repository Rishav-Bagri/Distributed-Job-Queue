package main

import (
	"log"
	"net/http"

	"github.com/rishavbagri/go-job-queue/internal/router"
)




func main() {
	r:=router.New()
	log.Fatal(http.ListenAndServe(":8000",r))
}
