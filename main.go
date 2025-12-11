package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Generador de manifiestos:")
	config := menu()
	fmt.Println(config.ProjectName)
	paths := rutasProyecto(config.ProjectName)
	crearEstructurasProyecto(paths)
	// pathDev := "./" + config.ProjectName + "/dev/sites/" + config.ProjectName
	// pathProd := "./" + config.ProjectName + "/prod/sites/" + config.ProjectName
	// pathBase := "./" + config.ProjectName + "/sitesfoca/base/" + config.ProjectName
	// pathSitesDev := "./" + config.ProjectName + "/sitesfoca/dev/" + config.ProjectName
	// pathSitesProd := "./" + config.ProjectName + "/sitesfoca/prod/" + config.ProjectName
	// crearCarpetas(pathDev)
	// crearCarpetas(pathProd)
	// crearCarpetas(pathBase)
	// crearCarpetas(pathSitesDev)
	// crearCarpetas(pathSitesProd)
	//pathDev := crearEstructurasProyecto(config.ProjectName)["pathDev"]
	//pathProd := crearEstructurasProyecto(config.ProjectName)["pathProd"]
	//fmt.Println("Carpetas creadas correctamente." + pathDev)

	pathDev := paths.PathDev
	pathProd := paths.PathProd

	pathsBaseCert := paths.PathBaseCert
	pathsBaseConfig := paths.PathBaseConfig
	pathsBasePvPvc := paths.PathBasePvPvc
	pathsBaseDatabase := paths.PathBaseDatabase
	pathsBaseBackend := paths.PathBaseBackend
	pathsBaseIngress := paths.PathBaseIngress
	pathsSitesDev := paths.PathSitesDev
	pathsSitesProd := paths.PathSitesProd

	//========================================================
	// Crear archivos de manifiestos dev
	//=========================================================
	// Crear ImageRepository
	imageRepoContent := imageRepository(config.ProjectName, config.ImageRepository)
	imageRepoPath := pathDev + "/imagerepository.yaml"
	os.WriteFile(imageRepoPath, []byte(imageRepoContent), 0664)
	fmt.Printf("Archivo de ImageRepository creado en: %s\n", imageRepoPath)
	// Crear GitRepository
	gitRepositoryContent := gitRepository(config.ProjectName, config.ImageRepository)
	gitRepositoryPath := pathDev + "/gitrepository.yaml"
	os.WriteFile(gitRepositoryPath, []byte(gitRepositoryContent), 0664)
	fmt.Printf("Archivo de GitRepository creado en: %s\n", gitRepositoryPath)
	// Crear ImagePolicy
	imagePolicyContent := imagePolicy(config.ProjectName)
	imagePolicyPath := pathDev + "/imagepolicy.yaml"
	os.WriteFile(imagePolicyPath, []byte(imagePolicyContent), 0664)
	fmt.Printf("Archivo de ImagePolicy creado en: %s\n", imagePolicyPath)
	// Crear kustomization Dev Sites
	kustomizationDevSitesContent := kustomizationDevSites()
	kustomizationDevSitesPath := pathDev + "/kustomization.yaml"
	os.WriteFile(kustomizationDevSitesPath, []byte(kustomizationDevSitesContent), 0664)
	fmt.Printf("Archivo de Kustomization Dev Sites creado en: %s\n", kustomizationDevSitesPath)
	//========================================================
	// Crear archivos de manifiestos prod
	//=========================================================
	// Crear ImageRepository
	imageRepoContentProd := imageRepositoryProd(config.ProjectName, config.ImageRepository)
	imageRepoPathProd := pathProd + "/imagerepository.yaml"
	os.WriteFile(imageRepoPathProd, []byte(imageRepoContentProd), 0664)
	fmt.Printf("Archivo de ImageRepository creado en: %s\n", imageRepoPathProd)
	// Crear GitRepository
	gitRepositoryContentProd := gitRepositoryProd(config.ProjectName, config.ImageRepository)
	gitRepositoryPathProd := pathProd + "/gitrepository.yaml"
	os.WriteFile(gitRepositoryPathProd, []byte(gitRepositoryContentProd), 0664)
	fmt.Printf("Archivo de GitRepository creado en: %s\n", gitRepositoryPathProd)
	// Crear ImagePolicy
	imagePolicyContentProd := imagePolicyProd(config.ProjectName)
	imagePolicyPathProd := pathProd + "/imagepolicy.yaml"
	os.WriteFile(imagePolicyPathProd, []byte(imagePolicyContentProd), 0664)
	fmt.Printf("Archivo de ImagePolicy creado en: %s\n", imagePolicyPathProd)
	// Crear kustomization Prod Sites
	kustomizationProdSitesContent := kustomizationProdSites()
	kustomizationProdSitesPath := pathProd + "/kustomization.yaml"
	os.WriteFile(kustomizationProdSitesPath, []byte(kustomizationProdSitesContent), 0664)
	fmt.Printf("Archivo de Kustomization Prod Sites creado en: %s\n", kustomizationProdSitesPath)
	//========================================================
	// Crear archivos de manifiestos base
	//=========================================================
	// Crear certs
	certsContent := certs(config.ProjectName, config.DNS)
	certsPath := pathsBaseCert + "/01-cert.yaml"
	os.WriteFile(certsPath, []byte(certsContent), 0664)
	fmt.Printf("Archivo de Certs creado en: %s\n", certsPath)
	// Crear config
	configContent := configMap(config.ProjectName)
	configPath := pathsBaseConfig + "/01-configmap.yaml"
	os.WriteFile(configPath, []byte(configContent), 0664)
	fmt.Printf("Archivo de ConfigMap creado en: %s\n", configPath)
	//Crear secret
	secretContent := secret(config.ProjectName)
	secretPath := pathsBaseConfig + "/01-secret.yaml"
	os.WriteFile(secretPath, []byte(secretContent), 0664)
	fmt.Printf("Archivo de Secret creado en: %s\n", secretPath)
	// Crear pv
	pvContent := pv(config.ProjectName)
	pvPath := pathsBasePvPvc + "/01-pv.yaml"
	os.WriteFile(pvPath, []byte(pvContent), 0664)
	fmt.Printf("Archivo de PV creado en: %s\n", pvPath)
	// Crear pvc
	pvcContent := pvc(config.ProjectName)
	pvcPath := pathsBasePvPvc + "/02-pvc.yaml"
	os.WriteFile(pvcPath, []byte(pvcContent), 0664)
	fmt.Printf("Archivo de PVC creado en: %s\n", pvcPath)
	// Crear headless service de la base de datos
	databaseHeadlessContent := databaseHeadless(config.ProjectName, config.DBImageName, config.DBPort)
	databaseHeadlessPath := pathsBaseDatabase + "/01-headless.yaml"
	os.WriteFile(databaseHeadlessPath, []byte(databaseHeadlessContent), 0664)
	fmt.Printf("Archivo de Database Headless creado en: %s\n", databaseHeadlessPath)
	// Cerar service database
	databaseServiceContent := databaseService(config.ProjectName, config.DBImageName, config.DBPort)
	databaseServicePath := pathsBaseDatabase + "/01-service.yaml"
	os.WriteFile(databaseServicePath, []byte(databaseServiceContent), 0664)
	fmt.Printf("Archivo de Database Server creado en: %s\n", databaseServicePath)
	// Crear statefulset database
	databaseStatefulSetContent := databaseStatefulSet(config.ProjectName, config.DBImageName, config.DBTagName, config.DBPort)
	databaseStatefulSetPath := pathsBaseDatabase + "/02-statefulset.yaml"
	os.WriteFile(databaseStatefulSetPath, []byte(databaseStatefulSetContent), 0664)
	fmt.Printf("Archivo de Database StatefulSet creado en: %s\n", databaseStatefulSetPath)
	// Crear service backend
	backendServiceContent := developmentService(config.ProjectName, config.APPport)
	backendServicePath := pathsBaseBackend + "/01-service.yaml"
	os.WriteFile(backendServicePath, []byte(backendServiceContent), 0664)
	fmt.Printf("Archivo de Backend creado en: %s\n", backendServicePath)
	// Crear deployment backend
	backendDeploymentContent := developmentDeployment(config.ProjectName, config.ImageRepository, config.APPport)
	backendDeploymentPath := pathsBaseBackend + "/02-deployment.yaml"
	os.WriteFile(backendDeploymentPath, []byte(backendDeploymentContent), 0664)
	fmt.Printf("Archivo de Backend Deployment creado en: %s\n", backendDeploymentPath)
	// Crear ingress
	ingressContent := ingress(config.ProjectName, config.DNS, config.APPport)
	ingressPath := pathsBaseIngress + "/01-ingress.yaml"
	os.WriteFile(ingressPath, []byte(ingressContent), 0664)
	fmt.Printf("Archivo de Ingress creado en: %s\n", ingressPath)
	// Crear kustomization Base Project
	kustomizationBaseProjectContent := kustomizationBaseProject()
	kustomizationBaseProjectPath := paths.PathBase + "/kustomization.yaml"
	os.WriteFile(kustomizationBaseProjectPath, []byte(kustomizationBaseProjectContent), 0664)
	fmt.Printf("Archivo de Kustomization Base Project creado en: %s\n", kustomizationBaseProjectPath)

	//========================================================
	// Crear overlays dev y prod
	//=========================================================
	// Crear kustomization Dev Project
	kustomizationDevProjectContent := kustomizationDevProject(config.ProjectName)
	kustomizationDevProjectPath := pathsSitesDev + "/kustomization.yaml"
	os.WriteFile(kustomizationDevProjectPath, []byte(kustomizationDevProjectContent), 0664)
	fmt.Printf("Archivo de Kustomization Dev Project creado en: %s\n", kustomizationDevProjectPath)
	// Crear kustomization Prod Project
	kustomizationProdProjectContent := kustomizationProdProject(config.ProjectName)
	kustomizationProdProjectPath := pathsSitesProd + "/kustomization.yaml"
	os.WriteFile(kustomizationProdProjectPath, []byte(kustomizationProdProjectContent), 0664)
	fmt.Printf("Archivo de Kustomization Prod Project creado en: %s\n", kustomizationProdProjectPath)
	//========================================================
	// Crear archivos de manifiestos patch
	//=========================================================

	// Crear cert dev
	patchCertDevContent := patchCertDev(config.ProjectName, config.DNS)
	patchCertDevPath := pathsSitesDev + "/01-patch-certs-dev.yaml"
	os.WriteFile(patchCertDevPath, []byte(patchCertDevContent), 0664)
	fmt.Printf("Archivo de Patch Cert Dev creado en: %s\n", patchCertDevPath)
	// Crear cert prod
	patchCertProdContent := patchCertProd(config.ProjectName, config.DNS)
	patchCertProdPath := pathsSitesProd + "/01-patch-certs-prod.yaml"
	os.WriteFile(patchCertProdPath, []byte(patchCertProdContent), 0664)
	fmt.Printf("Archivo de Patch Cert Prod creado en: %s\n", patchCertProdPath)
	// Crear configmap dev
	patchConfigMapDevContent := patchConfigMapDev(config.ProjectName)
	patchConfigMapDevPath := pathsSitesDev + "/01-patch-configmap-dev.yaml"
	os.WriteFile(patchConfigMapDevPath, []byte(patchConfigMapDevContent), 0664)
	fmt.Printf("Archivo de Patch ConfigMap Dev creado en: %s\n", patchConfigMapDevPath)
	// Crear configmap prod
	patchConfigMapProdContent := patchConfigMapProd(config.ProjectName)
	patchConfigMapProdPath := pathsSitesProd + "/01-patch-configmap-prod.yaml"
	os.WriteFile(patchConfigMapProdPath, []byte(patchConfigMapProdContent), 0664)
	fmt.Printf("Archivo de Patch ConfigMap Prod creado en: %s\n", patchConfigMapProdPath)
	// Crear secret dev
	patchSecretDevContent := patchSecretDev(config.ProjectName)
	patchSecretDevPath := pathsSitesDev + "/01-patch-secret-dev.yaml"
	os.WriteFile(patchSecretDevPath, []byte(patchSecretDevContent), 0664)
	fmt.Printf("Archivo de Patch Secret Dev creado en: %s\n", patchSecretDevPath)
	// Crear secret prod
	patchSecretProdContent := patchSecretProd(config.ProjectName)
	patchSecretProdPath := pathsSitesProd + "/01-patch-secret-prod.yaml"
	os.WriteFile(patchSecretProdPath, []byte(patchSecretProdContent), 0664)
	fmt.Printf("Archivo de Patch Secret Prod creado en: %s\n", patchSecretProdPath)
	// Crear database dev
	patchDatabaseDevContent := patchDatabaseDev(config.ProjectName, config.VolumeHandler)
	patchDatabaseDevPath := pathsSitesDev + "/02-patch-database-dev.yaml"
	os.WriteFile(patchDatabaseDevPath, []byte(patchDatabaseDevContent), 0664)
	fmt.Printf("Archivo de Patch Database Dev creado en: %s\n", patchDatabaseDevPath)
	// Crear database prod
	patchDatabaseProdContent := patchDatabaseProd(config.ProjectName, config.VolumeHandler)
	patchDatabaseProdPath := pathsSitesProd + "/02-patch-database-prod.yaml"
	os.WriteFile(patchDatabaseProdPath, []byte(patchDatabaseProdContent), 0664)
	fmt.Printf("Archivo de Patch Database Prod creado en: %s\n", patchDatabaseProdPath)
	// Crear backend dev
	patchBackendDevContent := patchBackendDev(config.ProjectName, config.ImageRepository, config.TagNameDev)
	patchBackendDevPath := pathsSitesDev + "/03-patch-backend-dev.yaml"
	os.WriteFile(patchBackendDevPath, []byte(patchBackendDevContent), 0664)
	fmt.Printf("Archivo de Patch Backend Dev creado en: %s\n", patchBackendDevPath)
	// Crear backend prod
	patchBackendProdContent := patchBackendProd(config.ProjectName, config.ImageRepository, config.TagNameProd)
	patchBackendProdPath := pathsSitesProd + "/03-patch-backend-prod.yaml"
	os.WriteFile(patchBackendProdPath, []byte(patchBackendProdContent), 0664)
	fmt.Printf("Archivo de Patch Backend Prod creado en: %s\n", patchBackendProdPath)

	// Crear ingress dev
	patchIngressDevContent := patchIngressDev(config.ProjectName, config.DNS, config.APPport)
	patchIngressDevPath := pathsSitesDev + "/04-patch-ingress-dev.yaml"
	os.WriteFile(patchIngressDevPath, []byte(patchIngressDevContent), 0664)
	fmt.Printf("Archivo de Patch Ingress Dev creado en: %s\n", patchIngressDevPath)
	// Crear ingress prod
	patchIngressProdContent := patchIngressProd(config.ProjectName, config.DNS, config.APPport)
	patchIngressProdPath := pathsSitesProd + "/04-patch-ingress-prod.yaml"
	os.WriteFile(patchIngressProdPath, []byte(patchIngressProdContent), 0664)
	fmt.Printf("Archivo de Patch Ingress Prod creado en: %s\n", patchIngressProdPath)
	// Finalización
	fmt.Println("\nFuncionando!")
}

