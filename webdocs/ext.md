# Defines
## EXT
### Purpose
EXT is used for checking weather input file are valid for the compiler. This is set as a macro variable becuase this compiler is meant to be easily hackable. The best way we found to do this was to make it customizable in the build chain.
#### Example
```add_compile_definitions(EXT={"foo"})```
>[!NOTE]
> Periods are added in compilation time.
