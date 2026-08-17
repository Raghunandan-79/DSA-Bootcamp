#include <bits/stdc++.h>
using namespace std;

void printUniqueElements(long long arr[], int n) {
    for (int i = 0; i < n; i++) {
        bool found = false;

        for (int j = 0; j < n; j++) {
            if (i != j && arr[i] == arr[j]) {
                found = true;
                break;
            }
        }

        if (!found) {
            cout << arr[i] << " ";
        }
    }
}

int main() {
    int n;
    cin >> n;
    long long arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    printUniqueElements(arr, n);

    return 0;
}