func menu() ManifestConfig {

	projectName := leerLinea("1. Escriba el nombre del proyecto")
	imageRepository := leerLinea("2. Coloque el repositorio de la imagen docker ej: docker.io/empresa-usuario/nombre-imagen")
	tagNameDev := leerLinea("3. Coloque el tag de la imagen docker ej: dev415, develop-2323")
	tagNameProd := leerLinea("4. Coloque el tag de la imagen docker ej: prod415, master-2323")
	appport := leerLinea("5. Coloque el puerto de la aplicación ej: 80, 8080, 3000")
	dbImageName := leerLinea("6. Coloque el nombre de la imagen de la base de datos ej: mongo, postgres, mariadb")
	dbTagName := leerLinea("7. Coloque el tag de la imagen de la base de datos ej: latest, 6.0, 14-alpine")
	dbPort := leerLinea("8. Coloque el puerto de la base de datos ej: 27017, 5432, 3306")
	dns := leerLinea("9. Coloque el dominio DNS del proyecto ej: miproyecto.misitio.com")
	volumeHandler := leerLinea("10. Coloque el volume handler ej: ocid1.volume.oc1.sa-santiago-1.xxx")
	fmt.Println("Iniciando configuración...")

	return ManifestConfig{
		ProjectName:     projectName,
		ImageRepository: imageRepository,
		DBImageName:     dbImageName,
		DBTagName:       dbTagName,
		DNS:             dns,
		VolumeHandler:   volumeHandler,
		DBPort:          dbPort,
		APPport:         appport,
		TagNameDev:      tagNameDev,
		TagNameProd:     tagNameProd,
	}

}
func leerLinea(prompt string) string {
	var reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Println(prompt)

		texto, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error leyendo la entrada:", err)
			os.Exit(1)
		}

		texto = strings.TrimSpace(texto)

		if texto == "" {
			fmt.Println("El valor no puede estar vacío, intente de nuevo.")
			continue
		}
		texto = strings.ToLower(texto)
		texto = strings.ReplaceAll(texto, " ", "")
		return texto
	}
}

