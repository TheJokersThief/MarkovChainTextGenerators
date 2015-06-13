file_ = open('../test_text.txt')

import markov

markov = markov.Markov(file_)

print markov.generate_markov_text()