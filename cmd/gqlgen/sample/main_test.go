package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mhamrah/gql/cmd/gqlgen/sample/sample"
	"github.com/mhamrah/gql/handler"
	"github.com/stretchr/testify/assert"
)

func TestGraphqlIntrospection(t *testing.T) {
	svc := sample.New(&svc{})
	h := handler.NewGqlHandler(svc)
	ts := httptest.NewServer(h)
	defer ts.Close()

	testCases := []struct {
		title        string
		body         string
		responseCode int
		response     string
	}{
		{"no query", "", http.StatusBadRequest, "no query present\n"},
		{"missing param", "{ human {} }", http.StatusBadRequest, "id is not an present in argument list to human\n"},
		{"simple", "{ human(id: \"101\") {} }", http.StatusOK, "{\"data\":{}}\n"},
		{"simple", "{ human(id: \"101\") { name } }", http.StatusOK, "{\"data\":{\"name\":\"foo\"}}\n"},
		//	{"simple", "{ __schema { queryType } }", http.StatusOK, "{\"data\":{\"name\":\"foo\"}}\n"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.title), func(t *testing.T) {
			res, err := http.Post(ts.URL, "application/json", bytes.NewBufferString(tc.body))
			assert.NoError(t, err)
			assert.Equal(t, tc.responseCode, res.StatusCode)
			response, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			assert.NoError(t, err)
			assert.Equal(t, tc.response, string(response))
		})
	}
}
