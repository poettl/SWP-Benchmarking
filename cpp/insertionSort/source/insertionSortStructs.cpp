//
// Created by flori on 16.06.2021.
//

#include "../header/insertionSortStructs.h"

void insertionSortStructs(struct Person data[], int size) {
    int i,j;
    struct Person key;
    for (i = 1; i < size; i++) {
        key = data[i];
        j = i - 1;

        while (j >= 0 && data[j].age > key.age) {
            data[j + 1] = data[j];
            j--;
        }
        data[j + 1] = key;
    }
}