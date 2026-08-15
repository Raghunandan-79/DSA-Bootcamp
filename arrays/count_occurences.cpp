#include <bits/stdc++.h>
using namespace std;

int countOccurences(long long arr[], long long x, int n) {
    int count = 0;

    for (int i = 0; i < n; i++) {
        if (arr[i] == x) {
            count++;
        }
    }

    return count;
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

    cout << countOccurences(arr, x, n) << endl;

    return 0;
}