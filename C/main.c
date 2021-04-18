
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "go.h"

int main()
{
    char *s = Hello();
    if (!s)
    {
        return 1;
    }
    printf("%s\n", s);
    free(s);

    uint64_t c;
    c = Add(10, 20);
    printf("%ld\n", c);
}
