#include <bits/stdc++.h>
using namespace std;

int main() {
    long long x, n;
    cin >> x >> n;
    long long ans = 1;

    for (int i = 0; i < n; i++) {
        ans *= x;
    }

    cout << ans;

    return 0;
}