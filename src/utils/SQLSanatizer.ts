export default (param: string) => {
    return param.replace(/[\t\r\n]|(--[^\r\n]*)|(\/\*[\w\W]*?(?=\*)\*\/)|['"*;\\]/gi, "");
}
