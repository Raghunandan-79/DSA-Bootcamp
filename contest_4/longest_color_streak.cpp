#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;
    vector<long long> arr(n);

    for (auto &it: arr) {
        cin >> it;
    }

    long long longest_streak = n > 0 ? 1 : 0;
    long long streak = 1;
    for (int i = 1; i < n; i++) {
        if (arr[i - 1] == arr[i]) {
            streak++;
            longest_streak = max(longest_streak, streak);
        }
        else {
            streak = 1;
        }
    }

    cout << longest_streak << "\n";

    return 0;
}