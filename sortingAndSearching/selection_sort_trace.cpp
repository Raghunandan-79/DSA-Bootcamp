#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;

    vector<int> a(n);
    for (int i = 0; i < n; i++) {
        cin >> a[i];
    }

    for (int i = 0; i < n - 1; i++) {
        int minIndex = i;
        int minVal = a[i];

        for (int j = i + 1; j < n; j++) {
            if (a[j] < minVal) {
                minVal = a[j];
                minIndex = j;
            }
        }

        swap(a[i], a[minIndex]);

        cout << "Pass " << i + 1 << ": ";
        for (int j = 0; j < n; j++) {
            if (j > 0) cout << ' ';
            cout << a[j];
        }
        cout << " , min_selected = " << minVal << '\n';
    }

    return 0;
}