import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HomeComponent } from './views/home/home.component';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatIconModule} from '@angular/material/icon';
import {MatCardModule} from '@angular/material/card';
import {MatButtonModule} from '@angular/material/button';
import {MatSidenavModule} from '@angular/material/sidenav';
import {LayoutModule} from '@angular/cdk/layout';
import {MatListModule} from '@angular/material/list';
import { StrategiesListComponent } from './views/home/strategies-list/strategies-list.component';
import { HttpClientModule } from '@angular/common/http';
import { FlexLayoutModule } from '@angular/flex-layout';
import { LocalDateTimePipe } from './shared/pipe/local-date-time.pipe';
import { StrategyComponent } from './views/strategy/strategy.component';
import { NavbarComponent } from './views/navbar/navbar.component';
import { SidenavComponent } from './views/strategy/sidenav/sidenav.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatNativeDateModule } from '@angular/material/core';
import { MaterialExampleModule } from './material.module';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    StrategiesListComponent,
    LocalDateTimePipe,
    StrategyComponent,
    NavbarComponent,
    SidenavComponent
  ],
  imports: [
    HttpClientModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatToolbarModule,
    MatIconModule,
    MatCardModule,
    MatButtonModule,
    FlexLayoutModule,
    MatSidenavModule,
    LayoutModule,
    MatListModule,
    FormsModule,
    MatNativeDateModule,
    ReactiveFormsModule,
    MaterialExampleModule
  ],
  providers: [
    LocalDateTimePipe
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
