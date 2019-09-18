# About

Golang diceware cli

# What is Diceware?

That's what we have Wikipedia for: https://en.wikipedia.org/wiki/Diceware

# Dependencies

This project is built using [bazel](https://bazel.build/). Check the [getting started page](https://docs.bazel.build/versions/master/getting-started.html) for installation instructions.

# Usage

Build using bazel:
```
$ bazel build //cmd/diceware
```

This will create a binary at `bazel-bin/cmd/diceware/${os}_${arch}_pure_stripped/diceware`

Running:
```
$ diceware -help
Usage of diceware:
  -delimiter string
        Character to separate the words (may be left out for none)
  -number int
        Number of words to generate
$ diceware -delimiter '/' -number 4
retreat/reset/illusion/outburst
```
