package main

import (
	"fmt"
	GitPR "test/PR/PR"
)

func main() {
	//确保当前目录是一个Git仓库
	repoPath := "/Users/lezh/Desktop/GO/src/miniblog"
	MyGit := GitPR.NewMyGit(repoPath)

	branch := MyGit.EchoBranch()
	MyGit.Pull()

	newBranch := MyGit.CreateBranch()
	MyGit.PushBranch(newBranch)

	GiteePR := GitPR.NewGiteePR()
	Url := GiteePR.SendGitee("test", branch, newBranch)
	fmt.Println(Url)
	//GiteePR.sendBot(fmt.Sprintf("PR: \n %s分支 合并到 %s分支 \n %s ", branch, newBranch, Url))
}
