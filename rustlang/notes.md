
### Type Morphing mechanisms
* core::convert::From, this return a owned type
* core::convert::AsRef, this return a reference
* core::ops::Deref, transform with & operator

* core::convert::Into, consume self, convert to T

## Quirks

### in rust most of objects is a resource, while in js or java, most of objects is data.

for example in rust, if `read_dir` return a reference type `&DirEntry`, and also you have a 
func that take that `&DirEntry`, your natural thought would be writing a function
`fn return_entry(path: Path)->&DirEntry` right ?, but that returned pointer would be
recycled before you can make any use of it. so you have to make that `return_entry`
function take a callback param with type `Fn(&DirEntry)`.

so, it's kind weird to me.


### 


    use std::fs::{read_dir, DirEntry};
    use std::path::{PathBuf, Path};
    use std::convert::AsRef;
    use std::error::Error;
    use std::convert::From;


    fn file_to_direntry<T: AsRef<Path>>(filepath: T) -> Result<DirEntry, Box<Error>> {

        let path = filepath.as_ref();
        let pf = path.to_path_buf();
        if !pf.is_file() {
            return Err(From::from("not a file"));
        }

        let parent = pf.parent();

        match parent {
            Some(parent) => {
                let filename = try!(pf.strip_prefix(parent));
                path_to_entry(parent, filename)
            }
            None => path_to_entry(".", path),
        }
    }
    
    /*

    fn path_to_entry<A: AsRef<Path>>(path: A, filename: A) -> Result<DirEntry, Box<Error>> { 

    if i change func signature to above, the code would failed to compile, even though from 
    my perspective, it should be totally legal, but some how, rust ties the two parameter's type,
    that is it sees first is a str, so the second param it infered out should also be a string type.

    error msg it emmit on that line:

    ^^^^ expected str, found struct `std::path::Path`
    */
    fn path_to_entry<A: AsRef<Path>, B: AsRef<Path>>(path: A, filename: B) -> Result<DirEntry, Box<Error>> {

        let filename: &Path = filename.as_ref();
        for entry in try!(read_dir(path.as_ref())) {
            let entry = try!(entry);
            if entry.path().is_file() && entry.path() == filename {
                return Ok(entry);
            }
        }

        Err(From::from("no file found"))
    }


    #[cfg(test)]
    mod test {
        use super::file_to_direntry;
        use std::path::{PathBuf, Path};

        #[test]
        fn test_file_to_direntry() {
            if let Ok(_) = file_to_direntry(PathBuf::from("tests/resources/some.txt")) {
                assert!(true);
            }
        }
    }


## miscellaneous

* where clause take `Type`, not `&Type`

    ```rust
    // F in where clause would cause compiler to fail
    fn file_to_direntry<T, F>(filepath: T, cb: F) -> () 
      where T: AsRef<Path>, F: &Fn(Result<&DirEntry, Box<Error>>)->(){
    
    // change to:
    fn file_to_direntry<T, F>(filepath: T, cb: &F) -> () 
      where T: AsRef<Path>, F: Fn(Result<&DirEntry, Box<Error>>)->(){

    ```
