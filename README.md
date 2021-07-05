# SocialNN
## Example of Neural Network models of social and personality psychology phenomena ##

This repository gathers a collection of neural network models of various social and personality psychology phenomena built by Stephen J. Read in the new *go* based version of the Leabra neural network modeling architecture.  This software architecture has been built by Randy O'Reilly at the University of California, Davis.  

Information about the Leabra modeling system can be found at the *emergent* Wiki:

https://github.com/emer/emergent/wiki

and in the Leabra repository:

https://github.com/emer/leabra


The repository currently has the following models
- **StereotypeModel** - This is a simple example model of stereotyping/classification of different individuals
- **PersonPerception** - This a *go* implementation of a constraint satisfaction model of PersonPerception based on Freeman and Ambady's person construal model in their 2011 *Psychological Review* Article.
- **IRmodel** - This is a neural network implemention of Cunningham and Zelazo's Iterative Reprocessing model of social evaluation.  
- **BeanFest** - This is an implementation of Eiser and Fazio's BeanFest reward learning game, which models reinforcement learning in a game in which individuals need to learn whether different beans are good or bad.
- **Personality** - This a version of Read and Millers (2021a,b) Virtual Personality Model, built in *go*.


Each folder includes all the files necessary to build and run the model for that folder. You can edit the *go* files and the data files to create variants of the program, or you can use it as the basis of a new model.  In order to build a program from *go* files you need to have *go* and *emergent* installed.  

In addition, the releases on this page contain already compiled executables for each of the programs.**NOTE:** *Currently there are only executables available for Mac, Windows executables are planned, but not available.* You do not need either *go* or *emergent* installed to run the executables.  To run the executables you need to download them to your computer and unzip them.  You then can execute them by using the *Terminal* program on Mac. You need to `cd` to the directory containing the files and then you can execute the programs. (On Mac you can type `cd` at the command line, type a space and then in the Finder drag the directory to the command line after the `cd` command.)  This will automatically fill in the path to the directory. Then hit return.  Once you have switched to the relevant directory, you then need to type

`./projectname`

at the command line(where `projectname` is the name of the program) and it will execute the program and load the relevant files in the directory.  The `./` before the projectname is necessariy to make sure that the program will load the relevant files. If you neglect the `./` it will start properly, but when you hit `Init` to initialize the program it will crash.  
