#include <bits/stdc++.h>
using namespace std;

int main() {
    int q;
    cin >> q;

    set<long long> st;

    while (q--) {
        int t;
        cin >> t;

        if (t == 1) {
            long long x;
            cin >> x;
            st.insert(x);
        }
        else if (t == 2) {
            long long x;
            cin >> x;
            st.erase(x);
        }
        else if (t == 3) {
            cout << st.size() << "\n";
        }
        else if (t == 4) {
            long long x;
            cin >> x;

            if (st.count(x))
                cout << "YES\n";
            else
                cout << "NO\n";
        }
    }

    return 0;
}