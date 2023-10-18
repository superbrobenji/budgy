export default (param: string) => {
    const string = param.toString();
    return string.replace(/[\t\r\n]|(--[^\r\n]*)|(\/\*[\w\W]*?(?=\*)\*\/)|['"*;\\]/gi, "");
}
