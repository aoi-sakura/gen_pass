# gen_pass: CLI for generating password

## Usage

    gen_pass [-i] <length> [-s|-S]
      -S       generate password without some symbols.
      -i int   password length. (default 16)
      -s       generate password without all symbols.

## Install

### From local repository

    % go build -o gen_pass main.go

### From remote repository

    % go get github.com/wataruSakurai/gen_pass
