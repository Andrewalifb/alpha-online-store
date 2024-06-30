package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Andrewalifb/alpha-online-store/cart-services/dto"
)

func CreateTransactions(transactions *dto.CreateTransactionsRequest) (*dto.CreateTransactionsResponse, error) {
	transactionsJson, err := json.Marshal(transactions)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	req, err := http.NewRequest("POST", os.Getenv("CREATE_NEW_TRANSACTION"), bytes.NewBuffer(transactionsJson))
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	var respDto dto.CreateTransactionsResponse
	err = json.Unmarshal(respBody, &respDto)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &respDto, nil
}
