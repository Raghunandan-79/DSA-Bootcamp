#include <bits/stdc++.h>
using namespace std;

int main() {
    string firstName1, lastName1;
    cin >> firstName1 >> lastName1;
    string firstName2, lastName2;
    cin >> firstName2 >> lastName2;

    if (lastName1 == lastName2) {
        cout << "YES" << endl;
    }
    else {
        cout << "NO" << endl;
    }

    return 0;
}