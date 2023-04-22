#include<stdio.h>
#include<algorithm>
#include<vector>
using namespace std;

int main() {
	int n,i,j,tmp, cntOne, cntTwo;

	scanf("%d", &n);

	for (i=2; i <= n; i++) {
		tmp = i;
		j = 2;
		while (1) {
			if (tmp % j == 0) {
				if (j == 2) cntOne++;
				else if (j == 5) cntTwo++;
				tmp = tmp / j;
			} 
			else j++;
			if (tmp ==1) break;
		}
	}

	if (cntOne > cntTwo) {
		printf("%d", cntOne);
	} else printf("%d", cntTwo);

	return 0;

}