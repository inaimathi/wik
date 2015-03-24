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

func (w *Wiki) Raw(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func (w *Wiki) Render(path string) ([]byte, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	unsafe := blackfriday.MarkdownCommon(body)
	return bluemonday.UGCPolicy().SanitizeBytes(unsafe), nil
}

func (w *Wiki) Local(path string) (string, error) {
	p := filepath.Clean(filepath.Join(w.Path, path))
	if (strings.HasPrefix(p, w.Path)) {
		return p, nil
	}
	return "", errors.New("path outside of repo")
}

func (w *Wiki) Create(path string) error {
	p, err := w.Local(path)
	if (err != nil) { return err }
	return ioutil.WriteFile(p, []byte("# " + path), 0600)
}

func (w *Wiki) Edit(path string, contents []byte) error {
	p, err := w.Local(path)
	if (err != nil) { return err }
	return ioutil.WriteFile(p, contents, 0600)
}

func (w *Wiki) Remove(path string) error {
	p, err := w.Local(path)
	if (err != nil) { return err }
	os.Remove(p)
	return nil
}

func (w *Wiki) ExecIn(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = w.Path
	return cmd.Run()
}

func (w *Wiki) Initialize() error {
	return w.ExecIn("git", "init")
}

func (w *Wiki) Commit(path string, message string) error {
	w.ExecIn("git", "add", "--all", path)
	w.ExecIn("git", "commit", "-m", message)
	return nil
}
