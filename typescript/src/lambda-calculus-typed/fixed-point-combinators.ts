import { LambdaFunction } from './function-types';

// export const Y = (g: UnaryFunction<any, any>) => ((x) => g(x(x)))((x: UnaryFunction<any, any>) => g(x(x)));
export const Y = (g: LambdaFunction) => ((x) => g(x(x)))((x) => g(x(x)));
export const Z = (g: LambdaFunction) => ((x) => g((v) => x(x)(v)))((x) => g((v) => x(x)(v)));
