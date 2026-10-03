#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n1;
    cin >> n1;
    vector<long long> arr1(n1);
    for (auto &it: arr1) {
        cin >> it;
    }

    multiset<long long> st;
    for (auto it: arr1) {
        st.insert(it);
    }

    long long n2;
    cin >> n2;
    vector<long long> arr2(n2);
    for (auto &it: arr2) {
        cin >> it;
    }

    multiset<long long> intersection;
    for (auto it : arr2) {
        auto pos = st.find(it);

        if (pos != st.end()) {
            intersection.insert(it);
            st.erase(pos);   // remove ONE occurrence
        }
    }

    cout << intersection.size() << "\n";
    for (auto it: intersection) {
        cout << it << " ";
    }

    return 0;
}