#include "RadixSortPointer.h"

int getMaxNumber(Person data[], int size) {

    // we need to know the biggest number in the dataset, so search it
    int maxNumber = data[0].age;

    for (int i = 0; i < size; i++) {
        if (data[i].age > maxNumber) {
            maxNumber = data[i].age;
        }
    }
    return maxNumber;
}

void sortDigits(Person data[], int exponent, int size) {
    vector<Person> buckets[10];

    for (int i = 0; i < size; i++) {

        int pos = (data[i].age / exponent) % 10;
        buckets[pos].push_back(data[i]);
    }

    vector<Person> mergedBuckets; 

    for (int i = 0; i < 10; i++) {
        vector<Person>::const_iterator iterator;

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

void radixSort(Person data[], int size) {
    int maxNumber = getMaxNumber(data, size);

    for (int exponent = 1; maxNumber/exponent > 0; exponent *= 10) {
        sortDigits(data, exponent, size);
    }

    for (int i = 0; i < size; i++) {
      printf("%d \n", data[i].age);
    }
}



int main(int argc, char ** argv) { 
    int foo [7] = { 16, 2, 77, 40, 200, 111, 112 }; 
   
    radixSort(foo, 7);

    return 0;
}





