package main

import (
	"bytes"
	"text/template"
)

func certs(projectName string, dns string) string {
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
    kind: ClusterIssuer
    name: letsencrypt-staging
`

	t, err := template.New("cert").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func configMap(projectName string) string {
	data := ManifestConfig{
		ProjectName: projectName,
	}

	tmpl := `apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .ProjectName }}-configmap
data:
  APP_ENV: "development"
`

	t, err := template.New("configMap").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func secret(projectName string) string {
	data := ManifestConfig{
		ProjectName: projectName,
	}

	tmpl := `apiVersion: v1
kind: Secret
metadata:
  labels:
    app: {{ .ProjectName }}
    tier: backend
    type: config
  name: {{ .ProjectName }}-secret
type: Opaque
stringData:
  SECRET_LOCAL: "XXX"

`

	t, err := template.New("secret").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func pv(projectName string) string {
	data := ManifestConfig{
		ProjectName: projectName,
	}

	tmpl := `apiVersion: v1
kind: PersistentVolume
metadata:
  name: {{ .ProjectName }}-pv
spec:
  capacity:
    storage: 50Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  volumeMode: Filesystem
  storageClassName: oci-bv-expandable-iscsi-sites-foca
  csi:
    driver: blockvolume.csi.oraclecloud.com
    volumeHandle: ocid1.volume.oc1.sa-santiago-1.CAMBIAME
    fsType: ext4
