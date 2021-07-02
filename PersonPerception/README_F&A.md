

# Introduction to Person Perception model

This project is based on a paper by Freeman, J. B., & Ambady, N. (2011). A Dynamic Interactive Theory of Person Construal. *Psychological Review, 118(2)*, 247-279. 

They outline a simple model of person perception that captures some of the cognitive dynamics of person judgments and relies heavily on constraint satisfaction processes.  This simulation is based on the model in their Figure 2. Note that in this simulation the weights are hard wired and stored in a file called `Person.wts`. You can edit the weights in this network, using a text editor, if you wish. This simulation does not have any learning.

This project explores a simple **semantic network** intended to represent a (very small) set of relationships among different features of a person and particular task instructions.  The features are Gender Features: male versus female features, and Race Features: Black, White, and Hispanic features. The categories are Gender: Male or Female and Race: Black, White, Asian.  The personality characteristis are Docile and Aggressive.  These two characteristics were chosen because they represent two characteristics that are differentially stereopically associated with different Genders and Races.  Men in general are viewed as more aggressive and less docile than women, while Blacks are viewed as most aggressive and least docile, and Asians are viewed as less aggressive and more docile.  Thus, characteristics that are more stereotypic of women (docile and not aggressive) are also more characteristic of Asians than of Black, or Whites.  

In the network we can represent 6 type of individuals: Black Female, Black Male, White Female, White Male, Asian Female, and Asian Male.  These individals vary in terms of how Docile and Aggressive they are.  In addition, we can represent the impact of two different tasks on the inferences that are made about the characters.  Freeman and Ambady have shown that asking subjects to focus on identifying the sex of an individual tends to bias the activation of characteristics that are Gender related (e.g., Docile), whereas asking subjects to focus on the Race of an individual biases the activation of Race related characteristics (e.g., Aggressive)

The network contains information about a number of individuals and is able to use this information to make inferences about the personal characteristics of the different individuals, as well as to make *generalizations* about what Men and Women (or different Races) in general have in common, and what is unique about them. It can also tell you about the *consistency* of a given feature with being Male versus Female, or Black, White or Asian -- this is where the harmony function can be useful in assessing the total constraint satisfaction level of the network with a particular configuration of feature inputs. The network can also be used to perform *pattern completion* as a way of retrieving information about a particular individual or individuals. 

This knowledge is encoded by simply setting weights of different strengths representing the strength of the relationships among the different features in the different layers.  (c.f., the Jets and Sharks model from McClelland & Rumelhart, 1988). You can explore the weights by clicking on the `r.wt` and `s.wt` buttons in the Netview and then click on various nodes to see what other nodes they are connected to and how strongly.  

Each of the groups of features are represented within distinct layers that have their own within-layer inhibition. We use the `FFFB` inhibitory function here which allows considerable flexibility in the actual number of active units per layer.

## Exploration

Take some time to examine the weights in the network, and verify that the weights implement the knowledge shown in the table. To do so, select the `r.Wt` value in the `NetView` and then click on individual neurons in the different layers. One can do the same thing with `s.Wt` button in the `Netview`.

We can first examine how the network responds to different patterns of Race and Gender cues and different tasks.  When you press on the `Test Trial` button it will sequentially go through the 18 items in the TestPats, which represent all combinations of the Race, Gender and Task cues.  If you click on the Act button in the Netview you can see how the activation develops in the network.  If you want to look at the activation for a single instance you can click on the TstCycPlot button to show the Cycle Plot, which will show the activation of different layers and different statistics over the 100 cycles of activation. Clicking the checkbox next to a variable in the Plot window will turn the display of that variable on and off.  

You should see that the network activates the appropriate features for the different individuals. You can think about this process as finding the most harmonious activation state given the input constraints of the input features, and the constraints in the network's weights. Equivalently, you can think about it as settling into the relevant attractor.

## Constraint Satisfaction

Now let's make use of some of the constraint satisfaction ideas. We can view the *harmony* of the network over cycles of settling using a graph view.

* Go back to testing just one instance in the TestPats table and select the `TstCycPlot` tab to view a plot of harmony over cycles of settling.

Notice that, as we expected, this value appears to monotonically increase over settling, indicating that the network is increasingly satisfying the constraints as the activations are updated.



NOTE: Could we modify the network so we could also ask it what the Gender and Race tend to be of someone who is Docile or Aggressive?  Might have to change the layers to all be Input layers, with no clamping


Now, let's see how this network can give us general information about men versus women.








