import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { TarifComponent } from './tarif/tarif.component';
import { ThemeComponent } from './theme/theme.component';

const routes: Routes = [
  { path: 'tarif', component: TarifComponent },
  { path: 'theme', component: ThemeComponent },
  {path: '**', redirectTo: 'tarif', pathMatch: 'full'},
  {path: '', redirectTo: 'tarif',  pathMatch: 'full'}
];
@NgModule({
  imports: [
    RouterModule,
    RouterModule.forRoot(routes)
   ],
  exports: [ RouterModule ]
})
export class AppRoutingModule { }
