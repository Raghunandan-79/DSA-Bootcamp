#include <bits/stdc++.h>
using namespace std;

void intersectionOfArrays() {
    int n;
    cin >> n;
    int arr1[n];

    for (int i = 0; i < n; i++) {
        cin >> arr1[i];
    }

    int m;
    cin >> m;
    vector<int> arr2(m);

    for (int i = 0; i < m; i++) {
        cin >> arr2[i];
    }

    sort(arr2.begin(), arr2.end());

    bool first = true;
    for (int k = 0; k < n; k++) {
        int val = arr1[k];
        auto it = lower_bound(arr2.begin(), arr2.end(), val);
        if (it != arr2.end() && *it == val) {
            if (!first) cout << " ";
            cout << val;
            first = false;
            arr2.erase(it);
        }
    }
    cout << '\n';
}

int main() {
    int t;
    cin >> t;

    while (t--) {
        intersectionOfArrays();
    }

    return 0;
}