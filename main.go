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
	crearEstructurasProyecto(config.ProjectName)
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
	pathDev := crearEstructurasProyecto(config.ProjectName)["pathDev"]
	pathProd := crearEstructurasProyecto(config.ProjectName)["pathProd"]
	//fmt.Println("Carpetas creadas correctamente." + pathDev)

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
	// Finalización
	fmt.Println("\nFuncionando!")
}

func menu() ManifestConfig {

	projectName := leerLinea("1. Escriba el nombre del proyecto")
	imageRepository := leerLinea("2. Coloque el repositorio de la imagen docker ej: docker.io/empresa-usuario/nombre-imagen")
	dbImageName := leerLinea("3. Coloque el nombre de la imagen de la base de datos ej: mongo:8, postgres:18")
	fmt.Println("Iniciando configuración...")

	return ManifestConfig{
		ProjectName:     projectName,
		ImageRepository: imageRepository,
		DBImageName:     dbImageName,
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
	err := os.MkdirAll(ruta, 0755)
	if err != nil {
		fmt.Println("Error al crear la carpeta:", err)
		return err
	}
	fmt.Println("Carpeta creada (o ya existía) en:", ruta)
	return nil
}
func crearEstructurasProyecto(projectName string) map[string]string {
	// rutas := []string{
	// 	"./" + projectName + "/dev/sites/" + projectName,
	// 	"./" + projectName + "/prod/sites/" + projectName,
	// 	"./" + projectName + "/sitesfoca/base/" + projectName,
	// 	"./" + projectName + "/sitesfoca/dev/" + projectName,
	// 	"./" + projectName + "/sitesfoca/prod/" + projectName,
	// }
	rutas := map[string]string{
		"pathDev":       "./" + projectName + "/dev/sites/" + projectName,
		"pathProd":      "./" + projectName + "/prod/sites/" + projectName,
		"pathBase":      "./" + projectName + "/sitesfoca/base/" + projectName,
		"pathSitesDev":  "./" + projectName + "/sitesfoca/dev/" + projectName,
		"pathSitesProd": "./" + projectName + "/sitesfoca/prod/" + projectName,
	}

	for _, ruta := range rutas {
		crearCarpetas(ruta)
		// err := os.MkdirAll(ruta, 0755)
		// if err != nil {
		// 	fmt.Println("Error al crear la carpeta:", err)
		// } else {
		// 	fmt.Println("Carpeta creada (o ya existía) en:", ruta)
		// }
	}
	return rutas
}
