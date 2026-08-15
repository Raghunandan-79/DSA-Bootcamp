#include <bits/stdc++.h>
using namespace std;

long long sumOfArray(long long arr[], int n) {
    long long sum = 0;

    for (int i = 0; i < n; i++) {
        sum += arr[i];
    }

    return sum;
}

int main() {
    int n;
    cin >> n;
    long long arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    cout << sumOfArray(arr, n) << endl;
    
    return 0;
}