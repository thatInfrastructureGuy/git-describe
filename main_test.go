package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestGenerateVersionFile(t *testing.T) {
	t.Parallel()
	buf := bytes.Buffer{}
	pkgName := "testpkg"
	variableName := "testVar"
	generateVersionFile(&buf, pkgName, variableName)
	if !strings.Contains(buf.String(), "package "+pkgName) {
		t.Errorf("Package name is not set to %v\n", pkgName)
	}

	if !strings.Contains(buf.String(), "const "+variableName) {
		t.Errorf("Constant variable is not set to %v\n", variableName)
	}
}

func TestCreate(t *testing.T) {
	t.Parallel()
	folder := "/tmp/xxyyzz"
	nonExistent := folder + "/d/a.txt"
	f, err := create(nonExistent)
	if err != nil {
		t.Errorf("Error creating file '%v': %v", nonExistent, err)
	}
	err = f.Close()
	if err != nil {
		t.Errorf("Error closing file '%v': %v", nonExistent, err)
	}
	_, err = os.ReadFile(nonExistent)
	if err != nil {
		t.Errorf("Error reading file '%v': %v", nonExistent, err)
	}

	err = os.RemoveAll(folder)
	if err != nil {
		t.Errorf("Error deleting folder '%v': %v", folder, err)
	}
}
