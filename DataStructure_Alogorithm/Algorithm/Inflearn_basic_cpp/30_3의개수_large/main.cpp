#include<stdio.h>
#include<algorithm>
#include<vector>
using namespace std;

int main() {
	int n, lt=1, cur, rt ,k=1, res=0;
	n = 5367;
	while (lt != 0) {
		lt = n / (k *10);
		rt = n % k;
		cur = (n / k) % 10;
		
		if (cur > 3) {
			res = res + (lt + 1) * k;
		} else if (cur < 3) {
			res = res + (lt *k);
		} else {
			res = res + (lt * k) +(rt + 1);
		}
		printf("%d %d %d", lt, cur, rt);
		k = k * 10;
	}
	printf("%d \n", res);
	// 각 자리 별로 얼마나 숫자가 나타나는지 파악.
	// 30 ~ 39 10개
	// 130 ~ 139 10개
	
	return 0;
}
