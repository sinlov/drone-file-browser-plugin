package main

import (
	"github.com/joho/godotenv"
	"github.com/sinlov/drone-file-browser-plugin"
	"github.com/sinlov/drone-file-browser-plugin/cmd/cli"
	"github.com/sinlov/drone-info-tools/pkgJson"
	"github.com/sinlov/drone-info-tools/template"
	"log"
	"os"
)

func main() {
	pkgJson.InitPkgJsonContent(drone_file_browser_plugin.PackageJson)
	template.RegisterSettings(template.DefaultFunctions)

	app := cli.NewCliApp()

	// kubernetes runner patch
	if _, err := os.Stat("/run/drone/env"); err == nil {
		errDotEnv := godotenv.Overload("/run/drone/env")
		if errDotEnv != nil {
			log.Fatalf("load /run/drone/env err: %v", errDotEnv)
		}
	}

	// app run as urfave
	if err := app.Run(os.Args); nil != err {
		log.Println(err)
	}
}
