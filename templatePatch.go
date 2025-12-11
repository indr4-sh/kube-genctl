package main

import (
	"bytes"
	"text/template"
)

func patchConfigMap(projectName string, newData map[string]string) string {
	data := ManifestConfig{
		ProjectName: projectName,
		//NewData:     newData,
	}

	tmpl := `apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .ProjectName }}-config
data:
{{- range $key, $value := .NewData }}
  {{ $key }}: "{{ $value }}"
{{- end }}
`

	t, err := template.New("patchConfigMap").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func kustomizationDevProject(projectName string) string {
	data := ManifestConfig{
		ProjectName: projectName,
		//NewData:     newData,
	}
	tmpl := `apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

namespace: sites-foca-dev

resources:
  - ../../base/{{ .ProjectName }}/

labels:
  - pairs:
      env: dev
    includeSelectors: true
    includeTemplates: true
    includeApps: true


patches:
#  - path: 01-patch-certs-dev.yaml
  - path: 01-patch-configmap-dev.yaml
  - path: 01-patch-secret-dev.yaml
  - path: 02-patch-database-dev.yaml
  - path: 03-patch-backend-dev.yaml
#  - path: 04-patch-ingress-dev.yaml
`
	t, err := template.New("kustomizationDevProject").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}

func kustomizationProdProject(projectName string) string {
	data := ManifestConfig{
		ProjectName: projectName,
		//NewData:     newData,
	}
	tmpl := `apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

namespace: sites-foca

resources:
  - ../../base/{{ .ProjectName }}/

labels:
  - pairs:
      env: dev
    includeSelectors: true
    includeTemplates: true
    includeApps: true


patches:
#  - path: 01-patch-certs-dev.yaml
  - path: 01-patch-configmap-dev.yaml
  - path: 01-patch-secret-dev.yaml
  - path: 02-patch-database-dev.yaml
  - path: 03-patch-backend-dev.yaml
#  - path: 04-patch-ingress-dev.yaml
`
	t, err := template.New("kustomizationProdProject").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
