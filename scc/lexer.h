/* lexer.h: defines lexer functions for c, cc, etc
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

#ifndef LEXER_H
#define LEXER_H

#ifdef __cplusplus
extern "C" {
#endif

/* MXTKLE
 * Max token lenth SCC can take
 */
#define MXTKLE 100

/* TokenType
 * types of tokens
 */
typedef enum {
  TOKEN_EOF,
  TOKEN_IDENT,
  TOKEN_KEYWORD,
  TOKEN_NUMBER,
  TOKEN_SYMBOL,
  TOKEN_NIL,
} TokenType;

/* Token
 * individual token
 */
typedef struct {
  TokenType type;
  char value[MXTKLE];
} Token;

/* lexer()
* @param char* The input string
* @return Token
*/
Token* lexer(const char *source, int* token_count);


#ifdef __cplusplus
}
#endif
#endif
