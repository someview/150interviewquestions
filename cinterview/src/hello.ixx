module;
#include <cstdio>

export module hello;

export namespace hello {
    void say(const char* str) {
        printf("%s\n", str);
    }
}