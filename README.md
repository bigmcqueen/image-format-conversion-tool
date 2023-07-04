# Image Format Converter (IFC)

## Overview

IFC (Image Format Converter) is a command line tool built with Go. 

It converts the format of an input image to a different format.

- PNG ⇔ JPEG

## Installation

The installation of this tool requires Go to be installed on your machine. Make sure you have Go installed.

After that, you can build this tool with the following commands:

```shell:bash
$ git clone https://github.com/BigMcQueen/image-format-conversion-tool.git
$ cd image-format-conversion-tool
$ go build -o ifc
```

## Usage

You can use this tool in the following format:

```shell:bash
$ ./ifc <inputfile> <outputfile>
```

For example, to convert sample.png to JPEG format:

```shell:bash
$ ./ifc sample.png sample.jpg
```

Similarly, to convert sample.jpg to PNG format:

```shell:bash
$ ./ifc sample.jpg sample.png
```

## License

MIT
