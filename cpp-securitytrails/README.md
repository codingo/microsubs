# cpp-securitytrails

This application queries the securitytrails.com and prints discovered subdomains.
## Requirement
- C++11
- Linux

## How to make

```
$ make securitytrails

g++  -std=c++11 -I/includes -L/includes -Lbuild/includes  -c main.cpp
g++  -std=c++11 -I/includes -L/includes -Lbuild/includes  main.o    -o securitytrails
```

# Command line arguments

```
$ ./securitytrails -h
This queries the securitytrails.com and prints the sub domain found.
Usage: ./securitytrails [OPTIONS]

Options:
  -h,--help                   Print this help message and exit
  -d,--domain TEXT            Target domain e.g apple.com (required)
  -o,--output TEXT            Output file (optional)
  -s,--silent UINT            To display only the subdomains
  -c,--config TEXT            Configuration file. (optional)
```


## To run silent
Set the silent flag
```
$ ./securitytrails -d apple.com -s 1
```
## Libraries Used
[CLI11](https://github.com/CLIUtils/CLI11)

[Json](https://github.com/nlohmann/json)

[HTTPRequest](https://github.com/elnormous/HTTPRequest)
