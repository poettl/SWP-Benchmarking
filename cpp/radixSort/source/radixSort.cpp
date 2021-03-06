#include "../header/radixSort.h"

int getMaxNumber(int data[], int size) {

    // we need to know the biggest number in the dataset, so search it
    int maxNumber = data[0];

    for (int i = 0; i < size; i++) {
        if (data[i] > maxNumber) {
            maxNumber = data[i];
        }
    }
    return maxNumber;
}

void sortDigits(int data[], int exponent, int size) {
    vector<int> buckets[10];

    for (int i = 0; i < size; i++) {

        int pos = (data[i] / exponent) % 10;
        buckets[pos].push_back(data[i]);
    }

    vector<int> mergedBuckets; 

    for (int i = 0; i < 10; i++) {
        vector<int>::const_iterator iterator;

        iterator = buckets[i].begin();
        while(iterator !=  buckets[i].end()) {
            mergedBuckets.push_back(*iterator);
            
            iterator++;             
        }
    }

    for (int i = 0; i < size; i++) {
        data[i] = mergedBuckets[i];
    }
   
}

void radixSort(int data[], int size) {
    int maxNumber = getMaxNumber(data, size);

    for (int exponent = 1; maxNumber/exponent > 0; exponent *= 10) {
        sortDigits(data, exponent, size);
    }
}




