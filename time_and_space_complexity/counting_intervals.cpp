#include <bits/stdc++.h>
using namespace std;

int main() {
    long long q;
    cin >> q;

    while (q--) {
        long long t, l, r;
        cin >> t >> l >> r;

        if (l > r) {
            cout << 0 << "\n";
            continue;
        }

        long long count = r - l + 1;
        if (t == 1 || t == 3) {
            count--;
        }

        if (t == 1 || t == 2) {
            count--;
        }

        cout << max((long long) 0, count) << "\n";
    }

    return 0;
}