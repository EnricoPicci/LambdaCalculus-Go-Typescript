// Types of functions
export type UnaryFunction<T, R> = (x: T) => R;
export type BinaryFunction<T, U, R> = (x: T) => (y: U) => R;
export type TrinaryFunction<T, U, V, R> = (x: T) => (y: U) => (z: V) => R;

export type LambdaFunction = (x: any) => any;
