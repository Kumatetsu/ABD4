import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { AppComponent } from './app.component';
import { WebbaseModule } from './webbase/webbase.module';
import { AppRoutingModule } from './app-routing.module';
import { WelcomeModule } from './welcome/welcome.module';
import { ContactModule } from './contact/contact.module';
import { ServpageModule } from './servpage/servpage.module';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    WebbaseModule,
    RouterModule,
    AppRoutingModule,
    WelcomeModule,
    ContactModule,
    ServpageModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
