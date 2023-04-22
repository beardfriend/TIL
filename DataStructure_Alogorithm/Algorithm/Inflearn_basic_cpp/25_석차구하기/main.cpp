#include<stdio.h>
#include<algorithm>
using namespace std;

int main() {

	int i, j, n, a[200], b[200];
	scanf("%d ", &n);
	for(i=1; i<=n; i++) {
		scanf("%d", &a[i]);
		b[i]=1;
	}

	for(i=1; i<=n; i++) {
		for(j=1; j<=n; j++) {
			if(a[j]>a[i]) b[i]++;
		}
	}

	for(i=1; i<=n; i++) {
		printf("%d ", b[i]);
	}
	

	return 0;
}