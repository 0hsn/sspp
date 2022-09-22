## Command reference

[See usages](usage.md)

```
NAME
       sspp -- Structured Stream Processing and Presentation

DESCRIPTION
        sspp parse and presentation utility for structured documents, e.g json.

        If sspp is used in pipe, then confirm that --data option is not present.
        Otherwise data moved from previous command by pipe will be omitted.

        Multiple selector options (--json and --xml) can not exists on same command
        execution.

SYNOPSIS
       sspp [OPTIONS...]

OPTIONS
        --or        Set a default value in case the expected not found

        --data      Set data to parse, and to select data from.
                    This option get priority over data send via pipe.

    -J, --json      Node selector for json data

    -X, --xml       Node selector for xml data
```
