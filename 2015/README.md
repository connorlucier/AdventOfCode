# Advent of Code 2015

Started this in my downtime during 2022 AoC for fun.

## Language Selection

Opted for C++ to grease the wheels a bit. Haven't written much C++ since college but want to keep things fresh. Also wanted to play around with shell and makefiles since that's always a doozy.

## Build and Run

The `makefile` in this directory has several commands:

- `make` - build code for a single day
- `make run` - build and run code for a single day
- `make all` - build code for all days (multiple executables)
- `make clean` - remove all executables

Both `make` and `make run` require you to specify which day to build and/or run. For example:

- `make day=01` will build the `01/` directory
- `make run day=10` will build and run the `10/` directory
