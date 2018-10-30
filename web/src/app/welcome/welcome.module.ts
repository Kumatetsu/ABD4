import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SlideshowModule } from 'ng-simple-slideshow';
import { RouterModule } from '@angular/router';
import { AppRoutingModule } from '../app-routing.module';

// Components
import { WelcomeComponent } from './welcome/welcome.component';
import { CarousselComponent } from './caroussel/caroussel.component';
import { IntroComponent } from './intro/intro.component';
import { TileComponent } from './tile/tile.component';

@NgModule({
  imports: [
    CommonModule,
    SlideshowModule,
    RouterModule,
    AppRoutingModule
  ],
  declarations: [WelcomeComponent, CarousselComponent, TileComponent, IntroComponent],
  exports: [WelcomeComponent]
})
export class WelcomeModule { }
