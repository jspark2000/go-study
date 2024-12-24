package router

import (
	"encoding/json"
	"net/http"

	"github.com/jspark2000/go-study/src/judger"
)

type JudgeRequest struct {
	Code string `json:"code"`
}

type JudgeRouter interface {
	HandleJudge()
}

type BasicJudgeRouter struct {
	judger judger.Judger
}

func NewBasicJudgeRouter() BasicJudgeRouter {
	judger := judger.NewBasicJudger()
	return BasicJudgeRouter{judger: &judger}
}

func (b *BasicJudgeRouter) HandleJudge(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var judgeRequest JudgeRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&judgeRequest)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	judgeResult := b.judger.Judge(judgeRequest.Code)

	response, err := json.Marshal(judgeResult)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
