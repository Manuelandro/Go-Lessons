package main

import (
	"bufio"
	"fmt"
	"my/cli/fileutils"
	"my/cli/resources"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want Webpack Config? (Y/n)")
	promptWebpack, _ := reader.ReadString('\n')
	fmt.Print("Do you want Babel Config? (Y/n)")
	promptBabel, _ := reader.ReadString('\n')
	fmt.Print("Do you want Eslint Config? (Y/n)")
	promptEslint, _ := reader.ReadString('\n')
	fmt.Print("Do you want a simple Redux Store Config? (Y/n)")
	promptRxstore, _ := reader.ReadString('\n')
	fmt.Print("Do you want a simple Redux Reducer Config? (Y/n)")
	promptRxreduce, _ := reader.ReadString('\n')
	fmt.Print("Do you want VsCode Config? (Y/n)")
	promptVscode, _ := reader.ReadString('\n')

	app := cli.NewApp()
	app.Name = "itly"
	app.Usage = "generate js boilerplate"
	app.Action = func(c *cli.Context) error {
		fmt.Println("bootstrapping...")
		if promptWebpack == 'Y' {
			fileutils.CreateFile(resources.WebpackPath, resources.Webpack)
		}

		if promptBabel == 'B' {
			fileutils.CreateFile(resources.BabelPath, resources.Babel)
		}

		fileutils.CreateFile(resources.EslintPath, resources.Eslint)
		fileutils.CreateFile(resources.ReduxStorePath, resources.ReduxStore)
		fileutils.CreateFile(resources.ReduxReducerPath, resources.ReduxReducer)
		return nil
	}

	app.Run(os.Args)
}
