#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;

    int total_lines = 2 * n - 1;
    
    for (int i = 0; i < total_lines; i++) {
        int line_num;
        if (i < n) {
            line_num = i;
        } 
        else {
            line_num = total_lines - i - 1;
        }
        
        for (int j = 0; j < line_num; j++) {
            cout << " ";
        }
        cout << ">";
        
        if (line_num > 0) {
            for (int j = 0; j < 2 * line_num - 1; j++) {
                cout << " ";
            }
            cout << ">";
        }
        
        cout << "\n";
    }

    return 0;
}