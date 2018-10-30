import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { WelcomeComponent } from './welcome/welcome/welcome.component';
import { ContactComponent } from './contact/contact/contact.component';
import { ServpageComponent } from './servpage/servpage/servpage.component';

const routes: Routes = [
  { path: 'welcome', component: WelcomeComponent },
  { path: 'contact', component: ContactComponent },
  { path: 'service', component: ServpageComponent },
  {path: '**', redirectTo: 'welcome', pathMatch: 'full'},
  {path: '', redirectTo: 'welcome',  pathMatch: 'full'}
];
@NgModule({
  imports: [
    RouterModule,
    RouterModule.forRoot(routes)
   ],
  exports: [ RouterModule ]
})
export class AppRoutingModule { }
