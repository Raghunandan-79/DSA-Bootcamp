#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;
    long long copy_n = n;

    long long reversed_n = 0;

    while (copy_n > 0) {
        long long current_digit = copy_n % 10;
        reversed_n = (reversed_n * 10) + current_digit;
        copy_n /= 10;
    }

    if (reversed_n == n) {
        cout << "YES";
    }
    else {
        cout << "NO";
    }

    return 0;
}