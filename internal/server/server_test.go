package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_longestSubstring(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "successfull_1",
			args: args{
				input: "abcabcbb",
			},
			want: "abc",
		},
		{
			name: "successfull_2",
			args: args{
				input: "aaabcda",
			},
			want: "abcd",
		},
		{
			name: "empty_string",
			args: args{},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubstring(tt.args.input); got != tt.want {
				t.Errorf("longestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSubstring(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/substring?input=abcabcbb", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetSubstring)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetSubstringFail(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/substring?", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetSubstring)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
