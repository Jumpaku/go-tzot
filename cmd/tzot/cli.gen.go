// Code generated by cyamli v0.0.12, DO NOT EDIT.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Func[Input any] func(subcommand []string, input Input, inputErr error) (err error)

type CLI struct {
	Gen CLI_Gen

	List CLI_List

	FUNC Func[CLI_Input]
}

func (CLI) DESC_Simple() string {
	return "tzot (v0.0.0):\nA command line tool to generate CLI for your app from YAML-based schema.\n\nUsage:\n    $ tzot [<option>]...\n\nOptions:\n    -help, -version\n\nSubcommands:\n    gen, list\n\n"
}
func (CLI) DESC_Detail() string {
	return "tzot (v0.0.0):\nA command line tool to generate CLI for your app from YAML-based schema.\n\nUsage:\n    $ tzot [<option>]...\n\n\nOptions:\n    -help[=<boolean>], -h[=<boolean>]  (default=false):\n        Shows description of this tool\n\n    -version[=<boolean>], -v[=<boolean>]  (default=false):\n        Shows version of this tool\n\n\nSubcommands:\n    gen:\n        Generates Go code to handle timezone offset transitions for specified timezone IDs.\n\n    list:\n        Lists of all available timezone IDs.\n\n"
}

type CLI_Input struct {
	Opt_Help bool

	Opt_Version bool
}

