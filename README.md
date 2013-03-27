# img-ext

Extensions for [img][].

The usage instructions will assume that you have [img][] installed, but that is
not strictly necessary. It is possible to run these with without you just need
to put a `-` between `img` and their name, for instance run `$ img-databend`
instead of `$ img databend`. The only negative is that you will not be able to
display help for the command.


## img databend

Install and use like,

``` bash
$ go install github.com/hawx/img-ext/img-databend
$ img databend < input.bmp > output.bmp
```

![](http://hawx.github.com/img-ext/databend.jpg)


## img wlsn

Install and use like,

``` bash
$ go install github.com/hawx/img-ext/img-wlsn
$ img wlsn < input.png > output.png
```

Named for Richard Wilson, creator of ["Turning the Place Over"][turning].

![](http://hawx.github.com/img-ext/wilson.jpg)

[img]: http://github.com/hawx/img
[turning]: http://www.richardwilsonsculptor.com/projects/turning_the%20_place_over.html
