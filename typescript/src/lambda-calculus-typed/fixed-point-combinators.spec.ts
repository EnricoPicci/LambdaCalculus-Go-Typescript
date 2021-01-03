import { expect } from 'chai';
import { church_1, church_4, isZero, mult, numberJs, sub } from './church-numbers';
import { Z } from './fixed-point-combinators';

describe(`Test Z fixed point combinator`, () => {
    it(`calculate factorial of a number with Z combinator`, () => {
        const F = (f) => (x) => (x === 0 ? 1 : x * f(x - 1));
        const F_Z = Z(F);
        const result = F_Z(4);
        expect(result).eq(24);
    });
    it(`calculate factorial of a number with Z combinator using only Lambda Calculus`, () => {
        // to avoid infinite loop that comes because of the eager nature or Typescript (Javascript),
        // rather than passing "church_1" and "mult(n)(f(sub(n)(church_1)))" as parameters of "isZero(n)"
        // we must pass 2 functions like these "n => church_1" and "n => mult(n)(f(sub(n)(church_1)))"
        // in this way we avoid Javascript trying to evaluate "mult(n)(f(sub(n)(church_1)))" at each iteration,
        // which is what causes the infinite loop
        // since it contains a reference to the "f" function which is passed in as parameter
        //
        const F = (f) => (n) => isZero(n)((_n) => church_1)((n) => mult(n)(f(sub(n)(church_1))))(n);
        const F_Z = Z(F);
        const result = numberJs(F_Z(church_4));
        expect(result).eq(24);
    });
});
