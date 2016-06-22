package main // import "github.com/harbur/tide/cmd/tide"

import (
	"bytes"

	"k8s.io/helm/cmd/tiller/environment"
	chartutil "k8s.io/helm/pkg/chartutil"
)

var env = environment.New()

func readManifest(chart string) (string, error) {
	chrt, err := chartutil.Load(chart)
	if err != nil {
		print_error("Chart cannot be loaded: %", err)
		return "", err
	}

	values_file := chart + "/values.yaml"
	if len(profile) > 0 {
		values_file = chart + "/values-" + profile + ".yaml"
	}
	vals, err := chartutil.ReadValuesFile(values_file)
	debug("Values loaded %s", vals)
	if err != nil {
		print_error("Values cannot be loaded: %s", err)
		return "", err
	}

	files, err := env.EngineYard.Default().Render(chrt, vals)

	b := bytes.NewBuffer(nil)
	for name, file := range files {
		if manifest_file == "" || manifest_file == name {
			// Ignore empty documents because the Kubernetes library can't handle
			// them.
			if len(file) > 0 {
				b.WriteString("\n---\n# Source: " + name + "\n")
				b.WriteString(file)
			}
		}
	}

	manifest := b.String()
	return manifest, nil
}
