/* main.cc: defines main() for c, cc, etc
    Copyright (C) 2024-2024 Silly Software Foundation.

This file is part of SCC.

SCC is free software; you can redistribute it and/or modify it under
the terms of the BSD 3-Clause License as shown in the LICENCE file.

SCC is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warrenty of MERCHANTABILITY or
FITNESS FOR A PARTICULAR PURPOSE. See the BSD 3-Clause Licence
for more details.

You should have received a copy of the BSD 3-Clause
along with SCC; see the file LICENCE. If not see
<https://raw.githubusercontent.com/sillysoftware/scc/refs/heads/master/LICENSE> */

#include "error.h"
#include <cstdio>
#include <cstring>
#include <optional>
#include <string>
#include <iostream>
#include <vector>

#ifndef EXT
#define EXT {""}
#endif

struct pflags {
    bool help {false};
    bool repl {false};
    std::optional<std::string> outfile;
};

int main(int argc, char *argv[]) {
    std::vector<std::string> args;
    for (int i = 0; i < argc; ++i) {
        args.push_back(std::string(argv[i]));
    }
    if (argc < 2) {
        fatal_error("no input files");
    }
    std::vector<std::string> ext = EXT;
    int found = 0;  
    for (int i = 1; i < argc; i++) {
        auto carg = args[i];
        for (const auto &str : ext) {
            if (carg.length() >= str.length() && 
                carg.compare(carg.length() - str.length(), str.length(), str) == 0) {
                std::cout << "Found: " << carg << std::endl;  
                found++;
                break; 
            }
        }
    }
    if (found == 0) {  
        fatal_error("file format not supported");
    }
    /* parser argv on ext */
    return 0;
}
