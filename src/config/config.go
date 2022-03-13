package config

import (
	"flag"
	"os"
)

type Config struct {
	NfdPrefix string
	GpuPrefix string
	Namespace string
}

var GlobalConfig Config

func printHelpAndExit() {
	flag.CommandLine.Usage()
	os.Exit(0)
}

func ProcessArgs() {
	ret := &GlobalConfig
	flag.StringVar(&ret.NfdPrefix, "nfd-csv-prefix", "node-feature-discovery-operator", "Prefix of nfd csv file")
	flag.StringVar(&ret.GpuPrefix, "gpu-csv-prefix", "nvidia-gpu-addon", "Prefix of gpu operator csv file")
	flag.StringVar(&ret.Namespace, "namespace", "redhat-nvidia-gpu", "Namespace")

	h := flag.Bool("help", false, "Help message")
	flag.Parse()
	if h != nil && *h {
		printHelpAndExit()
	}
}
