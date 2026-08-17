#include <bits/stdc++.h>
using namespace std;

void arrangeAndPrint(long long arr[], int n) {
    vector<long long> result;
    
    int left = n / 2 - 1;
    int right = n / 2;
    
    result.push_back(arr[left]);
    result.push_back(arr[right]);
    
    left--;
    right++;
    
    while (left >= 0 && right < n) {
        result.push_back(arr[left]);
        result.push_back(arr[right]);
        left--;
        right++;
    }
    
    for (int i = 0; i < result.size(); i++) {
        if (i > 0) cout << " ";
        cout << result[i];
    }
    cout << endl;
}

int main() {
    int n;
    cin >> n;
    long long arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    arrangeAndPrint(arr, n);

    return 0;
}