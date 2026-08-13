#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;

    if (n == 0) {
        cout << 0;
    }
    else {
        while (n > 0) {
            cout << n % 10;
            n /= 10;
        }
    }

    return 0;
}