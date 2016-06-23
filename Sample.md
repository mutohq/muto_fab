# Sample Input and Output


## Initial Directory structure
```
.
|-- doc-templates
|   |-- aa
|   |-- a.cpp
|   |-- assests
|   |   |-- abc
|   |   |-- File1
|   |   |   `-- one.txt
|   |   |-- File2
|   |   |   `-- one.txt
|   |   `-- xyz.cpp
|   |-- file1.tmpl
|   |-- file2.tmpl
|   |-- footer.tmpl
|   |-- header.tmpl
|   `-- yo
|       `-- done
|-- input.md
|-- lib
|   |-- abc
|   |   `-- des
|   |-- l1.txt
|   |-- l2.cpp
|   `-- xyz
|-- mainbkup.go
|-- main.go
|-- output.md
`-- README.md

12 directories, 15 files
```
## After Execution of source code
* go run main.go
```
.
|-- doc-templates
|   |-- aa
|   |-- a.cpp
|   |-- assests
|   |   |-- abc
|   |   |-- File1
|   |   |   `-- one.txt
|   |   |-- File2
|   |   |   `-- one.txt
|   |   `-- xyz.cpp
|   |-- file1.tmpl
|   |-- file2.tmpl
|   |-- footer.tmpl
|   |-- header.tmpl
|   `-- yo
|       `-- done
|-- document
|   |-- aa
|   |-- a.cpp
|   |-- assests
|   |   |-- abc
|   |   |-- File1
|   |   |   `-- one.txt
|   |   |-- File2
|   |   |   `-- one.txt
|   |   `-- xyz.cpp
|   |-- file1.html
|   |-- file2.html
|   |-- lib
|   |   |-- abc
|   |   |   `-- des
|   |   |-- l1.txt
|   |   |-- l2.cpp
|   |   `-- xyz
|   `-- yo
|       `-- done
|-- lib
|   |-- abc
|   |   `-- des
|   |-- l1.txt
|   |-- l2.cpp
|   `-- xyz
|-- mainbkup.go
|-- main.go
|-- output.md
`-- README.md

24 directories, 22 files
```