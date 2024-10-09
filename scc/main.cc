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
#include <regex>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    if (argc < 2) {
        fatal_error("no input files");
    }
    std::regex ext_regex(".c");
    int found;
    for (int i = 1; i < argc; i++) {
        auto carg = argv[i];
        if (std::regex_search(carg, ext_regex)) {
            puts("Found the extention for C/CXX");
            found++;
        }
    }
    if (found == 0) {
        fatal_error("file format not supported");
    }
    return 0;
}
