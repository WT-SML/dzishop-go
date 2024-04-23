import { pinyin } from 'pinyin-pro';

/**
 * 比较函数
 * @author Ash
 * @param {String} propertyName 比较的字段
 * @param {Boolean} isAscending 是否升序
 * @return {Number}
 */
export const compare = (propertyName, isAscending = 0, isFilename = false) => {
    if (isAscending === undefined) {
        isAscending = 1
    } else {
        isAscending = isAscending ? 1 : -1
    }
    return function (obj1, obj2) {
        let value1, value2
        if (isFilename) {
            let char1 = obj1[propertyName][0].toLocaleLowerCase()
            let char2 = obj1[propertyName][0].toLocaleLowerCase()
            value1 = isChineseCharacter(char1) ? pinyin(char1)[0] : char1
            value2 = isChineseCharacter(char2) ? pinyin(char2)[0] : char2
            console.log(value1)
            console.log(value2)
        } else {
            value1 = obj1[propertyName]
            value2 = obj2[propertyName]
        }
        if (value1 < value2) {
            return isAscending * -1
        } else if (value1 > value2) {
            return isAscending * 1
        } else {
            return 0
        }
    }
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
function isNumber(obj) {
    if (parseFloat(obj).toString() == "NaN") {
        return false;
    }
    return true;
}

/**
 * 比较两个数字数组
 *
 * @param numberArray1
 * @param numberArray2
 */
function compareNumberArray(numberArray1, numberArray2) {
    for (let i = 0; i < numberArray1.length; i++) {
        if (numberArray2.length < i + 1) { // 此时数字数组2比1短，直接返回
            return 1;
        }
        let compareResult = parseInt(numberArray1[i]) - parseInt(numberArray2[i]);
        if (compareResult != 0) {
            return compareResult;
        }
    }
    // 说明数组1比数组2短，返回小于
    return -1;
}
/**
 * 解析json字符串
 * @author Ash
 * @return {Number}
 */
export const p = (jsonstr) => {
    return JSON.parse(jsonstr)
}