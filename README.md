# score_averager

> *Written by: Joshua Fuentes <joshuarpf@gmail.com>*

This is a simple API backend implementation written in Go. This is my submission for a coding test that is aimed to demonstrate my knowledge in the domain, the approaches that I generally take in programing and the code that I produce. 

## Overview

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


