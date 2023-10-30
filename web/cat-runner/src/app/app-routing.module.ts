import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './views/home/home.component';
import { StrategyComponent } from './views/strategy/strategy.component';

const routes: Routes = [
  {
    path: '',
    component: HomeComponent
  },
  {
    path: 'strategy/:strategyName',
    component: StrategyComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
