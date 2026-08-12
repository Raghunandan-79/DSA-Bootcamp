#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n, f;
    cin >> n >> f;

    if (n % f == 0) {
        cout << "Yes" << endl;
    }
    else {
        cout << "No" << endl;
    }

    return 0;
}