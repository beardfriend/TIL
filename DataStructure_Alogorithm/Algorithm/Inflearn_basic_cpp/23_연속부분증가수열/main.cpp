#include<stdio.h>
#include<vector>

// 방법 1 . 이전 값과 비교 해서 크면 +를 해준다.
// max에 넣어주고 max랑 비교를 해준다.

int main() {
	int n, i, pre, now, cnt, max;
	scanf("%d", &n);
	scanf("%d", &pre);
	cnt = 1;
	max = 1;

	for(i=2; i<=n; i++) {
		scanf("%d", &now);
		if(now >= pre) {
			cnt++;
			if (cnt > max) {
				max = cnt;
			}
		} else cnt=1;
		pre=now;
	}
	
	printf("%d\n", max);


	return 0;
}