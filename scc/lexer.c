/* main.c: defines main() for c, cc, etc
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

#include <ctype.h>
#include <string.h>

#define MXTKLE 100

typedef enum {
  TOKEN_EOF,
  TOKEN_IDENT,
  TOKEN_KEYWORD,
  TOKEN_NUMBER,
  TOKEN_SYMBOL,
  TOKEN_NIL,
} TokenType;

typedef struct {
  TokenType type;
  char value[MXTKLE];
} Token;

const char* keywords[] = {
  "exit", // exit syscall
  NULL
};

int is_keyword(const char* str) {
  for (int i = 0; keywords[i] != NULL; i++) {
    if (strcmp(str, keywords[i]) == 0) {
      return 1;
    }
  }
  return 0;
}

Token next(const char** input) {
  Token internal_token;
  internal_token.type = TOKEN_NIL;
  internal_token.value[0] = '\0';
  while (isspace(**input)) {
    (*input)++;
  }
  while (**input == '\0') {
    internal_token.type = TOKEN_EOF;
    return internal_token;
  }
  if (isalpha(**input)) {
    int len = 0;
    while (isalnum(**input) && len < MXTKLE -1) {
      internal_token.value[len++] = *(*input)++;
    }
    internal_token.value[len] = '\0';
    if (is_keyword(internal_token.value)) {
      internal_token.type = TOKEN_KEYWORD;
    } else {
      internal_token.type = TOKEN_IDENT;
    }
  }
  else if (isdigit(**input)) {
    int len = 0;
    while (isdigit(**input) && len < MXTKLE -1) {
      internal_token.value[len++] = *(*input)++;
    }
    internal_token.value[len] = '\0';
    internal_token.type = TOKEN_NUMBER;
  } else {
    internal_token.value[0] = *(*input)++;
    internal_token.value[1] = '\0';
    internal_token.type = TOKEN_SYMBOL;
  }
  return internal_token;
}

Token token;
Token* lexer(const char *source, int* token_count) {
  while (1) {
    token = next(&source);
    token_count++;
    if (token.type == TOKEN_EOF) {
      break;
    }
    
  }
  return &token;
}
