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
$ go get hawx.me/code/img-ext/img-databend
$ img databend < input.bmp > output.bmp
```

## img wlsn

Install and use like,

``` bash
$ go get hawx.me/code/img-ext/img-wlsn
$ img wlsn < input.png > output.png
```

Inspired by Richard Wilson's, ["Turning the Place Over"][turning].

## img timeslice

Install and use like,

``` bash
$ go get hawx.me/code/img-ext/img-timeslice
$ img timeslice photos/* > output.png
```

## img lomo

Install and use like,

```bash
$ go get hawx.me/code/img-ext/img-lomo
$ img lomo < input.png > output.png
```

Based on [Lomography, UNIX Style][lomo].


[img]:     http://hawx.me/code/img
[turning]: http://www.richardwilsonsculptor.com/projects/turning_the%20_place_over.html
[lomo]:    http://the.taoofmac.com/space/blog/2005/08/23/2359
