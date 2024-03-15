package main

import (
	"encoding/json"
	"fmt"
	"github.com/Jumpaku/go-tzot"
	"github.com/Jumpaku/go-tzot/generate"
	"log"
	"os"
)

//go:generate go run "github.com/Jumpaku/cyamli/cmd/cyamli@latest" golang -schema-path=cli.yaml -out-path=cli.gen.go

var cli = NewCLI()

func main() {
	// Overwrite behaviors
	cli.FUNC = showHelp
	cli.List.FUNC = listZoneIDs
	cli.Fetch.FUNC = fetchZones
	cli.Gen.FUNC = generateAPI
	// Run with command line arguments
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}

func showHelp(subcommand []string, input CLI_Input, inputErr error) (err error) {
	if input.Opt_Version && !input.Opt_Help {
		fmt.Println(tzot.ModuleVersion())
		return nil
	}
	if input.Opt_Help {
		fmt.Println(cli.DESC_Detail())
		return nil
	}
	fmt.Println(cli.DESC_Simple())
	if inputErr != nil {
		fmt.Fprintln(os.Stderr, cli.DESC_Simple())
		log.Panicf("input error: %+v", inputErr)
	}
	return nil
}

func listZoneIDs(subcommand []string, input CLI_List_Input, inputErr error) (err error) {
	if input.Opt_Help {
		fmt.Println(cli.List.DESC_Detail())
		return nil
	}
	if inputErr != nil {
		fmt.Fprintln(os.Stderr, cli.List.DESC_Simple())
		log.Panicf("input error: %+v", inputErr)
	}
	for _, zoneID := range tzot.AvailableZoneIDs() {
		fmt.Println(zoneID)
	}
	return nil
}

func fetchZones(subcommand []string, input CLI_Fetch_Input, inputErr error) (err error) {
	if input.Opt_Help {
		fmt.Println(cli.List.DESC_Detail())
		return nil
	}
	if inputErr != nil {
		fmt.Fprintln(os.Stderr, cli.List.DESC_Simple())
		log.Panicf("input error: %+v", inputErr)
	}

	zoneIDs := input.Arg_TimezoneIdList
	for _, zoneID := range zoneIDs {
		zone, found := tzot.GetZone(zoneID)
		if !found {
			log.Panicf("zone %q is not available", zoneID)
		}
		j, err := json.Marshal(zone)
		if err != nil {
			log.Panicf("failed to marshal zone %q as a JSON value: %+v", zoneID, err)
		}
		fmt.Println(string(j))
	}

	return nil
}

func generateAPI(subcommand []string, input CLI_Gen_Input, inputErr error) (err error) {
	if input.Opt_Help {
		fmt.Println(cli.Gen.DESC_Detail())
		return nil
	}
	if inputErr != nil {
		fmt.Fprintln(os.Stderr, cli.Gen.DESC_Simple())
		log.Panicf("input error: %+v", inputErr)
	}
	out := os.Stdout
	if input.Opt_OutputPath != "" {
		f, err := os.Create(input.Opt_OutputPath)
		if err != nil {
			log.Panicf("fail to open %q: %+v", input.Opt_OutputPath, err)
		}
		defer f.Close()
		out = f
	}

	zones := []tzot.Zone{}
	zoneIDs := input.Arg_TimezoneIdList
	if input.Opt_All {
		zoneIDs = tzot.AvailableZoneIDs()
	}
	for _, zoneID := range zoneIDs {
		zone, found := tzot.GetZone(zoneID)
		if !found {
			log.Panicf("zone %q is not available", zoneID)
		}
		zones = append(zones, zone)
	}

	err = generate.Generate(input.Opt_Package, zones, out)
	if err != nil {
		log.Panicf("fail to generate code: %+v", err)
	}

	return nil
}
