package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Run it using
	// go run . -dir=-name=name -r -uc -c
	// Ex : go run . -name=financing_repayment -r -uc

	dir := flag.String("dir", "", "")
	name := flag.String("name", "", "")
	repository := flag.Bool("r", false, "")
	usecase := flag.Bool("uc", false, "")
	controller := flag.Bool("c", false, "")

	flag.Parse()
	fmt.Printf("name: \"%v\"\n", string(*name))
	fmt.Printf("repository: \"%v\"\n", bool(*repository))

	if *name != "" {
		if *repository {
			_, repoErr := os.Stat(*dir + "repository/" + *name + "_repository.go")

			if repoErr == nil {
				fmt.Printf("Repository exists\n")
			} else {
				generateRepository(*dir, *name)
			}
		}

		if *usecase {
			_, usecaseErr := os.Stat(*dir + "usecase/" + *name + "_usecase.go")
			if usecaseErr == nil {
				fmt.Printf("Usecase exists\n")
			} else {
				generateUseCase(*dir, *name, *repository)
			}
		}

		if *controller {
			_, controllerErr := os.Stat(*dir + "controller/" + *name + "_controller.go")
			if controllerErr == nil {
				fmt.Printf("Controller exists\n")
			} else {
				generateController(*dir, *name, *usecase)
			}
		}
	}
}
