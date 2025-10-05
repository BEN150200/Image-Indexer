# Image Indexer

This is a personal project to learn the basics of Go and help myself organize and search all the pictures acomulated over the years

The program don't modify the files in any way, it only indexes the files using it's absolute path and stores it's metadata to search them.


## Dependencies

This are the main dependencies of the project.

- [goexif](https://github.com/rwcarlsen/goexif) - Used to read the metadata of the images
- [bbolt](https://github.com/etcd-io/bbolt) - Simpel storage package used as a database
