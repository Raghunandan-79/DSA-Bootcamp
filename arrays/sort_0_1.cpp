#include <bits/stdc++.h>
using namespace std;

void sortArray(int arr[], int n) {
    int countZeros = 0;
    int countOnes = 0;

    for (int i = 0; i < n; i++) {
        if (arr[i] == 0) {
            countZeros++;
        }
        else {
            countOnes++;
        }
    }

    for (int i = 1; i <= countZeros; i++) {
        cout << "0 ";
    }

    for (int i = 1; i <= countOnes; i++) {
        cout << "1 ";
    }

    cout << endl;
}

void solve() {
    int n;
    cin >> n;
    int arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    sortArray(arr, n);
}

int main() {
    int t;
    cin >> t;

    while (t--) {
        solve();
    }

    return 0;
}