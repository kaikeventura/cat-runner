import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { StrategyBase } from 'src/app/shared/model/strategy.model';
import { StrategyService } from 'src/app/shared/service/strategy.service';

@Component({
  selector: 'app-strategies-list',
  templateUrl: './strategies-list.component.html',
  styleUrls: ['./strategies-list.component.css']
})
export class StrategiesListComponent {

  strategies: StrategyBase[] = []

  constructor(
    private strategyService: StrategyService,
    private router: Router
  ) { }

  ngOnInit() {
    this.getStrategies()
  }

  getStrategies() {
    this.strategyService.getAllStrategies().subscribe(data => {
      this.strategies = data
    })
  }

  navigateToStrategy(strategyName: string) {
    this.router.navigate(['/strategy', strategyName]);
  }
}
