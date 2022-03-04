const fs = require('fs'),
    path = require('path');

/**
 * Calls callback on every file under rootFolder tree.
 * If fileExtensions array not empty, it will only call the callback on files with the extensions listed.
 * The callback will be passed an object {path: <pathToFile>, isFolder: <if path specifies a folder>}.
 * Remove '..' folders from path string.
 * @param {string} rootFolder Path to folder.
 * @param {Array.<string>} fileExtensions Array of file extensions to exclusively include (Ex: ['.txt', '.doc']).
 * @param {function(Object)} callback Callback function passed object with members 'path' string to file and 'isFolder' boolean indicating if folder.
 */
const forFolder = function(rootFolder, fileExtensions, callback) {
    // For each file (and folder) in rootFolder.
    fs.readdirSync(rootFolder).forEach( file => {
        let filePath = path.join(rootFolder, file);

        // If directory, call forFolder() on it.
        if(fs.statSync(filePath).isDirectory()) {
            callback({path: filePath, isFolder: true});
            forFolder(filePath, fileExtensions, callback);
        }
        else {
            // If user specified file extensions, only pass those to the callback.
            if( (fileExtensions.length === 0) || (fileExtensions.includes(path.extname(filePath))) ) {
                callback({path: filePath, isFolder: false});
            }
        }
    });
};



/**
 * Remove '..' folders from path string.
 * @param {string} pathStr File path string.
 * @returns {string} Converted path with no '..' folders.
 */
const removeRelParents = function(pathStr) {
    // If Windows-style path separator.
    if(path.sep === '\\') {
        return pathStr.replace(new RegExp(/\.\.\\/, 'g'), '');
    }
    // Linux/Mac style separator.
    else if(path.sep === '/') {
        return pathStr.replace(new RegExp(/\.\.\//, 'g'), '');
    }
};



/**
 * Writes file, mirroring folder structure.
 * @param {string} outputRoot Root folder path to place output folder tree.
 * @param {string} pathToInputFile Path to input file.
 * @param {string} fileContents Contents to write.
 */
const writeFile = function (outputRoot, pathToInputFile, fileContents) {
    const outputPathWithFile = path.join(outputRoot, removeRelParents(pathToInputFile)),
        outputPathWithoutFile = path.dirname(outputPathWithFile);

    // If path up to file does not exist, create it.
    if(!fs.existsSync( outputPathWithoutFile )) {
        fs.mkdirSync( outputPathWithoutFile, { recursive: true } );
    }

    // Write output to outputPathWithFile.
    fs.writeFileSync( outputPathWithFile, fileContents);
};



module.exports = {
    forFolder: forFolder,
    writeFile: writeFile
};