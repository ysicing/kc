// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ysicing/go-utils/extime"
	"runtime"
)

var (
	Version   string
	BuildDate string
	GitSHA    string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Run: func(cmd *cobra.Command, args []string) {
		if len(Version) == 0 {
			Version = fmt.Sprintf("v%v", extime.GetToday())
		}
		if len(BuildDate) == 0 {
			BuildDate = extime.NowFormat()
		}
		if len(GitSHA) == 0 {
			GitSHA = "12345678"
		}
		fmt.Printf("--------\nkc version: %v \nGit Sha: %v \nBuildDate: %v \nGo version: %v \nGo Os/Arch: %v/%v\n--------\n",
			Version, GitSHA, BuildDate, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	},
}
