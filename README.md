# score_averager

> *Written by: Joshua Fuentes <joshuarpf@gmail.com>*

This is a simple API backend implementation written in Go. This is my submission for a coding test that is aimed to demonstrate my knowledge in the domain, the approaches that I generally take in programing and the code that I produce. 

## Developer Notes

*Environment*
For the sole purpose of this exercise, I chose to setup a containerized development environment. I did this to demonstrate the use of a few things, such as knowledge in utilizing docker configuration  and features and a minimal setup of a Go development environment. I also think that it gives an additional advantage to developers since this allows working  on a "stripped down" enviroment as compared to managing the resources in a local machine. That beign said, here are the dependencies for included: 

*Dependencies*
 - CompileDaemon - a package that will allow something likea "live reload". I'm a fan of working like that, than having to compile each time I save. 
 - Gorilla/mux - router package

*Coding*
I am a bigtime fan of the SOLID principles in software development, but as much as possible, I also did not want to over-engineer in this task. Each file has comments describing it's purpose and focus, and generally, this is patterened from the MVC architecture. 

I went ahead and took a strict approach in decoding the JSON from requests, rather than the lenient way of unmarshalling. This is because it matches the requirements of this exercise, and as much as possible,  I would like to be in control of what happens. 

Here's a breakdown of the packages

 - **router**: contains files that serve purpose to data routing. It also includes integraiton of the gorilla/mux router. 
 - **payload**: broken down into input and output, and they all contain the structure of the data that passes accordingly
 - **model**: provides the data layer of the application. They contain files that are named accordingly, and serve only one purpose. It also has a helper file that contains the methods that are relevant to these models. 
 - **helper**: a general package that contains various implementations. They are also grouped into files named acordingly
 - **handler**: package that is responsible for handling the routed information. 
 - **app**: the package that does the main tasks of the application. This includes the *controller*. 


## App Overview

The general task is to accept a JSON input via HTTP request and it should return a structured output that contains hte results of call or the errors if there are any. 

*Sample JSON input*

    {"scores":{"managers":[{"userId":1,"score":1},{"userId":2,"score":5}],"team":[{"userId":4,"score":1},{"userId":5,"score":5},{"userId":6,"score":3},{"userId":7,"score":2}],"others":[{"userId":8,"score":1},{"userId":9,"score":5}]}}

*JSON output: (success)*

    {
	    "Success":true,
	    "Data":
	    {
		    "Scores":{
			    "Managers":3,
			    "Others":0,
			    "Team":2.75
		    }
	    },
	    "Errors":[]
    }

*JSON output: (fail)*

    {
	    "Success":false,
	    "Data":null,
	    "Errors":[
		    "Invalid score input",
		    "User: 1 of managers has an invalid score of 1100."
	    ]
    }


## Running the application

*Requirements*

 1. Linux or Linux based OS
 2. docker 
 3. docker-compose

*Installation guide*
1. Clone the repository. ( *Note: I was using gitflow for the duration of this project, and I think I made a slight change on renaming the "master" branch into "main". Rest assured it will have the lastest in it.* )
2. Run "**make**" from the directory where you cloned the application. If your system is configured correctly, this will build the necessary containers and all for you. 
3. Since this environment was **intentionally** created  for development purposes, you will notice that upon running, you will see the a flow logs on from your terminal. That is how I would normaly do it for development environments.
4. Once everything is running, you can then start  sending **POST** request to the URL **localhost:5000**. Feel free to use the sample input included in this document, or exeriment on your own. ( *Note: I would recommend using a tool like postman on this one, so it will be easier for you* )
5. To stop, simply step out of the container by pressing "**ctrl+c**" and running "**make clean**". This command shall do some housekeeping and prune your unused controllers, including the image built.


