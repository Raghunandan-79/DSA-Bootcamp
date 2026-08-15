#include <bits/stdc++.h>
using namespace std;

bool searchElement(long long arr[], long long x, int n) {
    for (int i = 0; i < n; i++) {
        if (arr[i] == x) {
            return true;
        }
    }

    return false;
}

int main() {
    int n;
    cin >> n;
    long long x;
    cin >> x;
    long long arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    if (searchElement(arr, x, n)) {
        cout << "YES" << endl;
    }
    else {
        cout << "NO" << endl;
    }

    return 0;
}   