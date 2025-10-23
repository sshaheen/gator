# Blog Aggregator

## Dependencies

- Postgres
- Go

You will need Postgres and Go installed on your system to run this application

You also need to set up a .gatorconfig.json file in your home directory with the following contents:

{  
 DbURL string `json:"db_url"`  
 CurrentUserName string `json:"current_user_name"`  
}

## Build Process

To run the application you first need to build it using the build.sh script  
Then you should be able to run the executable like so gator [command]

## Usage
Some examples:

- gator addfeed [feed name] [url]
- gator register [username]
- gator follow [feed url]
- gator agg [00s00m00h]
- gator browse [num posts] (just a number)
