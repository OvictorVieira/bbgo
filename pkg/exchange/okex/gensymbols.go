//go:build ignore
// +build ignore

package main

import (
	"context"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/OvictorVieira/promeheux.api/pkg/exchange/okex/okexapi"
)

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package okex

var spotSymbolMap = map[string]string{
{{- range $k, $v := . }}
	{{ printf "%q" $k }}: {{ printf "%q" $v }},
{{- end }}
}

`))

func main() {
	ctx := context.Background()
	client := okexapi.NewClient()
	instruments, err := client.PublicDataService.NewGetInstrumentsRequest().InstrumentType(okexapi.InstrumentTypeSpot).Do(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var data = map[string]string{}
	for _, instrument := range instruments {
		symbol := strings.ReplaceAll(instrument.InstrumentID, "-", "")
		data[symbol] = instrument.InstrumentID
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
