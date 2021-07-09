package main

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestGetHello(t *testing.T) {
	t.Run("Returns app ver ip path", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/testing", nil)
		response := httptest.NewRecorder()

		HelloServer(response, request)

		got := response.Body.String()
		re := regexp.MustCompile(`^App ver: 1.0.0, Ip: (([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]), Path: /testing$`)

		if !re.MatchString(got) {
			t.Errorf("got %q", got)
			t.Errorf("want %q", re)
		}
	})
}
