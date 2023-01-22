//go:build e2e
// +build e2e

package sanple_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/takokun778/oreoreddd/e2e/helper"
)

func TestE2ESamplePut(t *testing.T) {
	t.Parallel()

	t.Run("200", func(t *testing.T) {
		t.Parallel()

		client := helper.NewClient(t)

		type Req struct {
			Name string `json:"name"`
		}

		param := "post"

		req := &Req{
			Name: param,
		}

		code, res := client.Request(t, http.MethodPost, "api/v1/samples/", req)
		if code != http.StatusOK {
			t.Errorf("status code is %d", code)
		}

		type Res struct {
			Sample struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"sample"`
		}

		var body Res

		json.Unmarshal(res, &body)

		if body.Sample.ID == "" {
			t.Errorf("sample id is empty")
		}
		if body.Sample.Name != param {
			t.Errorf("sample name is %s", body.Sample.Name)
		}

		id := body.Sample.ID

		url := "api/v1/samples/" + id

		param = "put"

		req = &Req{
			Name: param,
		}

		code, res = client.Request(t, http.MethodPut, url, req)
		if code != http.StatusOK {
			t.Errorf("status code is %d", code)
		}

		json.Unmarshal(res, &body)

		if body.Sample.ID != id {
			t.Errorf("sample id is empty")
		}
		if body.Sample.Name != param {
			t.Errorf("sample name is %s", body.Sample.Name)
		}
	})
}
