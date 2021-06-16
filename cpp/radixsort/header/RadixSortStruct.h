
#include <stdio.h>
#include <vector>
#include <iostream>
using namespace std;

struct Person {
    int age;
    string name;
    long height;
    long weight;
}

void getMaxNumber(Person data[], int size);
void sortDigits(Person data[], int exponent, int size);
void radixSort(Person data[], int size);
