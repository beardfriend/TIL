#include<iostream>
using namespace std;

// 자연수의 합
int main(){
	int a,b,i, sum=0;
	cin>>a>>b;
	for(i=a; i<b; i++) {
		cout<<i<<" + ";
		sum = sum + i;
	}
	cout<< i << " = ";
	cout<<sum+i;
	return 0;
}