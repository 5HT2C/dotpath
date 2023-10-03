<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

# godotpath-shell
Set your shell's PATH from a file, preserving uniqueness.

Why? So you can easily configure your shell's PATH and remove duplicate entries, quickly & without fuss.

## Usage

```bash
make
eval "$(./dotpath)"
# Test that it worked
echo $PATH | tr ':' '\n' | sort

# Optional: Add dotpath to your path
mkdir -p ~/.local/bin
cp dotpath ~/.local/bin/
```

```
Usage of ./dotpath:
  -allowInvalid
        Add invalid directories to PATH
  -allowMissing
        Add non-existent directories to PATH (default true)
  -file string
        The location of your paths file (default "~/.paths")
  -path string
        PATH to add onto (default "$PATH")
```

## Bash / ZSH / Fish / etc

Add `eval "$(./dotpath)"` to your `~/.bashrc` / whatever your shell's config is called.

## Configuration

Create a file called `~/.paths` (or any other name, this is just the default).

Each path has it's own line:
```
$HOME/.local/bin
$HOME/.local/bin/gradle/bin
$HOME/.cargo/bin
$HOME/.android/sdk/tools/bin
$HOME/.android/sdk/cmdline-tools/latest/bin
$GOPATH/bin
```
