export default (param: string) => {
    console.log(param, typeof param)
    const string = param.toString();
    return string.replace(/[\t\r\n]|(--[^\r\n]*)|(\/\*[\w\W]*?(?=\*)\*\/)|['"*;\\]/gi, "");
}
