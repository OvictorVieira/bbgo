//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/OvictorVieira/bbgo/pkg/exchange/kucoin/kucoinapi"
)

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package kucoin

var symbolMap = map[string]string{
{{- range $k, $v := . }}
	{{ printf "%q" $k }}: {{ printf "%q" $v }},
{{- end }}
}

func toLocalSymbol(symbol string) string {
	s, ok := symbolMap[symbol]
	if ok {
		return s
	}

	return symbol
}
`))

type Market struct {
	Symbol string `json:"symbol"`
}

type ApiResponse struct {
	Data []Market `json:"data"`
}

func main() {

	const apiUrl = kucoinapi.RestBaseURL + "/v1/symbols"

	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	r := &ApiResponse{}
	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		log.Fatal(err)
	}

	var data = map[string]string{}
	for _, m := range r.Data {
		key := strings.ReplaceAll(strings.ToUpper(strings.TrimSpace(m.Symbol)), "-", "")
		data[key] = m.Symbol
	}

	f, err := os.Create("symbols.go")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	err = packageTemplate.Execute(f, data)
	if err != nil {
		log.Fatal(err)
	}
}
