#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n1;
    cin >> n1;
    vector<long long> arr1(n1);
    for (auto &it: arr1) {
        cin >> it;
    }

    long long n2;
    cin >> n2;
    vector<long long> arr2(n2);
    for (auto &it: arr2) {
        cin >> it;
    }

    set<long long> st1;
    for (auto it: arr1) {
        st1.insert(it);
    }

    set<long long> st2;
    for (auto it: arr2) {
        st2.insert(it);
    }
    
    set<long long> intersection;
    for (auto it: st1) {
        if (st2.find(it) != st2.end()) {
            intersection.insert(it);
        }
    }

    cout << intersection.size() << '\n';
    for (auto it: intersection) {
        cout << it << " ";
    }
    cout << '\n';

    return 0;
}