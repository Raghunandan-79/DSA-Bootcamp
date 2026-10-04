#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n, q;
    cin >> n >> q;

    vector<long long> prefix(n + 1, 0);
    for (long long i = 1; i <= n; ++i) {
        long long value;
        cin >> value;
        prefix[i] = prefix[i - 1] ^ (value & 1);
    }

    while (q--) {
        long long l, r;
        cin >> l >> r;
        long long remainingParity = prefix[n] ^ (prefix[r] ^ prefix[l - 1]);
        cout << (remainingParity ? "YES" : "NO") << '\n';
    }

    return 0;
}