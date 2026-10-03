#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n, k;
    cin >> n >> k;
    vector<long long> arr(n);
    for (auto &it: arr) {
        cin >> it;
    }

    set<long long> st;
    for (auto x: arr) {
        if (st.count(k - x)) {
            cout << "TRUE" << "\n";
            return 0;
        }
        st.insert(x);
    }

    cout << "FALSE" << "\n";

    return 0;
}