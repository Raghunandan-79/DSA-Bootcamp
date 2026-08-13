#include <bits/stdc++.h>
using namespace std;

int main() {
	int n;
	cin >> n;

	for (int i = 1; i <= n; ++i) {
		char start = (i % 2 == 1) ? '0' : '1';
		
        for (int j = 0; j < i; ++j) {
			cout << ( (j % 2 == 0) ? start : (start == '0' ? '1' : '0') );
		}

		if (i < n) cout << '\n';
	}

	return 0;
}