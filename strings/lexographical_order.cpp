#include <bits/stdc++.h>
using namespace std;

int main() {
    string a, b;
    cin >> a >> b;

    if (a < b) {
        cout << "A" << endl;
    }
    else if (b < a) {
        cout << "B" << endl;
    }
    else {
        cout << "Equal" << endl;
    }

    return 0;
}