#include <bits/stdc++.h>
using namespace std;

bool binarySearch(vector<long long> arr, long long x) {
    int low = 0, high = arr.size() - 1;

    while (low <= high) {
        int mid = low + (high - low) / 2;

        if (arr[mid] == x) {
            return true;
        }
        else if (arr[mid] > x) {
            high = mid - 1;
        }
        else {
            low = mid + 1;
        }
    }

    return false;
}

int main() {
    int n;
    cin >> n;
    vector<long long> arr(n);
    
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    long long x;
    cin >> x;

    if (binarySearch(arr, x)) {
        cout << "YES" << endl;
    }
    else {
        cout << "NO" << endl;
    }

    return 0;
}