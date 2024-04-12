package GitPR

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestPR(t *testing.T) {
	//确保当前目录是一个Git仓库
	repoPath := "/Users/lezh/Desktop/GO/src/miniblog"

	MyGit := NewMyGit(repoPath)

	branch := MyGit.EchoBranch()
	MyGit.Pull()

	branchList := MyGit.EchoBranchList()
	fmt.Print("请输入要合并的分支: ")
	reader := bufio.NewReader(os.Stdin)
	branchName, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("读取输入失败"))
	}

	MyGit.BranchExists(strings.Split(branchList, "\n"), branchName)

	fmt.Println(branch)
	fmt.Println(branchList)

	GiteePR := NewGiteePR()
	Url := GiteePR.SendGitee("hotfix", branch, branchName, "lzh123789")
	fmt.Printf(Url)
}
