#include <bits/stdc++.h>
using namespace std;

void findMissing() {
    int n;
    cin >> n;
    int arr[n];
    
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    int xorr = 0;

    for (int i = 0; i < n; i++) {
        xorr ^= arr[i];
    }

    cout << xorr << endl;
}

int main() {
    int t;
    cin >> t;

    while (t--) {
        findMissing();
    }

    return 0;
}