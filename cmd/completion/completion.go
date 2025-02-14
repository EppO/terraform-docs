/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package completion

import (
	"github.com/spf13/cobra"

	"github.com/terraform-docs/terraform-docs/cmd/completion/bash"
	"github.com/terraform-docs/terraform-docs/cmd/completion/fish"
	"github.com/terraform-docs/terraform-docs/cmd/completion/zsh"
)

// NewCommand returns a new cobra.Command for 'completion' command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "completion SHELL",
		Short: "Generate shell completion code for the specified shell (bash, zsh, fish)",
		Long:  longDescription,
	}

	// subcommands
	cmd.AddCommand(bash.NewCommand())
	cmd.AddCommand(zsh.NewCommand())
	cmd.AddCommand(fish.NewCommand())

	return cmd
}

const longDescription = `Outputs terraform-doc shell completion for the given shell (bash, zsh, fish)
This depends on the bash-completion binary.  Example installation instructions:
# for bash users
	$ terraform-doc completion bash > ~/.terraform-doc-completion
	$ source ~/.terraform-doc-completion

# for zsh users
	% terraform-doc completion zsh > /usr/local/share/zsh/site-functions/_terraform-doc
	% autoload -U compinit && compinit
# or if zsh-completion is installed via homebrew
    % terraform-doc completion zsh > "${fpath[1]}/_terraform-doc"

# for fish users
	$ terraform-doc completion fish > ~/.config/fish/completions/terraform-docs.fish

Additionally, you may want to output the completion to a file and source in your .bashrc
Note for zsh users: [1] zsh completions are only supported in versions of zsh >= 5.2
`
