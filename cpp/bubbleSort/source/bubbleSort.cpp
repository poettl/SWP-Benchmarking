#include "../header/bubbleSort.h"

void swap(int* x, int*y){
    int temp = *x;
    *x = *y;
    *y = temp;
}

void bubbleSort(int data[], int size){
    int i, j;
    bool swapped;

    for(i = 0; i < size-1; i++){
        swapped = false;
        for(j = 0; j < size-1;j++){
            if(data[j] > data[j+1]){
                swap(&data[j], &data[j+1]);
                swapped = true;
            }
        }

        if(!swapped){
            break;
        }
    }
}