<?php
/**
 * Used to generate markov-chain-based text
 */
class Markov{
	// Create all of our private variables
	private $store;
	private $storeSize;
	private $file;
	private $words;
	private $wordsSize;

	/**
	 * @param String $file The contents of a text file
	 */
	function __construct( $file ) {
		// Setup our file and store, then process the file's contents
		// and create a simple probability array
		$this->file = $file;
		$this->store = [];
		$this->words = $this->splitWords( );
		$this->wordsSize = count( $this->words );
		$this->storeProbability( );
		$this->storeSize = count( $this->store );
	}

	/**
	 * Splits the string into an array of all the words delimited by a space
	 * @return Array
	 */
	private function splitWords( ){
		$wordSplit = explode(' ', $this->file);
		return $wordSplit;
	}

	/**
	 * Creates a simple 2nd Order probability table for the words
	 * @return Array
	 */
	private function storeProbability( ){
		if( $this->wordsSize < 3 ){
			// If the total number of words is less than 3, we can't use this text
			return;	
		} else {
			for ($i=0; $i < ($this->wordsSize - 2); $i++) { 
				// We'll use a comma-separated string as key 
				// and explode it later rather than messing 
				// with the SplObjectStorage class
				$key = $this->words[ $i ] .','. $this->words[ $i + 1 ];
				$this->store[ $key ][] = $this->words[ $i + 2 ];
			}
		}
	}

	/**
	 * Generate the markov chain text based on the calculated store
	 * @param  integer $wordLimit The amount of words to be returned
	 * @return String             The resulting text
	 */
	public function generateMarkovText( $wordLimit = 250 ){
		$output = '';

		if( !empty( $this->store ) ){
			// If store has more than three words
			
			// Grab a random key from our array
			$key = array_rand( $this->store );
			for ($i=0; $i < $wordLimit; $i++) {
				// For until we reach the word limit

				// Extract our individual words from the key and output the first
				$words = explode(',', $key);
				$output .= ' '.$words[ 0 ];

				if( isset( $this->store[ $key ] ) ){
					// If the new key exists

					// Randomly select one of the new words and get a new key
					$randomInteger = mt_rand( 0, count( $this->store[$key] ) - 1 );
					$key = $words[ 1 ].','.$this->store[ $key ][ $randomInteger ];
				} else {
					// If it doesn't exist, grab another random one
					$key = array_rand( $this->store );
				}
			}
			// Bit of beautification for the output string
			return ucfirst( trim( $output ) ).'.';
		}
		// If the store's empty, a text with less than 3 words was supplied
		return 'Please use a text with more than 3 words.';
	}
}

?>