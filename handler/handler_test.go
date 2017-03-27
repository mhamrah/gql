package handler

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mhamrah/gql"
	"github.com/mhamrah/gql/mocks"
	"github.com/stretchr/testify/assert"
)

func returnsSuccess(ctx context.Context, selection gql.Selection) (gql.Selectable, error) {
	return nil, nil
}

var handlers = map[string]gql.HandlerFunc{
	"success": returnsSuccess,
}

var schema = gql.Schema{}

func TestHandler(t *testing.T) {
	testCases := []struct {
		title        string
		body         string
		responseCode int
		response     string
	}{
		{"no query", "", http.StatusBadRequest, "no query present\n"},
		{"simple", "{ success {} }", http.StatusOK, "{\"data\":{}}\n"},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	svc := mocks.NewMockService(mockCtrl)
	svc.EXPECT().Handlers().Return(handlers)
	svc.EXPECT().Schema().Return(schema)
	h := NewGqlHandler(svc)

	ts := httptest.NewServer(h)
	defer ts.Close()

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
