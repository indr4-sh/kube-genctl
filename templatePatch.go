package main

import (
	"bytes"
	"text/template"
)

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
      env: prod
    includeSelectors: true
    includeTemplates: true
    includeApps: true


patches:
#  - path: 01-patch-certs-prod.yaml
  - path: 01-patch-configmap-prod.yaml
  - path: 01-patch-secret-prod.yaml
  - path: 02-patch-database-prod.yaml
  - path: 03-patch-backend-prod.yaml
#  - path: 04-patch-ingress-prod.yaml
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

func patchCertDev(projectName string, dns string) string {
	data := ManifestConfig{
		ProjectName: projectName,
		DNS:         dns,
	}

	tmpl := `apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: {{ .ProjectName }}-cert
spec:
  secretName: {{ .ProjectName }}-tls
  dnsNames:
    - {{ .DNS }}
  issuerRef:
    name: letsencrypt-staging
`

	t, err := template.New("patchCertDev").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func patchCertProd(projectName string, dns string) string {
	data := ManifestConfig{
		ProjectName: projectName,
		DNS:         dns,
	}

	tmpl := `apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: {{ .ProjectName }}-cert
spec:
  secretName: {{ .ProjectName }}-tls
  dnsNames:
    - {{ .DNS }}
  issuerRef:
    name: letsencrypt-prod
`

	t, err := template.New("patchCertProd").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}

func patchConfigMapDev(projectName string) string {
	data := ManifestConfig{
		ProjectName: projectName,
	}

	tmpl := `apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .ProjectName }}-configmap
data:
  ENV: "development" # <-- Colocar environment correspondiente

`

	t, err := template.New("patchConfigMapDev").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func patchConfigMapProd(projectName string) string {
	data := ManifestConfig{
		ProjectName: projectName,
	}

	tmpl := `apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .ProjectName }}-configmap
data:
  ENV: "production" # <-- Colocar environment correspondiente

`

	t, err := template.New("patchConfigMapProd").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func patchSecretDev(projectName string) string {
	data := ManifestConfig{
		ProjectName: projectName,
	}

	tmpl := `apiVersion: v1
kind: Secret
metadata:
  name: {{ .ProjectName }}-secret
type: Opaque
stringData:
  SECRET_LOCAL: "XXX" # <-- Colocar secret correspondiente
 
`

	t, err := template.New("patchSecretDev").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func patchSecretProd(projectName string) string {
	data := ManifestConfig{
		ProjectName: projectName,
	}

	tmpl := `apiVersion: v1
kind: Secret
metadata:
  name: {{ .ProjectName }}-secret
type: Opaque
stringData:
  SECRET_LOCAL: "XXX" # <-- Colocar secret correspondiente
 
`

	t, err := template.New("patchSecretProd").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func patchDatabaseDev(projectName string, volumeHandle string) string {
	data := ManifestConfig{
		ProjectName:   projectName,
		VolumeHandler: volumeHandle,
	}

	tmpl := `apiVersion: v1
kind: PersistentVolume
metadata:
  name: {{ .ProjectName }}-pv
spec:
  storageClassName: oci-bv-expandable-iscsi-sites-foca
  capacity:
    storage: 50Gi
  csi:
    volumeHandle: {{ .VolumeHandler }}
    fsType: ext4
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: {{ .ProjectName }}-pvc
spec:
  accessModes:
    - ReadWriteOnce
  volumeMode: Filesystem
  resources:
    requests:
      storage: 50Gi
  storageClassName: oci-bv-expandable-iscsi-sites-foca
  volumeName: {{ .ProjectName }}-pv

`

	t, err := template.New("patchDatabaseDev").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func patchDatabaseProd(projectName string, volumeHandle string) string {
	data := ManifestConfig{
		ProjectName:   projectName,
		VolumeHandler: volumeHandle,
	}

	tmpl := `apiVersion: v1
kind: PersistentVolume
metadata:
  name: {{ .ProjectName }}-pv
spec:
  storageClassName: oci-bv-expandable-iscsi-sites-foca
  capacity:
	storage: 50Gi
  csi:
    volumeHandle: xxx # <-- Colocar volume handle prod correspondiente
	fsType: ext4
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: {{ .ProjectName }}-pvc
spec:
  accessModes:
	- ReadWriteOnce
  volumeMode: Filesystem
  resources:
    requests:
      storage: 50Gi
  storageClassName: oci-bv-expandable-iscsi-sites-foca 
  volumeName: {{ .ProjectName }}-pv

`

	t, err := template.New("patchDatabaseProd").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}

func patchBackendDev(projectName string, imageRepository string, tagNameDev string) string {
	data := ManifestConfig{
		ProjectName:     projectName,
		ImageRepository: imageRepository,
		TagNameDev:      tagNameDev,
	}

	tmpl := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .ProjectName }}
spec:
  replicas: 1
  template:
    spec:
      containers:
      - name: backend-{{ .ProjectName }}
        image: {{ .ImageRepository }}:{{ .TagNameDev }} # {"$imagepolicy": "flux-system:{{ .ProjectName }}-imagepolicy"}

`

	t, err := template.New("patchBackendDev").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func patchBackendProd(projectName string, imageRepository string, tagNameProd string) string {
	data := ManifestConfig{
		ProjectName:     projectName,
		ImageRepository: imageRepository,
		TagNameProd:     tagNameProd,
	}

	tmpl := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .ProjectName }}
spec:
  replicas: 1
  template:
	spec:
	  containers:
	  - name: backend-{{ .ProjectName }}
		image: {{ .ImageRepository }}:{{ .TagNameProd }} # {"$imagepolicy": "flux-system:{{ .ProjectName }}-imagepolicy"}
`

	t, err := template.New("patchBackendProd").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func patchIngressDev(projectName string, dns string, appport string) string {
	data := ManifestConfig{
		ProjectName: projectName,
		DNS:         dns,
		APPport:     appport,
	}

	tmpl := `
apiVersion: traefik.io/v1alpha1
kind: IngressRoute
metadata:
  name: {{ .ProjectName }}-https
spec:
  routes:
    - match: Host("{{ .DNS }}")
      kind: Rule
      services:
        - name: {{ .ProjectName }}-service
          port: {{ .APPport }}
  tls:
    secretName: {{ .ProjectName }}-tls
---
apiVersion: traefik.io/v1alpha1
kind: IngressRoute
metadata:
  name: {{ .ProjectName }}-http
spec:
  routes:
    - match: Host("{{ .DNS }}")
      kind: Rule
      services:
        - name: {{ .ProjectName }}-service
          port: {{ .APPport }}
      middlewares:
        - name: redirect-https
          namespace: sites-foca-dev
`
	t, err := template.New("patchIngressDev").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func patchIngressProd(projectName string, dns string, appport string) string {
	data := ManifestConfig{
		ProjectName: projectName,
		DNS:         dns,
		APPport:     appport,
	}

	tmpl := `
apiVersion: traefik.io/v1alpha1
kind: IngressRoute
metadata:
  name: {{ .ProjectName }}-https
spec:
  routes:
	- match: Host("{{ .DNS }}") # <-- Agregar dominio prod correspondiente
	  kind: Rule
	  services:
		- name: {{ .ProjectName }}-service
		  port: {{ .APPport }}
  tls:
	secretName: {{ .ProjectName }}-tls
---
apiVersion: traefik.io/v1alpha1
kind: IngressRoute
metadata:
  name: {{ .ProjectName }}-http
spec:
  routes:
	- match: Host("{{ .DNS }}") # <-- Agregar dominio prod correspondiente
	  kind: Rule
	  services:
		- name: {{ .ProjectName }}-service
		  port: {{ .APPport }}
	  middlewares:
		- name: redirect-https
		  namespace: sites-foca
`
	t, err := template.New("patchIngressProd").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
