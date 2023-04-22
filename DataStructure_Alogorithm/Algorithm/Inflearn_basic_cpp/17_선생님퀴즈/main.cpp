#include<stdio.h>
using namespace std;

int main() {
	int n, sum=0, i, j, m, ans;
	scanf("%d", &n);
	for(i=1; i<=n; i++) {
		scanf("%d %d",&m, &ans);
		sum=0;
		for (j=1; j<=m; j++) {
			sum += j;
		}
		if (sum == ans) {
			printf("YES\n");
		} else {
			printf("NO\n");
		}
	}

	return 0;
}