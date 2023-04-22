#include<stdio.h>
using namespace std;

int ch[10];
int main() {
	int i, digit, max = -214700000, res;
	char a[101];
	scanf("%s", a);
	for (i=0; a[i] != '\0'; i++) {
		digit=a[i] - 48;
		ch[digit]++;
	}

	for(i=0; i<=9; i++) {
		if (ch[i] >= max) {
			max = ch[i];
			res = i;
		}
	}

	printf("%d\n",res);
	return 0;
}