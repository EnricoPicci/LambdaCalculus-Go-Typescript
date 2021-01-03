import { church_true, church_false } from './church-booleans';
import { BinaryFunction } from './function-types';

// Pair type definition
export type Pair<T, U> = (f: BinaryFunction<T, U, any>) => any;

// Pair constructor type definition
export type PairConst<T, U> = BinaryFunction<T, U, Pair<T, U>>;

// Pair selector type
export type PairSelector<T, U> = BinaryFunction<T, U, T | U>;

// generic Constructor for a Pair
export const newPair: PairConst<any, any> = (x) => (y) => (f: BinaryFunction<any, any, any>) => f(x)(y);

// Selectors
export const first: PairSelector<any, any> = church_true;
export const second: PairSelector<any, any> = church_false;
