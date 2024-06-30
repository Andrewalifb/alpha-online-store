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

func GetProduct(id string) (*dto.ReadProductApiResponse, error) {
	req, err := http.NewRequest("GET", os.Getenv("GET_PRODUCT_BY_ID_API")+id, nil)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non 200 response status code %d while fetching product", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	var respDto dto.ReadProductApiResponse
	err = json.Unmarshal(respBody, &respDto)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &respDto, nil
}

func UpdateProductById(id string, product *dto.UpdateProductRequest) (*dto.UpdateProductApiResponse, error) {
	productJson, err := json.Marshal(product)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	req, err := http.NewRequest("PUT", os.Getenv("UPDATE_PRODUCT_BY_ID_API")+id, bytes.NewBuffer(productJson))
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

	var respDto dto.UpdateProductApiResponse
	err = json.Unmarshal(respBody, &respDto)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &respDto, nil
}
