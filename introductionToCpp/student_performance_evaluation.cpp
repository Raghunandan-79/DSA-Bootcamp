#include <bits/stdc++.h>
using namespace std;

int main() {
    int marks;
    cin >> marks;

    if (marks > 90) {
        cout << "Excellent" << endl;
    }
    else if (marks > 80) {
        cout << "Good" << endl;
    }
    else if (marks > 70) {
        cout << "Fair" << endl;
    }
    else if (marks > 60) {
        cout << "Meets Expectations" << endl;
    }
    else {
        cout << "Below Par" << endl;
    }

    return 0;
}