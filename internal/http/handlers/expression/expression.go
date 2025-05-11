package handlers

import (
	"calculator_go/internal/storage"
	"calculator_go/internal/utils/agent/validator"
	"calculator_go/internal/utils/orchestrator/manager"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

// Handlers for operations with expressions

type Request struct {
	Expression string `json:"expression"`
}

type ResponseData struct {
	ID         int64  `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
	Date       string `json:"date"`
	Status     string `json:"status"`
}

// must be somewhere else than here
var (
	null   = "null"
	stored = "stored"
)

// CreateExpressionHandler - post method handler which stores an expression

func CreateExpressionHandler(ctx context.Context, expressionSaver storage.ExpressionInteractor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			http.Error(w, `{"error": Expected method POST}`, http.StatusMethodNotAllowed)
		}

		date := time.Now()

		jsonDec := json.NewDecoder(r.Body)

		var req Request
		if err := jsonDec.Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userID, ok := r.Context().Value("userid").(int64)
		if !ok {
			http.Error(w, "userID not received", http.StatusBadRequest)
			log.Printf("userID not received: %d", userID)
			return
		}

		if !validator.IsValidExpression(req.Expression) {
			http.Error(w, "Invalid expression", http.StatusBadRequest)
			return
		}

		var expressionStruct = storage.Expression{
			UserID:     userID,
			Expression: req.Expression,
			Answer:     null,
			Date:       date.Format("2006/01/02 15:04:05"),
			Status:     stored,
		}

		id, err := expressionSaver.InsertExpression(ctx, &expressionStruct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		go manager.Manage(ctx, expressionSaver, agentAddress())

		response := map[string]int64{"id": id}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
		log.Printf("Successful CreateExpressionHandler operation; id = %d", id)
	}
}

func GetExpressionsHandler(ctx context.Context, expressionSaver storage.ExpressionInteractor) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodGet {
			http.Error(w, `{"error": Expected method GET}`, http.StatusMethodNotAllowed)
		}

		userID, ok := r.Context().Value("userid").(int64)
		if !ok {
			http.Error(w, "userID not received", http.StatusBadRequest)
			log.Printf("userID not received: %d", userID)
			return
		}

		go manager.Manage(ctx, expressionSaver, agentAddress())

		allExpressions, err := expressionSaver.SelectExpressionsByID(ctx, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var respData []ResponseData

		for _, expr := range allExpressions {
			resp := ResponseData{
				ID:         expr.ID,
				Expression: expr.Expression,
				Result:     expr.Answer,
				Date:       expr.Date,
				Status:     expr.Status,
			}

			respData = append(respData, resp)
		}

		json.NewEncoder(w).Encode(respData)
		log.Print("Successful GetExpressionsHandler operation")
	}
}

func DeleteExpressionHandler(ctx context.Context, expressionSaver storage.ExpressionInteractor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		go manager.Manage(ctx, expressionSaver, agentAddress())

		pathValues := strings.Split(r.URL.Path, "/")
		if len(pathValues) < 3 || pathValues[2] == "" {
			http.Error(w, "Invalid URL format", http.StatusBadRequest)
			return
		}

		expressionID, err := strconv.ParseInt(pathValues[4], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Проверка на существование выражения
		expression, err := expressionSaver.SelectExpressionByID(ctx, int64(expressionID))
		if err != nil {
			http.Error(w, "Expression with this id was not found", http.StatusNotFound)
			return
		}

		// Удаление выражения
		err = expressionSaver.DeleteExpression(ctx, int64(expression.ID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		response := map[string]int64{"Expression with this id was successfully deleted": expression.ID}
		json.NewEncoder(w).Encode(response)
		log.Print("Successful DeleteExpressionHandler operation")
	}
}

func agentAddress() string {
	agentHost, ok := os.LookupEnv("AGENT_HOST")
	if !ok {
		log.Print("AGENT_HOST not set, using 0.0.0.0")
		agentHost = "0.0.0.0"
	}

	agentPort, ok := os.LookupEnv("AGENT_PORT")
	if !ok {
		log.Print("AGENT_PORT not set, using 5000")
		agentPort = "5000"
	}
	return fmt.Sprintf("%s:%s", agentHost, agentPort)
}
