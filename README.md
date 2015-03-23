# logrus-ltsv
[LTSV](http://ltsv.org/) Formatter for [Logrus(github.com/Sirupsen/logrus)](https://github.com/Sirupsen/logrus)

``` go
// sample

log := logrus.New()
log.Formatter = new(logrusltsv.Formatter)

// ...

log.WithFields(logrus.Fields{"foo": "bar", "fiz": feez}).Info("test")
```

Such an example might produce a log record like this:

``` go
time:2015-03-23T11:37:13+09:00  level:info  msg:test  foo:bar  fiz:feez
```
