package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func generateController(directory string, name string, hasUseCase bool) {

	camelName := snakeCaseToCamelCase(name)
	titleName := strings.Title(camelName)
	path := directory + "controller/"

	f, err := os.Create(path + name + "_controller.go")
	check(err)

	defer f.Close()

	var d1 []byte
	if hasUseCase {
		d1 = []byte(`
package controller

type ` + titleName + `HttpController struct {
	` + camelName + `UseCase usecase.` + titleName + `UseCaseInterface
}

func New` + titleName + `Controller(` + camelName + `UseCase usecase.` + titleName + `UseCaseInterface) *` + titleName + `HttpController {
	return &` + titleName + `HttpController{
		` + camelName + `UseCase: ` + camelName + `UseCase,
	}
}
			`)
	} else {
		d1 = []byte(`
package controller

type ` + titleName + `HttpController struct {
	` + camelName + `UseCase usecase.` + titleName + `UseCaseInterface
}

func New` + titleName + `Controller(` + camelName + `UseCase usecase.` + titleName + `UseCaseInterface) *` + titleName + `HttpController {
	return &` + titleName + `HttpController{}
}
					`)
	}
	err = ioutil.WriteFile(path+name+"_controller.go", d1, 0777)
	check(err)
}
