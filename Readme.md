# Linear Regression Package in Go

Simple least squares linear regression implemented in Go.

## TODO

* [x] Write tests
* [x] Split in packages (linearRegression, Vec, Arr)
* [ ] Make input data more generic.
    * [ ] Make input data either 1D or more
    * [ ] Use an interface that runs `AddOnes()` on either a `Vec` or an `Arr`.
* [ ] Multiple optimizers
* [ ] Be able to save the model
* [ ] Add metrics to estimate model quality: variance, R**2, F-score, P-value
