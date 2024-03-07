import { readdirSync } from 'fs';
import path = require('path');
const getFileNamesAndDirectories = (dir: string): { fileNames: string[], directories: string[] } => {
    const lambdaDirectories = getDirectories(dir);
    let fileNames: string[] = [];
    let directories: string[] = [];
    for (const directory of lambdaDirectories) {
        const filesInDirectory = readdirSync(path.join(dir, directory));
        const file = filesInDirectory[0];
        const filePath = path.join(dir, directory );

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
