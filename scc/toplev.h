/* error.h: defines error handel functions for c, cc, etc
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

#ifndef ERROR_H
#define ERROR_H

#include <vector>
#include <string>

/* toplev()
* @param argc Argument count
* @param argv Value of each element in the vector of strings of the initial input args
* @return void
*/
void toplev(int argc, std::vector<std::string> argv);

#endif
