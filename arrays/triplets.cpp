#include <bits/stdc++.h>
using namespace std;

void findTriplets() {
    int n;
    cin >> n;
    long long arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    long long x;
    cin >> x;

    int totalTriplets = 0;
    for (int i = 0; i < n - 2; i++) {
        for (int j = i + 1; j < n - 1; j++) {
            for (int k = j + 1; k < n; k++) {
                if (arr[i] + arr[j] + arr[k] == x) {
                    totalTriplets++;
                }
            }
        }
    }

    cout << totalTriplets << endl;
}

int main() {
    int t;
    cin >> t;

    while (t--) {
        findTriplets();
    }

    return 0;
}