func crearCarpetas(ruta string) error {
	err := os.MkdirAll(ruta, 0744)
	if err != nil {
		fmt.Println("Error al crear la carpeta:", err)
		return err
	}
	fmt.Println("Carpeta creada (o ya existía) en:", ruta)
	return nil
}
func crearEstructurasProyecto(rutas RutasConfig) {
	// rutas := []string{
	// 	"./" + projectName + "/dev/sites/" + projectName,
	// 	"./" + projectName + "/prod/sites/" + projectName,
	// 	"./" + projectName + "/sitesfoca/base/" + projectName,
	// 	"./" + projectName + "/sitesfoca/dev/" + projectName,
	// 	"./" + projectName + "/sitesfoca/prod/" + projectName,
	// }
	// rutas := map[string]string{
	// 	"pathDev":       "./" + projectName + "/dev/sites/" + projectName,
	// 	"pathProd":      "./" + projectName + "/prod/sites/" + projectName,
	// 	"pathBase":      "./" + projectName + "/sitesfoca/base/" + projectName,
	// 	"pathSitesDev":  "./" + projectName + "/sitesfoca/dev/" + projectName,
	// 	"pathSitesProd": "./" + projectName + "/sitesfoca/prod/" + projectName,
	// }

	// for _, ruta := range rutas {
	// 	crearCarpetas(ruta)
	// 	// err := os.MkdirAll(ruta, 0755)
	// 	// if err != nil {
	// 	// 	fmt.Println("Error al crear la carpeta:", err)
	// 	// } else {
	// 	// 	fmt.Println("Carpeta creada (o ya existía) en:", ruta)
	// 	// }
	// }
	//rutas := rutasProyecto(projectName)

	rutasSlice := []string{
		rutas.PathDev,
		rutas.PathProd,
		rutas.PathBase,
		rutas.PathBaseCert,
		rutas.PathBaseConfig,
		rutas.PathBasePvPvc,
		rutas.PathBaseDatabase,
		rutas.PathBaseBackend,
		rutas.PathBaseIngress,
		rutas.PathSitesDev,
		rutas.PathSitesProd,
	}
	for _, ruta := range rutasSlice {
		crearCarpetas(ruta)
	}
}
func rutasProyecto(projectName string) RutasConfig {
	return RutasConfig{
		PathDev:          "./" + projectName + "/dev/sites/" + projectName,
		PathProd:         "./" + projectName + "/prod/sites/" + projectName,
		PathBase:         "./" + projectName + "/sitesfoca/base/" + projectName,
		PathBaseCert:     "./" + projectName + "/sitesfoca/base/" + projectName + "/01-certs",
		PathBaseConfig:   "./" + projectName + "/sitesfoca/base/" + projectName + "/01-config",
		PathBasePvPvc:    "./" + projectName + "/sitesfoca/base/" + projectName + "/01-pv-pvc",
		PathBaseDatabase: "./" + projectName + "/sitesfoca/base/" + projectName + "/02-database",
		PathBaseBackend:  "./" + projectName + "/sitesfoca/base/" + projectName + "/03-backend",
		PathBaseIngress:  "./" + projectName + "/sitesfoca/base/" + projectName + "/04-ingress",
		PathSitesDev:     "./" + projectName + "/sitesfoca/dev/" + projectName,
		PathSitesProd:    "./" + projectName + "/sitesfoca/prod/" + projectName,
	}
}
