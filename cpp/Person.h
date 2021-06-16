//
// Created by flori on 16.06.2021.
//

#include <string>

#ifndef SORTING_PERSON_H
#define SORTING_PERSON_H

struct Person{
    int age;
    std::string name;
    long height;
    long weight;

    Person& operator =(const Person& a){
        age = a.age;
        name = a.name;
        height = a.height;
        weight = a.weight;
        return *this;
    }
};
#endif //SORTING_PERSON_H