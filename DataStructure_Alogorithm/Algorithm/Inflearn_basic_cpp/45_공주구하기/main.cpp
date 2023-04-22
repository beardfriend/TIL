#include<stdio.h>
#include<vector>
#include<algorithm>
using namespace std;
int main(){
	freopen("input.txt", "rt", stdin);
	int n, k, p=0, i, bp=0, cnt=0;
	scanf("%d %d", &n, &k);
	vector<int> prince(n+1);
	while(1){
		p++;
		if(p>n) p=1;
		if(prince[p]==0){
			cnt++;
			if(cnt==k){
				prince[p]=1;
				cnt=0;
				bp++;
			}
		}
		if(bp==n-1) break;
	}
	for(i=1; i<=n; i++){
		if(prince[i]==0){
			printf("%d\n", i);
			break;
		}
	}
	return 0;
}