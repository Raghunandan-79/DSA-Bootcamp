#include <bits/stdc++.h>
using namespace std;

void printSpecialFactors(int n) {
    bool found = false;

    for (int i = 1; i <= n; i++) {
        if (n % i == 0 && (i % 10 == 2 || i % 10 == 7)) {
            cout << i << " ";
            found = true;
        }
    }

    if (!found) {
        cout << -1;
    }
}

int main() {
    int n;
    cin >> n;

    printSpecialFactors(n);

    return 0;
}