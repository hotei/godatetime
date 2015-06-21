<center>
# godatetime
</center>

## OVERVIEW

I use this to embed the date and time when a program is compiled into the
go source code just before compilation.  This allows me to use it when 
printing version info. I tried to find something similar as a Go built-in but
stopped looking when I got this idea. I suppose it could be used to capture
other parts of the compile-time environment like memory size and CPU architecture
of the system doing the compiling, but date and time was all I was
interested in at the moment.

It's almost trivial but it has proven very useful so I'm sharing it. 

### Installation

If you have a working go installation on a Unix-like OS:

> ```go get github.com/hotei/godatetime```

Will copy github.com/hotei/godatetime to the first entry of your $GOPATH

or if go is not installed yet :

> ```cd DestinationDirectory```

> ```git clone https://github.com/hotei/godatetime.git```

### Features
* simple
* works as intended

### Limitations

* non known

### Usage

Typical usage is in two parts.  Note __goprog__ is the program being compiled;
N.B. the name given to output of godatetime isn't important, but should be
chosen to avoid overwriting other files.

First, in the goprog Makefile:

```
# Makefile for goprog

all:
	godatetime -pkg=main > compileDate.go
	go build
```

Somewhere in goprog.go you'd have a line similar to:

```go
fmt.Printf("Compiled on %s\n", CompileDateTime)
```

### BUGS
* none known

### To-Do

* Essential:
 * TBD
* Nice:
 * TBD
* Nice but no immediate need:
 * TBD

### Change Log
* 2015-06-12 validated with 1.4.2
* 2014-xx-xx Started

 
### Resources

* [go language reference] [1] 
* [go standard library package docs] [2]
* [Source for godatetime] [3]

[1]: http://golang.org/ref/spec/ "go reference spec"
[2]: http://golang.org/pkg/ "go package docs"
[3]: http://github.com/hotei/program "github.com/hotei/godatetime"

Comments can be sent to <hotei1352@gmail.com> or to user "hotei" at github.com.
License is BSD-two-clause, in file "LICENSE"

License
-------
The 'godatetime' go package/program is distributed under the Simplified BSD License:

> Copyright (c) 2015 David Rook. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

Documentation (c) 2014-2015 David Rook 

// EOF README-godatetime.md  (this is a markdown document and tested OK with blackfriday)
