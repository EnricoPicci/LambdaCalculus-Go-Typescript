import { expect } from 'chai';
import { and, boolJs, church_false, church_true, or } from './church-booleans';

describe(`Test boolean "and" operator on Church encoded booleans`, () => {
    it(`true && true should be true`, () => {
        const result = and(church_true)(church_true);
        const resJs = boolJs(result);
        expect(resJs).to.be.true;
    });
    it(`true && false should be false`, () => {
        const result = and(church_true)(church_false);
        const resJs = boolJs(result);
        expect(resJs).to.be.false;
    });
    it(`false && true should be false`, () => {
        const result = and(church_false)(church_true);
        const resJs = boolJs(result);
        expect(resJs).to.be.false;
    });
    it(`false && false should be false`, () => {
        const result = and(church_false)(church_false);
        const resJs = boolJs(result);
        expect(resJs).to.be.false;
    });
});

describe(`Test boolean "or" operator on Church encoded booleans`, () => {
    it(`true || true should be true`, () => {
        const result = or(church_true)(church_true);
        const resJs = boolJs(result);
        expect(resJs).to.be.true;
    });
    it(`true || false should be true`, () => {
        const result = or(church_true)(church_false);
        const resJs = boolJs(result);
        expect(resJs).to.be.true;
    });
    it(`false || true should be true`, () => {
        const result = or(church_false)(church_true);
        const resJs = boolJs(result);
        expect(resJs).to.be.true;
    });
    it(`false || false should be false`, () => {
        const result = or(church_false)(church_false);
        const resJs = boolJs(result);
        expect(resJs).to.be.false;
    });
});
