package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCafeWhenOk(t *testing.T) {
	handler := http.HandlerFunc(mainHandle)

	requests := []struct {
		request      string
		expectedCode int
		expectedBody string
	}{
		{request: "/cafe?count=2&city=moscow", expectedCode: http.StatusOK, expectedBody: "Мир кофе,Сладкоежка"},
		{request: "/cafe?city=tula", expectedCode: http.StatusOK, expectedBody: "Пир и мир,Красиво есть не запретишь,Поздний завтрак"},
		{request: "/cafe?city=moscow&search=ложка", expectedCode: http.StatusOK, expectedBody: "Ложка и вилка"},
		{request: "/cafe", expectedCode: http.StatusBadRequest, expectedBody: "unknown city"},
		{request: "/cafe?city=omsk", expectedCode: http.StatusBadRequest, expectedBody: "unknown city"},
		{request: "/cafe?city=tula&count=na", expectedCode: http.StatusBadRequest, expectedBody: "incorrect count"},
	}

	for _, v := range requests {
		response := httptest.NewRecorder()
		req := httptest.NewRequest("GET", v.request, nil)

		handler.ServeHTTP(response, req)

		assert.Equal(t, v.expectedCode, response.Code)
		// пока сравнивать не будем, а просто выведем ответы
		// удалите потом этот вывод
		body := strings.TrimSpace(response.Body.String())
		fmt.Println(body)
		assert.Equal(t, v.expectedBody, body)
	}
}
