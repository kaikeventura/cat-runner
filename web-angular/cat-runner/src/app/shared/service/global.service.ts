import { Injectable } from '@angular/core';
import { Strategy } from '../model/strategy.model';

@Injectable({
  providedIn: 'root'
})
export class GlobalService {
  strategyGlobal!: Strategy;

  setGlobalVariable(value: any) {
    this.strategyGlobal = value;
  }

  getGlobalVariable() {
    return this.strategyGlobal;
  }
}
