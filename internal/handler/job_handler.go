package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rishavbagri/go-job-queue/internal/model"
	"github.com/rishavbagri/go-job-queue/internal/queue"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Job Queue API 🚀")
}


type JobHandler struct{
	producer *queue.Producer
}

func NewJobHandler(p *queue.Producer) *JobHandler{
	return &JobHandler{
		producer: p,
	}
}

func (h *JobHandler) CreateJob(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data model.Job

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}



	ctx:= context.Background()

	job:= model.Job{
		ID:uuid.New().String(),
		Type: data.Type,
		Payload: data.Payload,
		Status: "queued",
		Attempts: 0,
		CreatedAt: time.Now(),
	}

	err = h.producer.Enqueue(ctx,job)

	if err != nil {
		http.Error(w, "failed to enqueue job", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	err = json.NewEncoder(w).Encode(map[string]string{
		"jobId": job.ID,
		"status": job.Status,
	})

	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}