package main

import (
	GitPR "test/PR/PR"
)

func main() {
	//确保当前目录是一个Git仓库
	repoPath := "/Users/lezh/Desktop/GO/src/miniblog"

	// 检查是否是Git仓库
	GitPR.GitBase(repoPath, []string{"rev-parse", "--git-dir"})

	MyGit := GitPR.NewMyGit(repoPath)
	//MyGit.Pull()
	MyGit.EchoBranch()
	MyGit.CreateBranch()
	MyGit.EchoBranch()
}
