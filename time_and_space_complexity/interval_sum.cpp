#include <bits/stdc++.h>
using namespace std;

int main() {
    long long q;
    cin >> q;

    while (q--) {
        long long l, r;
        cin >> l >> r;

        cout << ((r - l + 1) * (l + r) / 2) << "\n";
    }

    return 0;
}