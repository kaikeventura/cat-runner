import { Component, OnInit } from '@angular/core';
import { StrategyService } from 'src/app/shared/service/strategy.service';

@Component({
  selector: 'app-strategies-list',
  templateUrl: './strategies-list.component.html',
  styleUrls: ['./strategies-list.component.css']
})
export class StrategiesListComponent implements OnInit {

  strategies: String[] = []

  constructor(
    public strategyService: StrategyService
  ) { }

  ngOnInit() {
    this.getStrategies()
  }

  getStrategies() {
    this.strategyService.getAllStrategies().subscribe(data => {
      this.strategies = data
      console.log(this.strategies)
    })
  }
}
