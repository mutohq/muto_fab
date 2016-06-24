#muto_fab

This repo copies the content of source directory to destination directory.
It also includes the following operation :
Let we have :
 * [header.xyz] (https://github.com/mutohq/muto_fab/blob/master/doc-templates/header.tmpl) 
 * [footer.xyz] (https://github.com/mutohq/muto_fab/blob/master/doc-templates/footer.tmpl)
 * [one of body file.xyz] (https://github.com/mutohq/muto_fab/blob/master/doc-templates/file1.tmpl)
Then output will be :  
 * [output.abc] (https://github.com/mutohq/muto_fab/blob/master/document/file1.html)

 * [Sample I/O] (https://github.com/mutohq/muto_fab/blob/master/Sample.md)

## How it works 
 * [configfile] (https://github.com/mutohq/muto_fab/blob/master/config.json)
 ```
 {
    "sourceDir" : "/home/ashu/Desktop/intern/git/muto-fab/doc-templates",
    "destDir" : "/home/ashu/Desktop/intern/git/muto-fab/document",
    "headerTmpl" : "/home/ashu/Desktop/intern/git/muto-fab/doc-templates/header.tmpl",
    "footerTmpl" : "/home/ashu/Desktop/intern/git/muto-fab/doc-templates/footer.tmpl",
    "tmplExtInput" : "tmpl", 
    "tmplExtOutput" : "html",
    "otherfolderstocopy" : [
        "/home/ashu/Desktop/intern/git/muto-fab/lib"
    ]
}

 ```
* `sourceDir`:
    It is the path having header, footer, body and assests files.
* `destDir`:
    It is the path of directory where the result of the operation will be stored.
* `headerTmpl`:
    It is the path of header file .
* `footerTmpl`:
    It is the path of footer file .
* `tmplExtInput`:
    It is the the extension of the file that is supposed to be merged in one file. 
* `tmplExtOutput`:
    It is the extension of the output file.
    e.g. `header.tmpl + f1.tmpl + footer.tmpl = f1.html`
* `otherfolderstocopy`: 
    It is an array that contains list of different directories those need to be copied 
    in the `destDir`.

##Procedure for using
* clone / download [muto_fab] (https://github.com/mutohq/muto_fab)
* set the path in config.json as : 
    * `/..../muto-fab/(sourcedir) `
    * `/..../muto-fab/(destdir) `
    * `/..../muto-fab/[xyz]/(header) `
    * `/..../muto-fab/[xyz]/(footer) `
* run on terminal `go run sourceToDest.go config.json`
* [Sample I/O] (https://github.com/mutohq/muto_fab/blob/master/Sample.md)    