# Advent of Code 2015

Started this in my downtime during 2022 AoC for fun.

## Language Selection

Opted for C++ since I haven't written much C++ since college but want to grease those wheels a bit. Also wanted to play around with shell and makefiles since that's always a doozy.

## Build and Run

The `makefile` in this directory has several commands:

- `make` - build code for a single day
- `make run` - build and run code for a single day
- `make all` - build code for all days (multiple executables)
- `make new` - create folder + files for a new day
- `make clean` - remove all executables

The `make`, `make run` and `make new` commands require you to specify a `day` parameter. For example:

- `make day=01` will build the `01/` directory
- `make run day=10` will build and run the `10/` directory
- `make new day=26` will create a new directory called `26/` with the following files:
    - `main.cpp`
    - `part1.h`
    - `part2.h`
