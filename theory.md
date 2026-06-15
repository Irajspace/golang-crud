## Comparison with node
```
1. Dependency Management
Node.js: package.json and package-lock.json
Go: go.mod and go.sum

In Node, you run npm install express.

In Go, you run go get github.com/go-chi/chi/v5. It updates your go.mod (the list of packages) and go.sum (the exact cryptographic hashes of those packages to ensure security).


```
# internal vs pkg
```
internal = private to this repository/application

pkg = public library for other repositories
```

## overriding the env file
```
yaml:"env": "When you read the local.yaml file, look for a key named env and put its value here."

env:"ENV": "Also, check the computer's OS environment variables for something called ENV. If it exists, override what was in the YAML file."

env-required:"true": "If you check both the YAML and the OS, and this value is still empty, crash the program right now. Do not let the server start without this!"

```

## what does os.stat does 

```
. Checking if the file actually exists
You said: "os.stat so wecall os.fatal()it is just like printf but with an exit at last"
Verdict: 100% Correct.

Go
if _, err := os.Stat(configPath); os.IsNotExist(err) {
    log.Fatalf("config file does not exist...")
}
All the code before this step only gave us a String (the name of the path). Go doesn't know if that file actually exists on your hard drive yet! os.Stat goes to the hard drive and checks. If someone passed a path to a file that got deleted, log.Fatal prints the error and runs os.Exit(1) to kill the server.



```
##
```
Here is what this single line does:
1. It physically opens the `local.yaml` file.
2. It reads the text inside.
3. It looks at your `&cfg` struct and reads the struct tags (`yaml:"env"`, etc.).
4. It maps the text from the file into the variables inside your Go code.
5. It enforces your `env-required:"true"` rules. 

If all of that succeeds, it hands you back a fully populated `Config` object, ready for your HTTP server to use.


```
## if something is interface and we want to declare something of that type we never do * to the interface
## under the hood it is already pointing to some data

## the panic and 3 errors
```
The 3 Signals & Goroutine Panics
You asked what happens if the server panics, and what those three signals were.

If the server panics: (For example, port 8080 is already being used by another app). The ListenAndServe() function returns an error immediately. Your code says panic(err). The entire program instantly crashes and prints the error to your screen. It never even tries to send a message to the channel.

The 3 Signals: If it doesn't panic and runs smoothly, signal.Notify waits for:

os.Interrupt: A generic interrupt signal.

syscall.SIGINT: The specific signal sent when a human presses Ctrl+C in the terminal.

syscall.SIGTERM: The specific signal sent by Docker or Kubernetes when it wants to scale down your server.

context.WithTimeout(context.Background(), 5*time.Second), you take that blank walkie-talkie and attach a 5-second self-destruct timer to it. You hand that to the server and say, "You have 5 seconds to finish your active web requests, or I am pulling the plug."

```

## yaml file
```
The Go Code: Your MustLoad function is flawlessly checking the OS, falling back to the terminal flags, verifying the file exists, and handing it to cleanenv.

The Tags: Your struct tags (yaml:"address", yaml:"http_server", etc.) are perfectly set up to catch the data.

The Crash: When cleanenv physically opens the local.yaml file to read it, the YAML parser chokes on env:"dev" because without a space, it doesn't recognize it as a standard Key-Value pair.
```
## even u put configpath to be empty the cleanenv can read from the terminal itself
```
Read the File: It opens local.yaml and reads address: "localhost:8080". It temporarily writes 8080 into your struct.

Scan the Tags: It looks at your Go code and sees env:"http_server".

The Secret OS Check: It says, "Ah! The developer wants me to check the OS for this specific field!" cleanenv then automatically runs os.Getenv("http_server") itself, behind the scenes.

The Override: It finds your "localhost:9090" from the terminal, realizes it is more important than the YAML file, and overwrites the 8080
```
## hamdle multiple clients
```
If you actually open the source code for Go's standard library and look inside the net/http package, you will find an infinite for loop that accepts connections. Inside that loop, there is literally a line that looks like this:

Go
go c.serve(connCtx)
The net/http package is doing exactly what we did manually in the TCP Server lesson! It accepts the raw network connection, slaps the go keyword in front of the handler function, and spins it off into the background.

```