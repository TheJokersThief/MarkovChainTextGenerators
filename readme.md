#Simple Markov Chain Text Generators
These are a set of simple 2nd order Markov Chain Text Generators I've made in different languages.


Whenever I'm bored, I like to learn the basics of a new language and I've found that one of the best ways to do that is go through the process of creating a Markov Chain. It just happens to incorporate a lot of different basic elements and it's pretty rewarding to finish each time too. The emphasis of this project isn't necessarily to create the most efficient or correct version of the program but to learn how to use all the elements of the new language I'm learning.


### Test Text
The test text I use is 1984 by George Orwell. If you would like your own copy, the text is in the public domain and can either be downloaded here or from [Project Gutenberg Australia](http://gutenberg.net.au/ebooks01/0100021.txt)

## What is a Markov Chain Text Generator?

A markov chain is basically a very basic form of "machine learning" that predicts a future action based on the results of the past. Then a text generator is just a something practical on top of it to demonstrate its functionality. 

Each of the programs reads in the given text, sets up a table of predictions and then randomly chooses two words to start with. After that word, it goes to the table and randomly picks one of the words that's likely to follow and uses that along with the second word to predict the next word and so on.

### How is randomly choosing a word from the possibilities, "predicting" the word?

Well, we've scanned the text and taken all pairs of words and then created a list of all the words that immediately follow that pair - even if there are duplicates. 

Why do we leave in duplicates? So that if a word is more likely to occur, its probability is weighted in order to make it more probable to be picked. 

It's like going to a local community raffle. Everyone's name is put in the box exactly once. However, you've recently won a competition to put your name in the box 4 times. Now you're 4 times more likely to be picked but there's still a chance that 1 of the other names will be picked too.