func resolve_CLI_Input(input *CLI_Input, restArgs []string) error {
	*input = CLI_Input{

		Opt_Help: false,

		Opt_Version: false,
	}

	var arguments []string
	for idx, arg := range restArgs {
		if arg == "--" {
			arguments = append(arguments, restArgs[idx+1:]...)
			break
		}
		if !strings.HasPrefix(arg, "-") {
			arguments = append(arguments, arg)
			continue
		}
		optName, lit, cut := strings.Cut(arg, "=")
		consumeVariables(optName, lit, cut)

		switch optName {
		default:
			return fmt.Errorf("unknown option %q", optName)

		case "-help", "-h":
			if !cut {
				lit = "true"

			}
			if err := parseValue(&input.Opt_Help, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-version", "-v":
			if !cut {
				lit = "true"

			}
			if err := parseValue(&input.Opt_Version, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		}
	}

	return nil
}

type CLI_Gen struct {
	FUNC Func[CLI_Gen_Input]
}

func (CLI_Gen) DESC_Simple() string {
	return "Generates Go code to handle timezone offset transitions for specified timezone IDs.\n\nUsage:\n    $ <program> gen [<option>|<argument>]... [-- [<argument>]...]\n\nOptions:\n    -help, -output-path, -package\n\nArguments:\n    <timezone_id_list>...\n\n"
}
func (CLI_Gen) DESC_Detail() string {
	return "Generates Go code to handle timezone offset transitions for specified timezone IDs.\n\nUsage:\n    $ <program> gen [<option>|<argument>]... [-- [<argument>]...]\n\n\nOptions:\n    -help[=<boolean>], -h[=<boolean>]  (default=false):\n        Shows description of this subcommand.\n\n    -output-path=<string>, -o=<string>  (default=\"\"):\n        Specifies output path of gen subcommand. If not specified, stdout is used.\n\n    -package=<string>, -p=<string>  (default=\"tzot_gen\"):\n        Specifies package that output API belongs to.\n\n\nArguments:\n    [0:] [<timezone_id_list:string>]...\n        specifies timezone IDs for which Go code is generated.\n\n"
}

type CLI_Gen_Input struct {
	Opt_Help bool

	Opt_OutputPath string

	Opt_Package string

	Arg_TimezoneIdList []string
}

func resolve_CLI_Gen_Input(input *CLI_Gen_Input, restArgs []string) error {
	*input = CLI_Gen_Input{

		Opt_Help: false,

		Opt_OutputPath: "",

		Opt_Package: "tzot_gen",
	}

	var arguments []string
	for idx, arg := range restArgs {
		if arg == "--" {
			arguments = append(arguments, restArgs[idx+1:]...)
			break
		}
		if !strings.HasPrefix(arg, "-") {
			arguments = append(arguments, arg)
			continue
		}
		optName, lit, cut := strings.Cut(arg, "=")
		consumeVariables(optName, lit, cut)

		switch optName {
		default:
			return fmt.Errorf("unknown option %q", optName)

		case "-help", "-h":
			if !cut {
				lit = "true"

			}
			if err := parseValue(&input.Opt_Help, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-output-path", "-o":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_OutputPath, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		case "-package", "-p":
			if !cut {
				return fmt.Errorf("value is not specified to option %q", optName)

			}
			if err := parseValue(&input.Opt_Package, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		}
	}

	if len(arguments) <= 0-1 {
		return fmt.Errorf("too few arguments")
	}
	if err := parseValue(&input.Arg_TimezoneIdList, arguments[0:]...); err != nil {
		return fmt.Errorf("values [%s] are not assignable to arguments at [%d:]", strings.Join(arguments[0:], " "), 0)
	}

	return nil
}

type CLI_List struct {
	FUNC Func[CLI_List_Input]
}

func (CLI_List) DESC_Simple() string {
	return "Lists of all available timezone IDs.\n\nUsage:\n    $ <program> list [<option>]...\n\nOptions:\n    -help\n\n"
}
func (CLI_List) DESC_Detail() string {
	return "Lists of all available timezone IDs.\n\nUsage:\n    $ <program> list [<option>]...\n\n\nOptions:\n    -help[=<boolean>], -h[=<boolean>]  (default=false):\n        Shows description of this subcommand.\n\n"
}

type CLI_List_Input struct {
	Opt_Help bool
}

func resolve_CLI_List_Input(input *CLI_List_Input, restArgs []string) error {
	*input = CLI_List_Input{

		Opt_Help: false,
	}

	var arguments []string
	for idx, arg := range restArgs {
		if arg == "--" {
			arguments = append(arguments, restArgs[idx+1:]...)
			break
		}
		if !strings.HasPrefix(arg, "-") {
			arguments = append(arguments, arg)
			continue
		}
		optName, lit, cut := strings.Cut(arg, "=")
		consumeVariables(optName, lit, cut)

		switch optName {
		default:
			return fmt.Errorf("unknown option %q", optName)

		case "-help", "-h":
			if !cut {
				lit = "true"

			}
			if err := parseValue(&input.Opt_Help, lit); err != nil {
				return fmt.Errorf("value %q is not assignable to option %q", lit, optName)
			}

		}
	}

	return nil
}

func NewCLI() CLI {
	return CLI{}
}

func Run(cli CLI, args []string) error {
	subcommandPath, restArgs := resolveSubcommand(args)
	switch strings.Join(subcommandPath, " ") {

	case "":
		funcMethod := cli.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.FUNC not assigned", "")
		}
		var input CLI_Input
		err := resolve_CLI_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	case "gen":
		funcMethod := cli.Gen.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.Gen.FUNC not assigned", "gen")
		}
		var input CLI_Gen_Input
		err := resolve_CLI_Gen_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	case "list":
		funcMethod := cli.List.FUNC
		if funcMethod == nil {
			return fmt.Errorf("%q is unsupported: cli.List.FUNC not assigned", "list")
		}
		var input CLI_List_Input
		err := resolve_CLI_List_Input(&input, restArgs)
		return funcMethod(subcommandPath, input, err)

	}
	return nil
}

func resolveSubcommand(args []string) (subcommandPath []string, restArgs []string) {
	if len(args) == 0 {
		panic("command line arguments are too few")
	}
	subcommandSet := map[string]bool{
		"":    true,
		"gen": true, "list": true,
	}

	for _, arg := range args[1:] {
		if arg == "--" {
			break
		}
		pathLiteral := strings.Join(append(append([]string{}, subcommandPath...), arg), " ")
		if !subcommandSet[pathLiteral] {
			break
		}
		subcommandPath = append(subcommandPath, arg)
	}

	return subcommandPath, args[1+len(subcommandPath):]
}

func parseValue(dstPtr any, strValue ...string) error {
	switch dstPtr := dstPtr.(type) {
	case *[]bool:
		val := make([]bool, len(strValue))
		for idx, str := range strValue {
			if err := parseValue(&val[idx], str); err != nil {
				return fmt.Errorf("fail to parse %#v as []bool: %w", str, err)
			}
		}
		*dstPtr = val
	case *[]float64:
		val := make([]float64, len(strValue))
		for idx, str := range strValue {
			if err := parseValue(&val[idx], str); err != nil {
				return fmt.Errorf("fail to parse %#v as []float64: %w", str, err)
			}
		}
		*dstPtr = val
	case *[]int64:
		val := make([]int64, len(strValue))
		for idx, str := range strValue {
			if err := parseValue(&val[idx], str); err != nil {
				return fmt.Errorf("fail to parse %#v as []int64: %w", str, err)
			}
		}
		*dstPtr = val
	case *[]string:
		val := make([]string, len(strValue))
		for idx, str := range strValue {
			if err := parseValue(&val[idx], str); err != nil {
				return fmt.Errorf("fail to parse %#v as []string: %w", str, err)
			}
		}
		*dstPtr = val
	case *bool:
		val, err := strconv.ParseBool(strValue[0])
		if err != nil {
			return fmt.Errorf("fail to parse %q as bool: %w", strValue[0], err)
		}
		*dstPtr = val
	case *float64:
		val, err := strconv.ParseFloat(strValue[0], 64)
		if err != nil {
			return fmt.Errorf("fail to parse %q as float64: %w", strValue[0], err)
		}
		*dstPtr = val
	case *int64:
		val, err := strconv.ParseInt(strValue[0], 0, 64)
		if err != nil {
			return fmt.Errorf("fail to parse %q as int64: %w", strValue[0], err)
		}
		*dstPtr = val
	case *string:
		*dstPtr = strValue[0]
	}

	return nil
}

func consumeVariables(...any) {}
