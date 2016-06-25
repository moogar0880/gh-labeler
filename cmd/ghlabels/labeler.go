package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/moogar0880/ghlabeler"
)

const (
	// Version constant for printing the version information
	Version = "0.0.0-alpha"

	// DefaultFile is the default file name to load label config data from
	DefaultFile = "labels.json"
)

var (
	flHelp         = flag.Bool("help", false, "Print usage")
	flVersion      = flag.Bool("version", false, "Print version information and quit")
	flRemoveAbsent = flag.Bool("remove", false, "Remove labels that are absent from the config file")
	flFile         = flag.String("file", DefaultFile, "Specify Config File to Load")
	flToken        = flag.String("token", os.Getenv("GH_ACCESS_TOKEN"), "The Github Access Token to use")
)

// handleFlags parses commandline flags from stdin and handles printing help or
// version information, as well as doing validation on provided commandline args
func handleFlags() {
	flag.Parse()
	if *flHelp {
		fmt.Println("HALP")
		os.Exit(0)
	} else if *flVersion {
		fmt.Println(Version)
		os.Exit(0)
	} else if *flToken == "" {
		fmt.Println("Must Specify '-token' flag or set 'GH_ACCESS_TOKEN' environment variable")
		os.Exit(1)
	}
}

func main() {
	handleFlags()
	config := ghlabels.LoadConfig(*flFile)
	labeler := ghlabels.NewLabeler(*flToken, config)
	existingLabels := labeler.GetLabels()
	labeler.SetLabels(existingLabels, *flRemoveAbsent)
}
