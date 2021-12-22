package cmd

import (
	"github.com/spf13/pflag"
	"os"
)

func main() {

	// parse Args
	flags := pflag.NewFlagSet("", pflag.ExitOnError)
	flags.Parse(os.Args)

	// Init K8s Client

}
