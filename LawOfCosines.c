#include<stdio.h>
#include<math.h>

#define PI 3.14159265

int main() {
  char input[256];
  int side1,side2,angle;
  float ans; 

  printf("Side 1: ");
  fgets(input,256,stdin);
  side1 = atoi(input);

  printf("Side 2: ");
  fgets(input,256,stdin);
  side2 = atoi(input);
  
  printf("Angle: ");
  fgets(input,256,stdin);
  angle = atoi(input);

  ans = (side1*side1)+(side2*side2)-(2*side1*side2*cos(PI * angle / 180));
  ans = sqrt(ans);

  printf("Answer: %.6f\n",ans);

  return 0;

}
