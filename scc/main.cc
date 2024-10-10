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
#include <string>
#include <iostream>
#include <vector>

struct pflags {
    bool help {false};
    bool repl {false};
    option<string> outfile;
};

pflags parse_pflags(int argc, const char* argv[]) {
    pflags enabled;
    for(int i {1}; i < argc; i++) {
    string opt {argv[i]};
    if(auto j {NoArgs.find(opt)}; j != NoArgs.end())
      j->second(settings);
    else if(auto k {OneArgs.find(opt)}; k != OneArgs.end())
      if(++i < argc)
        k->second(settings, {argv[i]});
      else
        throw std::runtime_error {"missing param after " + opt};
    else if(!settings.infile)
      settings.infile = argv[i];
    else
      cerr << "unrecognized command-line option " << opt << endl;
  }
  return enabled;
}
}

int main(int argc, char *argv[]) {
    std::vector<std::string> args;
    for (int i = 0; i < argc; ++i) {
        args.push_back(std::string(argv[i]));
    }
    if (argc < 2) {
        fatal_error("no input files");
    }
    std::vector<std::string> ext = {"bsc", "b", "basic"};
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
