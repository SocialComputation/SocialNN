# Description of Iterative Reprocessing(IR) Model 


This model is based on the one used in the published paper, although it does not have some of the more advanced features of that model:

Ehret, P. J., Monroe, B. M., & Read, S. J. (2015). Modeling the Dynamics of Evaluation: A Multilevel Neural Network Implementation of the Iterative Reprocessing Model. *Personality and Social Psychology Review, 19(2)*, 148-176.

This version differs from the published version in several ways. First, the current version does not use Lesioning. The results seem comparable to the published version, but they are not identical since they do not train the network in two stages and lesion the higher layers during learning of the early evaluations. Second, the model does not train in two stages, race and gender, then higher order information. Instead it trains everything at once. Third, it does not allow one to simulate the context effects that were simulated in the published paper.  

You can examine the connections among the different layers, by clicking on s.Wt next to the network graph and then clicking on nodes in the diagram.  Nodes that the clicked on node *sends to* will be highlighted.  You can also click on r.Wt and then nodes that the clicked on node *receives from* will be highlighted.  

This program can train and test multiple different networks at one time. 

## Testing of Iterative Processing of Evaluation

The purpose of the major simulation is to show that evaluations can develop over time as a result of iterative reprocessing in a network of neural systems. We have the model learn to identify a number of different kinds of individuals and to learn their evaluation. The model is constructed so that there is early, quick evaluative processing based on the analysis of early perceptual cues that allow one to quickly identify such things as Race and Gender, followed by later evaluative processing based on more detailed semantic processing of the type of individual. For example, if the model sees a Black Male Doctor, it first quickly processes Race and Gender cues and generates the associated evaluation, followed by more detailed semantic processing that identifies the individual as a Doctor in a Hospital, with a number of specific attributes, and the associated evaluation.

## Key Layers Features

**Input** - Features of the situation and the individuals

Female features, Male features, Dark Skin, Light Skin, Lab coat, Tie, Dress shoes/Slacks, Gang color band, Sweat pants, Gang Tattoos, Suit jacket, Street signs, Cars, Patient beds, Nursing station, Receptionist, Cubicles.

**Race and Gender** - classification of Race and Gender

Male, Female. Black, White.

**Conjunction layer** - Identification of four types of individuals: 

Black Female and Male; White Female and Male

**Context** - Clothing types and types of Situations

Street corner, Office building, Professional Clothing, Athletic clothes, Gang clothes, Hospital

**Profession** - Doctor, Businessman, Athlete, Gangster

Businessperson, Athlete, Doctor, Gangster

**Attributes** - attributes and personality traits

Caring, Athletic, Popular, Rich, Greedy, Violent, Intelligent, Unintelligent

**Positive Evaluation** (Scalar Layer)

**Negative Evaluation** (Scalar Layer)

The network has separate representations of Positive and Negative evaluation as there is psychological evidence that these are largely orthogonal dimensions in the evaluation process.

## Training

The model is set up so that it is possible to automatically train multiple instances, test them, and save the results, so that it is possible to look at results that are averaged across a number of different random starting values for the network. You can specify how many instances you can train by using the MaxRuns field on the left hand side of the gui.  The default number of Runs is 10.

There are two types of training cues in the Training file.

First, there are training items for four types of individuals: Black Male, White Male, Black Female and White Female.  The network is given the relevant Gender and Race cues and learns the associated category. Each of these individuals has an associated evaluation. The evaluations of Black Males and Females is negative, whereas the evaluation of White Males and Females is positive. 

Second, the network also learns about a number of new individuals for whom there is considerably more information. It learns about Doctors, Businessman, Athletes, and Gangsters,  with their Race, Gender, associated clothing, physical situations, personal attributes and their associated evaluations. There is a total of 20 different kinds of individuals.

This training information is in the IRTrain.tsv DataTable. In order to roughly capture the ecological frequency of these different kinds of individuals in the real world, we varied the frequency of occurrence of the different kinds of individuals. Obviously, the current frequencies could be edited as the user wishes, by changing copying and pasting additional instances in the DataTable. 

All the DataTables are tab-delimited text files and can be read in a spreadsheet program, such as Excel or text editor.

## Testing

Once the network has learned this detailed information it is then tested by presenting the network with each of the 20 individuals and then recording the change in evaluation over cycles of processing for each individual. The DataTable used is: IRTest.tsv.  

There are two ways to examine the development of evaluation over time.  First, one can hit the Test Trial button which will step through the testing information, one trial at a time in sequential order.  One can click on the Act and ActM buttons to show that in the NetView.  In order to see the development of activation more slowly one can used the 'VCR' type buttons on the bottom right of the NetView window to scroll back and forth through the activation for one trial.  


The current version of the program only records the data and plots it for one testing instance at a time. Results for single instances of the network are saved into the TstCycLog and plotted in the TstCycPlot graph window.  Each new testing trial will overwrite the previous results in the TstCycPlot window.


