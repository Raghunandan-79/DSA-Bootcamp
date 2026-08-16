#include <bits/stdc++.h>
using namespace std;

void pairSum() {
    int n;
    cin >> n;
    long long arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    long long x;
    cin >> x;

    int countPairs = 0;
    for (int i = 0; i < n; i++) {
        for (int j = i + 1; j < n; j++) {
            if (arr[i] + arr[j] == x) {
                countPairs++;
            }
        }
    }

    cout << countPairs << endl;
}

int main() {
    int t;
    cin >> t;

    while (t--) {
        pairSum();
    }

    return 0;
}