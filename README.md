## path values

The three examples to go with the blog post at https://hawkins.id.au/posts/go-mux/path-values

These should all work identically, each just shows a different way of achieving
the same thing - extracing a user id from the path and making the user struct
(instantiated from some sort of database, faked in common/db/db.go for this
example) available to the route handler.

    go run go-mux/pathvalues/1-basic

    go run go-mux/pathvalues/2-slightly-less-basic

    go run go-mux/pathvalues/3-middleware