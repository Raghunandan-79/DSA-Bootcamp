#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;
    vector<long long> arr(n);

    for (auto &it : arr) {
        cin >> it;
    }

    long long total = 0;
    for (int i = 0; i < n; i++) {
        total += arr[i];
    }

    long long average = total / n;
    long long maximum = *max_element(arr.begin(), arr.end()); 
    long long minimum = *min_element(arr.begin(), arr.end());
    
    cout << total << " " << maximum << " " << minimum << " " << average << endl;

    return 0;
}