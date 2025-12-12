# Gotree

A small Go program that replicates the classic `tree` command. It scans a directory and prints a clean hierarchical view of its files and folders.

## Installation

### Install via go

If you have **Go** installed:

```sh
go install github.com/anas-shakeel/gotree/cmd/gotree@latest
```

### Download Precompiled Binaries

Visit the [Releases Page](https://github.com/Anas-Shakeel/gotree/releases/latest) to download precompiled binaries. **(Linux, MacOS, Windows)**.

-   After downloading, unpack and move the binary to a folder in your `$PATH`.

### Build from Source

If you have **Go** installed, you can build the project by running:

```sh
git clone https://github.com/anas-shakeel/gotree.git
cd gotree
go build
```

## Usage

```sh
gotree [OPTIONS] <path>
```

For example:

```sh
gotree .

# OR for help message
gotree -h
```




## License

MIT
