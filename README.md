# lsgo

A tool to find and list files in a folder you don't really know the location of.

## Build the binary

1. Clone the repo
2. Go in the repo
3. Execute the following command :
`go build .`
4. Move the build binary to a folder that's in your _$PATH_ variable 

## Usage

To find the folder _test_ in your machine use the following command. If you have
more than 10 results it will show 10 results and you will be able to browse to a
next page. If you enter on a folder, it will list the entire folder.

`lsgo test`