`

	t, err := template.New("pv").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func pvc(projectName string) string {
	data := ManifestConfig{
		ProjectName: projectName,
	}

	tmpl := `apiVersion: v1
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

	t, err := template.New("pvc").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func databaseHeadless(projectName string, dbname string, dbport string) string {
	data := ManifestConfig{
		ProjectName: projectName,
		DBImageName: dbname,
		DBPort:      dbport,
	}

	tmpl := `apiVersion: v1
kind: Service
metadata:
  name: {{ .ProjectName }}-headless
  labels:
    app: {{ .ProjectName }}
    tier: database
    type: headless
spec:
  clusterIP: None
  selector:
    app: {{ .ProjectName }}
    tier: database
  ports:
    - name: {{ .DBImageName }}db
      port: {{ .DBPort }}
      targetPort: {{ .DBPort }}

`

	t, err := template.New("databaseHeadless").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func databaseService(projectName string, dbname string, dbport string) string {
	data := ManifestConfig{
		ProjectName: projectName,
		DBImageName: dbname,
		DBPort:      dbport,
	}

	tmpl := `apiVersion: v1
kind: Service
metadata:
  name: {{ .ProjectName }}-{{ .DBImageName }}-service
  labels:
    app: {{ .ProjectName }}
    tier: database
    type: service
spec:
  ports:
    - name: "{{ .DBPort }}"
      port: {{ .DBPort }}
      targetPort: {{ .DBPort }}
  selector:
    app: {{ .ProjectName }}
    tier: database
    type: service

`

	t, err := template.New("databaseService").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func databaseStatefulSet(projectName string, dbname string, dbtagname string, dbport string) string {
	data := ManifestConfig{
		ProjectName: projectName,
		DBImageName: dbname,
		DBPort:      dbport,
		DBTagName:   dbtagname,
	}

	tmpl := `apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    app: {{ .ProjectName }}
    tier: database
    type: statefulset
  name: {{ .ProjectName }}-{{ .DBImageName }}-db
spec:
  serviceName: {{ .ProjectName }}-{{ .DBImageName }}-headless
  updateStrategy:
    type: RollingUpdate 
  replicas: 1
  selector:
    matchLabels:
      app: {{ .ProjectName }}
      tier: database
      type: service
  template:
    metadata:
      labels:
        app: {{ .ProjectName }}
        tier: database
        type: service
    spec:
      containers:
        - name: {{ .DBImageName }}
          image: docker.io/library/{{ .DBImageName }}:{{ .DBTagName }}
          envFrom:
		        - configMapRef:
		            name: {{ .ProjectName }}-configmap
		        - secretRef:
		            name: {{ .ProjectName }}-secret
          ports:
            - containerPort: {{ .DBPort }}
              name: {{ .DBImageName }}
          volumeMounts:
            - name: db-{{ .DBImageName }}-{{ .ProjectName }}
              mountPath: /data/db # <-- Asegúrate de que esta ruta es correcta para tu base de datos
      volumes:
        - name: db-{{ .DBImageName }}-{{ .ProjectName }}
          persistentVolumeClaim:
            claimName: {{ .ProjectName }}-pvc
`

	t, err := template.New("databaseStatefulSet").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func developmentService(projectName string, appport string) string {
	data := ManifestConfig{
		ProjectName: projectName,
		APPport:     appport,
	}

	tmpl := `apiVersion: v1
kind: Service
metadata:
  labels:
    app: {{ .ProjectName }}
    tier: backend
    type: service
  name: {{ .ProjectName }}-service
spec:
  ports:
    - port: {{ .APPport }}
      targetPort: {{ .APPport }}
  selector:
    app: {{ .ProjectName }}
    tier: backend

`

	t, err := template.New("developmentService").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func developmentDeployment(projectName string, imageRepository string, appport string) string {
	data := ManifestConfig{
		ProjectName:     projectName,
		ImageRepository: imageRepository,
		APPport:         appport,
	}

	tmpl := `apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: {{ .ProjectName }}
    tier: backend
    type: deployment
  name: {{ .ProjectName }}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{ .ProjectName }}
      tier: backend
  template:
    metadata:
      labels:
        app: {{ .ProjectName }}
        tier: backend
    spec:
      containers:
        - image: {{ .ImageRepository }}:latest
          name: backend-{{ .ProjectName }}
          ports:
            - containerPort: {{ .APPport }}
          envFrom:
          - configMapRef:
              name: {{ .ProjectName }}-configmap
          - secretRef:
              name: {{ .ProjectName }}-secret
      imagePullSecrets:
        - name: focasoft
      restartPolicy: Always
`

	t, err := template.New("developmentDeployment").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}
func ingress(projectName string, dns string, appport string) string {
	data := ManifestConfig{
		ProjectName: projectName,
		APPport:     appport,
		DNS:         dns,
	}

	tmpl := `---
apiVersion: traefik.io/v1alpha1
kind: IngressRoute
metadata:
  name: {{ .ProjectName }}-http
spec:
  entryPoints:
    - web
  routes:
    - match: Host("{{ .DNS }}") # <-- Colocar backticks
      kind: Rule 
      services:
        - name: {{ .ProjectName }}-service
          port: {{ .APPport }}
      middlewares:
        - name: redirect-https
---
apiVersion: traefik.io/v1alpha1
kind: IngressRoute
metadata:
  name: {{ .ProjectName }}-https
spec:
  entryPoints:
    - websecure
  routes:
  - match: Host("{{ .DNS }}") # <-- Colocar backticks
    kind: Rule
    services:
    - name: {{ .ProjectName }}-service
      port: {{ .APPport }}
  tls:
    secretName: CAMBIAME


`

	t, err := template.New("ingress").Parse(tmpl)
	if err != nil {
		return "error parseando template"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "error ejecutando template"
	}

	return buf.String()
}

func kustomizationBaseProject() string {

	tmpl := `apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
  - 01-certs/01-cert.yaml
  - 01-config/01-configmap.yaml
  - 01-config/01-secret.yaml
  - 01-pv-pvc/01-pv.yaml
  - 01-pv-pvc/02-pvc.yaml
  - 02-database/01-headless.yaml
  - 02-database/01-service.yaml
  - 02-database/02-statefulset.yaml
  - 03-backend/01-service.yaml
  - 03-backend/02-deployment.yaml
  - 04-ingress/01-ingress.yaml
`

	return tmpl
}
