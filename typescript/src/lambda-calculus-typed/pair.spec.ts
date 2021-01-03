import { expect } from 'chai';
import { first, newPair, second } from './pair';

describe(`Test selector on pairs`, () => {
    const firstElement = 'a';
    const secondElement = 'b';
    const pair = newPair(firstElement)(secondElement);
    it(`first should return the first element of the pair`, () => {
        const result = pair(first);
        expect(result).eq(firstElement);
    });
    it(`second should return the second element of the pair`, () => {
        const result = pair(second);
        expect(result).eq(secondElement);
    });
});
