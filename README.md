# gen_pass: CLI for generating password

## Usage

    gen_pass [-i] <length> [-S string|-s]
      -i int     password length. (default 16)
      -S string  generate password with specified symbols.
      -s         generate password without all symbols.

## Install

### From local repository

    % go build -o gen_pass main.go

### From remote repository

    % go get github.com/wataruSakurai/gen_pass
