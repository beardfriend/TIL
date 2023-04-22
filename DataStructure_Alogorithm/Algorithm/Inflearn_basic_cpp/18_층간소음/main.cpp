#include<stdio.h>

using namespace std;



int main() {
	int n,a, val, i, cnt=0, max=-2147000000;
	scanf("%d %d", &n, &val);

	for(i=0; i<n; i++) {
		scanf("%d", &a);
		if(a>val) cnt++;
		else cnt =0;
		if(cnt>max) max=cnt;
	}

	if(max==0) printf("-1\n");
	else printf("%d", &max);
	return 0;
}