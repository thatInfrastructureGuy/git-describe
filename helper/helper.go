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

// Helper package contains useful functions for implementation abstraction.
package helper

import (
	"log"
	"os/exec"
	"strings"
)

// ExecCommand runs shell commands.
// It takes string as input. The binary should be present in the path.
// Command returns "UNKNOWN" if the binary errors out.
func ExecCommand(cmd string) string {
	cmdWithFlags := strings.Fields(cmd)

	cli, err := exec.LookPath(cmdWithFlags[0])
	if err != nil {
		log.Println("Error finding binary:", err)
		return "UNKNOWN"
	}

	output, err := exec.Command(cli, cmdWithFlags[1:]...).Output()
	if err != nil {
		log.Println("Error executing binary:", err)
		return "UNKNOWN"
	}

	return strings.Trim(string(output), " \n\t\r")
}
