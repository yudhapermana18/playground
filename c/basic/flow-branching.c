#include <stdio.h>

void ifelse(int num) {
  printf("pointer num in ifelse: %p\n", &num);

  if (num >= 10)
  {
    printf("%d equal greather than 10\n", num);
  } else
  {
    printf("%d less than 10\n", num);
  }
}

void switchCase(int num) {
  switch (num)
  {
  case 1:
    printf("this one\n");
    break;

  case 2:
    printf("this two\n");
    break;
  
  default:
    printf("default\n");
    break;
  }
}

int main(void) {
  int a = 5;
  int b = 2;

  ifelse(a);
  switchCase(b);

  printf("pointer a : %p\n", &a);

  printf("done!\n");
  return 0;
}