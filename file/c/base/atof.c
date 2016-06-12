/* This one doesn't check for bad syntax or overflow,
   and is slow and inaccurate.
   But it's good enough for the occasional string literal... */

#include "base.h"
#include <ctype.h>

double atof(char *s)
{
    double a = 0.0;
    int e = 0;
    int c;
    while ((c=*s++) != '\0' && isdigit(c)) {
        a = a*10 + (c - '0');
    }
    if (c == '.') {
        while((c=*s++) != '\0' && isdigit(c)) {
            a = a*10 + (c - '0');
            e -= 1;
        }
    }
    if (c == 'e' || c == 'E') {
        int sign = 1;
        int ex = 0;
        c = *s++;
        if (c == '+') {
            c = *s++;
        }
        else if (c == '-') {
            c = *s++;
            sign = -1;
        }
        while (isdigit(c)) {
            ex = ex*10 + (c - '0');
        }
        e += sign * ex;
    }
    while (e>0) {
        a *= 10;
        e --;
    }
    while (e<0) {
        a *= 0.1;
        e ++;
    }
    return a;
}
