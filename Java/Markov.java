import java.util.*;
import java.io.*;
import java.util.Random;

/**
 * Generates text based on a Markov Chain
 */
public class Markov{
	public static void main( String[] args ){
		Markov markov = new Markov( );
		System.out.println( markov.generateMarkovText( 250 ) );
	}

	private final String[] words;
	private final HashMap<String, ArrayList<String>> store;
	private final String file;
	private final Object[] keys;

	/**
	 * Opens a file and calculates the probability of words 
	 */
	public Markov( ){
		this.store = new HashMap<String, ArrayList<String>>( );
		String content = "";

		try{
			content = new Scanner(new File("../test_text.txt")).useDelimiter("\\Z").next();
		} catch (FileNotFoundException e) {
			// If file not found, let someone know and exit out
		    System.out.println( "The specified file was not found" );
		    e.printStackTrace();
		    System.exit( 1 );
		}

		// Put the file's contents into an instance variable
		this.file = content;
		// Split the contents into an array of words
		this.words = content.split(" ");
		this.storeProbablity( );
		this.keys = this.store.keySet( ).toArray( );
	}

	/**
	 * Calculates the probability of succeeding words in a text
	 */
	public void storeProbablity( ){
		int wordSize = this.words.length;
		if( wordSize < 3 ){
			// If the text is too small, error out
			System.out.println( "Please use a larger sample text" );
			System.exit( 1 );
		} else {
			for (int i = 0; i < (wordSize-2); i++) {
				// For each triple of words
				String key = this.words[i] + ">>" + this.words[i+1];

				if( ! ( this.store.containsKey( key ) ) ){
					// If the store doesn't contain the key already, initialise it
					// as an empty ArrayList
					this.store.put( key, new ArrayList<String>( ) );
				}
				// Add the succeeding word to the store
				this.store.get( key ).add( this.words[i+2] );
			}
		}
	}

	/**
	 * Generates text based on a Markov Chain
	 * @param  wordLimit The number of words to be output
	 * @return           The generated text
	 */
	public String generateMarkovText( int wordLimit ){
		String output = "";

		// Get a random key to begin with
		Random random = new Random( );
		int randomInt = random.nextInt( this.store.size( ) - 1 );
		String key = (String) this.keys[ randomInt ];
		
		for (int i=0; i < wordLimit; i++) {

			// Split the words into individual words
			String[] words = key.split(">>");

			output += words[0] + " ";

			// The number of possibilities of the size of a key's array
			int possibilities = this.store.get( key ).size( );
			if( possibilities < 1 ){
				// If there is more than 1 possibility, randomly get a succeeding word
				// and use that for the next key
				int randomPossibility = random.nextInt( this.store.size( ) - 1 );
				key = words[1] + ">>" + this.store.get( key ).get( randomPossibility );
			} else {
				// If there are no possibilities, randomly get another key
				int randomPossibility = random.nextInt( this.store.size( ) - 1 );
				key = (String) this.keys[ randomPossibility ];
			}
		}
		return output.replace( "\n", " " );
	}
}