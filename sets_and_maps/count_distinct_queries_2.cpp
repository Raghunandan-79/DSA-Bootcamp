#include <bits/stdc++.h>
using namespace std;

int main() {
    int q;
    cin >> q;

    unordered_map<long long, int> freq;

    while (q--) {
        int t;
        cin >> t;

        if (t == 1) {
            long long x;
            cin >> x;
            freq[x]++;
        }
        else if (t == 2) {
            long long x;
            cin >> x;

            if (freq.count(x)) {
                freq[x]--;

                if (freq[x] == 0) {
                    freq.erase(x);
                }
            }
        }
        else if (t == 3) {
            cout << freq.size() << "\n";
        }
        else if (t == 4) {
            long long x;
            cin >> x;

            cout << (freq.count(x) ? "YES" : "NO") << "\n";
        }
    }

    return 0;
}