#include <bits/stdc++.h>
using namespace std;

int main() {
    int hour, minute;
    cin >> hour;
    cin.ignore();
    cin >> minute;

    if (hour == minute) {
        cout << "Mogambo is happy" << endl;
    }
    else {
        cout << "Mogambo is sad" << endl;
    }

    return 0;
}