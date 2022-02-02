# simple-ps
more simple "ps"

To find the PID of a program usally you would have to do somthing like this...

```PID=`ps -ef | grep program | grep -v grep | awk '{print $2}'```

now you can do this...

```PID=`./better-ps -p program```

You can also you it to echo programs with a certain name

```./better-ps -n program```
