/* toplev.c: defines toplev main for c, cc, etc
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
#define EXT {""} /* set in CMakeLists.txt */
#endif

struct pflags {
    bool help {false};
    bool repl {false};
    bool assembly {false};
    bool verbose {false};
    bool expand {false};
    bool obj {false};
    std::optional<std::string> outfile;
};

pflags parse_args(const std::vector<std::string>& argv) {
    pflags flags;
    for (size_t i = 1; i < argv.size(); ++i) {
        const std::string& arg = argv[i];
        if (arg == "--help" || arg == "-h") {
            flags.help = true;
        } 
        else if (arg == "--repl" || arg == "-R") {
            flags.repl = true;
        } 
        else if (arg == "-S") {
            flags.assembly = true;
        }
        else if (arg == "-E") {
            flags.expand = true;
        }
        else if (arg == "-c") {
            flags.obj = true;
        }
        else if (arg == "-o") {
            if (i + 1 < argv.size()) {
                flags.outfile = argv[i + 1];
                i++;
            } else {
                error("missing filename after '-o'");
                return flags;
            }
        }
    }

    return flags;
}

void toplev(int argc, std::vector<std::string> argv) {
    if (argc < 2) {
        fatal_error("no input files");
    }
    pflags flags = parse_args(argv);
    if (flags.help) {
        std::string help = "Usage: scc [options] file...\n\n";
        help += "Options:\n";
        help += "  -h, --help\tPrints out the help and exit.\n";
        help += "  -R, --repl\tEnables the REPL.\n";
        help += "  -o <file>\tPlace the output into <file>.\n";
        help += "  -S\t\tCompile only; do not assemble or link.\n";
        help += "  -E\t\tPreproccess only; do not compile, assemble or link.\n";
        help += "  -c\t\tCompile and assemble, but do not link.\n";
        help += "\nIf any bugs are found during use of this software report them to:\n<https://github.com/sillysoftware/scc/issues>";
        std::cout << help << std::endl;
        exit(0);
    }
    std::vector<std::string> ext = EXT;
    int found = 0;  
    int kargc = argc;
    for (int i = 1; i < kargc; i++) {
        auto carg = argv[i];
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
        error("file format not supported");
    }
    if (flags.repl) {
        std::cout << "REPL mode activated.\n";
    }

    if (flags.outfile) {
        std::cout << "Output file: " << *flags.outfile << '\n';
    }
}
