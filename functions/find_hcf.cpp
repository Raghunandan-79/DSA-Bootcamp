#include <bits/stdc++.h>
using namespace std;

int hcf(int a, int b) {
    while (a != 0 && b != 0) {
        if (a > b) a %= b;
        else b %= a;
    }

    return (b == 0) ?  a : b;
}

int main() {
    int a, b;
    cin >> a >> b;

    cout << hcf(a, b) << endl;

    return 0;
}