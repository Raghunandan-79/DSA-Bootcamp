#include <bits/stdc++.h>
using namespace std;

int fastestRacer(long long arr[], int n) {
    int fastest = 0;
    int prevFastest = arr[0];

    for (int i = 1; i < n; i++) {
        if (arr[i] < prevFastest) {
            fastest = i;
            prevFastest = arr[i];
        }
        else if (arr[i] == prevFastest) {
            fastest = max(fastest, i);
            prevFastest = arr[i];
        }
    }

    return fastest + 1;
}

int main() {
    int n;
    cin >> n;
    long long arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    cout << fastestRacer(arr, n) << endl;

    return 0;
}