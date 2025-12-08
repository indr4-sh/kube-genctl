package main

import (
	"fmt"
)

func main() {
	//filename := "example.txt"
	fmt.Print(imageRepository())
	// file, err := os.Create(filename)
	// if err != nil {
	// 	fmt.Println("Error al crear el archivo:", err)
	// 	return
	// }
	// defer file.Close()
	fmt.Println("\nFuncionando!")
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
