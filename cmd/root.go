package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Version is the version number
	Version = "unset"

	// BuildTag set during build to git tag, if any
	BuildTag = "unset"

	// BuildSHA is the git sha set during build
	BuildSHA = "unset"

	force bool
)

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:                "cober",
		Short:              "practice CLI",
		SilenceErrors:      true,
		DisableSuggestions: false,
		Version:            fmt.Sprintf("%s (%s/%s)", Version, BuildTag, BuildSHA),
		RunE: func(cmd *cobra.Command, args []string) error {
			force, err := cmd.PersistentFlags().GetBool("force")
			if err != nil {
				return err
			}

			viper.BindPFlag("force", cmd.PersistentFlags().Lookup("force"))
			fmt.Printf("force: %v\n", viper.GetBool("force"))
			fmt.Println("force: ", force)
			return nil

		},
	}

	rootCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "force update")
	fmt.Printf("force: %v\n", force)
	return rootCmd

}

// Execute run CLI
func Execute() error {

	log.Printf("[INFO] pkg version: %s", Version)
	log.Printf("[INFO] Go runtime version: %s", runtime.Version())
	log.Printf("[INFO] Build tag/SHA: %s/%s", BuildTag, BuildSHA)
	log.Printf("[INFO] CLI args: %#v", os.Args)

	defer log.Printf("[DEBUG] root command execution finished")

	rootCmd := newRootCmd()
	return rootCmd.Execute()
}
