#include <bits/stdc++.h>
using namespace std;

int main() {
    int n, m;
    cin >> n >> m;

    vector<long long> a(n), b(m);
    for (long long& x : a) cin >> x;
    for (long long& x : b) cin >> x;

    vector<long long> c;
    c.reserve(n + m);

    int i = 0, j = 0;
    while (i < n || j < m) {
        if (j == m || (i < n && a[i] <= b[j])) {
            c.push_back(a[i++]);
        } 
        else {
            c.push_back(b[j++]);
        }
    }

    for (int k = 0; k < n + m; ++k) {
        if (k > 0) cout << ' ';
        cout << c[k];
    }
    cout << '\n';

    return 0;
}