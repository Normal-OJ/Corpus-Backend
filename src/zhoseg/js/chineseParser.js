const PARAMS = {
  MAX_SEGMENT_CHARS: 10,
  PATH_TO_SIMPLIFIED_DICT: './dict/cedictBySimp.js',
  PATH_TO_TRAD_DICT: './dict/cedictByTrad.js',
  PATH_TO_DICTREMOVE_CSV: './dictRemove.txt',
  PATH_TO_DICTADD_CSV: './dictAdd.txt'
};

const fs = require('fs'),
  EOL = require('os').EOL; // Get OS-dependent end-of-line character.


// Load dictionaries.
const dictBySimp = require(PARAMS.PATH_TO_SIMPLIFIED_DICT).cedictBySimp,
  dictByTrad = require(PARAMS.PATH_TO_TRAD_DICT).cedictByTrad;
  


/**
 * Remove items from dictionary listed in dictRemove.txt.
 */
const dictRemove = function() {
  const itemsToRemove = fs.readFileSync(PARAMS.PATH_TO_DICTREMOVE_CSV, 'utf8').split(EOL);

  // Remove an item from dict if exists.
  function removeItem(dict, item) {
    if( dict.hasOwnProperty(item) ) {
      delete dict[ item ];
    }
  }

  for(let c=0; c < itemsToRemove.length; c++) {
    // Remove any whitespace.
    itemsToRemove[c] = itemsToRemove[c].replace(/\s/g, '');

    removeItem(dictBySimp, itemsToRemove[c]);
    removeItem(dictByTrad, itemsToRemove[c]);
  }
};


/**
 * Add items to dictionary listed in dictAdd.txt.
 */
const dictAdd = function() {
  const itemsToAdd = fs.readFileSync(PARAMS.PATH_TO_DICTADD_CSV, 'utf8').split(EOL);

  // Remove an item from dict if exists.
  function addItem(dict, dictType, item) {
    let simpTrad = item.split(','),
      newItem = {"s": simpTrad[0], "t": simpTrad[1], "p": "", "e": ""};

    if(dictType === 's') {
      if(dict.hasOwnProperty(simpTrad[0])) {
        dict[ simpTrad[0] ].push(newItem);
      }
      else {
        dict[ simpTrad[0] ] = [newItem];
      }
    }
    else {
      if(dict.hasOwnProperty(simpTrad[1])) {
        dict[ simpTrad[1] ].push(newItem);
      }
      else {
        dict[ simpTrad[1] ] = [newItem];
      }
    }
  }

  for(let c=0; c < itemsToAdd.length; c++) {
    // Remove any whitespace.
    itemsToAdd[c] = itemsToAdd[c].replace(/\s/g, '');

    addItem(dictBySimp, 's', itemsToAdd[c]);
    addItem(dictByTrad, 't', itemsToAdd[c]);
  }
};



/**
 * Segment and/or convert simplified/traditional Chinese characters in string.
 * @param {string} inputStr String to convert Chinese characters.
 * @param {string} fromFormat Format of Chinese characters in inputFolder (simplified ('s') or traditional ('t').
 * @param {string} toFormat Desired format of Chinese characters in outputFolder (simplified ('s') or traditional ('t').
 * @param {segmentize} boolean Whether or not to segment Chinese in outputFolder.
 * @returns {string} Converted string.
 */
const segmentString = function(inputStr, fromFormat, toFormat, segmentize) {
  let numChars = 0, segmentedStr = '',
    prevChinese = false,
    maxSeg, spaceChar,
    cedict; // Reference to dictBySimp or dictByTrad depending on user-specififed "from" parameter.


  // Set dictionary we use (dictBySimp or dictByTrad) based on fromFormat.
  if(fromFormat === 's') {
    cedict = dictBySimp;
  }
  else if(fromFormat === 't') {
    cedict = dictByTrad;
  }

  //ignoreDictItems(cedict);

  // Whether or not to add segments (spaces) on word boundaries to output.
  if(segmentize) {
    spaceChar = ' ';
  }
  else {
    spaceChar = '';
  }

  // Get the maximal segment at start of string.
  function maximalSeg(inputStr) {
    let maxSeg = '';

    // If inputStr begins with an alphanumeric char, return empty segment (empty string).
    if (inputStr.charAt(0).match(/^[0-9a-z]+$/i)) {
      return maxSeg;
    }

    for(let c=0; ((c <= PARAMS.MAX_SEGMENT_CHARS) && (c <= inputStr.length )); c++) {
      if( cedict.hasOwnProperty( inputStr.slice(0, c) )) {
        maxSeg = inputStr.slice(0, c);
      }
    }

    return maxSeg;
  }

  // Process each character in inputStr.
  while (numChars <= inputStr.length) {
    maxSeg = maximalSeg(inputStr.slice(numChars-1));

    // If inputStr beginning at numChars-1 does not match the start of any string in dictionary, add as is to segmentedStr.
    if(maxSeg.length === 0) {
      segmentedStr += inputStr.charAt(numChars-1);
      numChars++;
      prevChinese = false;
    }
    // Add a Chinese segment.
    else {
      // If previous thing added to segmentedStr was not a Chinese segment, prepend space before adding Chinese segment.
      if(!prevChinese) {
        segmentedStr += spaceChar;
      }

      // Append segment in format if exists.
      if(cedict[maxSeg][0][toFormat] !== '') {
        segmentedStr += cedict[maxSeg][0][toFormat] + spaceChar;
      }
      // Else append segment in original format.
      else {
        segmentedStr += maxSeg + spaceChar;
      }
      
      numChars += maxSeg.length;
      prevChinese = true;
    }
  }

  return segmentedStr.trim();
};



module.exports = {
  segmentString: segmentString,
  dictAdd: dictAdd,
  dictRemove: dictRemove
};



