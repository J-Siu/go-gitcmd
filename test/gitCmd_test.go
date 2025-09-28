package gitcmd

import (
	"fmt"
	"path"
	"testing"

	"github.com/J-Siu/go-gitcmd"
	"github.com/J-Siu/go-helper/v2/cmd"
	"github.com/J-Siu/go-helper/v2/ezlog"
)

func TestGitRoot_hasSubmodule(t *testing.T) {
	// Debug = true

	// Pull submodules
	var args []string
	var myCmd *cmd.Cmd
	title := "myCmd"
	//		git submodule add -f https://github.com/J-Siu/submodule_test_1
	args = []string{"submodule", "add", "-f", "https://github.com/J-Siu/submodule_test_1"}
	myCmd = cmd.Run("git", &args, nil)
	ezlog.Log().NameLn(title).Msg(myCmd)
	if myCmd.Stderr.Len() > 0 {
		ezlog.Log().Name(title).Name("Stderr").Msg(myCmd.Stderr).Out()
	}
	if myCmd.Stdout.Len() > 0 {
		ezlog.Log().Name(title).Name("Stdout").Msg(myCmd.Stdout).Out()
	}

	//   git submodule update --recursive --init
	args = []string{"submodule", "update", "--recursive", "--init"}
	myCmd = cmd.Run("git", &args, nil)
	ezlog.Log().NameLn(title).Msg(myCmd)
	if myCmd.Stderr.Len() > 0 {
		ezlog.Log().Name(title).Name("Stderr").Msg(myCmd.Stderr).Out()
	}
	if myCmd.Stdout.Len() > 0 {
		ezlog.Log().Name(title).Name("Stdout").Msg(myCmd.Stdout).Out()
	}

	// Test
	//		testDir should be set to a git submodule directory.
	var testDir string = "submodule_test_1/submodule_test_2/submodule_test_3/submodule_test_4/submodule_test_5"
	var wanted string = "go-gitcmd"
	var msg string = path.Base(gitcmd.GitRoot(&testDir))

	// Clean up
	//		git rm -rf submodule_test_1 ../.gitmodules
	args = []string{"rm", "-rf", "submodule_test_1", "../.gitmodules"}
	myCmd = cmd.Run("git", &args, nil)

	// Result
	fmt.Println("Git root of " + testDir + " is `" + msg + "`")
	if msg != wanted {
		t.Fatalf(`GitRoot(rtSp(%s) = ""`, testDir)
	}
}

func TestGitRoot_fail(t *testing.T) {
	// Assuming root directory is not a git directory
	var testDir string = "/"
	var msg string = gitcmd.GitRoot(&testDir)
	fmt.Println("Git root of " + testDir + " is `" + msg + "`")
	if msg != "" {
		t.Fatalf(`GitRoot(rtSp(%s) = ""`, testDir)
	}
}
