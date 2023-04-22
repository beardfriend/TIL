#include<stdio.h>
using namespace std;

int main() {
	int i,N,cnt=0,tmp;
	scanf("%d",&N);
	for(i=1; i<=N; i++) {
		tmp = i;
		while(tmp>0){
			tmp = tmp/10;
			cnt++;
		}
	}
	printf("%d\n",cnt);
	return 0;
}