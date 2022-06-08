package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print Version.",
	}

	app := &cli.App{
		Name:    "discordgo",
		Version: "v 1.0.0",
		Commands: []*cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create A new Project",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
					},
				},
				Action: CreateProject,
			},
		},
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateFile(file_name string) error {
	file, err := os.Create(file_name)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
func CreateDir(name string) error {
	err := os.Mkdir(name, 0750)
	if err != nil {
		return err
	}
	return nil
}
func CreateProject(ctx *cli.Context) error {
	name := ctx.String("name")
	err := CreateDir(name)
	file, err := os.Create("./" + name + "/main.go")
	if err != nil {
		return err
	}
	color.Red("go mod init " + name + ".me command runnings")
	command := exec.Command("go", "mod", "init", name+".me")
	command.Dir = name
	cmdOut, err := command.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(cmdOut))
	err = os.WriteFile("./"+name+"/main.go", []byte("func main() {\n\treturn nil\n}"), 0660)

	fmt.Println("Create Project.", file.Name())
	return nil
}
