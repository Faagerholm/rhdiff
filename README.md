# Rdiff tool made with go

Rdiff is a tool for generating a signature file from files/folders. With the help of a signature file and a modified version of the original state(files) the original files can be reproduces. A Delta file can be generated with the help of the signature file and the modified state(files).

With the help of the delta file the original state can be synchronized to match the modified version. 
With the help of these signature/delta files less files need to travel over the internet thus saving bandwidth. 

For more information, see the implementation of [rdiff](https://linux.die.net/man/1/rdiff) and [rdiff-backup](https://rdiff-backup.net/)


This project is heavily based on c implementation of librsync  
https://librsync.github.io/


# Get started

Download a copy of this repository  
`git clone github.com/faagerholm/rhdiff`  
or  
install with `go install github.com/faagerholm/rhdiff`

### Test

Test the code with `go test ./... -v`  
if all tests pass (some are skipped for now) you can go ahead and build

### Build
you can then build it with 
`go build -o rhdiff cmd/local/*.go`



## Commands

```bash

> rhdiff signature -i <INPUT_FILE> -s <SIGNATURE_OUTPUT> (default: signature.bin)
> rhdiff delta -i <INPUT_FILE> -s <SIGNATURE_FILE> -d <DELTA_OUTPUT> (default: delta.sin) 

```

## The rolling sum hash algorithm

The rolling sum hash algorithm is implemented and the "local tool" uses this implementation untill the bugs gets fixed with the RK rolling hash algorithm. See Below for more information

## The Rabin Karp rolling hash algorithm

Rabin Karp hash algorithm uses a rolling hash instead of a 2 part rolling sum. As of now the Rotation of two bytes has a bug where the wrong end Hash is produced. This is a major bug which prevents us to use the RK algorithm of this implementation.

For more information how the rotation is implemented in the older version, have a look at [This file](https://librsync.github.io/rabinkarp_8h_source.html)

## Tests
No tests has been produces for the Delta conversion. These should be implemented to verify the functionality of our code.

# Contribution & Disclaimer

I'm happy to receive PR's and comments on part that could be improved.
You can find my contact information on my profile :)