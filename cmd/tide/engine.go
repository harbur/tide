package main

import (
	"bytes"

	"github.com/kubernetes/helm/cmd/tiller/environment"
	chartutil "github.com/kubernetes/helm/pkg/chart"

	"github.com/kubernetes/helm/pkg/helm"
)

var env = environment.New()

func readManifest(installArg string) (string, error) {
	chfi, err := chartutil.LoadChart(installArg)
	if err != nil {
		return "", err
	}

	chpb, err := helm.ChartToProto(chfi)
	if err != nil {
		return "", err
	}

	vals, err := helm.ValuesToProto(chfi)
	if err != nil {
		return "", err
	}

	overrides := map[string]interface{}{
		"Release": map[string]interface{}{
			"Name":      "name",
			"Time":      "ts",
			"Namespace": "s.env.Namespace",
			"Service":   "Tide",
		},
		"Chart": "req.Chart.Metadata",
	}

	files, err := env.EngineYard.Default().Render(chpb, vals, overrides)

	b := bytes.NewBuffer(nil)
	for name, file := range files {
		// Ignore empty documents because the Kubernetes library can't handle
		// them.
		if len(file) > 0 {
			b.WriteString("\n---\n# Source: " + name + "\n")
			b.WriteString(file)
		}
	}

	manifest := b.String()
	return manifest, nil
}
