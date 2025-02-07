/*
Copyright AppsCode Inc. and Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"

	"sigs.k8s.io/yaml"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: yfmt a.yaml")
		os.Exit(1)
	}

	filename := os.Args[1]
	cur, err := read(filename)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to read file %q: %v", filename, err)
		os.Exit(1)
	}
	data, err := yaml.Marshal(cur)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to marshal: %v", err)
		os.Exit(1)
	}
	err = os.WriteFile(filename, data, 0o644)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to write file %q: %v", filename, err)
		os.Exit(1)
	}
}

func read(filename string) (map[string]any, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var obj map[string]any
	if err := yaml.Unmarshal(data, &obj); err != nil {
		return nil, err
	}
	return obj, nil
}
