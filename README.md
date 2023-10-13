# Flow - A Chaining Mechanism for Go

_Flow is an experiment to check if go generics give more room for code obfuscation and a known problem from other languages - toxic senior++ code, only readable by experts with proper context, overcomplicated to be scary for newcomers. The answer is undoubtedly yes; it gives developers a chance to write such sophisticated constructs, so it's terrifing but remail "legal" code. Anyway, I am still determining if it's necessary to be so minimalistic to disallow doing bad things._

> Also, see my old Python experiment on the same topic: https://github.com/machineandme/sneact __(it was a joke)__. TLDR; Python with HTML embeddings. That is a pure Python React-JSX-like library that tests the limits of syntax parser of Python interpreter by ["operator overloading"](https://en.wikipedia.org/wiki/Operator_overloading).

Flow is a simple and lightweight Go package that provides a chaining mechanism for functions (Functors) that operate on values and errors. It allows you to create chains of functions and apply them to a value, making it easy to implement flexible error handling and value transformation pipelines.

## Features

- Create chains of functions that process values and errors sequentially.
- Conditionally apply functions based on the presence or absence of errors.
- Define default values or actions to handle errors in a customizable way.
- Easily combine and extend chains of functions.
