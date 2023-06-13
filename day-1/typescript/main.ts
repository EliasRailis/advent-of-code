import * as fs from 'fs';

let total:number = 0;
let arr:number[] = [];

const file = fs.readFileSync('../inputs/input.txt', 'utf-8');
const numbers = file.split('\r\n');
console.log(numbers)