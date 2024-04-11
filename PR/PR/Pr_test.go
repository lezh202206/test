package GitPR

import (
	"testing"
)

func TestPR(t *testing.T) {
	//确保当前目录是一个Git仓库
	//repoPath := "/Users/lezh/Desktop/GO/src/miniblog"
	//
	//// 检查是否是Git仓库
	//stdout := GitBase(repoPath, []string{"rev-parse", "--git-dir"})
	//fmt.Println(stdout.String())
	//
	//MyGit := NewMyGit(repoPath)
	//MyGit.EchoBranch()

	Url := NewGiteePR().sendGitee("test", "xxx", "master")
	NewGiteePR().sendBot(Url)
}
