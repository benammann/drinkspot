package graphql

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpHandler(c *gin.Context) {

	schema := getSchema()

	ctx := context.WithValue(c.Request.Context(), "gin_context", c)

	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&params); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	response := schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
	c.Writer.Write(responseJSON)

}
