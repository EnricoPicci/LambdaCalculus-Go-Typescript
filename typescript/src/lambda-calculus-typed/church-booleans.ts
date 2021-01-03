import { BinaryFunction } from './function-types';

// Church encoded boolean type definition
export type ChurchBoolean = BinaryFunction<any, any, any>;

// Boolean operators
export const and = (p: ChurchBoolean) => (q: ChurchBoolean) => p(q)(p);
export const or = (p: ChurchBoolean) => (q: ChurchBoolean) => p(p)(q);

// true and false
export const church_true: ChurchBoolean = (x: any) => (_y: any) => x;
export const church_false: ChurchBoolean = (_x: any) => (y: any) => y;

// boolJs converts a church encoded boolean to a javascript boolean
export const boolJs = (b: ChurchBoolean) => b(true)(false);
