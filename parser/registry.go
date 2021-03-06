package parser

import (
	"github.com/deis/workflow-cli/cmd"
	docopt "github.com/docopt/docopt-go"
)

// Registry routes registry commands to their specific function
func Registry(argv []string) error {
	usage := `
Valid commands for registry:

registry:list        list registry info for an app
registry:set         set registry info for an app
registry:unset       unset registry info for an app

Use 'deis help [command]' to learn more.
`

	switch argv[0] {
	case "registry:list":
		return registryList(argv)
	case "registry:set":
		return registrySet(argv)
	case "registry:unset":
		return registryUnset(argv)
	default:
		if printHelp(argv, usage) {
			return nil
		}

		if argv[0] == "registry" {
			argv[0] = "registry:list"
			return registryList(argv)
		}

		PrintUsage()
		return nil
	}
}

func registryList(argv []string) error {
	usage := `
Lists registry information for an application.

Usage: deis registry:list [options]

Options:
  -a --app=<app>
    the uniquely identifiable name of the application.
`

	args, err := docopt.Parse(usage, argv, true, "", false, true)

	if err != nil {
		return err
	}

	return cmd.RegistryList(safeGetValue(args, "--app"))
}

func registrySet(argv []string) error {
	usage := `
Sets registry information for an application.

key/value pairs used to configure / authenticate against external registries

Usage: deis registry:set [options] <key>=<value>...

Arguments:
  <key> the registry key, for example: "username" or "password"
  <value> the registry value, for example: "bob" or "s3cur3pw1"

Options:
  -a --app=<app>
    the uniquely identifiable name for the application.
`

	args, err := docopt.Parse(usage, argv, true, "", false, true)

	if err != nil {
		return err
	}

	app := safeGetValue(args, "--app")
	info := args["<key>=<value>"].([]string)

	return cmd.RegistrySet(app, info)
}

func registryUnset(argv []string) error {
	usage := `
Unsets registry information for an application.

Usage: deis registry:unset [options] <key>...

Arguments:
  <key> the registry key to unset, for example: "username" or "password"

Options:
  -a --app=<app>
    the uniquely identifiable name for the application.
`

	args, err := docopt.Parse(usage, argv, true, "", false, true)

	if err != nil {
		return err
	}

	app := safeGetValue(args, "--app")
	key := args["<key>"].([]string)

	return cmd.RegistryUnset(app, key)
}
