package main

import (
	"flag"
	"os"
)

type args struct {
	appType string
	appName string
	saveDir string
}

func parseArgs() *args {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	appType := os.Args[1]
	appName := os.Args[2]

	saveDir := flag.String("save", ".", "Directory to save new application to")

	return &args{
		appType: appType,
		appName: appName,
		saveDir: *saveDir,
	}
}
