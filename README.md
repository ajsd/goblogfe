# Blog Frontend

An API-only blog frontend server on [App Engine][appengine]. Blog entries are
stored as [Markdown][md] and it's up to a client to render HTML.

The service is currently limited to a few servers (CORS).

## Service methods:

### blog.Search
```json
{"q": "..."}
```

Queries for (lists) entries.

Queries can be:
1. `"*"`: Lists all
2. `"id:xxxxxxx"`: Lists exactly 1 entry with the given `id`, if one exists.
3. `"from:yyyy-mm-dd[ to:yyyy-mm-dd]"`: Lists entries within the date range for
   last updated date.

Each of these are mutually exclusive; specify exactly one. The `"*"` query is
matched first, then `id` and then `from[to]`. If `id` exists in the query
string, the remaining query is discarded.

### blog.Save
```json
{"md": "..."[, "id": xxxxxxx]}
```

`"md"` is the required Markdown string/text. `"id"` should only be specified
when updating an existing entity (the `id` must already exist).

### blog.Delete
```json
{"id": xxxxxxx}
```

Deletes an entry with the given `id`, if one exists.

Currently, only admins can use the `blog.Save` and `blog.Delete` methods.

[appengine]: https://developers.google.com/appengine/docs/go
[md]: http://daringfireball.net/projects/markdown/
