#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    long long k;
    cin >> n >> k;

    vector<long long> a(n);

    unordered_map<long long, int> pos;

    for (int i = 0; i < n; i++) {
        cin >> a[i];

        long long need = k - a[i];

        if (pos.count(need)) {
            cout << pos[need] << " " << i + 1 << "\n";
            return 0;
        }

        pos[a[i]] = i + 1;
    }

    cout << -1 << "\n";
    return 0;
}