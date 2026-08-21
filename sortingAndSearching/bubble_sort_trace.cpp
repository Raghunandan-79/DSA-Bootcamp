#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;

    vector<int> a(n);
    for (int &value : a) {
        cin >> value;
    }

    for (int pass = 1; pass < n; ++pass) {
        int swaps = 0;

        for (int i = 0; i < n - pass; ++i) {
            if (a[i] > a[i + 1]) {
                swap(a[i], a[i + 1]);
                ++swaps;
            }
        }

        cout << "Pass " << pass << ": ";
        for (int i = 0; i < n; ++i) {
            if (i > 0) cout << ' ';
            cout << a[i];
        }
        cout << " , swaps = " << swaps << '\n';

        if (swaps == 0) break;
    }

    return 0;
}