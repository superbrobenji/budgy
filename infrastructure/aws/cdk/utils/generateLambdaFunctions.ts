import { readdirSync } from 'fs';
import path = require('path');

const getFileNamesAndDirectories = (dir: string): { fileNames: string[], directories: string[] } => {
    const serviceDirectories = getDirectories(dir);
    let fileNames: string[] = [];
    let directories: string[] = [];
    for (const serviceDirectory of serviceDirectories) {
        const file = serviceDirectory + 'Lambda' + '.go';
        const filePath = path.join(dir, serviceDirectory);

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
