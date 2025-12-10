package main

import (
	"bytes"
	"text/template"
)

func certs(projectName string, imageRepository string) string {
	data := ManifestConfig{
		ProjectName:     projectName,
		ImageRepository: imageRepository,
	}

	tmpl := `apiVersion: image.toolkit.fluxcd.io/v1
kind: ImageRepository
metadata:
  name: {{ .ProjectName }}-imagerepository
  namespace: flux-system
spec:
  image: {{ .ImageRepository }}
  interval: 1m0s
  secretRef:
    name: focasoft
`

	t, err := template.New("imageRepository").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
