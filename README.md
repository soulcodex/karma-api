<table align="center">
    <tr style="text-align: center;">
        <td align="center" width="9999">
            <img src="./.etc/karma.png" alt="Project icon" style="max-width: 250px; margin: 30px auto 15px auto; display: inline-block">

 <h1 style="color: black;">Karma API</h1>

<p style="color: black">A `simple` API capable to manage a Karma counter by user and its assigner</p>
</td>
</tr>
</table>

### How to use this repository ❓

As we saw during the lecture this repository has been divided by meaningful blocks
per branches in order to keep the focus on what really matter `simplicity`. 

That said, feel free to navigate throughout the repository branches and you'll be able to 
see the evolution of the API components and entrypoints starting from `feat/block-01` 
up to `feat/block-03-testing`.

### Meaningful branches 💡

* `feat/block-01` ➡️ Basic setup with environments vars and a basic JSON logger.

```bash 
git checkout feat/block-01
```

* `feat/block-02` ➡️ Added the HTTP router and another meaningful components to our API

```bash 
git checkout feat/block-02
```

* `feat/block-03-database` ➡️ Added database connections and SQL migrations

```bash 
git checkout feat/block-03-database
```

* `feat/block-03-testing` ➡️ Added integration, unitary tests and mocks generator

```bash 
git checkout feat/block-03-testing
```

### Bonus Tracks 🎁 

* [OpenAPI (openapi-go)](https://github.com/swaggest/openapi-go)
* [MongoDB Driver](https://github.com/mongodb/mongo-go-driver)
* [Apache Cassandra Driver](https://github.com/apache/cassandra-gocql-driver)
* [Firebase Firestore](https://pkg.go.dev/cloud.google.com/go/firestore)
* [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
* [Justfile](https://github.com/casey/just)
* [CloudEvents SDK](https://github.com/cloudevents/sdk-go)