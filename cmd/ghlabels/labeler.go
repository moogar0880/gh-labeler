package main

import (
	"fmt"
	"os"

	"github.com/moogar0880/ghlabeler"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	// Version constant for printing the version information
	Version = "0.2.0"

	// DefaultFile is the default file name to load label config data from
	DefaultFile = "labels.json"
)

// Options is the set of command line options to provide to the ghlabels cmd
type Options struct {
	ghlabels.Config
	Version bool
	File    string
	Token   string
	Remove  bool
}

// NewOptions returns a new Options
func NewOptions() *Options {
	return &Options{File: DefaultFile}
}

// SetFlags adds flags for the common options on the FlagSet
func (opts *Options) SetFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&opts.Token,
		"token",
		"t",
		os.Getenv("GH_ACCESS_TOKEN"),
		"The Github Access Token to use")
	flags.StringVarP(&opts.File,
		"file",
		"f",
		DefaultFile,
		"Specify Config File to Load")
	flags.BoolVarP(&opts.Version,
		"version",
		"v",
		false,
		"Print version information and quit")
	flags.BoolVarP(&opts.Remove,
		"remove",
		"r",
		false,
		"Remove labels that are not present in the config file")
}

// NewGHlabelsCommand creates a cobra command for all common config operations
func NewGHlabelsCommand(version string) *cobra.Command {
	opts := NewOptions()
	var flags *pflag.FlagSet

	cmd := &cobra.Command{
		Use:           "ghlabels [OPTIONS]",
		Short:         "Define more useful labels for your Github issues.",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Version {
				fmt.Printf("ghlabels version %s\n", version)
				os.Exit(0)
				return nil
			}
			config := ghlabels.LoadConfig(opts.File)
			labeler := ghlabels.NewLabeler(opts.Token, config)

			for _, repo := range config.Repos {
				existingLabels := labeler.GetLabels(repo)
				labeler.SetLabels(existingLabels, repo, opts.Remove)
			}
			return nil
		},
	}

	flags = cmd.Flags()
	opts.SetFlags(flags)
	return cmd
}

func main() {
	rootCmd := NewGHlabelsCommand(Version)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
