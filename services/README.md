# [Golang.hr] Platform service package
Base service package for [Golang.hr Platform].

**NOTICE: Package under active development. No ETA atm...**

### Design Concept / TODO

  - [] Should have some service defaults such as name, version, description.
  - [] Should have router. Router in this case should support [gRPC] and HTTP adapters.
  Point is to have one handler definition that is accessible by various protocols
  - [] Should accept options
  - [] Should extend manager
  - []

### License

```
The MIT License (MIT)

Copyright (c) 2015 Golang Croatia

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

[Golang.hr]: <https://github.com/golanghr>
[Golang.hr Platform]: <https://github.com/golanghr/platform>
