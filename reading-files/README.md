# Reading from files

* Reading the contents of a file can be achieved in two different ways
  - reading whole file at a time
  - reading files by chunks, lines, etc

## Reading whole file

* The `ReadFile()` function in `io/ioutil` package can be used to achieve this.
* The `ReadFile()` opens the file, reads it and closes it, all at once. This does this by reading the whole contents of the file into memory.
* Reading a file at once has its perks and quirks, for example, if the file size is large or unknown - reading the contents into the memory might cause memory shortages and runtime errors.

## Reading in chunks

* This type of functions read chunks of the file and use the same memory allocated repeatedly. This helps in getting the information from large files or the files that cannot be processed in single go.
