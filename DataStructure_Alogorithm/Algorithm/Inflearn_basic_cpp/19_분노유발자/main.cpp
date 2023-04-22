#include<stdio.h>

using namespace std;

int main() {
	int n, i, cnt=0, max, h[101];
	scanf("%d", &n);
	
	for(i=1; i<n; i++) {
		scanf("%d", &h[i]);
	}
	max=h[n];
	for(i=n; i>=1; i--) {
		if(h[i] > max) {
			max = h[i];
			cnt++;
		}
	}

	printf("%d\n",cnt);

	return 0;
}