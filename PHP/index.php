<?php
	include( 'markov.php' );

	$filePath = "../test_text.txt";

	$file = fopen( $filePath, "r" ) or die( "Unable to open file!" );
	$openFile = fread( $file, filesize( $filePath ) );

	$markov = new Markov( $openFile );
	echo $markov->generateMarkovText( );
?>