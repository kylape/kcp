/*
Copyright 2022 The KCP Authors.

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
	"os"

	"github.com/spf13/pflag"

	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/kcp-dev/kcp/pkg/cliplugins/workspace/cmd"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-kcp-workspace", pflag.ExitOnError)
	pflag.CommandLine = flags

	root, err := cmd.NewCmdWorkspace(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err != nil {
		os.Exit(1)
	}
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
