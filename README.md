# JD

`jd` (as in "journal day") is a simple cli utility writen in Go to assist in the creation of daily journal journal pages in markdown format.

Created as a toy project to learn about using the `flag` and `template` packages to create cli utilities with go.

## Installation

Run the `INSTALL.sh` script, which will automatically install the jd binary to `/usr/local/bin`, unless `JD_INSTALL_PATH` has been set, in which case, it will install to the path specified.

## Usage

```shell
jd [OPTIONS] [TITLE]
```

Creates a partially filled journaling template in the current directory, or if set, the directory specified in the `JD_PATH` environment variable. Prints the name of the file created to stdout.

## Options

- `-f` `-filename`: Changes the filename from `DD-MM-YYYY-jd.md` to the string given as the flag's value.

- `-q` `--quiet`: Silences the output of the program (the filename of the page created).

- `-t` `--tag`: Adds the given string as a tag to the file which can be easily targeted by the grep utility. `grep -Er Tags:.*(<tag1>|<tag2>) .` will list the files in the current directory where `<tag1>` and `<tag2>` are given. This flag can be specified multiple times.

## Other Information

- The template is installed at `/usr/local/etc/jd/journalTemplate.txt`, and modifications can be made to its structure; however, care should be taken when modifying template components.
