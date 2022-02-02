# simple-ps
more simple "ps" (unix)

Works on windows and linux but be mindfull that for windows you have to include ".exe" in the name to use the -n arg

# Arguments

```better-ps [--name|-n <name>]```

or 

```Usage: better-ps [--pid|-p <name>]```

# Use Cases
To find the PID of a program usally you would have to do somthing like this...

```PID='ps -ef | grep program | grep -v grep | awk '{print $2}''```

now you can do this...

```PID='./better-ps -p program'```

You can also you it to echo programs with a certain name

```./better-ps -n program```
