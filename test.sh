#!/bin/sh

curl -v -X POST \
  http://localhost:8080/discountSettings \
  -H 'content-type: application/json' \
  -d '
    {
      "bookCost": 800,
      "discountableBookIds": [3, 47, 83, 133, 194],
      "discountScaling": [
        { "nbOfBooks": 2, "discountPercentage": 5},
        { "nbOfBooks": 3, "discountPercentage": 10},
        { "nbOfBooks": 4, "discountPercentage": 20},
        { "nbOfBooks": 5, "discountPercentage": 25}
      ]
  }
'

curl -v -X POST \
  http://localhost:8080/cost \
  -H 'content-type: application/json' \
  -d '
    { "basketBookIds": [1, 3, 3, 47, 83, 133, 194] }
'
