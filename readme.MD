# Challenge Test Ecommerce Backend

## Ecommerce documentation
- Entity Relationship Digram (erd): https://app.diagrams.net/#G1MJin66fmBRSNOQWQ_QxVG_Fu1zFG89iX

## Explanation
### conf folder
- The conf folder contains the load .env file folder, inside which there are several init envs, such as getenv for ```postgres``` or ```user admin```. This folder makes it easier for me when I want integration with vault or other things.
### infra folder
- I use the infra folder itself to carry out initialization when the application is run, such as ```postgres, zerolog``` or other things such as ```kafka redis and anything``` related to external connections
### layering folder (repository, usecase, presentation)
- File naming at the repository layer and at the usecase layer follows the Java naming, the difference is that it follows the interface name. 
- Other than the file name, this is just a tool for another repository or usecase. example repository/read.go. This is a tool for other repositories that does readQuery so there isn't a lot of duplicate code
  - #### repository layer
    - This repository is the innermost layer, I usually use this to connect it to the database, or integrate with other parties. be it communication between services or like other third parties
  - #### usecase layer
    - This usecase is the business logic layer, everything is managed here. just like when checking out we have to check whether the user is valid or not, check stock and whatever. 
  - #### presentation layer
    - This presentation, it is the outermost layer, inside it there is a ```rapi``` folder or restfull api. why did I do that? I think it makes it easier to separate the types of API protocols. The folder schema itself is a struct for receiving request bodies from clients
### test folder
- test file naming, repository_NameInterface.go. The goal is to make it easier to differentiate between repositories and usecases.
- for integration tests, all repository layer or usecase tests are run on 01_init_test.go. This is because integration tests depend on existing tests, so it aims to make it easier to sort which ones should be run first.
- Why do his names start with 00 or 01, that's because he's at the top
### helper folder
- Folder helper is a tool used for all layers, be it repo, usecase or presentation
### migration
- I have created a goose migration that runs on cmd from the go binary directly, which aims to make it easier to migrate without having to run other cmd such as ```goose up``` or others. If you want to check it, please run ```make migrate``` or ```go run main.go migrate```
- You can even run the compiled go binary to carry out the migration. for example, run ```./main.go migrate```
### ci/cd
- I have prepared ci/cd on github action. ci itself aims to carry out integration tests and unit tests, cd itself is for carrying out deployment
- if ci fails then cd will not run