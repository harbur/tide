package main // import "github.com/harbur/tide/cmd/tide"

import (
	"bytes"
	"k8s.io/helm/cmd/tiller/environment"
	chartutil "k8s.io/helm/pkg/chartutil"
	chart "k8s.io/helm/pkg/proto/hapi/chart"
	"os"
)

var env = environment.New()

func readManifest(chartname string) (string, error) {
	chrt, err := chartutil.Load(chartname)
	if err != nil {
		print_error("Chart cannot be loaded: %", err)
		return "", err
	}

	// values.yaml is optional
	vals := chartutil.Values{}
	if _, err := os.Stat(chartname + "/values.yaml"); err == nil {
		vals, err = chartutil.ReadValuesFile(chartname + "/values.yaml")
		debug("Values loaded %v", vals)
		if err != nil {
			print_error("Values cannot be loaded: %s", err)
			return "", err
		}
	}

	// input_file is optional
	overrides := chartutil.Values{}
	if _, err := os.Stat(input_file); err == nil {
		overrides, err = chartutil.ReadValuesFile(input_file)
		if err != nil {
			print_error("Overrides cannot be loaded: %s", err)
			return "", err
		}
	}

	valsYAML, _ := vals.YAML()
	valss := &chart.Config{Raw: valsYAML}
	cfg, err := chartutil.CoalesceValues(chrt, valss, overrides.AsMap())
	debug("Coalsce loaded %#v %v", cfg, cfg)
	if err != nil {
		return "", err
	}

	files, err := env.EngineYard.Default().Render(chrt, cfg)

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
