#include<stdio.h>

int main() {
	int A[101], N, tmp;
	scanf("%d", &N);

	for (int i = 0; i < N; i++) {
		scanf("%d", &A[i]);
	}

	for (int i = 0; i< N - 1; i++) {
		for (int j=0; j <N-i-1; j++){
			if (A[j] > 0 && A[j +1 ] <0 ){
				tmp = A[j];
				A[j] = A[j+1];
				A[j+1] = tmp;
			}
		}
	}

	for(int i = 0; i<N; i++) printf("%d ",A[i]);

	return 0;
}