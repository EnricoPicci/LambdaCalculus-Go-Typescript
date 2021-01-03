import { church_true, church_false, ChurchBoolean } from './church-booleans';
import { ChurchNumber, church_0, succ } from './church-numbers';
import { Z } from './fixed-point-combinators';
import { BinaryFunction, LambdaFunction, UnaryFunction } from './function-types';
import { first, second } from './pair';

// List type definitions
export type List<T> = (f: ListSelector<T>) => T | List<T>;

// List constructor type definition
export type ListConst<T> = BinaryFunction<T, List<T>, List<T>>;

// Selectors for List definition it returns either head or tail
export type Head<T> = UnaryFunction<List<T>, T>;
export type Tail<T> = UnaryFunction<List<T>, List<T>>;
export type IsEmptyChecker = BinaryFunction<any, any, ChurchBoolean>;
export type ListSelector<T> = BinaryFunction<T, List<T>, T | List<T>>;

// Type of the ElementAt
export type ElementAt<T> = (n: ChurchNumber) => (l: List<T>) => List<T>;

// the empty list
export const emptyList: List<any> = (_f) => church_true;

// Functions operating on Lists
export const isEmpty = (l: List<any>) => l((_h: any) => (_t: any) => church_false);
export const insertOnTop: ListConst<any> = (x) => (l) => (f) => f(x)(l);
export const value: Head<any> = (l) => l(first);
export const next: Tail<any> = (l) => isEmpty(l)(emptyList)(l(second));
export const elementAt: ElementAt<any> = (n) => (l) => n(next)(l);

// F_listSize and F_lastElement are implemented passing 2 functions as parameters to "isEmpty(l)"
// we pass functions instead of values, to avoid Javascript (given its eager nature) to evaluate the expressions returning the values
// the second of these expressions, i.e. the one that is returned by "isEmpty(l)" when the list is not empty, contains a reference
// to the function "f" passed as parameter, which is what causes the infinte loop
const F_listSize = (f: LambdaFunction) => (l: List<any>) => (n: ChurchNumber) =>
    isEmpty(l)((n: ChurchNumber) => n)((n: ChurchNumber) => f(next(l))(succ(n)))(n);
const _listSize = Z(F_listSize);
export const listSize = (l: List<any>) => _listSize(l)(church_0);

const F_lastElement = (f: LambdaFunction) => (l: List<any>) => (n: ChurchNumber) => (prevL: List<any>) =>
    isEmpty(l)((_n: ChurchNumber) => prevL)((n: ChurchNumber) => f(next(l))(succ(n))(l))(n);
const _lastElement = Z(F_lastElement);
export const lastElement = (l: List<any>) => _lastElement(l)(church_0)(emptyList);
