import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

import { AppModule } from './app/app.module';
import { bootstrapApplication } from '@angular/platform-browser';
import { provideAnimations } from '@angular/platform-browser/animations';
import { provideHttpClient } from '@angular/common/http';
import { importProvidersFrom } from '@angular/core';
import { MatNativeDateModule } from '@angular/material/core';
// import { SidenavComponent } from './app/views/strategy/sidenav/sidenav.component';


platformBrowserDynamic().bootstrapModule(AppModule)
  .catch(err => console.error(err));

// bootstrapApplication(SidenavComponent, {
//   providers: [
//     provideAnimations(),
//     provideHttpClient(),
//     importProvidersFrom(MatNativeDateModule)
//   ]
// }).catch(err => console.error(err));