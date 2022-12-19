# Flamegraph Generator

This is a Go script that generates a flamegraph of the performance of a process with a given PID when run with the `--pid` flag. It does nothing intelligent - just a wrapper for perf and flamegraph.

## Prerequisites

Before you can use this script, you will need to have the following tools installed:

- `perf`: a performance analysis tool for Linux
- `flamegraph`: a tool for generating flamegraphs from performance data
- `stackcollapse-perf.pl`: a script that is part of the `flamegraph` toolkit and is used to process the output of `perf`

You can install these tools by following the instructions at the following links:

- `perf`: https://perf.wiki.kernel.org/index.php/Main_Page
- `flamegraph`: https://github.com/brendangregg/FlameGraph

## Usage

To use this script, first compile it with the following command:

```sh
go build -o flamegraph
```

This will create an executable file called `flamegraph`.

To generate a flamegraph of the performance of a process with a given PID, run the `flamegraph` executable with the `--pid` flag, followed by the PID of the process you want to generate a flamegraph for. For example:

```sh
./flamegraph --pid 12345
```

This will generate a flamegraph of the performance of the process with PID 12345 and print it to stdout. You can redirect the output to a file if you want to save the flamegraph to a file, like this:

```sh
./flamegraph --pid 12345 > flamegraph.svg
``

This will save the flamegraph to a file called `flamegraph.svg`. You can then open the file in a web browser or a viewer to view the flamegraph.

## Notes

- This script assumes that the `perf`, `flamegraph`, and `stackcollapse-perf.pl` tools are in the PATH.
- If the `perf` or `flamegraph` command is not found, the script will exit with an error message indicating that the command was not found.
- The script requires the `--pid` flag to be specified. If the flag is not provided, the script will exit with an error message.
