import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Strategy } from '../model/strategy.model';

@Injectable({
  providedIn: 'root'
})
export class StrategyService {

  apiUrl = "http://localhost:8080/v1/strategy"
  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    })
  }

  constructor(
    private httpClient: HttpClient
  ) { }

  public getAllStrategies(): Observable<Strategy[]> {
    return this.httpClient.get<Strategy[]>(this.apiUrl)
  }
}
