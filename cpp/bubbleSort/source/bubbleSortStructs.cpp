//
// Created by flori on 16.06.2021.
//

#include "../header/bubbleSortStructs.h"
void swapStructs(Person* x, Person*y){
    Person temp = *x;
    *x = *y;
    *y = temp;
}

void bubbleSort(Person data[], int size) {
    int i, j;
    bool swapped;

    for (i = 0; i < size - 1; i++) {
        swapped = false;
        for (j = 0; j < size - 1; j++) {
            if (data[j].age > data[j + 1].age) {
                swapStructs(&data[j], &data[j + 1]);
                swapped = true;
            }
        }

        if (!swapped) {
            break;
        }
    }
}