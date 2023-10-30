import { Injectable } from '@angular/core';
import { Strategy } from '../model/strategy.model';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class GlobalService {
  private strategyGlobal = new BehaviorSubject<Strategy | null>(null)

  setStrategyGlobal(value: Strategy) {
    this.strategyGlobal.next(value);
  }

  getStrategyGlobal(): Observable<Strategy | null> {
    return this.strategyGlobal.asObservable();
  }
}
