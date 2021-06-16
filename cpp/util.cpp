//
// Created by flori on 15.06.2021.
//

#include "../header/util.h"
#include <iostream>

using namespace std;

void printArray(int arr[], int n) {
    int i;
    for (i = 0; i < n; i++)
        cout << arr[i] << " ";
    cout << endl;
}

void printPointerArray(int *arr[], int n) {
    int i;
    for (i = 0; i < n; i++)
        cout << *arr[i] << " ";
    cout << endl;
}

void printStructArray(Person arr[], int n){
    int i;
    for (i = 0; i < n; i++){
        cout <<  "age: " << arr[i].age << " name: " << arr[i].name << " | ";
    }

    cout << endl;
}