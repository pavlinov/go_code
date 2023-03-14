## File System Manipulation task

1. Create bash script that will create initial file structure for File System Manipulation task
2. Modify script to take the number of folders to be created as script argument, use default value=2 if not provided
3. Modify script to read environment variable with the number of files, use default value=2 if it is not set
4. Modify script to :
   - check if at least one of parameters (folder cnt and files cnt) is provided
   - check if the number of input arguments is at most 1
   - show clear error message and stop execution if above conditions are not met
5. Modify script to take optional parameter –remove=<folder_name>:
   - remove folder under the folder_src directory
   - show message to clarify whether folder was actually remove or not

## Build
```shell
go build create_file_structure.go
```

## Running

```shell
FOLDERS_CNT=6 ./create_file_structure --remove folder_4 --foldersCnt=5 --filesCnt=2
```
```
Going to create 5 folders with 2 files
Created: 5 folders and 2 files structure created successfully!
Folder folder_src/folder_4 removed successfully!
```
or

```shell
FOLDERS_CNT=6 ./create_file_structure --remove folder_4 --foldersCnt=2 --filesCnt=3       
```
```
Going to create 2 folders with 3 files
Created: 2 folders and 3 files structure created successfully!
Error in remove
stat folder_src/folder_4: no such file or directory
```



### Running tests

```shell
go test -v  
```

```
=== RUN   TestGetIntEnv
--- PASS: TestGetIntEnv (0.00s)
=== RUN   TestRemoveFolder
Folder test_dir removed successfully!
--- PASS: TestRemoveFolder (0.00s)
=== RUN   TestCreateFileStructure
=== RUN   TestCreateFileStructure/0-test_create_2,2
Going to create 2 folders with 2 files
Created: 2 folders and 2 files structure created successfully!
=== RUN   TestCreateFileStructure/1-test_create_1,1
Going to create 1 folders with 1 files
Created: 1 folders and 1 files structure created successfully!
=== RUN   TestCreateFileStructure/2-test_create_2,1
Going to create 2 folders with 1 files
Created: 2 folders and 1 files structure created successfully!
=== RUN   TestCreateFileStructure/3-test_create_0,0
--- PASS: TestCreateFileStructure (0.00s)
    --- PASS: TestCreateFileStructure/0-test_create_2,2 (0.00s)
    --- PASS: TestCreateFileStructure/1-test_create_1,1 (0.00s)
    --- PASS: TestCreateFileStructure/2-test_create_2,1 (0.00s)
    --- PASS: TestCreateFileStructure/3-test_create_0,0 (0.00s)
PASS
ok      create_file_structure_p 0.394s

```
