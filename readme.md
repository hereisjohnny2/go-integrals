# Go Integrals

## About

Go Integrals is an implementation of numerical integration algorithms in GoLang. It's mainly an library containing mathematical basic 2D-functions and the integrations algorithms it self. 

So far the following functions are available:
- Polynomial *(f(x) = c0 + c1\*x +c2\*x^2 + ... + cn\*x^(n-1))*
- Exponential *(f(x) = a\*e^(b\*x))*

And the following integration algorithms:
    - Trapezoid Rule

Any other function can be use in the algorithm as long as they implement the `Function` interface.the same is valid for the integration algorithms, that should implement `Integral`.

### To Do
- [ ] Create develop more functions and add math operations between then
- [ ] Plot graphs
- [ ] Save and Load points from a graph
- [ ] Integrate over an slice or array of points

## Getting Started

Add the package to your project:
```bash
go get github.com/hereisjohnny2/go-integrals
```

The `main.go` file in the project root is only for testing integration propose. It will be removed in the future.

## Tests 

On each polynomial function or integration folder run:
```bash
cd functions/exponential
go test
```
