#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;
    vector<long long> arr(n);
    
    for (auto &it: arr) cin >> it;

    set<long long> st;
    for (auto it: arr) {
        st.insert(it);
    }

    cout << st.size() << "\n";

    return 0;
}