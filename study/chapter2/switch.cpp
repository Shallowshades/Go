#include <iostream>

using namespace std;

int main() {

    int val = 5;
    //按照代码顺序排列(包括default)，当满足某个case条件时进入，
    //若该case没有break，顺序执行其下面的case(无需条件判断)，直到遇到break或执行完全部剩余case终止
    switch (val) {
        default:
            cout << "default" << endl;
        case 1:
            cout << "1" << endl;
        case 2:
            cout << "2" << endl;
        case 3:
            cout << "3" << endl;
            break;
        case 4:
            cout << "4" << endl;
    }
    return 0;
}