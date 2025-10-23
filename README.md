# Blog Aggregator  

## Description  
This program allows you to create users which can follow RSS feeds and get blog posts from them.

## Dependencies

- Postgres
- Go

You will need Postgres and Go installed on your system to run this application

You also need to set up a .gatorconfig.json file in your home directory with the following contents:

{  
 "db_url":"[VALUE]",  
 "current_user_name":"[VALUE]"    
}

## Build Process

To run the application you first need to build it using the build.sh script  
Then you should be able to run the executable like so gator [command]

## Usage

In order to truly test the program and its capabilities you will need to run commands in a specific order for the best experience.  
I suggest you follow these steps:  

- gator addfeed [feed name] [url]
- gator register [username]
- gator follow [feed url]
- gator agg [time duration] (format: 00s00m00h)
- gator browse [num posts] (just a number)
