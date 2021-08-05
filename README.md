## Aesthetic Garden

This is an implementation of a solution for finding the number of ways which a garden of trees could be modified
to make it aesthetically pleasing. The trees in the garden are aesthetically pleasing when they alternately increase 
and decrease in height. 


#### NOTES: 

* The smallest number of trees to create an aesthetic garden is 3. A group of 3 trees is aesthetic if the middle tree 
is either the shortest or the tallest of the group. 

* Create a function isAesthetic which checks if a group of 3 trees is aesthetic.
  
* Search the garden in batches of 3s, move by 1 tree to create a new group of trees. 

* If a group that isn't aesthetic is encountered, to get the number of ways it could be made aesthetic; 
  1. remove a tree from the group
   2. create a new temporary group by including the next tree after that group 
   3. use the `isAesthetic` method to check the group. Do this for each tree in the group.
    
* The number of ways gotten in that group is the maximum number of ways a tree can be cut to make the garden aesthetic since
only one tree is legally allowed to be cut down. Keep searching the garden in batches of 3s.
* If any other unaesthetic 
group is encountered return -1. If no other unaesthetic group is encountered, return the number of ways. If all the 
groups are aesthetic, return 0.