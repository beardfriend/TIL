#include <stdio.h>

int main() {
	int a[100], i, j, n, tmp;

	scanf("%d", &n);

	for (i = 0; i < n; i++) {
		scanf("%d", &a[i]);
    }

    for (i = 0; i < n; i++) {
		tmp = a[i];
		for(j = i-1; j >= 0; j--) {
			if (a[j] > tmp) {
				a[j+1] = a[j];
            } else {
				break;
			}
		}
		a[j+1] = tmp;
	}
	
	for (i=0; i< n; i++) {
		printf("%d ", a[i]);
    }

	return 0;
}