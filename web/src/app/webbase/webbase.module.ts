import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LayoutComponent } from './layout/layout.component';
import { HeaderComponent } from './header/header.component';
import { FooterComponent } from './footer/footer.component';
import { RouterModule } from '@angular/router';
import { AppRoutingModule } from '../app-routing.module';

@NgModule({
  imports: [
    CommonModule,
    RouterModule,
    AppRoutingModule
  ],
  declarations: [
    LayoutComponent,
    HeaderComponent,
    FooterComponent
  ],
  exports: [LayoutComponent],
})
export class WebbaseModule { }
