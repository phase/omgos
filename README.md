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