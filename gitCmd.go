/*
The MIT License

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package gitcmd

import (
	"os/exec"
	"strings"

	"github.com/J-Siu/go-helper/v2/cmd"
	"github.com/J-Siu/go-helper/v2/file"
	"github.com/J-Siu/go-helper/v2/str"
)

// Run "git <optionsP>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func Git(workPathP *string, optionsP *[]string) *cmd.Cmd {
	return cmd.Run("git", optionsP, workPathP)
}

// Run "git clone <optionsP>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func GitClone(workPathP *string, optionsP *[]string) *cmd.Cmd {
	args := []string{"clone"}
	if optionsP != nil {
		args = append(args, *optionsP...)
	}
	return Git(workPathP, &args)
}

// Check git executable exist.
func GitExecExist() bool {
	return GitExecPath() != ""
}

// Get git executable path.
//   - Return empty string if not found.
func GitExecPath() string {
	path, err := exec.LookPath("git")
	if err != nil {
		return ""
	}
	return path
}

// Run "git init".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func GitInit(workPathP *string) *cmd.Cmd {
	return Git(workPathP, &[]string{"init"})
}

// Run "git branch --show-current".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func GitBranchCurrent(workPathP *string) *cmd.Cmd {
	return Git(workPathP, &[]string{"branch", "--show-current"})
}

// Run "git pull <optionsP>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func GitPull(workPathP *string, optionsP *[]string) *cmd.Cmd {
	args := []string{"pull"}
	if optionsP != nil {
		args = append(args, *optionsP...)
	}
	return Git(workPathP, &args)
}

// Run "git push <optionsP>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func GitPush(workPathP *string, optionsP *[]string) *cmd.Cmd {
	args := []string{"push"}
	if optionsP != nil {
		args = append(args, *optionsP...)
	}
	return Git(workPathP, &args)

}

// Run "git remote".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return remotes in string array.
func GitRemote(workPathP *string, v bool) *[]string {
	var args []string
	if v {
		args = []string{"remote", "-v"}
	} else {
		args = []string{"remote"}
	}
	output := cmd.Run("git", &args, workPathP).Stdout.String()
	return str.LnSplit(&output)
}

// Run "git remote add <name> <git>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func GitRemoteAdd(workPathP *string, name string, git string) *cmd.Cmd {
	return Git(workPathP, &[]string{"remote", "add", name, git})
}

// Check if a git remote(by name) exist in workPath.
//   - If <workPathP> is empty/nil, current directory is used.
func GitRemoteExist(workPathP *string, name string) bool {
	r := GitRemote(workPathP, false)
	return str.ArrayContains(r, &name)
}

// Run "git remote remove".
//   - If <workPathP> is empty/nil, current directory is used.
//   - If remote exist Return a cmd.Cmd pointer for execution information.
//   - If remote does not exit, return nil.(Nothing to remove)
func GitRemoteRemove(workPathP *string, name string) *cmd.Cmd {
	return Git(workPathP, &[]string{"remote", "remove", name})
}

// Run "git remote remove" all git remotes
//   - If <workPathP> is empty/nil, current directory is used.
func GitRemoteRemoveAll(workPathP *string) {
	gr := GitRemote(workPathP, false)
	for _, r := range *gr {
		GitRemoteRemove(workPathP, r)
	}
}

// Get git root from current directory.
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return empty string if not a git dir.
func GitRoot(workPathP *string) string {
	if *workPathP == "" {
		file.CurrentPath()
	}
	// Check submodule path repeatedly
	var submodulePath string = *workPathP
	var currentPath string = *workPathP
	for submodulePath != "" {
		submodulePath = GitRootSubmodule(&submodulePath)
		if submodulePath != "" {
			currentPath = submodulePath
		}
	}
	// Check git root
	var opts []string = []string{"rev-parse", "--show-toplevel"}
	var c = cmd.Run("git", &opts, &currentPath)
	if c.Err != nil {
		return ""
	}
	return strings.TrimSpace(c.Stdout.String())
}

// Get git submodule root from `workPath`.
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return empty string if not a submodule dir.
func GitRootSubmodule(workPathP *string) string {
	var opts []string = []string{"rev-parse", "--show-superproject-working-tree"}
	var c = cmd.Run("git", &opts, workPathP)
	if c.Err != nil {
		return ""
	}
	return strings.TrimSpace(c.Stdout.String())
}
