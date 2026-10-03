#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;
    vector<long long> arr1(n);
    for (auto &it: arr1) cin >> it;

    long long m;
    cin >> m;
    vector<long long> arr2(m);
    for (auto &it: arr2) cin >> it;

    set<long long> st;

    for (auto it: arr1) {
        st.insert(it);
    }

    for (auto it: arr2) {
        st.insert(it);
    }


    cout << st.size() << "\n";
    for (auto it: st) {
        cout << it << " ";
    }
    cout << "\n";

    return 0;
}