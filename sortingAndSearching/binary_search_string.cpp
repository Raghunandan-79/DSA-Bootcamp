#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;

    vector<string> names(n);
    for (int i = 0; i < n; ++i) {
        cin >> names[i];
    }

    string x;
    cin >> x;

    int left = 0, right = n - 1;
    bool found = false;

    while (left <= right) {
        int mid = left + (right - left) / 2;
        if (names[mid] == x) {
            found = true;
            break;
        } 
        else if (names[mid] < x) {
            left = mid + 1;
        } 
        else {
            right = mid - 1;
        }
    }

    cout << (found ? "YES" : "NO");

    return 0;
}