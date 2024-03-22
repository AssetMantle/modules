package e2e

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/kinbiko/jsonassert"
)

func TestClientKeyAdd(t *testing.T) {
	url := "http://localhost:1317"
	client := &http.Client{}

	keyBody := []byte(`{"name":"test-key"}`)
	req, err := http.NewRequest("POST", url+"/keys/add", bytes.NewBuffer(keyBody))

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	jsonassert.New(t).Assertf(string(body), `{"height":"<<PRESENCE>>","result":{"success":true,"error":null,"keyOutput":{"name":"<<PRESENCE>>","type":"<<PRESENCE>>","address":"<<PRESENCE>>","pubkey":"<<PRESENCE>>","mnemonic":"<<PRESENCE>>"}}}`)
}
