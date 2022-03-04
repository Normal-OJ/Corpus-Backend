This folder includes a JavaScript program (zhoseg.js) for segmenting Mandarin words in a CHAT file.  It uses CEDict which is included in JS format in /js/lib.

To run, you need to first install Node from https://nodejs.org.  
Check for installation by typing:  node --version


=================================================================
To segment all the files in the chatFiles folder, make sure your current directory is where zhoseg.js is located and then type this command into a terminal:

node zhoseg.js MODE SEGMENT PATH_TO_CHAT_FILE_FOLDER OUTPUT_FOLDER

Where:
MODE = ss (Simplified char input -> Simplified char output
MODE = tt (Traditional char input -> Traditional char output
MODE = st (Simplified char input -> Traditional char output
MODE = ts (Traditional char input -> Simplified char output
MODE = sp (Simplified char input -> Pinyin output
MODE = tp (Traditional char input -> Pinyin output

SEGMENT = seg (Segment Chinese characters
SEGMENT = noseg (Do not segment Chinese characters


This will recursively descend the PATH_TO_CHAT_FILE_FOLDER and create the same folder tree under OUTPUT_FOLDER with segmented versions of your files based on maximal match lookup from CEDict.

The zhoseg program works by parsing a string left to right, finding the longest sequence of characters that has an entry in the dictionary.  Once it finds an matching entry, it places space characters around it in the output file. If a character in the file has no match in the dictionary, it is included as is.


If there are Chinese words that zhoseg segments that you wish to not be segmented, include them in the dictRemove.txt file.  Each item listed here (each followed by a return char) is removed from the dictionary.


If there are Chinese words that zhoseg DOES NOT group into a segment that you wish to be segmented, include them in the dictAdd.txt file.  Items listed here will be added to the dictionary.  The dictAdd.txt file is a comma-separated file where the first entry in each row is the simplifed character string, followed by the traditional character equivalent.

simplifiedCharacters, traditionalCharacters