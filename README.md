# omgos
`omgos` is a small HTTP Server used for running commands for certain files
types.

If your `config.json` contains:

```json
{
    "py": "python $file"
}
```

Then accessing `server:80/hello.py` will run `python hello.py` and the output
will be sent to the client.