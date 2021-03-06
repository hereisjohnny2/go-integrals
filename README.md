# Go Integrals

## About

Go Integrals is an implementation of numerical integration algorithms in GoLang. It's mainly an library containing mathematical basic 2D-functions and the integrations algorithms it self.

So far the following functions are available:

- Polynomial _(f(x) = c0 + c1\*x +c2\*x^2 + ... + cn\*x^(n-1))_
- Exponential _(f(x) = a\*e^(b\*x))_

And the following integration algorithms:

- [Trapezoidal Rule](https://en.wikipedia.org/wiki/Trapezoidal_rule)

There is also a Line "function" to work with a collection of bi-dimensional points. It's possible to calculate the area under a curve formed by those lines, plot a graphic, read and save data, and calculate its length.

Any other function can be use in the algorithm as long as they implement the `Function` interface.

### To Do

- [ ] Create develop more functions and add math operations between then
- [X] Plot graphs
- [X] Save and Load points
- [X] Integrate over an slice or array of points

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
