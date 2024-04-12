package main

import (
	"fmt"
	"os"
	GitPR "test/PR/PR"

	"github.com/spf13/cobra"
)

func main() {
	var (
		title string
		audit string
	)

	var rootCmd = &cobra.Command{
		Use:   "PR-GO",
		Short: "A CLI app with parameters using Cobra",
		Run: func(cmd *cobra.Command, args []string) {
			run(title, audit)
		},
	}

	rootCmd.Flags().StringVarP(&title, "title", "t", "", "Set PR title")
	rootCmd.Flags().StringVarP(&audit, "audit", "a", "", "Set PR auditName")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(title, audit string) {
	//确保当前目录是一个Git仓库
	repoPath := "/Users/lezh/Desktop/GO/src/miniblog"
	MyGit := GitPR.NewMyGit(repoPath)

	branch := MyGit.EchoBranch()
	MyGit.Pull()

	newBranch := MyGit.CreateBranch()
	MyGit.PushBranch(newBranch)

	GiteePR := GitPR.NewGiteePR()
	Url := GiteePR.SendGitee(title, branch, newBranch, audit)
	fmt.Println(Url)
	// lzh123789
}
