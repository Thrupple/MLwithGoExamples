
# MLwithGoExamples

This repository contains practical examples from the book 'Machine Learning with Go'. All examples are designed to provide hands on experience and understanding of the core machine learning concepts.

## Chapter 3

The data file `time_series.csv` includes two columns, one for the predicted value and one for the observed value (floating point numbers). You can calculate three metrics from this data set:

 * Mean-squared error
 * Mean Absolute error
 * R-squared

To compute, run the following:

```go
 go run chapter3/mean_squared_error.go chapter3/time_series.csv
```

Alternatively, consider scenarios where observed and predicted data sets are categories. The categories in `labeled.csv` are 0, 1, and 2. You can see the count of true positives and false positives (where the observed category matches the predicted category) as follows:

```go
  go run chapter3/category_accuracy.go chapter3/labeled.csv
```

Consider testing against the '0' category:

 * If the predicted value and observed value are both 0, this is a true positive (TP)
 * If the predicted value is 0 but the observed value isn't, this is a false positive (FP)
 * If the predicted value isn't 0 but the observed is, this is a false negative (FN)
 * If neither the predicted value nor the observed value is 0, this is a true negative (TN)

You can thus create metrics:

 * Accuracy: The ratio of true predictions vs false predictions: (TP+TN)/(FP+FN+TP+TN)
 * Precision: The ratio of true predictions over all predictions: TP/(TP+FP)
 * Recall

To evaluate data, consider creating training and testing sets. You can subsample one data set into two distinct sets using the following command:

```
  go run chapter3/subsample.go chapter3/time_series.csv
```