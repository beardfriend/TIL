#include<stdio.h>
#include<algorithm>
#include<vector>
using namespace std;

int main() {
	int i, n, tmp, digit, cnt; 
	scanf("%d", &n);

	for (i = 1; i <= n; i++) {
		tmp = i;

		while (tmp > 0) {
			digit = tmp % 10;
			if (digit == 3) cnt++;
			tmp = tmp / 10;
		}
	}

	printf("%d", cnt);


	return 0;

}