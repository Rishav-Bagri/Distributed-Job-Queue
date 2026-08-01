package router

import (
	"net/http"

	"github.com/rishavbagri/go-job-queue/internal/handler"
	"github.com/rishavbagri/go-job-queue/internal/queue"
)

func New() *http.ServeMux{
	mux:=http.NewServeMux()

	client := queue.NewRedisClient()

	producer := queue.NewProducer(client)

	pHandler:= handler.NewJobHandler(producer)

	mux.HandleFunc("/",handler.Home)
	mux.HandleFunc("/job",pHandler.CreateJob)

	return mux
}