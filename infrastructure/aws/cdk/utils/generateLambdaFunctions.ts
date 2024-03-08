import { readdirSync } from 'fs';
import path = require('path');

const getFileNamesAndDirectories = (dir: string): { fileNames: string[], directories: string[] } => {
    const handlerDirectories = getDirectories(dir);
    let fileNames: string[] = [];
    let directories: string[] = [];
    for (const handlerDirectory of handlerDirectories) {
        const file = handlerDirectory + '.go';
        const filePath = path.join(dir, handlerDirectory);

        const fileNameLowerCase = file.split('.')[0];
        const fileName = fileNameLowerCase.charAt(0).toUpperCase() + fileNameLowerCase.slice(1);

        fileNames.push(fileName);
        directories.push(filePath);
    }
    return { fileNames, directories };
}

const getDirectories = (source: string): string[] =>
    readdirSync(source, { withFileTypes: true })
        .filter(dirent => dirent.isDirectory())
        .map(dirent => dirent.name)

export default getFileNamesAndDirectories;
