#include<stdio.h>
using namespace std;

// 내 작성
// int main() {
// 	int i, j, N, cnt=0;
// 	scanf("%d", &N);

// 	for(i=1; i<=N; i++) {
// 		cnt = 0;
// 		for(j=1; j<=i; j++) {
// 			if (i % j == 0) {
// 				cnt++;
// 			}
// 		}
// 		printf("%d ",cnt);
// 	}
// }
// 전역변수 선언시 모든 배열의 값은 0으로 초기화.
int cnt[50001];
int main() {
	int n,i,j;
	scanf("%d", &n);
	for(i =1; i<=n; i++) {
		for(j=i; j<=n; j=j+i) {
			cnt[j]++;
		}
	}
	for(i=1; i<=n; i++) {
		printf("%d ", cnt[i]);
	}

	return 0;
}