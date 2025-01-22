<table align="center">
    <tr style="text-align: center;">
        <td align="center" width="9999">
            <img src="./.etc/karma.png" alt="Project icon" style="max-width: 500px; margin: 30px auto 15px auto; display: inline-block">

 <h1 style="color: black;">Karma API</h1>

<p style="color: black">A `simple` API capable to manage a Karma counter by user and its assigner</p>
</td>
</tr>
</table>

## Requirements to start developing ▶▶

1. You'll need also an IDE, take a look to [Goland](https://www.jetbrains.com/go/)
   or [Visual Studio Code](https://code.visualstudio.com/).

2. You need to have at least *Go v1.22* installed in your machine by following
the [download and install guide](https://go.dev/doc/install).

3. [casey/just](https://github.com/casey/just) to be able to run pre-made commands
which lives in the [justfile](./justfile)

4. [docker](https://docs.docker.com/get-started/get-docker/) to be able to bootstrap
the service whit its infrastructure dependencies

5. Install all the dependencies by running

```bash
go mod tidy
```
If something goes wrong try running
```bash
go mod tidy -compat=1.21
```

6. Bootstrap all those infrastructure pieces needed by running
```bash
just start
```
This command will generate the starter environment file and using
docker it will spin up all the infrastructure services.

7. Let's run our service by running
```bash
just run-karma-api
```

8. If you want to know all the available commands just run
```bash
just help
```
