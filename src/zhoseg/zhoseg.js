const fs = require('fs'),
  fileUtil = require('./js/fileUtil.js'),
  chineseParser = require('./js/chineseParser.js');


/**
 * Converts a folder (recursively) of CHAT files with Chinese characters.
 * @param {string} inputFolder Path to CHAT files.
 * @param {string} outputFolder Folder to place converted file tree.
 * @param {string} fromFormat Format of Chinese characters in inputFolder (simplified ('s') or traditional ('t').
 * @param {string} toFormat Desired format of Chinese characters in outputFolder (simplified ('s') or traditional ('t').
 * @param {segmentize} boolean Whether or not to segment Chinese in outputFolder.
 */
const convertFolder = function( inputFolder, outputFolder, fromFormat, toFormat, segmentize ) {
  // Callback passed to fileUtil.forFolder().
  const folderCB = function (currentFile) {
    if(!currentFile.isFolder) {
      console.log('Processing: ' + currentFile.path);

      // Get original file contents.
      const origFile = fs.readFileSync(currentFile.path, 'utf-8');

      // Convert file.
      let convertedFile = chineseParser.segmentString(origFile, fromFormat, toFormat, segmentize);

      // Since spaces are placed before and after Chinese segments, speaker lines beginning with Chinese will have '\t' and space.
      // Replace these with just tab.
      convertedFile = convertedFile.replace(/\t /g, '\t');

      // Replace any multi-space strings with single space.
      convertedFile = convertedFile.replace(/  +/g, ' ');

      // CHAT expects a final return char.
      convertedFile += '\n';

      // Write converted file.
      fileUtil.writeFile(outputFolder, currentFile.path, convertedFile);
    }
  };

  fileUtil.forFolder(inputFolder, ['.cha'], folderCB);
};




/**
 * Verifies command line params.
 * Shows message if malformed, calling convertFolder() with params otherwise.
 */
const runZhoseg = function () {
  const usageMessage = `
    Usage:
    node zhoseg.js MODE SEGMENT PATH_TO_CHAT_FILE_FOLDER OUTPUT_FOLDER

    Where:
    MODE = ss (Simplified char input -> Simplified char output)
    MODE = tt (Traditional char input -> Traditional char output)
    MODE = st (Simplified char input -> Traditional char output)
    MODE = ts (Traditional char input -> Simplified char output)
    MODE = sp (Simplified char input -> Pinyin output)
    MODE = tp (Traditional char input -> Pinyin output)

    SEGMENT = seg (Segment Chinese characters)
    SEGMENT = noseg (Do not segment Chinese characters)
    `;

  if( !(process.argv[2] && process.argv[3] && process.argv[4] && process.argv[5])
    || ((process.argv[2] !== "ss") && (process.argv[2] !== "tt") && (process.argv[2] !== "ts")
      && (process.argv[2] !== "st") && (process.argv[2] !== "sp")  && (process.argv[2] !== "tp"))
    || ((process.argv[3] !== "seg") && (process.argv[3] !== "noseg"))) {
      console.log(usageMessage);
  }
  else {
    let segmentize, fromFormat, toFormat;

    if(process.argv[3] === "seg") {
      segmentize = true;
    }
    else {
      segmentize = false;
    }

    fromFormat = process.argv[2].split('')[0].toLowerCase();
    toFormat = process.argv[2].split('')[1].toLowerCase();

    chineseParser.dictRemove();
    chineseParser.dictAdd();

    convertFolder(process.argv[4], process.argv[5], fromFormat, toFormat, segmentize);
  }
};


runZhoseg();
