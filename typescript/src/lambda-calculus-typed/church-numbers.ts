import { church_false, church_true } from './church-booleans';
import { first, newPair, Pair } from './pair';

// Church encoded natural number type definition
export type ChurchNumber = (f: (x: any) => any) => (x: any) => any;

// Arithmetic operations on Church encoded natural numbers
export const succ = (n: ChurchNumber) => (f: (x: any) => any) => (x: any) => f(n(f)(x)) as ChurchNumber;
export const sum = (n: ChurchNumber) => (m: ChurchNumber) => (f: (x: any) => any) => (x: any) =>
    n(succ(m))(f)(x) as ChurchNumber;
export const mult = (n: ChurchNumber) => (m: ChurchNumber) => (f: (x: any) => any) => (x: any) =>
    n(m(f))(x) as ChurchNumber;
export const power = (n: ChurchNumber) => (m: ChurchNumber) => (f: (x: any) => any) => (x: any) =>
    n(m)(f)(x) as ChurchNumber;

// Sutraction implemented using pairs
const nextPair = (pair: Pair<ChurchNumber, ChurchNumber>) => pair((_x) => (y) => newPair(y)(succ(y)));
export const pred = (n: ChurchNumber) => n(nextPair)(newPair(church_0)(church_0))(first) as ChurchNumber;
export const sub = (n: ChurchNumber) => (m: ChurchNumber) => m(pred)(n);

const always_false = ((x) => (_y: any) => x)(church_false);
export const isZero = (n: ChurchNumber) => n(always_false)(church_true);

// first 10 natural numbers
export const church_0: ChurchNumber = (_f) => (x) => x;
export const church_1 = succ(church_0);
export const church_2 = succ(church_1);
export const church_3 = succ(church_2);
export const church_4 = succ(church_3);
export const church_5 = succ(church_4);
export const church_6 = succ(church_5);
export const church_7 = succ(church_6);
export const church_8 = succ(church_7);
export const church_9 = succ(church_8);

// numberJs converts a church encoded number to a javascript number
export const numberJs = (n: ChurchNumber) => n((x) => x + 1)(0);
