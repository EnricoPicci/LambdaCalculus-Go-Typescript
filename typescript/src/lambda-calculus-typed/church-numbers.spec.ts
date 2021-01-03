import { expect } from 'chai';
import { boolJs } from './church-booleans';
import {
    church_0,
    church_1,
    church_2,
    church_3,
    church_4,
    isZero,
    mult,
    numberJs,
    power,
    pred,
    sub,
    sum,
} from './church-numbers';

describe(`Test arithmetic functions on Church encoded natural numbers`, () => {
    it(`1 + 3 should be 4`, () => {
        const result = sum(church_1)(church_3);
        const resJs = numberJs(result);
        expect(resJs).eq(4);
    });
    it(`2 * 3 should be 6`, () => {
        const result = mult(church_2)(church_3);
        const resJs = numberJs(result);
        expect(resJs).eq(6);
    });
    it(`2 ** 3 should be 8`, () => {
        const result = power(church_3)(church_2);
        const resJs = numberJs(result);
        expect(resJs).eq(8);
    });
    it(`predecessor of 3 should be 2`, () => {
        const result = pred(church_3);
        const resJs = numberJs(result);
        expect(resJs).eq(2);
    });
    it(`predecessor of 1 should be 0`, () => {
        const result = pred(church_1);
        const resJs = numberJs(result);
        expect(resJs).eq(0);
    });
    it(`predecessor of 0 should be 0`, () => {
        const result = pred(church_0);
        const resJs = numberJs(result);
        expect(resJs).eq(0);
    });
    it(`3 - 1 should be 2`, () => {
        const result = sub(church_3)(church_1);
        const resJs = numberJs(result);
        expect(resJs).eq(2);
    });
    it(`3 - 0 should be 3`, () => {
        const result = sub(church_3)(church_0);
        const resJs = numberJs(result);
        expect(resJs).eq(3);
    });
    it(`3 - 4 should be 0`, () => {
        const result = sub(church_3)(church_4);
        const resJs = numberJs(result);
        expect(resJs).eq(0);
    });
    it(`0 - 4 should be 0`, () => {
        const result = sub(church_3)(church_4);
        const resJs = numberJs(result);
        expect(resJs).eq(0);
    });
});

describe(`When using isZero`, () => {
    it(`with 0 it should return true`, () => {
        const result = isZero(church_0);
        const resJs = boolJs(result);
        expect(resJs).to.be.true;
    });
    it(`with  number not 0 it should return false`, () => {
        const result = isZero(church_3);
        const resJs = boolJs(result);
        expect(resJs).to.be.false;
    });
});
