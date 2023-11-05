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
import {MatTooltipModule} from '@angular/material/tooltip';
import {MatTabsModule} from '@angular/material/tabs';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import {MatTableModule} from '@angular/material/table';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatSelectModule} from '@angular/material/select';
import {MatDialogModule} from '@angular/material/dialog';
import {MatSnackBarModule} from '@angular/material/snack-bar';
import { StrategiesListComponent } from './views/home/strategies-list/strategies-list.component';
import { HttpClientModule } from '@angular/common/http';
import { FlexLayoutModule } from '@angular/flex-layout';
import { LocalDateTimePipe } from './shared/pipe/local-date-time.pipe';
import { StrategyComponent } from './views/strategy/strategy.component';
import { NavbarComponent } from './views/navbar/navbar.component';
import { SidenavComponent } from './views/strategy/sidenav/sidenav.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatNativeDateModule } from '@angular/material/core';
import { HttpMethodColorDirective } from './shared/directive/http-method-color.directive';
import { RequesterComponent } from './views/strategy/requester/requester.component';
import { HttpFormComponent } from './views/strategy/requester/http-form/http-form.component';
import { KeyValueDialogComponent } from './views/strategy/requester/http-form/key-value-dialog/key-value-dialog.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    StrategiesListComponent,
    LocalDateTimePipe,
    StrategyComponent,
    NavbarComponent,
    SidenavComponent,
    HttpMethodColorDirective,
    RequesterComponent,
    HttpFormComponent,
    KeyValueDialogComponent
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
    MatTooltipModule,
    MatTabsModule,
    MatGridListModule,
    MatSlideToggleModule,
    MatTableModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatDialogModule,
    MatSnackBarModule
  ],
  providers: [
    LocalDateTimePipe
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
