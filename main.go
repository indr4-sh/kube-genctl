package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ManifestConfig struct {
	ProjectName     string
	ImageRepository string
	DBImageName     string
}

func main() {
	fmt.Println("Generador de manifiestos:")
	config := menu()
	fmt.Printf(config.ProjectName)
	crearCarpetas("./" + config.ProjectName + "/dev/sites/" + config.ProjectName)
	crearCarpetas("./" + config.ProjectName + "/prod/sites/" + config.ProjectName)
	crearCarpetas("./" + config.ProjectName + "/sites/base/" + config.ProjectName)
	crearCarpetas("./" + config.ProjectName + "/sites/dev/" + config.ProjectName)
	crearCarpetas("./" + config.ProjectName + "/sites/prod/" + config.ProjectName)
	fmt.Println("\nFuncionando!")
}

func menu() ManifestConfig {

	projectName := leerLinea("1. Escriba el nombre del proyecto")
	imageRepository := leerLinea("2. Coloque el repositorio de la imagen docker ej: docker.io/empresa/nombre-imagen")
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
		texto = strings.TrimSpace(texto)
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

func imageRepository( /*proyect_name string, image_name string*/ ) string {
	proyect_name := "xxx"
	image_name := "xxx"
	yamlContent := `apiVersion: image.toolkit.fluxcd.io/v1
kind: ImageRepository
metadata:
  name: %s-imagerepository
  namespace: flux-system
spec:
  image: docker.io/focasoftware/%s
  interval: 1m0s
  secretRef:
    name: focasoft`
	//fmt.Print(yamlContent, proyect_name, image_name)
	yaml := fmt.Sprintf(yamlContent, proyect_name, image_name)
	return yaml
}
