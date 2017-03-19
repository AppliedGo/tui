/*
<!--
Copyright (c) 2017 Christoph Berger. Some rights reserved.

Use of the text in this file is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

Use of the code in this file is governed by a BSD 3-clause license that can be found
in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

<!--
NOTE: The comments in this file are NOT godoc compliant. This is not an oversight.

Comments and code in this file are used for describing and explaining a particular topic to the reader. While this file is a syntactically valid Go source file, its main purpose is to get converted into a blog article. The comments were created for learning and not for code documentation.
-->

+++
title = ""
description = ""
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2017-00-00"
publishdate = "2017-00-00"
draft = "true"
domains = [""]
tags = ["", "", ""]
categories = ["Tutorial"]
+++

Want to equip your terminal application with a nice visual user interface? Three TUI libraries got into a (small) contest to show their abilities.

<!--more-->

Console applications usually take some parameters at start, and maybe some more input through basic console I/O. And that's ok in most cases, though sometimes it would be great to have a visual user interface but without the size and complexity of a full-blown Web app.

[Text-Based User Interface](https://en.wikipedia.org/wiki/Text-based_user_interface) (or TUI) libraries meet this need. They bring panes, input, output, sometimes also mouse support, graphics, and audio, to your terminal.

This post evaluates three TUI libraries for Go. A small sample application with a few basic UI requirements gets implemented with each of these libraries. Then we can see, by comparing the code side-by-side, how each library works - how easy each step is, and how much code is required.


## The selection process

To get selected for the test, the libraries had to meet some basic requirements:

* They should be in a halfway stable state (that is, no "alpha" or "experimental" warning tags attached)
* They should not be abandoned (or look like being abandoned)
* They should provide a decent feature set (that is, they should not specialize on one specific functionality, like, for example, colored output or progress bars)
* They should have a reasonably good documentation or at least good demo code (as I don't like doing a trial-and-error style development based on reading just some API docs and trying to figure out how all these API functions work together)

Based on this I searched [GitHub](https://github.com), [LibHunt](https://go.libhunt.com), and [GolangLibs](https://golanglibs.com), checked through the [Awesome Go](https://awesome-go.com) list, and also did a sanity check via Google search.

The result consists of three projects that exist for quite some time and should be quite known by now (as far as I can tell):

* `nsf/termbox-go`
* `jroimartin/gocui` (based on termbox-go)
* `gdamore/tcell`


## The test application




## The code
*/

// ## Imports and globals
package main

/*
## How to get and run the code

Step 1: `go get` the code. Note the `-d` flag that prevents auto-installing
the binary into `$GOPATH/bin`.

    go get -d github.com/appliedgo/TODO:

Step 2: `cd` to the source code directory.

    cd $GOPATH/src/github.com/appliedgo/TODO:

Step 3. Run the binary.

    go run TODO:.go


## Odds and ends
## Some remarks
## Tips
## Links


**Happy coding!**

*/
