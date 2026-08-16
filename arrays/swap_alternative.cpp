#include <bits/stdc++.h>
using namespace std;

void swapAndPrint() {
    int n;
    cin >> n;
    long long arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    for (int i = 0; i < n - 1; i += 2) {
        swap(arr[i], arr[i + 1]);
    }

    for (int i = 0; i < n; i++) {
        cout << arr[i] << " ";
    }  
    cout << endl;
}

int main() {
    int t;
    cin >> t;

    while (t--) {
        swapAndPrint();
    }

    return 0;
}