#include <bits/stdc++.h>
using namespace std;

void findMin(long long arr[], int n) {
    long long minimum = arr[0];
    int minIndex = 0;

    for (int i = 1; i < n; i++) {
        if (arr[i] < minimum) {
            minimum = arr[i];
            minIndex = i;
        }
    }

    cout << minimum << " " << (minIndex + 1) << endl;
}

int main() {
    int n;
    cin >> n;
    long long arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    findMin(arr, n);

    return 0;
}