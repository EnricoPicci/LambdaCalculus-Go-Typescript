import { expect } from 'chai';
import { boolJs } from './church-booleans';
import { church_0, church_1, church_2, church_3, church_8, numberJs } from './church-numbers';
import { elementAt, emptyList, insertOnTop, isEmpty, lastElement, listSize, next, value } from './list';

describe(`when a list is empty`, () => {
    it(`it should result empty`, () => {
        const result = isEmpty(emptyList);
        const resJs = boolJs(result);
        expect(resJs).eq(true);
    });
    it(`its value should be true`, () => {
        const val = value(emptyList);
        const resJs = boolJs(val);
        expect(resJs).eq(true);
    });
    it(`its next element should be the empty list`, () => {
        const nextElement = next(emptyList);
        const result = isEmpty(nextElement);
        const resJs = boolJs(result);
        expect(resJs).eq(true);
    });
    it(`its size should be church number 0`, () => {
        const size = listSize(emptyList);
        const resJs = numberJs(size);
        expect(resJs).eq(0);
    });
    it(`its last element should be the empty list`, () => {
        const lastEl = lastElement(emptyList);
        const result = isEmpty(lastEl);
        const resJs = boolJs(result);
        expect(resJs).eq(true);
    });
});

describe(`when a list has one element`, () => {
    const aValue = 0;
    const aList = insertOnTop(aValue)(emptyList);
    it(`it should result NOT empty`, () => {
        const result = isEmpty(aList);
        const resJs = boolJs(result);
        expect(resJs).eq(false);
    });
    it(`its value should be the value added to it`, () => {
        const result = value(aList);
        expect(result).eq(aValue);
    });
    it(`its next element should be the empty list`, () => {
        const result = next(aList);
        const resJs = boolJs(isEmpty(result));
        expect(resJs).eq(true);
    });
    it(`its size should be church number 1`, () => {
        const size = listSize(aList);
        const resJs = numberJs(size);
        expect(resJs).eq(1);
    });
    it(`its last element should have the value added to it`, () => {
        const lastEl = lastElement(aList);
        const result = value(lastEl);
        expect(result).eq(aValue);
    });
});

describe(`when a list has three elements`, () => {
    const val_0 = 'a';
    const val_1 = 'b';
    const val_2 = 'c';
    const aList = insertOnTop(val_0)(insertOnTop(val_1)(insertOnTop(val_2)(emptyList)));
    it(`the first element should be the list itself`, () => {
        const result = value(aList);
        expect(result).eq(val_0);
    });
    it(`the next element after the first should be the second element added`, () => {
        const el_1 = next(aList);
        const result = value(el_1);
        expect(result).eq(val_1);
    });
    it(`the next of the next element after the first should be the third element added`, () => {
        const el_2 = next(next(aList));
        const result = value(el_2);
        expect(result).eq(val_2);
    });
    it(`applying 3 times next should return the empty list`, () => {
        const el_3 = next(next(next(aList)));
        const resJs = boolJs(isEmpty(el_3));
        expect(resJs).eq(true);
    });
    it(`applying 4 times next should return the empty list`, () => {
        const el_3 = next(next(next(next(aList))));
        const resJs = boolJs(isEmpty(el_3));
        expect(resJs).eq(true);
    });
    it(`the first element should be retrieved using elementAt with Church number 0`, () => {
        const element = elementAt(church_0)(aList);
        const result = value(element);
        expect(result).eq(val_0);
    });
    it(`the second element should be retrieved using elementAt with Church number 1`, () => {
        const element = elementAt(church_1)(aList);
        const result = value(element);
        expect(result).eq(val_1);
    });
    it(`the third element should be retrieved using elementAt with Church number 2`, () => {
        const element = elementAt(church_2)(aList);
        const result = value(element);
        expect(result).eq(val_2);
    });
    it(`the fourth element retrieved using elementAt with Church number 3 should be the empty list`, () => {
        const element = elementAt(church_3)(aList);
        const resJs = boolJs(isEmpty(element));
        expect(resJs).eq(true);
    });
    it(`any element retrieved using elementAt with Church number greater than 2 should be the empty list`, () => {
        const element = elementAt(church_8)(aList);
        const resJs = boolJs(isEmpty(element));
        expect(resJs).eq(true);
    });
    it(`its size should be church number 3`, () => {
        const size = listSize(aList);
        const resJs = numberJs(size);
        expect(resJs).eq(3);
    });
    it(`its last element should have the value equal to the first value added to it`, () => {
        const lastEl = lastElement(aList);
        const result = value(lastEl);
        expect(result).eq(val_2);
    });
});
