#include <stdio.h>
#include <string.h>

int main(void) {
  char c;
  char *s;
  int i;
  double d;
  float f;
  float* p;

  printf("char: %c, address: %p\n", c, &c);
  printf("string: %s, address: %p\n", s, &s);
  printf("integer: %i, address: %p\n", i, &i);
  printf("double: %f, address: %p\n", d, &d);
  printf("float: %f, address: %p\n", f, &f);
  printf("pointer: %p, address: %p, value: %f\n", p, &p, *p);

  printf("\n---------------\n\n");

  c = 'c';
  s = "Hello World!";
  i = 4;
  d = 3.23;
  f = 1.32;
  p = &f;
  
  printf("char: %c, address: %p\n", c, &c);
  printf("string: %s, address: %p\n", s, &s);
  printf("integer: %i, address: %p\n", i, &i);
  printf("double: %f, address: %p\n", d, &d);
  printf("float: %f, address: %p\n", f, &f);
  printf("pointer: %p, address: %p, value: %f\n", p, &p, *p);
}