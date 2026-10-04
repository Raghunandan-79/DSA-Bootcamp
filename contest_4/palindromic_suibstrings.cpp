#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;
    string s;
    cin >> s;

    for (long long start = 0; start < n; ++start) {
        for (long long end = start; end < n; ++end) {
            bool palindrome = true;
            for (long long left = start, right = end; left < right; ++left, --right) {
                if (s[left] != s[right]) {
                    palindrome = false;
                    break;
                }
            }

            if (palindrome) {
                cout << s.substr(start, end - start + 1) << '\n';
            }
        }
    }

    return 0;
}