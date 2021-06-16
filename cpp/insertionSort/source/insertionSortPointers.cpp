//
// Created by flori on 16.06.2021.
//

#include "../header/insertionSortPointers.h"

void insertionSortPointers(int *data[], int size) {
    int i, key, j;
    for (i = 1; i < size; i++) {
        key = *data[i];
        j = i - 1;

        while (j >= 0 && *data[j] > key) {
            *data[j + 1] = *data[j];
            j--;
        }
        *data[j + 1] = key;
    }
}