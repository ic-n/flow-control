# Flow - A Chaining Mechanism for Go

_Flow is an experiment to check if go generics give more room for code obfuscation and a known problem from other languages with many more features - master-only code, only readable by experts with proper context. The answer is undoubtedly yes; it gives developers a chance to write such sophisticated constructs, but I am still determining if it's necessary to be so minimalistic to disallow doing bad things. (also, see my old experiment https://github.com/machineandme/sneact on the same topic, but more visible for Python.)_

Flow is a simple and lightweight Go package that provides a chaining mechanism for functions (Functors) that operate on values and errors. It allows you to create chains of functions and apply them to a value, making it easy to implement flexible error handling and value transformation pipelines.

## Features

- Create chains of functions that process values and errors sequentially.
- Conditionally apply functions based on the presence or absence of errors.
- Define default values or actions to handle errors in a customizable way.
- Easily combine and extend chains of functions.
