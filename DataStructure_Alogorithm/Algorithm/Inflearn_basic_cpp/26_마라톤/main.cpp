#include<stdio.h>
#include<algorithm>
#include<vector>
using namespace std;
int main() {
	int n, i, j, cnt;
	scanf("%d", &n);
	vector<int> a(n+1);

	printf("1 ");
	for (i=2; i<=n; i++) {
		cnt = 0;
		for (j=i-1; j>=1; j--) {
			if(a[j] >= a[i]) cnt++;
		}
		printf("%d ", &cnt);
	}
	
	return 0;
}