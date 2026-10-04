#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n, k;
    cin >> n >> k;
    vector<long long> arr(n);

    for (auto &it: arr) {
        cin >> it;
    }

    unordered_map<long long, long long> frequency;
    long long answer = 0;
    long long processed = 0;

    for (long long value : arr) {
        if (k == 0) {
            if (value == 0) {
                answer += processed;
            }
        } 
        else if (value % k == 0) {
            answer += frequency[value / k];
        }

        frequency[value]++;
        processed++;
    }

    cout << answer << '\n';

    return 0;
}