package main

import (
	"bytes"
	"text/template"
)

func imageRepositoryProd(projectName string, imageRepository string) string {
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

func gitRepositoryProd(projectName string, imageRepository string) string {
	data := ManifestConfig{
		ProjectName:     projectName,
		ImageRepository: imageRepository,
	}

	tmpl := `apiVersion: kustomize.toolkit.fluxcd.io/v1
kind: Kustomization
metadata:
  name: {{.ProjectName}}-gitrepository
  namespace: flux-system
spec:
  interval: 1m
  path: ./clusters/sitesfoca/prod/{{.ProjectName}}
  prune: true
  sourceRef:
    kind: GitRepository
    name: flux-system
`

	t, err := template.New("gitRepository").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func imagePolicyProd(projectName string) string {
	data := ManifestConfig{
		ProjectName: projectName,
	}

	tmpl := `apiVersion: image.toolkit.fluxcd.io/v1
kind: ImagePolicy
metadata:
  name: {{ .ProjectName }}-imagepolicy
  namespace: flux-system
spec:
  imageRepositoryRef:
    name: {{ .ProjectName }}-imagerepository
  filterTags:
    pattern: '^master*'
  policy:
    alphabetical:
      order: asc
`

	t, err := template.New("imagePolicy").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}

func kustomizationProdSites() string {

	tmpl := `apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
  - gitrepository.yaml
  - imagerepository.yaml
  - imagepolicy.yaml
`
	return tmpl
}
