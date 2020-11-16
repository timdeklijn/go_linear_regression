# Linear Regression Package in Go

Simple least squares linear regression.

## Example

How to run a linear regression:

``` go
func main() {
	// Data
	length := Vec{5, 5.5, 5.8, 6, 6.3, 6.5, 6.9, 7.1, 7.4, 7.7}
	weight := Vec{7, 7.5, 7.3, 8, 7.7, 7.9, 8.2, 8.8, 8.4, 8.3}

	// Setup
	x := length.AddOnes()
	lr := NewLinearRegression(x, weight)
	// Run
	lr.Fit()
	fmt.Println("\nThetas:", lr.thetas)

	var sample = Arr{Vec{1, 10}}
	fmt.Println("Prediction:", lr.Predict(sample))
}
```

## TODO

* [x] Write tests
* [x] Split in packages (linearRegression, Vec, Arr)
* [ ] Make input data more generic.
    * [ ] Make input data either 1D or more
    * [ ] Use an interface that runs `AddOnes()` on either a `Vec` or an `Arr`.
* [ ] Multiple optimizers
* [ ] Be able to save the model
* [ ] Add metrics to estimate model quality: variance, R**2, F-score, P-value
