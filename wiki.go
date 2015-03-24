package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"io/ioutil"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type Wiki struct {
	Path string
}

////////// Mutating operations

// Create creates a new file in the given wiki
// TODO - create all intervening directories
func (w *Wiki) Create(path string) error {
	p, err := w.Local(path)
	if (err != nil) { return err }
	err = ioutil.WriteFile(p, []byte("# " + path), 0600)
	if (err != nil) { return err }
	return w.Commit(p, "Created " + path)
}

// Edit changes the contents of a file in the given wiki
func (w *Wiki) Edit(path string, contents []byte) error {
	p, err := w.Local(path)
	if (err != nil) { return err }
	err = ioutil.WriteFile(p, contents, 0600)
	if (err != nil) { return err }
	return w.Commit(p, "Edit to " + path)
}

// Remove removes a file in the given wiki
// TODO - remove the containing directory if empty
func (w *Wiki) Remove(path string) error {
	p, err := w.Local(path)
	if (err != nil) { return err }
	err = os.Remove(p)
	if (err != nil) { return err }
	return w.Commit(p, "Deleted " + path)
}

////////// Viewing operations

// Raw takes a path to a file in the given wiki and
// returns its raw contents.
func (w *Wiki) Raw(path string) ([]byte, error) {
	p, err := w.Local(path)
	if (err != nil) { return nil, err }
	return ioutil.ReadFile(p)
}

// Render takes a path to a file in the given wiki and
// returns its contents passed through a markdown converter
// followed by an XSS sanitizer.
func (w *Wiki) Render(path string) ([]byte, error) {
	p, err := w.Local(path)
	if (err != nil) { return nil, err }
	body, err := ioutil.ReadFile(p)
	if (err != nil) { return nil, err }
	unsafe := blackfriday.MarkdownCommon(body)
	return bluemonday.UGCPolicy().SanitizeBytes(unsafe), nil
}

////////// Git commands and various utility

// Initialize runs git-init in the directory of the given wiki
func (w *Wiki) Initialize() error {
	return w.ExecIn("git", "init")
}

// Commit runs a git-add/git-commit with the given message and file
func (w *Wiki) Commit(path string, message string) error {
	w.ExecIn("git", "add", "--all", path)
	w.ExecIn("git", "commit", "-m", message)
	return nil
}

// ExecIn executes a command with the wiki directory as CWD.
func (w *Wiki) ExecIn(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = w.Path
	return cmd.Run()
}

// Local takes a path and checks if it would fall within the given
// repo if joined with it. Returns either 
//   [sanitized path], nil    // if the given path is valid
//   "", error                // otherwise
// TODO exclude files present in the .git subdirectory
func (w *Wiki) Local(path string) (string, error) {
	p := filepath.Clean(filepath.Join(w.Path, path))
	if (strings.HasPrefix(p, w.Path)) {
		return p, nil
	}
	return "", errors.New("path outside of repo")
}
