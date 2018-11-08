import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { FormsModule} from '@angular/forms';

import { HttpClientModule} from '@angular/common/http';
import { NotifierModule, NotifierService } from 'angular-notifier';


import { AppComponent } from './app.component';
import { WebbaseModule } from './webbase/webbase.module';
import { AppRoutingModule } from './app-routing.module';

import { TarifComponent } from './tarif/tarif.component';   
import { TarifService } from './services/tarif.service';

import { ThemeComponent} from './theme/theme.component'
import { ThemeService} from './services/theme.service'
 

@NgModule({
  declarations: [
    AppComponent,
    TarifComponent,
    ThemeComponent
  ],
  imports: [
    BrowserModule,
    WebbaseModule,
    RouterModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule, 
    NotifierModule
  ],
  providers: [
    TarifService,
    NotifierService,
    ThemeService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
