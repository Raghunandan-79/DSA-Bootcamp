#include <bits/stdc++.h>
using namespace std;

bool linearSearch(vector<long long> arr, long long x) {
    for (int i = 0; i < arr.size(); i++) {
        if (arr[i] == x) {
            return true;
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

    if (linearSearch(arr, x)) {
        cout << "YES" << endl;
    }
    else {
        cout << "NO" << endl;
    }

    return 0;
}