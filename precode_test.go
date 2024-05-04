package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	statusCode := responseRecorder.Code

	assert.Equal(t, http.StatusOK, statusCode)
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenIncorrectCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=rostov", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	statusCode := responseRecorder.Code
	incorrectCity := "wrong city value"

	require.Equal(t, http.StatusBadRequest, statusCode)
	require.Equal(t, responseRecorder.Body.String(), incorrectCity)
}

func TestMainHandlerCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=9&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	responseBody := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, totalCount, len(responseBody))
}
