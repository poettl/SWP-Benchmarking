//
// Created by flori on 15.06.2021.
//

#include "../header/insertionSort.h"

void insertionSort(int data[], int size) {
    int i, key, j;
    for (i = 1; i < size; i++) {
        key = data[i];
        j = i - 1;

        while (j >= 0 && data[j] > key) {
            data[j + 1] = data[j];
            j--;
        }
        data[j + 1] = key;
    }
}