#include<stdio.h>
#include<vector>
#include<algorithm>
using namespace std;
int main() {
	int a, b=1, cnt=0, tmp, i;
	scanf("%d",&a);
	tmp = a;
	a--;
	while(a>0){
		b++;
		a = a-b;
		if(a % b ==0) {
			for(i =1; i<b; i++) {
				printf("%d + ", (a/b)+i);
			}
		}
	}


	return 0;
}
