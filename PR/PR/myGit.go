package GitPR

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type MyGit struct {
	rootPath string
}

func NewMyGit(rootPath string) MyGit {
	GitBase(rootPath, []string{"rev-parse", "--git-dir"})
	return MyGit{rootPath: rootPath}
}

func GitBase(rootPath string, args []string) strings.Builder {
	// 执行 Git 命令获取当前分支
	cmd := exec.Command("git", args...)
	cmd.Dir = rootPath // 替换为实际的 Git 仓库路径

	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Git command failed %s AND Git stderr %s: :", err, stderr.String()))
	}
	return stdout
}

func (mGit MyGit) Pull() {
	stdout := GitBase(mGit.rootPath, []string{"pull"})
	fmt.Println("Git 分支拉取状态:", stdout.String())
}

func (mGit MyGit) CreateBranch() string {
	fmt.Print("请输入要创建的新分支名称: ")
	reader := bufio.NewReader(os.Stdin)
	branchName, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// 清理输入的分支名称，去除换行符和空格
	branchName = strings.TrimSpace(branchName)

	stdout := GitBase(mGit.rootPath, []string{"checkout", "-b", branchName, "master"})
	// 打印 Git 命令的输出信息
	fmt.Println("Git stdout:", stdout.String())
	fmt.Printf("已成功创建并切换到分支 %s\n", branchName)
	return branchName
}

func (mGit MyGit) PushBranch(branchName string) {
	stdout := GitBase(mGit.rootPath, []string{"push", "--set-upstream", "origin", branchName})
	// 打印 Git 命令的输出信息
	fmt.Println("Git stdout:", stdout.String())
	fmt.Printf("已成功推送到远程 %s\n", branchName)
}

func (mGit MyGit) EchoBranch() string {
	stdout := GitBase(mGit.rootPath, []string{"rev-parse", "--abbrev-ref", "HEAD"})
	// 打印当前分支信息
	fmt.Printf("当前所在分支: %s\n", strings.TrimSpace(stdout.String()))
	return strings.TrimSpace(stdout.String())
}

func (mGit MyGit) EchoBranchList() string {
	stdout := GitBase(mGit.rootPath, []string{"branch", "--list"})
	return stdout.String()
}

func (mGit MyGit) BranchExists(branchList []string, branchName string) {
	for _, branch := range branchList {
		// 去除分支名中的空格和星号
		trimmedBranch := strings.TrimSpace(strings.TrimPrefix(branch, "*"))
		if trimmedBranch == strings.TrimSpace(branchName) {
			return
		}
	}
	panic("当前输入分支不存在")
}
