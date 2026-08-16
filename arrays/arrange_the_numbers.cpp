#include <bits/stdc++.h>
using namespace std;

void arrangeNumbersAndPrint() {
    int n;
    cin >> n;
    int arr[n];

    int left = 0;
    int right = n - 1;
    
    for (int i = 1; i <= n; i++) {
        if (i % 2 == 1) {
            arr[left] = i;
            left++;
        } 
        else {
            arr[right] = i;
            right--;
        }
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
        arrangeNumbersAndPrint();
    }

    return 0;
}