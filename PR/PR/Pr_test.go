package GitPR

import (
	"fmt"
	"testing"
)

func TestPR(t *testing.T) {
	//确保当前目录是一个Git仓库
	repoPath := "/Users/lezh/Desktop/GO/src/miniblog"

	MyGit := NewMyGit(repoPath)
	branch := MyGit.EchoBranch()

	newBranch := MyGit.CreateBranch()
	MyGit.PushBranch(newBranch)
	fmt.Println(branch)

	GiteePR := NewGiteePR()
	Url := GiteePR.SendGitee("test", branch, newBranch)
	fmt.Println(Url)
	GiteePR.sendBot(fmt.Sprintf("PR: \n %s分支 合并到 %s分支 \n %s ", branch, newBranch, Url))
}
