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
	options  []string
	runArgs  []string
	workPath string
}

func (t *GitCmd) New(workPath string) *GitCmd {
	t.Cmd = new(cmd.Cmd)
	t.options = nil
	t.runArgs = nil
	t.workPath = workPath
	return t
}

// Run "git <optionsP>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) Run() *cmd.Cmd {
	t.runArgs = append(t.runArgs, t.options...)
	return t.Cmd.New("git", &t.runArgs, &t.workPath).Run()
}

// -- cmd need Run()

// Run "git clone <optionsP>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) Clone(options []string) *GitCmd {
	t.runArgs = []string{"clone"}
	t.options = options
	return t
}

// Run "git init".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) Init() *GitCmd {
	t.runArgs = []string{"init"}
	return t
}

// Run "git branch --show-current".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) BranchCurrent() *GitCmd {
	t.runArgs = []string{"branch"}
	t.options = []string{"--show-current"}
	return t
}

// Run "git pull <optionsP>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) Pull(options []string) *GitCmd {
	t.runArgs = []string{"pull"}
	t.options = options
	return t
}

// Run "git push <optionsP>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) Push(options []string) *GitCmd {
	t.runArgs = []string{"push"}
	t.options = options
	return t

}

// Run "git remote add <name> <git>".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func (t *GitCmd) RemoteAdd(name string, gitUrl string) *GitCmd {
	t.runArgs = []string{"remote"}
	t.options = []string{"add", name, gitUrl}
	return t
}

func (t *GitCmd) RevParse(options []string) *GitCmd {
	t.runArgs = []string{"rev-parse"}
	t.options = options
	return t
}

func (t *GitCmd) Tag(options []string) *GitCmd {
	t.runArgs = []string{"tag"}
	t.options = options
	return t
}

// -- Utils

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

// Run "git remote".
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return remotes in string array.
func (t *GitCmd) Remote(v bool) *[]string {
	t.runArgs = []string{"remote"}
	if v {
		t.options = []string{"-v"}
	}
	output := t.Run().Stdout.String()
	return str.LnSplit(&output)
}

// Check if a git remote(by name) exist in workPathP.
//   - If <workPathP> is empty/nil, current directory is used.
func (t *GitCmd) RemoteExist(name string) bool {
	r := t.Remote(false)
	return str.ArrayContains(r, name, false)
}

// Run "git remote remove".
//   - If <workPathP> is empty/nil, current directory is used.
//   - If remote exist Return a cmd.Cmd pointer for execution information.
//   - If remote does not exit, return nil.(Nothing to remove)
func (t *GitCmd) RemoteRemove(name string) *GitCmd {
	t.runArgs = []string{"remote"}
	t.options = []string{"remove", name}
	return t
}

// Run "git remote remove" all git remotes
//   - If <workPathP> is empty/nil, current directory is used.
func (t *GitCmd) RemoteRemoveAll() {
	gr := t.Remote(false)
	for _, r := range *gr {
		t.RemoteRemove(r).Run()
	}
}

// Get git root from current directory.
//   - If <workPathP> is empty/nil, current directory is used.
//   - Return empty string if not a git dir.
func (t *GitCmd) Root() string {
	var (
		cmd           *cmd.Cmd
		git           = new(GitCmd)
		options       = []string{"--show-toplevel"}
		currentPath   string
		submodulePath string
		workPath      = t.workPath
	)
	if workPath == "" {
		workPath = *file.CurrentPath()
	}
	// Check submodule path repeatedly
	submodulePath = workPath
	currentPath = workPath
	for submodulePath != "" {
		git.New(submodulePath)
		submodulePath = git.RootSubmodule()
		if submodulePath != "" {
			currentPath = submodulePath
		}
	}
	// Check git root
	cmd = git.New(currentPath).RevParse(options).Run()
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
		options = []string{"--show-superproject-working-tree"}
		cmd     = t.RevParse(options).Run()
	)
	if cmd.Err != nil {
		return ""
	}
	return strings.TrimSpace(cmd.Stdout.String())
}

func (t *GitCmd) TagList() (out []string) {
	var (
		cmd = t.Tag(nil).Run()
	)
	if cmd.Err == nil {
		for _, l := range strings.Split(strings.Trim(t.Stdout.String(), "\n"), "\n") {
			l = strings.TrimSpace(l)
			if l != "" {
				out = append(out, l)
			}
		}
	}
	return out
}
