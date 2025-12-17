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

// Provide git command using exec.Cmd wrapper go-helper/v2/cmd
package gitcmd

import (
	"os/exec"
	"strings"

	"github.com/J-Siu/go-helper/v2/cmd"
	"github.com/J-Siu/go-helper/v2/file"
	"github.com/J-Siu/go-helper/v2/str"
)

type GitCmd struct {
	*cmd.Cmd
	workPathP *string
}

func (t *GitCmd) New(workPathP *string) *GitCmd {
	t.Cmd = new(cmd.Cmd)
	t.workPathP = workPathP
	return t
}

// Run "git clone <optionsP>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) Clone(optionsP *[]string) *cmd.Cmd {
	args := []string{"clone"}
	if optionsP != nil {
		args = append(args, *optionsP...)
	}
	return t.run(&args)
}

// Check git executable exist.
func (t *GitCmd) ExecExist() bool {
	return t.ExecPath() != ""
}

// Get git executable path.
//   - Return empty string if not found.
func (t *GitCmd) ExecPath() string {
	path, err := exec.LookPath("git")
	if err != nil {
		return ""
	}
	return path
}

// Run "git init".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) Init() *cmd.Cmd {
	return t.run(&[]string{"init"})
}

// Run "git branch --show-current".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) BranchCurrent() *cmd.Cmd {
	return t.run(&[]string{"branch", "--show-current"})
}

// Run "git pull <optionsP>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) Pull(optionsP *[]string) *cmd.Cmd {
	args := []string{"pull"}
	if optionsP != nil {
		args = append(args, *optionsP...)
	}
	return t.run(&args)
}

// Run "git push <optionsP>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) Push(optionsP *[]string) *cmd.Cmd {
	args := []string{"push"}
	if optionsP != nil {
		args = append(args, *optionsP...)
	}
	return t.run(&args)

}

// Run "git remote".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return remotes in string array.
func (t *GitCmd) Remote(v bool) *[]string {
	var args []string
	if v {
		args = []string{"remote", "-v"}
	} else {
		args = []string{"remote"}
	}
	output := t.run(&args).Stdout.String()
	return str.LnSplit(&output)
}

// Run "git remote add <name> <git>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) RemoteAdd(name string, gitUrl string) *cmd.Cmd {
	return t.run(&[]string{"remote", "add", name, gitUrl})
}

// Check if a git remote(by name) exist in workPathP.
//   - If <workPathP> is empty/nil, current directory is used.
func (t *GitCmd) RemoteExist(name string) bool {
	r := t.Remote(false)
	return str.ArrayContains(r, &name, false)
}

// Run "git remote remove".
//   - If <workPathP> is empty/nil, current directory is used.
//   - If remote exist Return a cmd.Cmd pointer for execution information.
//   - If remote does not exit, return nil.(Nothing to remove)
func (t *GitCmd) RemoteRemove(name string) *cmd.Cmd {
	return t.run(&[]string{"remote", "remove", name})
}

// Run "git remote remove" all git remotes
//   - If <workPathP> is empty/nil, current directory is used.
func (t *GitCmd) RemoteRemoveAll() {
	gr := t.Remote(false)
	for _, r := range *gr {
		t.RemoteRemove(r)
	}
}

// Get git root from current directory.
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return empty string if not a git dir.
func (t *GitCmd) Root() string {
	var (
		cmd  *cmd.Cmd
		git  = new(GitCmd)
		opts = []string{"rev-parse", "--show-toplevel"}

		currentPath   string
		submodulePath string
		workPathP     = t.workPathP
	)
	if *t.workPathP == "" {
		workPathP = file.CurrentPath()
	}
	// Check submodule path repeatedly
	submodulePath = *workPathP
	currentPath = *workPathP
	for submodulePath != "" {
		git.New(&submodulePath)
		submodulePath = git.RootSubmodule()
		if submodulePath != "" {
			currentPath = submodulePath
		}
	}
	// Check git root
	cmd = git.New(&currentPath).run(&opts)
	if cmd.Err != nil {
		return ""
	}
	return strings.TrimSpace(cmd.Stdout.String())
}

// Get git submodule root from `workPathP`.
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return empty string if not a submodule dir.
func (t *GitCmd) RootSubmodule() string {
	var (
		opts = []string{"rev-parse", "--show-superproject-working-tree"}
		cmd  = t.run(&opts)
	)
	if cmd.Err != nil {
		return ""
	}
	return strings.TrimSpace(cmd.Stdout.String())
}

// run "git <optionsP>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) run(optionsP *[]string) *cmd.Cmd {
	return t.Cmd.New("git", optionsP, t.workPathP).Run()
}
