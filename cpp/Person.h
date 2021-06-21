//
// Created by flori on 16.06.2021.
//

#include <string>
#include<vector>

#ifndef SORTING_PERSON_H
#define SORTING_PERSON_H

struct Person{
    int age;
    std::string name;
    long height;
    long weight;

    unsigned long long int array1[1023];

    Person& operator =(const Person& a){
        age = a.age;
        name = a.name;
        height = a.height;
        weight = a.weight;

        std::copy(std::begin(a.array1), std::end(a.array1), std::begin(array1));

        return *this;
    }
};
#endif //SORTING_PERSON_H