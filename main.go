/*
 * Copyright 2022 Kulkarni, Ashish <thatInfrastructureGuy@gmail.com>
 * Author: Ashish Kulkarni
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

//go:generate go run main.go

package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/thatInfrastructureGuy/generate-version/version"
)

const (
	defaultFilename     = "version/const.go"
	defaultPackageName  = "version"
	defaultVariableName = "Version"
)

func main() {
	fileName, packageName, variableName := parseFlags()

	f, err := create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	generateVersionFile(f, packageName, variableName)
}

func generateVersionFile(f *os.File, packageName, variableName string) {
	packageTemplate := template.Must(template.New("").Parse(`// THIS IS GENERATED CODE; DO NOT EDIT.
// Generated at {{ .Timestamp }}
package {{.PackageName}}

const {{.VariableName}} = "{{ .Version }}"
`))

	packageTemplate.Execute(f, struct {
		Timestamp    time.Time
		PackageName  string
		VariableName string
		Version      string
	}{
		Timestamp:    time.Now(),
		PackageName:  packageName,
		VariableName: variableName,
		Version:      version.GetVersion(),
	})
}

func create(filename string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return nil, err
	}
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func parseFlags() (string, string, string) {
	filePtr := flag.String("filepath", defaultFilename, "Filepath where contents should get generated. It should be a .go file.\nDefault: "+defaultFilename)
	packageNamePtr := flag.String("package", defaultPackageName, "Name of the package where file should be generated.\nDefault: "+defaultPackageName)
	variablePtr := flag.String("variable", defaultVariableName, "Variable name used to store version.\nDefault: "+defaultVariableName)

	return *filePtr, *packageNamePtr, *variablePtr
}
