package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	GitPR "test/PR/PR"
)

func main() {
	var (
		title string
		audit string
		flag  string
	)

	var rootCmd = &cobra.Command{
		Use:   "PR-GO",
		Short: "A CLI app with parameters using Cobra",
		Run: func(cmd *cobra.Command, args []string) {
			run(flag, title, audit)
		},
	}

	rootCmd.Flags().StringVarP(&flag, "flag", "f", "PR", "hotfix｜PR")
	rootCmd.Flags().StringVarP(&title, "title", "t", "", "Set PR title")
	rootCmd.Flags().StringVarP(&audit, "audit", "a", "", "Add AuditName")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(flag, title, audit string) {
	//确保当前目录是一个Git仓库
	repoPath := "/Users/lezh/Desktop/GO/src/miniblog"
	MyGit := GitPR.NewMyGit(repoPath)

	if flag == "PR" || flag == "pr" {
		PR(MyGit, title, audit)
	}
	if flag == "hotfix" {
		hotfix(MyGit, title, audit)
	}
}

func PR(MyGit GitPR.MyGit, title, audit string) {
	branch := MyGit.EchoBranch()
	MyGit.Pull()

	newBranch := MyGit.CreateBranch()
	MyGit.PushBranch(newBranch)

	GiteePR := GitPR.NewGiteePR()
	Url := GiteePR.SendGitee(title, branch, newBranch, audit)
	GiteePR.SendBot(fmt.Sprintf("PR: \n %s分支 合并到 %s分支 \n %s ", branch, newBranch, Url))
}

func hotfix(MyGit GitPR.MyGit, title, audit string) {
	branch := MyGit.EchoBranch()
	MyGit.Pull()

	branchList := MyGit.EchoBranchList()

	fmt.Print("请输入要合并的分支: ")
	reader := bufio.NewReader(os.Stdin)
	branchName, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("读取输入失败"))
	}
	branchName = strings.TrimSpace(branchName)
	MyGit.BranchExists(strings.Split(branchList, "\n"), branchName)

	GiteePR := GitPR.NewGiteePR()
	Url := GiteePR.SendGitee(title, branch, branchName, audit)
	GiteePR.SendBot(fmt.Sprintf("PR: \n %s分支 合并到 %s分支 \n %s ", branch, branchName, Url))
}
