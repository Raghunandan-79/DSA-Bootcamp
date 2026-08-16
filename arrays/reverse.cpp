#include <bits/stdc++.h>
using namespace std;

void reversePrint(long long arr[], int n) {
    for (int i = n - 1; i >= 0; i--) {
        cout << arr[i] << " ";
    }
}

int main() {
    int n;
    cin >> n;
    long long arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    reversePrint(arr, n);

    return 0;
}