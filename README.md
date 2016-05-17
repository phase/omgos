# omgos
`omgos` is a small HTTP Server used for running commands for certain files
types.

If your `commands.json` contains:

```json
{
    "py": "python $file"
}
```

Then accessing `server:80/hello.py` will run `python hello.py` and the output
will be sent to the client.

## config.json
`omgos` can be configured with a local `config.json`.

* `blocked`: Requests matching any files in the array will return a _403 Forbidden_.

  ```
  "blocked": ["config.json", "commands.json", "src/(.*)"]
  ```

  This blocks `config.json`, `commands.json`, and anything in the `src` folder. Full regex is supported.

## pygos
Included in this repository is a Python embedder called `pygos`. This will
read a file and run any Python code within `<% %>`.

```python
<h1>Awesome stuff</h1>
<p> Python Version: 
<%
import sys
print(sys.version,end='')
%>
</p>
```

This will output the Python version as well as the surrounding HTML.