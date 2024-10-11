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

int main(int argc, char *argv[]) {
    std::vector<std::string> args;
    for (int i = 0; i < argc; ++i) {
        args.push_back(std::string(argv[i]));
    }
    if (argc < 2) {
        fatal_error("no input files");
    }
    pflags flags = parse_args(args);
    if (flags.help) {
        std::string helptxt = "Usage: scc [options] files...\n\nOptions:\n\n  -h, --help        Prints out the help and exit.\n  -R, --repl        Enables the REPL.\n  -o                Sets the output file.\n  -S                Compile only; do not assemble or link.\n  -E                Preproccess only; do not compile, assemble or link.\n  -c                Compile and assemble, but do not link.";
        helptxt += "\n\nIf any bugs are found during use of this software report them to:\n<https://github.com/sillysoftware/scc/issues>";
        std::cout << helptxt << std::endl;
        exit(0);
    }
    std::vector<std::string> ext = EXT;
    int found = 0;  
    int kargc = argc;
    for (int i = 1; i < kargc; i++) {
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
        error("file format not supported");
    }
    if (flags.repl) {
        std::cout << "REPL mode activated.\n";
    }

    if (flags.outfile) {
        std::cout << "Output file: " << *flags.outfile << '\n';
    }
    return 0;
}
