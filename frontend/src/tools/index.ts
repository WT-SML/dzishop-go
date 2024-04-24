/**
 * 对数组按key排序
 * @author Ash
 * @param {Array} array 数组
 * @param {String} key key
 * @param {String} ascending 是否升序
 * @return {Number}
 */
export const sortByKey = (array, key, ascending = true) => {
    return array.sort((a, b) => {
        if (a[key] < b[key]) {
            return ascending ? -1 : 1;
        }
        if (a[key] > b[key]) {
            return ascending ? 1 : -1;
        }
        return 0; // 如果值相等，则保持原有顺序
    });
}

// 判断是否为汉字
export const isChineseCharacter = (char) => {
    return /^[\u4E00-\u9FFF]+$/.test(char);
}

/**
 * 判断是否为数字
 * @param obj
 * @returns
 */
export const isNumber = (obj) => {
    if (parseFloat(obj).toString() == "NaN") {
        return false;
    }
    return true;
}

/**
 * 解析json字符串
 * @author Ash
 * @return {Number}
 */
export const p = (jsonstr) => {
    return JSON.parse(jsonstr)
}