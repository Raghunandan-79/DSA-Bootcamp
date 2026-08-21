#include <bits/stdc++.h>
using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int n;
    cin >> n;
    vector<int> a(n);
    for (int &value : a) cin >> value;

    for (int i = 1; i < n; ++i) {
        int key = a[i];
        int j = i - 1;
        int shifts = 0;

        while (j >= 0 && a[j] > key) {
            a[j + 1] = a[j];
            --j;
            ++shifts;
        }
        a[j + 1] = key;

        cout << "Pass " << i << ": ";
        for (int k = 0; k < n; ++k) {
            if (k) cout << ' ';
            cout << a[k];
        }

        cout << " , ";
        for (int k = 0; k <= i; ++k) {
            if (k) cout << ' ';
            cout << a[k];
        }
        
        cout << " |";
        
        for (int k = i + 1; k < n; ++k) cout << ' ' << a[k];
        
        cout << " , shifts = " << shifts << '\n';
    }

    return 0;
}