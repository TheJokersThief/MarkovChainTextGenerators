import random

class Markov( object ):

	"""
		Create the object and begin to process the given text
	"""
	def __init__(self, open_file):
		# Create a place for us to store our values
	    self.store = {}

	    # Create an instance variable of the open file
	    self.open_file = open_file

	    # Extract all the words as triples
	    self.words = self.extract_words()

	    self.word_size = len(self.words)

	    # Store our prefixes and suffixes in our "store"
	    self.store_word_probability()

	"""
		Opens the file for reading and splits it into an array of
		words
	"""
	def extract_words( self ):
		self.open_file.seek( 0 )
		data = self.open_file.read( )
		return data.split( )

	"""
		Divides the words into sets of three
	"""
	def divide_to_sets( self ):
		if len( self.words ) < 3:
			# If there are less than 3 words, don't even bother
			return
		else:
			for i in range( len( self.words ) - 2 ):
				# For every set of 3 words we can make
				
				# Pack all of the triples into a generator
				yield(self.words[i], self.words[i+1], self.words[i+2])
	"""
		Store all of our words in a dictionary
	"""
	def store_word_probability( self ):
		for word_1, word_2, word_3 in self.divide_to_sets( ):
			# We'll use the first two words as a "prefix"/key
			key = ( word_1, word_2 )

			# Then we add the remaining word based on its key
			if key in self.store:
				self.store[ key ].append( word_3 )
			else:
				self.store[ key ] = [ word_3 ]

	"""
		(The fun part)
		Generates the predicted text
	"""
	def generate_markov_text( self, word_count = 250 ):
		# Generate a random first word and get its next counterpart
		random_index = random.randint( 0, self.word_size - 3 )
		word, next_word = self.words[ random_index ], self.words[ random_index + 1 ]

		output = []
		for i in xrange( word_count ):
			# Add our first word to the output
			output.append( word )
			
			# The following is necessary to avoid KeyErrors in smaller-sized sample texts
			if ( word, next_word ) in self.store:
				word, next_word = next_word, random.choice( self.store[ ( word, next_word ) ] )
			else:
				# Pick another starting point
				random_index = random.randint( 0, self.word_size - 3 )
				word, next_word = self.words[ random_index ], self.words[ random_index + 1 ]

		output.append( next_word )
		return ' '.join( output )
