### Ideas

- Include faker?
- `-config` can also fetch urls that hold the configs (easier setup)
- [Go frameworks by stars](https://github.com/mingrammer/go-web-framework-stars)

### Tests

```
echo with modules
  empty-body fails, because returns validation field errors
gin with modules
  can't run tests because of the repeating url paths that crash the thing by default
```


### Commands

```bash
# update on save
reflex -r '\.go' -s -- sh -c "go run main.go"

# check ports
sudo lsof -P -i TCP -s TCP:LISTEN

# stop server on port 4444
kill $(lsof -t -i:4444)

# rename placeholder sql files from the root of the generated server
for i in ./db/sql/*.sql.gen.txt;  do mv "$i" "${i/.sql.gen.txt}.sql"; done

# test generated queries from db dir
cd .. 
for i in ./db/sql/*.sql.gen.txt;  do mv "$i" "${i/.sql.gen.txt}.sql"; done
cd db
sqlc generate

```

## References

- [sqlite example](https://github.com/bopbi/simple-todo/blob/master/simple-todo.go)
- [Go types](https://golangbyexample.com/all-basic-data-types-golang/)
- [fiber bodyparser thing](https://docs.gofiber.io/api/ctx#bodyparser)
- [software versioning](https://stackoverflow.com/questions/2864448/best-practice-software-versioning)
- [detecting last item inside for loop](https://stackoverflow.com/questions/50085038/detect-last-item-inside-an-array-using-range-inside-go-templates)
- [db conn thing](https://www.alexedwards.net/blog/organising-database-access)
- [db conn example](https://github.com/teten-nugraha/go-dev-productify/blob/60b53e9e7f985b9e349889e8ffdb37a11ab1bd7f/config/dbConnect.go)
- [new orm thing](https://bun.uptrace.dev/)
- [bash command to rename file extensions](https://osxdaily.com/2017/05/12/change-all-file-extensions-command-line/)