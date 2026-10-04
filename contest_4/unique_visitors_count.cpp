#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;
    vector<long long> arr(n);

    for (auto &it: arr) {
        cin >> it;
    }

    set<long long> st;
    for (long long visitor : arr) {
        st.insert(visitor);
        cout << st.size() << ' ';
    }
    cout << '\n';

    return 0;
}