# gen_pass: CLI for generating password

## Usage

    gen_pass [-i] <length> [-s|-S]
      -S       generate password without some symbols.
      -i int   password length. (default 16)
      -s       generate password without all symbols.

## Install

    % go build -o gen_pass main.go
