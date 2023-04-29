package main

import (
	"io"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	cases := []struct {
		desc    string
		inPath  string
		wantMsg string
	}{
		{
			desc:    "A simple greeting",
			inPath:  "/hello?name=Alice",
			wantMsg: "👋, Alice!\n",
		},
		{
			desc:    "A harder case with a hex escape",
			inPath:  "/hello?name=Bets%20Platform",
			wantMsg: "👋, Bets Platform!\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			req := httptest.NewRequest("GET", tc.inPath, nil)
			w := httptest.NewRecorder()
			handler(w, req)

			resp := w.Result()
			if resp.StatusCode != 200 {
				t.Errorf("got HTTP status code %d, want 200", resp.StatusCode)
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("failed reading response body: %v", err)
			}
			if string(body) != tc.wantMsg {
				t.Errorf("response body was %q, wanted %q", string(body), tc.wantMsg)
			}
		})
	}
}
