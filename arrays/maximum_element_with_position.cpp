#include <bits/stdc++.h>
using namespace std;

void findMax(long long arr[], int n) {
    long long maximum = arr[0];
    long long maximumIndex = 0;

    for (int i = 1; i < n; i++) {
        if (arr[i] > maximum) {
            maximum = arr[i];
            maximumIndex = i;
        }
    }

    cout << maximum << " " << maximumIndex + 1 << endl;
}

int main() {
    int n;
    cin >> n;
    long long arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    findMax(arr, n);

    return 0;
}