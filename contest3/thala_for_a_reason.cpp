#include <bits/stdc++.h>
using namespace std;

void solve() {
    long long r;
    cin >> r;

    int sum = 0;
    do {
        sum += r % 10;
        r /= 10;
    } while (r > 0);

    if (sum == 7) {
        cout << "Thala for a reason" << endl;
    }
    else {
        cout << "Blocked for no reason" << endl;
    }
}

int main() {
    int t;
    cin >> t;

    while (t--) {
        solve();
    }

    return 0;
}